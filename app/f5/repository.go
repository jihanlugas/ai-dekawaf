package f5

import (
	"context"
	"dekawaf/httpclient"
	"dekawaf/request"
	"dekawaf/response"
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
	Upload(ctx context.Context, fileReq *multipart.FileHeader) (data response.F5Upload, err error)
	InstallCertOrChain(ctx context.Context, req request.F5Install) (data response.F5Install, err error)
	InstallKey(ctx context.Context, req request.F5Install) (data response.F5Install, err error)
	ListCertAndChain(ctx context.Context) (data []response.F5CertAndChain, err error)
	GetCertAndChain(ctx context.Context, name string) (data response.F5CertAndChain, err error)
	ListKey(ctx context.Context) (data []response.F5Key, err error)
	GetKey(ctx context.Context, name string) (data response.F5Key, err error)
	CreateSslSecurityProfile(ctx context.Context, req request.F5CreateSslSecurityProfile) (data response.F5CreateSslSecurityProfile, err error)
	ActiveSNI(ctx context.Context, name string, req request.F5ActiveSNI) (data response.F5ActiveSNI, err error)
	ListSNI(ctx context.Context) (data []response.F5SNI, err error)
	GetSNI(ctx context.Context, name string) (data response.F5SNI, err error)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		if err := json.Unmarshal(res.Body, &f5Error); err != nil {
			return fmt.Errorf("login failed with status code: %d", res.StatusCode)
		}
		return fmt.Errorf("login failed: %s", f5Error.Message)
	}

	if err := json.Unmarshal(res.Body, &data); err != nil {
		return fmt.Errorf("failed to unmarshal login response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("create node failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("create node failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal create node response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("create pool failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("create pool failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal create pool response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		if err := json.Unmarshal(res.Body, &f5Error); err != nil {
			return data, fmt.Errorf("upload failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("upload failed: %s", f5Error.Message)
	}

	if err := json.Unmarshal(res.Body, &data); err != nil {
		return data, fmt.Errorf("failed to unmarshal upload response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("install cert or chain failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("install cert or chain failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal install cert or chain response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("install key failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("install key failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal install key response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) ListCertAndChain(ctx context.Context) (data []response.F5CertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("list cert and chain failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("list cert and chain failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal list cert and chain response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) GetCertAndChain(ctx context.Context, name string) (data response.F5CertAndChain, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/cert/%s", os.Getenv("F5_URL"), name), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("get cert and chain failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("get cert and chain failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal get cert and chain response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) ListKey(ctx context.Context) (data []response.F5Key, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("list key failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("list key failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal list key response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) GetKey(ctx context.Context, name string) (data response.F5Key, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/sys/crypto/key/%s", os.Getenv("F5_URL"), name), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("get key failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("get key failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal get key response: %w", err)
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

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("create ssl security profile failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("create ssl security profile failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal create ssl security profile response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) ActiveSNI(ctx context.Context, name string, req request.F5ActiveSNI) (data response.F5ActiveSNI, err error) {
	res, err := httpclient.Patch(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), name), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, req)
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("active sni failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("active sni failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal active sni response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) ListSNI(ctx context.Context) (data []response.F5SNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("list sni failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("list sni failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal list sni response: %w", err)
	}

	return data, nil
}

func (f *f5Repository) GetSNI(ctx context.Context, name string) (data response.F5SNI, err error) {
	res, err := httpclient.Get(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/profile/client-ssl/%s", os.Getenv("F5_URL"), name), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	})
	if err != nil {
		return data, err
	}

	if res.StatusCode >= 400 {
		var f5Error response.F5Error
		err = json.Unmarshal(res.Body, &f5Error)
		if err != nil {
			return data, fmt.Errorf("get sni failed with status code: %d", res.StatusCode)
		}
		return data, fmt.Errorf("get sni failed: %s", f5Error.Message)
	}

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal get sni response: %w", err)
	}

	return data, nil
}
