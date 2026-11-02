package f5

import (
	"context"
	"dekawaf/httpclient"
	"dekawaf/request"
	"dekawaf/response"
	"encoding/json"
	"fmt"
	"os"
)

type Repository interface {
	Login(ctx context.Context) (err error)
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

func (f *f5Repository) CreateNode(ctx context.Context, node request.F5CreateNode) (data response.F5CreateNode, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/node", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, node)
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

func (f *f5Repository) CreatePool(ctx context.Context, pool request.F5CreatePool) (data response.F5CreatePool, err error) {
	res, err := httpclient.Post(ctx, fmt.Sprintf("%s/mgmt/tm/ltm/pool", os.Getenv("F5_URL")), map[string]string{
		"Content-Type":    "application/json",
		"X-F5-Auth-Token": Token,
	}, pool)
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
