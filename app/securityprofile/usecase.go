package securityprofile

import (
	"ai-dekawaf/app/auditlog"
	"ai-dekawaf/app/f5"
	"ai-dekawaf/helper"
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	CreateSecurityProfile(c *gin.Context, req request.CreateSecurityProfile) (err error)
}

type usecase struct {
	repo         Repository
	auditlogRepo auditlog.Repository
	f5Repo       f5.Repository
}

func NewUsecase(repo Repository, auditlogRepo auditlog.Repository, f5Repo f5.Repository) Usecase {
	return &usecase{
		repo:         repo,
		auditlogRepo: auditlogRepo,
		f5Repo:       f5Repo,
	}
}

func (u *usecase) CreateSecurityProfile(c *gin.Context, req request.CreateSecurityProfile) (err error) {

	go u.createSecurityProfile(c, req)

	return
}

func (u *usecase) createSecurityProfile(c *gin.Context, req request.CreateSecurityProfile) {
	var err error
	resNode, err := helper.LogF5Request("CreateNode", func() (response.F5CreateNode, error) {
		return u.f5Repo.CreateNode(c, request.F5CreateNode{
			Name:    req.Node,
			Address: helper.GetIP(c),
		})
	})
	if err != nil {
		return
	}

	resPool, err := helper.LogF5Request("CreatePool", func() (response.F5CreatePool, error) {
		return u.f5Repo.CreatePool(c, request.F5CreatePool{
			Name: req.Pool,
			Members: []struct {
				Name string `json:"name"`
			}{
				{
					Name: fmt.Sprintf("%s:%s", resNode.FullPath, req.Port),
				},
			},
		})
	})
	if err != nil {
		return
	}

	crtName := fmt.Sprintf("%s.crt", helper.Slugify(req.Name))
	resUploadCert, err := helper.LogF5Request("UploadCert", func() (response.F5Upload, error) {
		return u.f5Repo.Upload(c, crtName, req.CertificateFile)
	})
	if err != nil {
		return
	}
	_, err = helper.LogF5Request("InstallCert", func() (response.F5Install, error) {
		return u.f5Repo.InstallCertOrChain(c, request.F5Install{
			Command:       "install",
			Name:          crtName,
			FromLocalFile: resUploadCert.LocalFilePath,
		})
	})
	if err != nil {
		return
	}

	chainName := fmt.Sprintf("%s-chain.crt", helper.Slugify(req.Name))
	if req.ChainFile != nil {
		resUploadChain, err := helper.LogF5Request("UploadChain", func() (response.F5Upload, error) {
			return u.f5Repo.Upload(c, chainName, req.ChainFile)
		})
		if err != nil {
			return
		}
		_, err = helper.LogF5Request("InstallChain", func() (response.F5Install, error) {
			return u.f5Repo.InstallCertOrChain(c, request.F5Install{
				Command:       "install",
				Name:          chainName,
				FromLocalFile: resUploadChain.LocalFilePath,
			})
		})
		if err != nil {
			return
		}
	}

	keyName := fmt.Sprintf("%s.key", helper.Slugify(req.Name))
	resUploadKey, err := helper.LogF5Request("UploadKey", func() (response.F5Upload, error) {
		return u.f5Repo.Upload(c, keyName, req.KeyFile)
	})
	if err != nil {
		return
	}
	_, err = helper.LogF5Request("InstallKey", func() (response.F5Install, error) {
		return u.f5Repo.InstallKey(c, request.F5Install{
			Command:       "install",
			Name:          keyName,
			FromLocalFile: resUploadKey.LocalFilePath,
		})
	})
	if err != nil {
		return
	}

	var CertKeyChain []struct {
		Name  string `json:"name"`
		Cert  string `json:"cert"`
		Key   string `json:"key"`
		Chain string `json:"chain"`
	}
	CertKeyChain = append(CertKeyChain, struct {
		Name  string `json:"name"`
		Cert  string `json:"cert"`
		Key   string `json:"key"`
		Chain string `json:"chain"`
	}{
		Name:  helper.Slugify(req.Name),
		Cert:  crtName,
		Key:   keyName,
		Chain: chainName,
	})
	if req.ChainFile == nil {
		CertKeyChain[0].Chain = ""
	}
	resCreateSslSecurityProfile, err := helper.LogF5Request("CreateSslSecurityProfile", func() (response.F5CreateSslSecurityProfile, error) {
		return u.f5Repo.CreateSslSecurityProfile(c, request.F5CreateSslSecurityProfile{
			Name:         fmt.Sprintf("%s-client-ssl", helper.Slugify(req.Name)),
			DefaultsFrom: "clientssl",
			CertKeyChain: CertKeyChain,
		})
	})
	if err != nil {
		return
	}

	_, err = helper.LogF5Request("ActiveSNI", func() (response.F5ActiveSNI, error) {
		return u.f5Repo.ActiveSNI(c, helper.FormatF5Path(resCreateSslSecurityProfile.FullPath), request.F5ActiveSNI{
			ServerName: req.Domain,
			SniDefault: "true",
		})
	})
	if err != nil {
		return
	}

	resCreateVS, err := helper.LogF5Request("CreateVS", func() (response.F5VS, error) {
		return u.f5Repo.CreateVS(c, request.F5CreateVS{
			Name:        fmt.Sprintf("%s-vs", helper.Slugify(req.Name)),
			Partition:   "Common",
			Destination: "103.125.100.112:443", // sumber tidak diketahui, perlu ditanyakan lagi
			Mask:        "255.255.255.255",
			IPProtocol:  "tcp",
			Pool:        resPool.FullPath,
			Profiles: []request.F5VSProfile{
				{
					Name:    fmt.Sprintf("%s-client-ssl", helper.Slugify(req.Name)),
					Context: "tcp",
				},
				{
					Name:    resCreateSslSecurityProfile.FullPath,
					Context: "clientside",
				},
			},
		})
	})
	if err != nil {
		return
	}

	reqF5 := request.F5CreateSecurityPolicy{}
	switch req.EnforcementMode {
	case "blocking":
		reqF5 = request.F5CreateSecurityPolicy{
			Name:                 fmt.Sprintf("%s-policy", helper.Slugify(req.Name)),
			EnforcementMode:      "blocking",
			ApplicationLanguage:  "auto-detect",
			CaseInsensitive:      false,
			ProtocolIndependent:  false,
			Active:               true,
			PolicyBuilderEnabled: true,
		}
	case "transparent":
		reqF5 = request.F5CreateSecurityPolicy{
			Name:                 fmt.Sprintf("%s-policy", helper.Slugify(req.Name)),
			EnforcementMode:      "transparent",
			ApplicationLanguage:  "auto-detect",
			CaseInsensitive:      false,
			ProtocolIndependent:  false,
			Active:               true,
			PolicyBuilderEnabled: true,
		}
	default:
		return
	}

	resCreateSecurityPolicy, err := helper.LogF5Request("CreateSecurityPolicy", func() (response.F5CreateSecurityPolicy, error) {
		return u.f5Repo.CreateSecurityPolicy(c, reqF5)
	})
	if err != nil {
		return
	}

	_, err = helper.LogF5Request("AssignSecurityPolicyToVS", func() (response.F5AssignSecurityPolicyToVS, error) {
		return u.f5Repo.AssignSecurityPolicyToVS(c, resCreateSecurityPolicy.ID, request.F5AssignSecurityPolicyToVS{
			VirtualServers: []string{
				resCreateVS.FullPath,
			},
		})
	})
	if err != nil {
		return
	}

	for _, serverTecnology := range req.ServerTecnologies {
		_, err = helper.LogF5Request(fmt.Sprintf("AssignServerTecnologyToSecurityPolicy-%s", serverTecnology), func() (response.F5AssignServerTecnologyToSecurityPolicy, error) {
			return u.f5Repo.AssignServerTecnologyToSecurityPolicy(c, resCreateSecurityPolicy.ID, request.F5AssignServerTecnologyToSecurityPolicy{
				ServerTechnologyReference: request.F5Link{
					Link: serverTecnology,
				},
			})
		})
		if err != nil {
			return
		}
	}

	_, err = helper.LogF5Request("SetLearningModePolicy", func() (response.F5SetLearningModePolicy, error) {
		return u.f5Repo.SetLearningModePolicy(c, resCreateSecurityPolicy.ID, request.F5SetLearningModePolicy{
			LearningMode: "automatic",
		})
	})
	if err != nil {
		return
	}

	reqAttachBotDefense := request.F5AttachBotDefenseandLogProfileToVS{
		SecurityLogProfiles: []string{
			"/Common/local-bot-defense",
		},
	}
	if req.BotProtection {
		reqAttachBotDefense.BotDefenseProfile.Link = "https://localhost/mgmt/tm/security/bot-defense/profile/~Common~bot-defense"
	}

	_, err = helper.LogF5Request("AttachBotDefenseandLogProfileToVS", func() (response.F5AttachBotDefenseandLogProfileToVS, error) {
		return u.f5Repo.AttachBotDefenseandLogProfileToVS(c, resCreateVS.FullPath, reqAttachBotDefense)
	})
	if err != nil {
		return
	}

	return
}
