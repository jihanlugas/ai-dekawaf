package f5

import (
	"context"
	"ai-dekawaf/httpclient"
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/url"
	"os"
)

type Repository interface {
	Login(ctx context.Context) (err error)
	CreateNode(ctx context.Context, req request.F5CreateNode) (data response.F5CreateNode, err error)
	CreatePool(ctx context.Context, req request.F5CreatePool) (data response.F5CreatePool, err error)
	GetMemberPool(ctx context.Context, poolFullPath, memberFullPath string) (data response.F5MemberPool, err error)
	ListMemberPool(ctx context.Context, poolFullPath string) (data response.F5ListMemberPool, err error)
	CreateSecurityPolicy(ctx context.Context, req request.F5CreateSecurityPolicy) (data response.F5CreateSecurityPolicy, err error)
	ListSecurityPolicy(ctx context.Context) (data response.F5ListSecurityPolicy, err error)
	ListServerTecnology(ctx context.Context) (data response.F5ListServerTecnology, err error)
	ListServerTecnologyInPolicy(ctx context.Context, policyID string) (data response.F5ListServerTecnologyInPolicy, err error)
	AssignSecurityPolicyToVS(ctx context.Context, policyID string, req request.F5AssignSecurityPolicyToVS) (data response.F5AssignSecurityPolicyToVS, err error)
	AssignServerTecnologyToSecurityPolicy(ctx context.Context, policyID string, req request.F5AssignServerTecnologyToSecurityPolicy) (data response.F5AssignServerTecnologyToSecurityPolicy, err error)
	SetLearningModePolicy(ctx context.Context, policyID string, req request.F5SetLearningModePolicy) (data response.F5SetLearningModePolicy, err error)
	GetLearningModePolicy(ctx context.Context, policyID string) (data response.F5GetLearningModePolicy, err error)
	GetSecurityPolicy(ctx context.Context, policyID string) (data response.F5GetSecurityPolicy, err error)
	GetBotDefense(ctx context.Context, fullpath string) (data response.F5BotDefense, err error)
	ListBotDefense(ctx context.Context) (data response.F5ListBotDefense, err error)
	GetLogProfile(ctx context.Context, fullpath string) (data response.F5LogProfile, err error)
	ListLogProfile(ctx context.Context) (data response.F5ListLogProfile, err error)
	AttachBotDefenseandLogProfileToVS(ctx context.Context, fullpath string, req request.F5AttachBotDefenseandLogProfileToVS) (data response.F5AttachBotDefenseandLogProfileToVS, err error)
	Upload(ctx context.Context, fileReq *multipart.FileHeader) (data response.F5Upload, err error)
	InstallCertOrChain(ctx context.Context, req request.F5Install) (data response.F5Install, err error)
	InstallKey(ctx context.Context, req request.F5Install) (data response.F5Install, err error)
	ListCertAndChain(ctx context.Context) (data response.F5ListCertAndChain, err error)
	GetCertAndChain(ctx context.Context, fullpath string) (data response.F5CertAndChain, err error)
	ListKey(ctx context.Context) (data response.F5ListKey, err error)
	GetKey(ctx context.Context, fullpath string) (data response.F5Key, err error)
	CreateSslSecurityProfile(ctx context.Context, req request.F5CreateSslSecurityProfile) (data response.F5CreateSslSecurityProfile, err error)
	ActiveSNI(ctx context.Context, fullpath string, req request.F5ActiveSNI) (data response.F5ActiveSNI, err error)
	ListSNI(ctx context.Context) (data response.F5ListSNI, err error)
	GetSNI(ctx context.Context, fullpath string) (data response.F5SNI, err error)
	CreateVS(ctx context.Context, req request.F5CreateVS) (data response.F5VS, err error)
	ListVS(ctx context.Context) (data response.F5ListVS, err error)
	GetVS(ctx context.Context, fullpath string) (data response.F5VS, err error)
	VerifySSLInVS(ctx context.Context, fullpath string) (data response.F5VerifySSLInVS, err error)
}

type f5Repository struct {
}

func NewF5Repository() Repository {
	return &f5Repository{}
}

var Token string

func (f *f5Repository) Login(ctx context.Context) (err error) {
	var data response.F5Login
	loginRequest := request.F5Login{
		Username:          os.Getenv("F5_USERNAME"),
		Password:          os.Getenv("F5_PASSWORD"),
		LoginProviderName: "tmos",
	}

	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/shared/authn/login", os.Getenv("F5_URL")), map[string]string{
		"Content-Type": "application/json",
	}, loginRequest)
	if err != nil {
		return err
	}

	if err := decodeF5Response(res, "login", &data); err != nil {
		return err
	}

	Token = data.Token.Token

	return nil
}

