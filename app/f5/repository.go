package f5

import (
	"ai-dekawaf/helper"
	"ai-dekawaf/httpclient"
	"ai-dekawaf/request"
	"ai-dekawaf/response"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
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
	Upload(ctx context.Context, name string, fileReq *multipart.FileHeader) (data response.F5Upload, err error)
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
	ListVirtualAddress(ctx context.Context) (data response.F5ListVirtualAddress, err error)
	GetVirtualAddress(ctx context.Context, fullpath string) (data response.F5VirtualAddress, err error)
	AttachBotDefenseToVSProfile(ctx context.Context, fullpath string, req request.F5VSProfile) (data response.F5VSProfile, err error)
	DeleteVS(ctx context.Context, fullpath string) (err error)
	DeleteSecurityPolicy(ctx context.Context, policyID string) (err error)
	DeleteSslSecurityProfile(ctx context.Context, fullpath string) (err error)
	DeletePool(ctx context.Context, fullpath string) (err error)
	DeleteNode(ctx context.Context, fullpath string) (err error)
	DeleteCertOrChain(ctx context.Context, fullpath string) (err error)
	DeleteKey(ctx context.Context, fullpath string) (err error)
}

type repository struct {
	loginMu      sync.Mutex
	refreshMu    sync.Mutex
	refreshTimer *time.Timer
}

func NewRepository() Repository {
	return &repository{}
}

var Token string

var tokenMu sync.RWMutex

const (
	tokenRefreshMaxLead    = time.Minute
	tokenRefreshRetryDelay = 5 * time.Second
)

func currentToken() string {
	tokenMu.RLock()
	defer tokenMu.RUnlock()
	return Token
}

func updateToken(token string) {
	tokenMu.Lock()
	defer tokenMu.Unlock()
	Token = token
}

func (r *repository) Login(ctx context.Context) (err error) {
	r.loginMu.Lock()
	defer r.loginMu.Unlock()

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

	if data.Token.Token == "" {
		return fmt.Errorf("login response did not include token")
	}

	expiresAt := time.Time{}
	if data.Token.ExpirationMicros > 0 {
		expiresAt = time.UnixMicro(data.Token.ExpirationMicros)
	} else if data.Token.Timeout > 0 {
		expiresAt = time.Now().Add(time.Duration(data.Token.Timeout) * time.Second)
	}
	if expiresAt.IsZero() {
		return fmt.Errorf("login response did not include token expiration")
	}

	updateToken(data.Token.Token)
	r.scheduleTokenRefresh(expiresAt)
	return nil
}

func (r *repository) scheduleTokenRefresh(expiresAt time.Time) {
	delay := time.Until(expiresAt)
	if delay <= 0 {
		delay = time.Second
	} else {
		lead := delay / 10
		if lead > tokenRefreshMaxLead {
			lead = tokenRefreshMaxLead
		}
		delay -= lead
	}
	r.scheduleLogin(delay)
}

func (r *repository) scheduleLogin(delay time.Duration) {
	r.refreshMu.Lock()
	defer r.refreshMu.Unlock()

	if r.refreshTimer != nil {
		r.refreshTimer.Stop()
	}
	r.refreshTimer = time.AfterFunc(delay, func() {
		if err := r.Login(context.Background()); err != nil {
			log.Printf("F5 token refresh failed; retrying in %s: %v", tokenRefreshRetryDelay, err)
			r.scheduleLogin(tokenRefreshRetryDelay)
		}
	})
}

func (r *repository) stopTokenRefresh() {
	r.refreshMu.Lock()
	defer r.refreshMu.Unlock()

	if r.refreshTimer != nil {
		r.refreshTimer.Stop()
		r.refreshTimer = nil
	}
}

