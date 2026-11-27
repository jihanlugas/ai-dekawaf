package securityprofile

import (
	"ai-dekawaf/app/auditlog"
	"ai-dekawaf/app/f5"
	"ai-dekawaf/app/permission"
	"ai-dekawaf/app/project"
	"ai-dekawaf/app/websocket"
	"ai-dekawaf/constant"
	"ai-dekawaf/helper"
	"ai-dekawaf/middleware"
	"ai-dekawaf/model"
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Usecase interface {
	PageSecurityProfile(c *gin.Context, req request.PageSecurityProfile) (mSecurityProfiles []model.SecurityProfile, count int64, err error)
	CreateSecurityProfile(c *gin.Context, req request.CreateSecurityProfile) (err error)
	GetSecurityProfile(c *gin.Context, id string) (mSecurityProfile model.SecurityProfile, err error)
	DeleteSecurityProfile(c *gin.Context, id string) (err error)
}

type usecase struct {
	repo              Repository
	auditlogRepo      auditlog.Repository
	f5Repo            f5.Repository
	projectRepo       project.Repository
	permissionUsecase permission.Usecase
}

func NewUsecase(repo Repository, auditlogRepo auditlog.Repository, f5Repo f5.Repository, projectRepo project.Repository, permissionUsecase permission.Usecase) Usecase {
	return &usecase{
		repo:              repo,
		auditlogRepo:      auditlogRepo,
		f5Repo:            f5Repo,
		projectRepo:       projectRepo,
		permissionUsecase: permissionUsecase,
	}
}

func (u *usecase) PageSecurityProfile(c *gin.Context, req request.PageSecurityProfile) (mSecurityProfiles []model.SecurityProfile, count int64, err error) {
	// Authorize the user.
	authorized := u.permissionUsecase.CheckUserPermission(c, req.ProjectID, constant.PrivilegeDekaWaf, false)
	if !authorized {
		return nil, 0, fmt.Errorf(constant.ErrForbidden)
	}

	// Check project availability status.
	availability, err := u.projectRepo.ProjectAvailabilityStatus(req.ProjectID)
	if err != nil {
		return nil, 0, fmt.Errorf(constant.ErrCheckProjectAvailability)
	}

	if !availability {
		return nil, 0, fmt.Errorf(constant.ErrProjectNotAvailable)
	}

	return u.repo.Page(req)
}

func (u *usecase) GetSecurityProfile(c *gin.Context, id string) (mSecurityProfile model.SecurityProfile, err error) {
	mSecurityProfile, err = u.repo.GetByID(id)
	if err != nil {
		return mSecurityProfile, err
	}

	// Authorize the user.
	authorized := u.permissionUsecase.CheckUserPermission(c, mSecurityProfile.ProjectID, constant.PrivilegeDekaWaf, true)
	if !authorized {
		return mSecurityProfile, fmt.Errorf(constant.ErrForbidden)
	}

	// Check project availability status.
	availability, err := u.projectRepo.ProjectAvailabilityStatus(mSecurityProfile.ProjectID)
	if err != nil {
		return mSecurityProfile, fmt.Errorf(constant.ErrCheckProjectAvailability)
	}

	if !availability {
		return mSecurityProfile, fmt.Errorf(constant.ErrProjectNotAvailable)
	}

	return mSecurityProfile, nil
}

func (u *usecase) CreateSecurityProfile(c *gin.Context, req request.CreateSecurityProfile) (err error) {
	userID := middleware.GetUserID(c)

	// Authorize the user.
	authorized := u.permissionUsecase.CheckUserPermission(c, req.ProjectID, constant.PrivilegeDekaWaf, true)
	if !authorized {
		return fmt.Errorf(constant.ErrForbidden)
	}

	// Check project availability status.
	availability, err := u.projectRepo.ProjectAvailabilityStatus(req.ProjectID)
	if err != nil {
		return fmt.Errorf(constant.ErrCheckProjectAvailability)
	}

	if !availability {
		return fmt.Errorf(constant.ErrProjectNotAvailable)
	}

	mProject, err := u.projectRepo.GetByID(req.ProjectID)
	if err != nil {
		return err
	}

	mLog := model.Auditlog{
		UserID:         userID,
		OrganizationID: mProject.OrganizationID,
		ProjectID:      mProject.ID,
		IPAddress:      helper.GetIP(c),
		ServiceType:    "Deka Waf",
		ServiceName:    "Security Profile",
		MethodName:     "Create Security Profile",
	}

	mSecurityProfile := model.SecurityProfile{
		ID:              helper.GetUniqueID(),
		OrganizationID:  mProject.OrganizationID,
		ProjectID:       req.ProjectID,
		Name:            req.Name,
		Description:     req.Description,
		Domain:          req.Domain,
		Node:            req.Node,
		Pool:            req.Pool,
		Destination:     req.Destination,
		EnforcementMode: req.EnforcementMode,
		BotProtection:   req.BotProtection,
		Status:          model.SecurityProfileStatusPending,
	}
	err = u.repo.Create(mSecurityProfile)
	if err != nil {
		go u.auditlogRepo.ErrorToLog(mLog, "Error Create Security Profile", err.Error())
		return err
	}

	go func() {
		mSecurityProfile, err = u.createF5SecurityProfile(c, req, mSecurityProfile)
		if err != nil {
			mSecurityProfile.Status = model.SecurityProfileStatusInactive
			websocket.SendWebsocketMessage(mProject.ID, map[string]string{
				"security_profile_id":   mSecurityProfile.ID,
				"security_profile_name": mSecurityProfile.Name,
				"action":                "create_security_profile",
				"status":                "failed",
				"error":                 err.Error(),
			})
			go u.auditlogRepo.ErrorToLog(mLog, "Error Create Security Profile", err.Error())
		} else {
			mSecurityProfile.Status = model.SecurityProfileStatusActive
			websocket.SendWebsocketMessage(mProject.ID, map[string]string{
				"security_profile_id":   mSecurityProfile.ID,
				"security_profile_name": mSecurityProfile.Name,
				"action":                "create_security_profile",
				"status":                "success",
			})
			go u.auditlogRepo.InfoToLog(mLog, "Create Security Profile", fmt.Sprintf("Security Profile %s created successfully", mSecurityProfile.Name))
		}

		err = u.repo.Save(mSecurityProfile)
		if err != nil {
			fmt.Printf("saveSecurityProfile %v\n", err)
		}
	}()

	go u.auditlogRepo.InfoToLog(mLog, "Create Security Profile", fmt.Sprintf("Security Profile %s is being processed", mSecurityProfile.Name))

	return nil
}

func (u *usecase) createF5SecurityProfile(c *gin.Context, req request.CreateSecurityProfile, mSecurityProfile model.SecurityProfile) (mNewSecurityProfile model.SecurityProfile, err error) {
	mNewSecurityProfile = mSecurityProfile
	resNode, err := helper.LogF5Request("CreateNode", func() (response.F5CreateNode, error) {
		return u.f5Repo.CreateNode(c, request.F5CreateNode{
			Name:    req.Node,
			Address: helper.GetIP(c),
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5NodeFullPath = resNode.FullPath

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
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5PoolFullPath = resPool.FullPath

	crtName := fmt.Sprintf("%s.crt", helper.Slugify(req.Name))
	resUploadCert, err := helper.LogF5Request("UploadCert", func() (response.F5Upload, error) {
		return u.f5Repo.Upload(c, crtName, req.CertificateFile)
	})
	if err != nil {
		return mNewSecurityProfile, err
	}

	_, err = helper.LogF5Request("InstallCert", func() (response.F5Install, error) {
		return u.f5Repo.InstallCertOrChain(c, request.F5Install{
			Command:       "install",
			Name:          crtName,
			FromLocalFile: resUploadCert.LocalFilePath,
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5CertificatePath = f5CommonPath(crtName)

	chainName := fmt.Sprintf("%s-chain.crt", helper.Slugify(req.Name))
	if req.ChainFile != nil {
		resUploadChain, err := helper.LogF5Request("UploadChain", func() (response.F5Upload, error) {
			return u.f5Repo.Upload(c, chainName, req.ChainFile)
		})
		if err != nil {
			return mNewSecurityProfile, err
		}
		_, err = helper.LogF5Request("InstallChain", func() (response.F5Install, error) {
			return u.f5Repo.InstallCertOrChain(c, request.F5Install{
				Command:       "install",
				Name:          chainName,
				FromLocalFile: resUploadChain.LocalFilePath,
			})
		})
		if err != nil {
			return mNewSecurityProfile, err
		}
		mNewSecurityProfile.F5ChainPath = f5CommonPath(chainName)
	}

	keyName := fmt.Sprintf("%s.key", helper.Slugify(req.Name))
	resUploadKey, err := helper.LogF5Request("UploadKey", func() (response.F5Upload, error) {
		return u.f5Repo.Upload(c, keyName, req.KeyFile)
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	_, err = helper.LogF5Request("InstallKey", func() (response.F5Install, error) {
		return u.f5Repo.InstallKey(c, request.F5Install{
			Command:       "install",
			Name:          keyName,
			FromLocalFile: resUploadKey.LocalFilePath,
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5KeyPath = f5CommonPath(keyName)

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
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5ClientSSLFullPath = resCreateSslSecurityProfile.FullPath

	_, err = helper.LogF5Request("ActiveSNI", func() (response.F5ActiveSNI, error) {
		return u.f5Repo.ActiveSNI(c, resCreateSslSecurityProfile.FullPath, request.F5ActiveSNI{
			ServerName: req.Domain,
			SniDefault: "true",
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}

	resCreateVS, err := helper.LogF5Request("CreateVS", func() (response.F5VS, error) {
		return u.f5Repo.CreateVS(c, request.F5CreateVS{
			Name:        fmt.Sprintf("%s-vs", helper.Slugify(req.Name)),
			Partition:   "Common",
			Destination: fmt.Sprintf("%s:%s", req.Destination, req.Port),
			Mask:        "255.255.255.255",
			IPProtocol:  "tcp",
			Pool:        resPool.FullPath,
			Profiles: []request.F5VSProfile{
				{
					Name:    "tcp",
					Context: "all",
				},
				{
					Name:    resCreateSslSecurityProfile.FullPath,
					Context: "clientside",
				},
			},
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5VirtualServerFullPath = resCreateVS.FullPath

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
		return mNewSecurityProfile, fmt.Errorf("unsupported enforcement mode: %s", req.EnforcementMode)
	}

	resCreateSecurityPolicy, err := helper.LogF5Request("CreateSecurityPolicy", func() (response.F5CreateSecurityPolicy, error) {
		return u.f5Repo.CreateSecurityPolicy(c, reqF5)
	})
	if err != nil {
		return mNewSecurityProfile, err
	}
	mNewSecurityProfile.F5PolicyFullPath = resCreateSecurityPolicy.FullPath
	mNewSecurityProfile.F5PolicyID = resCreateSecurityPolicy.ID

	_, err = helper.LogF5Request("AssignSecurityPolicyToVS", func() (response.F5AssignSecurityPolicyToVS, error) {
		return u.f5Repo.AssignSecurityPolicyToVS(c, resCreateSecurityPolicy.ID, request.F5AssignSecurityPolicyToVS{
			VirtualServers: []string{
				resCreateVS.FullPath,
			},
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
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
			return mNewSecurityProfile, err
		}
	}

	_, err = helper.LogF5Request("SetLearningModePolicy", func() (response.F5SetLearningModePolicy, error) {
		return u.f5Repo.SetLearningModePolicy(c, resCreateSecurityPolicy.ID, request.F5SetLearningModePolicy{
			LearningMode: "automatic",
		})
	})
	if err != nil {
		return mNewSecurityProfile, err
	}

	if req.BotProtection {
		reqAttachBotDefense := request.F5AttachBotDefenseandLogProfileToVS{
			SecurityLogProfiles: []string{
				"/Common/local-bot-defense",
			},
		}

		_, err = helper.LogF5Request("AttachBotDefenseandLogProfileToVS", func() (response.F5AttachBotDefenseandLogProfileToVS, error) {
			return u.f5Repo.AttachBotDefenseandLogProfileToVS(c, resCreateVS.FullPath, reqAttachBotDefense)
		})
		if err != nil {
			return mNewSecurityProfile, err
		}

		_, err = helper.LogF5Request("AttachBotDefenseProfileToVS", func() (response.F5VSProfile, error) {
			return u.f5Repo.AttachBotDefenseToVSProfile(c, resCreateVS.FullPath, request.F5VSProfile{
				Name:    "/Common/bot-defense",
				Context: "all",
			})
		})
		if err != nil {
			return mNewSecurityProfile, err
		}
	}

	return mNewSecurityProfile, nil
}

func (u *usecase) DeleteSecurityProfile(c *gin.Context, id string) (err error) {
	userID := middleware.GetUserID(c)

	mSecurityProfile, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	// Authorize the user.
	authorized := u.permissionUsecase.CheckUserPermission(c, mSecurityProfile.ProjectID, constant.PrivilegeDekaWaf, true)
	if !authorized {
		return fmt.Errorf(constant.ErrForbidden)
	}

	// Check project availability status.
	availability, err := u.projectRepo.ProjectAvailabilityStatus(mSecurityProfile.ProjectID)
	if err != nil {
		return fmt.Errorf(constant.ErrCheckProjectAvailability)
	}

	if !availability {
		return fmt.Errorf(constant.ErrProjectNotAvailable)
	}

	mProject, err := u.projectRepo.GetByID(mSecurityProfile.ProjectID)
	if err != nil {
		return err
	}

	mLog := model.Auditlog{
		UserID:         userID,
		OrganizationID: mProject.OrganizationID,
		ProjectID:      mProject.ID,
		IPAddress:      helper.GetIP(c),
		ServiceType:    "Deka Waf",
		ServiceName:    "Security Profile",
		MethodName:     "Delete Security Profile",
	}

	mSecurityProfile.Status = model.SecurityProfileStatusPending
	err = u.repo.Save(mSecurityProfile)
	if err != nil {
		go u.auditlogRepo.ErrorToLog(mLog, "Delete Security Profile", fmt.Sprintf("Failed to update Security Profile %s", mSecurityProfile.Name))
		return err
	}

	go func() {
		err := u.deleteF5SecurityProfile(c, &mSecurityProfile)
		if err != nil {
			mSecurityProfile.Status = model.SecurityProfileStatusInactive
			if saveErr := u.repo.Save(mSecurityProfile); saveErr != nil {
				fmt.Printf("saveSecurityProfile %v\n", saveErr)
			}
			websocket.SendWebsocketMessage(mProject.ID, map[string]string{
				"security_profile_id":   mSecurityProfile.ID,
				"security_profile_name": mSecurityProfile.Name,
				"action":                "delete_security_profile",
				"status":                "failed",
				"error":                 err.Error(),
			})
			go u.auditlogRepo.ErrorToLog(mLog, "Delete Security Profile", fmt.Sprintf("Failed to delete Security Profile resources %s: %s", mSecurityProfile.Name, err.Error()))
		} else {
			err = u.repo.DeleteByID(mSecurityProfile.ID)
			if err != nil {
				mSecurityProfile.Status = model.SecurityProfileStatusInactive
				if saveErr := u.repo.Save(mSecurityProfile); saveErr != nil {
					fmt.Printf("saveSecurityProfile %v\n", saveErr)
				}
				websocket.SendWebsocketMessage(mProject.ID, map[string]string{
					"security_profile_id":   mSecurityProfile.ID,
					"security_profile_name": mSecurityProfile.Name,
					"action":                "delete_security_profile",
					"status":                "failed",
					"error":                 fmt.Sprintf("Failed to delete from database: %s", err.Error()),
				})
				go u.auditlogRepo.ErrorToLog(mLog, "Delete Security Profile", fmt.Sprintf("Failed to delete Security Profile from database %s: %s", mSecurityProfile.Name, err.Error()))
			} else {
				websocket.SendWebsocketMessage(mProject.ID, map[string]string{
					"security_profile_id":   mSecurityProfile.ID,
					"security_profile_name": mSecurityProfile.Name,
					"action":                "delete_security_profile",
					"status":                "success",
				})
				go u.auditlogRepo.InfoToLog(mLog, "Delete Security Profile", fmt.Sprintf("Security Profile %s deleted successfully", mSecurityProfile.Name))
			}
		}
	}()

	go u.auditlogRepo.InfoToLog(mLog, "Delete Security Profile", fmt.Sprintf("Security Profile %s is being deleted", mSecurityProfile.Name))

	return nil
}

func f5CommonPath(name string) string {
	return fmt.Sprintf("/Common/%s", name)
}

// deleteF5Resource skips empty paths and treats an already-missing resource as deleted.
func deleteF5Resource(name, path string, del func() error) error {
	if path == "" {
		return nil
	}
	_, err := helper.LogF5Request(name, func() (interface{}, error) {
		return nil, del()
	})
	if err != nil && !errors.Is(err, f5.ErrNotFound) {
		return err
	}
	return nil
}

// deleteF5SecurityProfile removes F5 resources in the reverse order of creation.
// The Virtual Server goes first because it holds the policy assignment.
// After each resource is deleted, its F5 path is cleared in the database so a
// failed delete can be retried from where it stopped.
func (u *usecase) deleteF5SecurityProfile(c *gin.Context, m *model.SecurityProfile) (err error) {
	steps := []struct {
		name  string
		field *string
		del   func() error
	}{
		{"DeleteVS", &m.F5VirtualServerFullPath, func() error { return u.f5Repo.DeleteVS(c, m.F5VirtualServerFullPath) }},
		{"DeleteSecurityPolicy", &m.F5PolicyID, func() error { return u.f5Repo.DeleteSecurityPolicy(c, m.F5PolicyID) }},
		{"DeleteSslSecurityProfile", &m.F5ClientSSLFullPath, func() error { return u.f5Repo.DeleteSslSecurityProfile(c, m.F5ClientSSLFullPath) }},
		{"DeleteKey", &m.F5KeyPath, func() error { return u.f5Repo.DeleteKey(c, m.F5KeyPath) }},
		{"DeleteChain", &m.F5ChainPath, func() error { return u.f5Repo.DeleteCertOrChain(c, m.F5ChainPath) }},
		{"DeleteCert", &m.F5CertificatePath, func() error { return u.f5Repo.DeleteCertOrChain(c, m.F5CertificatePath) }},
		{"DeletePool", &m.F5PoolFullPath, func() error { return u.f5Repo.DeletePool(c, m.F5PoolFullPath) }},
		{"DeleteNode", &m.F5NodeFullPath, func() error { return u.f5Repo.DeleteNode(c, m.F5NodeFullPath) }},
	}
	for _, step := range steps {
		if *step.field == "" {
			continue
		}
		if err := deleteF5Resource(step.name, *step.field, step.del); err != nil {
			return err
		}
		*step.field = ""
		if step.name == "DeleteSecurityPolicy" {
			m.F5PolicyFullPath = ""
		}
		if err := u.repo.Save(*m); err != nil {
			return err
		}
	}
	return nil
}