func (f *f5Repository) CreateNode(ctx context.Context, req request.F5CreateNode) (data response.F5CreateNode, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/node", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create node", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) CreatePool(ctx context.Context, req request.F5CreatePool) (data response.F5CreatePool, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create pool", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) GetMemberPool(ctx context.Context, poolFullPath, memberFullPath string) (data response.F5MemberPool, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool/%s/members/%s", os.Getenv("F5_URL"), url.PathEscape(poolFullPath), url.PathEscape(memberFullPath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get member pool", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListMemberPool(ctx context.Context, poolFullPath string) (data response.F5ListMemberPool, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool/%s/members", os.Getenv("F5_URL"), url.PathEscape(poolFullPath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list member pool", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) CreateSecurityPolicy(ctx context.Context, req request.F5CreateSecurityPolicy) (data response.F5CreateSecurityPolicy, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "create security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListSecurityPolicy(ctx context.Context) (data response.F5ListSecurityPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListServerTecnology(ctx context.Context) (data response.F5ListServerTecnology, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/server-technologies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list server tecnology", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListServerTecnologyInPolicy(ctx context.Context, policyID string) (data response.F5ListServerTecnologyInPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/server-technologies", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list server tecnology in policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) AssignSecurityPolicyToVS(ctx context.Context, policyID string, req request.F5AssignSecurityPolicyToVS) (data response.F5AssignSecurityPolicyToVS, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "Assign security policy to vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) AssignServerTecnologyToSecurityPolicy(ctx context.Context, policyID string, req request.F5AssignServerTecnologyToSecurityPolicy) (data response.F5AssignServerTecnologyToSecurityPolicy, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/server-technologies", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "Assign server tecnology to security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) SetLearningModePolicy(ctx context.Context, policyID string, req request.F5SetLearningModePolicy) (data response.F5SetLearningModePolicy, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/policy-builder", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "set learning mode policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) GetLearningModePolicy(ctx context.Context, policyID string) (data response.F5GetLearningModePolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/policy-builder", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get learning mode policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) GetSecurityPolicy(ctx context.Context, policyID string) (data response.F5GetSecurityPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListBotDefense(ctx context.Context) (data response.F5ListBotDefense, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/bot-defense/profile", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get bot defense", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) GetBotDefense(ctx context.Context, fullpath string) (data response.F5BotDefense, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/bot-defense/profile/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get bot defense", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListLogProfile(ctx context.Context) (data response.F5ListLogProfile, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/log/profile", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list log profile", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) GetLogProfile(ctx context.Context, fullpath string) (data response.F5LogProfile, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/log/profile/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get log profile", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) AttachBotDefenseandLogProfileToVS(ctx context.Context, fullpath string, req request.F5AttachBotDefenseandLogProfileToVS) (data response.F5AttachBotDefenseandLogProfileToVS, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "attach bot defense and log profile to vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) Upload(ctx context.Context, fileReq *multipart.FileHeader) (data response.F5Upload, err error) {
	if fileReq == nil {
		return data, fmt.Errorf("upload failed: file is nil")
	}

	file, err := fileReq.Open()
	if err != nil {
		return data, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	res, err := httpclient.PostReader(ctx, fmt.Sprintf("%s/mgmt/shared/file-transfer/uploads/%s", os.Getenv("F5_URL"), url.PathEscape(fileReq.Filename)),
		map[string]string{
			"Content-Type":    "application/octet-stream",
			"Content-Range":   fmt.Sprintf("0-%d/%d", fileReq.Size, fileReq.Size),
			"X-F5-Auth-Token": Token,
		},
		file,
		fileReq.Size,
	)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "upload", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) InstallCertOrChain(ctx context.Context, req request.F5Install) (data response.F5Install, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "install cert or chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) InstallKey(ctx context.Context, req request.F5Install) (data response.F5Install, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "install key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) ListCertAndChain(ctx context.Context) (data response.F5ListCertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list cert and chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) GetCertAndChain(ctx context.Context, fullpath string) (data response.F5CertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get cert and chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) ListKey(ctx context.Context) (data response.F5ListKey, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) GetKey(ctx context.Context, fullpath string) (data response.F5Key, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) CreateSslSecurityProfile(ctx context.Context, req request.F5CreateSslSecurityProfile) (data response.F5CreateSslSecurityProfile, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create ssl security profile", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) ActiveSNI(ctx context.Context, fullpath string, req request.F5ActiveSNI) (data response.F5ActiveSNI, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "active sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) ListSNI(ctx context.Context) (data response.F5ListSNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) GetSNI(ctx context.Context, fullpath string) (data response.F5SNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (f *f5Repository) CreateVS(ctx context.Context, req request.F5CreateVS) (data response.F5VS, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "create vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) ListVS(ctx context.Context) (data response.F5ListVS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) GetVS(ctx context.Context, fullpath string) (data response.F5VS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (f *f5Repository) VerifySSLInVS(ctx context.Context, fullpath string) (data response.F5VerifySSLInVS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s/profiles", os.Getenv("F5_URL"), url.PathEscape(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "verify ssl in vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func decodeF5Response(res *httpclient.Response, operation string, data interface{}) error {
	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		if err := json.Unmarshal(res.Body, &f5Error); err != nil || f5Error.Message == "" {
			return fmt.Errorf("%s failed with status code: %d", operation, res.StatusCode)
		}
		return fmt.Errorf("%s failed: %s", operation, f5Error.Message)
	}
	if err := json.Unmarshal(res.Body, data); err != nil {
		return fmt.Errorf("failed to unmarshal %s response: %w", operation, err)
	}
	return nil
}