func (r *repository) CreateNode(ctx context.Context, req request.F5CreateNode) (data response.F5CreateNode, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/node", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create node", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) CreatePool(ctx context.Context, req request.F5CreatePool) (data response.F5CreatePool, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create pool", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetMemberPool(ctx context.Context, poolFullPath, memberFullPath string) (data response.F5MemberPool, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool/%s/members/%s", os.Getenv("F5_URL"), helper.FormatF5Path(poolFullPath), helper.FormatF5Path(memberFullPath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get member pool", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListMemberPool(ctx context.Context, poolFullPath string) (data response.F5ListMemberPool, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool/%s/members", os.Getenv("F5_URL"), helper.FormatF5Path(poolFullPath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list member pool", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) CreateSecurityPolicy(ctx context.Context, req request.F5CreateSecurityPolicy) (data response.F5CreateSecurityPolicy, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "create security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListSecurityPolicy(ctx context.Context) (data response.F5ListSecurityPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListServerTecnology(ctx context.Context) (data response.F5ListServerTecnology, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/server-technologies", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list server tecnology", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListServerTecnologyInPolicy(ctx context.Context, policyID string) (data response.F5ListServerTecnologyInPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/server-technologies", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list server tecnology in policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) AssignSecurityPolicyToVS(ctx context.Context, policyID string, req request.F5AssignSecurityPolicyToVS) (data response.F5AssignSecurityPolicyToVS, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "Assign security policy to vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) AssignServerTecnologyToSecurityPolicy(ctx context.Context, policyID string, req request.F5AssignServerTecnologyToSecurityPolicy) (data response.F5AssignServerTecnologyToSecurityPolicy, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/server-technologies", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "Assign server tecnology to security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) SetLearningModePolicy(ctx context.Context, policyID string, req request.F5SetLearningModePolicy) (data response.F5SetLearningModePolicy, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/policy-builder", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "set learning mode policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetLearningModePolicy(ctx context.Context, policyID string) (data response.F5GetLearningModePolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s/policy-builder", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get learning mode policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetSecurityPolicy(ctx context.Context, policyID string) (data response.F5GetSecurityPolicy, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get security policy", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListBotDefense(ctx context.Context) (data response.F5ListBotDefense, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/bot-defense/profile", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get bot defense", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetBotDefense(ctx context.Context, fullpath string) (data response.F5BotDefense, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/bot-defense/profile/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get bot defense", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListLogProfile(ctx context.Context) (data response.F5ListLogProfile, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/log/profile", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list log profile", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetLogProfile(ctx context.Context, fullpath string) (data response.F5LogProfile, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/security/log/profile/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get log profile", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) AttachBotDefenseandLogProfileToVS(ctx context.Context, fullpath string, req request.F5AttachBotDefenseandLogProfileToVS) (data response.F5AttachBotDefenseandLogProfileToVS, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "attach bot defense and log profile to vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) Upload(ctx context.Context, name string, fileReq *multipart.FileHeader) (data response.F5Upload, err error) {
	if fileReq == nil {
		return data, fmt.Errorf("upload failed: file is nil")
	}

	file, err := fileReq.Open()
	if err != nil {
		return data, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	// Read all file content to get actual size (multipart.FileHeader.Size may be inaccurate)
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return data, fmt.Errorf("read file: %w", err)
	}
	fileSize := int64(len(fileContent))

	contentRange := fmt.Sprintf("0-%d/%d", fileSize-1, fileSize)
	fmt.Printf("Content-Range: %s\n", contentRange)

	res, err := httpclient.PostReader(ctx, fmt.Sprintf("%s/mgmt/shared/file-transfer/uploads/%s", os.Getenv("F5_URL"), name),
		map[string]string{
			"Content-Type":    "application/octet-stream",
			"Content-Range":   contentRange,
			"X-F5-Auth-Token": currentToken(),
		},
		bytes.NewReader(fileContent),
		fileSize,
	)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "upload", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) InstallCertOrChain(ctx context.Context, req request.F5Install) (data response.F5Install, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "install cert or chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) InstallKey(ctx context.Context, req request.F5Install) (data response.F5Install, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "install key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ListCertAndChain(ctx context.Context) (data response.F5ListCertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list cert and chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetCertAndChain(ctx context.Context, fullpath string) (data response.F5CertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get cert and chain", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ListKey(ctx context.Context) (data response.F5ListKey, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetKey(ctx context.Context, fullpath string) (data response.F5Key, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get key", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) CreateSslSecurityProfile(ctx context.Context, req request.F5CreateSslSecurityProfile) (data response.F5CreateSslSecurityProfile, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "create ssl security profile", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ActiveSNI(ctx context.Context, fullpath string, req request.F5ActiveSNI) (data response.F5ActiveSNI, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "active sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) ListSNI(ctx context.Context) (data response.F5ListSNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "list sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) GetSNI(ctx context.Context, fullpath string) (data response.F5SNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}

	if err := decodeF5Response(res, "get sni", &data); err != nil {
		return data, err
	}

	return data, nil
}

func (r *repository) CreateVS(ctx context.Context, req request.F5CreateVS) (data response.F5VS, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "create vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) ListVS(ctx context.Context) (data response.F5ListVS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetVS(ctx context.Context, fullpath string) (data response.F5VS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get vs", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) VerifySSLInVS(ctx context.Context, fullpath string) (data response.F5VerifySSLInVS, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s/profiles", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
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
		if res.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("F5 Unauthorized")
		}
		var f5Error response.F5Error
		if err := json.Unmarshal(res.Body, &f5Error); err != nil || f5Error.Message == "" {
			return fmt.Errorf("%s failed with status code: %d", operation, res.StatusCode)
		}
		return fmt.Errorf("%s failed: %s", operation, f5Error.Message)
	}
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(res.Body, data); err != nil {
		return fmt.Errorf("failed to unmarshal %s response: %w", operation, err)
	}
	return nil
}

func (r *repository) ListVirtualAddress(ctx context.Context) (data response.F5ListVirtualAddress, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual-address", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "list virtual address", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) GetVirtualAddress(ctx context.Context, fullpath string) (data response.F5VirtualAddress, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual-address/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "get virtual address", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) AttachBotDefenseToVSProfile(ctx context.Context, fullpath string, req request.F5VSProfile) (data response.F5VSProfile, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s/profiles", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	}, req)
	if err != nil {
		return data, err
	}
	if err := decodeF5Response(res, "attach bot defense to vs profile", &data); err != nil {
		return data, err
	}
	return data, nil
}

func (r *repository) DeleteVS(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/virtual/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete virtual server", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteSecurityPolicy(ctx context.Context, policyID string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/asm/policies/%s", os.Getenv("F5_URL"), url.PathEscape(policyID)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete security policy", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteSslSecurityProfile(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/client-ssl/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete ssl security profile", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeletePool(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete pool", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteNode(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/node/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete node", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteCertOrChain(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/sys/certificate/ssl-cert/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete certificate or chain", nil); err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteKey(ctx context.Context, fullpath string) (err error) {
	res, err := httpclient.Delete(ctx, fmt.Sprintf("%s/mgmt/tm/sys/certificate/ssl-key/%s", os.Getenv("F5_URL"), helper.FormatF5Path(fullpath)), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": currentToken(),
	})
	if err != nil {
		return err
	}
	if err := decodeF5Response(res, "delete key", nil); err != nil {
		return err
	}
	return nil
}
