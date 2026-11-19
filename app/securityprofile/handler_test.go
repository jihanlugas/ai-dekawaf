package securityprofile

import (
	"ai-dekawaf/request"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type captureUsecase struct {
	called bool
	req    request.CreateSecurityProfile
}

func (u *captureUsecase) CreateSecurityProfile(_ *gin.Context, req request.CreateSecurityProfile) error {
	u.called = true
	u.req = req
	return nil
}

func TestCreateMultipart(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fields := map[string][]string{
		"name": {"profile"}, "description": {"description"}, "domain": {"example.com"},
		"node": {"node"}, "pool": {"pool"}, "port": {"443"},
		"serverTechnologies": {"Flask", "Python"}, "enforcementMode": {"blocking"},
		"botProtection": {"true"},
	}
	for name, values := range fields {
		for _, value := range values {
			if err := writer.WriteField(name, value); err != nil {
				t.Fatal(err)
			}
		}
	}
	files := map[string]string{"certificateFile": "certificate", "keyFile": "key", "chainFile": "chain"}
	for name, contents := range files {
		part, err := writer.CreateFormFile(name, name+".pem")
		if err != nil {
			t.Fatal(err)
		}
		if _, err := io.WriteString(part, contents); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}

	u := &captureUsecase{}
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodPost, "/security-profile", &body)
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	(&Handler{Usecase: u}).Create(c)

	if recorder.Code != http.StatusOK || !u.called {
		t.Fatalf("status = %d, usecase called = %v, response = %s", recorder.Code, u.called, recorder.Body.String())
	}
	got := u.req
	if got.Name != "profile" || got.Description != "description" || got.Domain != "example.com" ||
		got.Node != "node" || got.Pool != "pool" || got.Port != "443" ||
		!reflect.DeepEqual(got.ServerTecnologies, []string{"Flask", "Python"}) ||
		got.EnforcementMode != "blocking" || !got.BotProtection {
		t.Fatalf("unexpected request: %+v", got)
	}
	for name, header := range map[string]*multipart.FileHeader{
		"certificateFile": got.CertificateFile, "keyFile": got.KeyFile, "chainFile": got.ChainFile,
	} {
		if header == nil || header.Filename != name+".pem" {
			t.Fatalf("%s: unexpected file header: %+v", name, header)
		}
		file, err := header.Open()
		if err != nil {
			t.Fatal(err)
		}
		content, err := io.ReadAll(file)
		file.Close()
		if err != nil || string(content) != files[name] {
			t.Fatalf("%s: content = %q, error = %v", name, content, err)
		}
	}
}

func TestCreateBinding(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		body        string
		status      int
		called      bool
	}{
		{"json remains supported", "application/json", `{"name":"profile"}`, http.StatusOK, true},
		{"invalid form value", "application/x-www-form-urlencoded", "botProtection=invalid", http.StatusBadRequest, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &captureUsecase{}
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)
			c.Request = httptest.NewRequest(http.MethodPost, "/security-profile", strings.NewReader(tt.body))
			c.Request.Header.Set("Content-Type", tt.contentType)
			(&Handler{Usecase: u}).Create(c)
			if recorder.Code != tt.status || u.called != tt.called {
				t.Fatalf("status = %d, usecase called = %v, response = %s", recorder.Code, u.called, recorder.Body.String())
			}
			if tt.called && u.req.Name != "profile" {
				t.Fatalf("name = %q", u.req.Name)
			}
		})
	}
}
