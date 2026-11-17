package f5

import (
	"bytes"
	"context"
	"dekawaf/request"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestVirtualServerOperations(t *testing.T) {
	const virtualPath = "/mgmt/tm/ltm/virtual"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-F5-Auth-Token") != "test-token" {
			t.Errorf("missing authentication header")
		}

		switch {
		case r.Method == http.MethodPost && r.URL.Path == virtualPath:
			var body request.F5CreateVS
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode create request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if body.Name != "test-vs" || body.Partition != "Common" || body.Destination != "/Common/192.0.2.1:443" ||
				body.Mask != "255.255.255.255" || body.IPProtocol != "tcp" || body.Pool != "/Common/pool" ||
				len(body.Profiles) != 1 || body.Profiles[0].Name != "client-ssl" || body.Profiles[0].Context != "clientside" {
				t.Errorf("unexpected create request: %+v", body)
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"name":"test-vs","enabled":true,"poolReference":{"link":"pool"},"profilesReference":{"link":"profiles","isSubcollection":true}}`))
		case r.Method == http.MethodGet && r.URL.Path == virtualPath:
			_, _ = w.Write([]byte(`{"kind":"collection","selfLink":"list","items":[{"name":"test-vs","rules":["rule"],"securityLogProfilesReference":[{"link":"log"}]}]}`))
		case r.Method == http.MethodGet && r.URL.Path == virtualPath+"/~Common~test-vs":
			_, _ = w.Write([]byte(`{"name":"test-vs","destination":"/Common/192.0.2.1:443"}`))
		case r.Method == http.MethodGet && r.URL.Path == virtualPath+"/~Common~test-vs/profiles":
			_, _ = w.Write([]byte(`{"kind":"profiles","selfLink":"profile-list","items":[{"name":"client-ssl","context":"clientside"},{"name":"tcp","context":"all"}]}`))
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.EscapedPath())
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	oldToken := Token
	Token = "test-token"
	t.Cleanup(func() { Token = oldToken })
	repo := NewF5Repository()
	ctx := context.Background()

	created, err := repo.CreateVS(ctx, request.F5CreateVS{
		Name: "test-vs", Partition: "Common", Destination: "/Common/192.0.2.1:443",
		Mask: "255.255.255.255", IPProtocol: "tcp", Pool: "/Common/pool",
		Profiles: []request.F5VSProfile{{Name: "client-ssl", Context: "clientside"}},
	})
	if err != nil || created.Name != "test-vs" || !created.Enabled || created.PoolReference == nil || created.PoolReference.Link != "pool" || !created.ProfilesReference.IsSubcollection {
		t.Fatalf("create: %+v, %v", created, err)
	}
	list, err := repo.ListVS(ctx)
	if err != nil || list.Kind != "collection" || list.SelfLink != "list" || len(list.Items) != 1 || len(list.Items[0].Rules) != 1 || len(list.Items[0].SecurityLogProfilesReference) != 1 {
		t.Fatalf("list: %+v, %v", list, err)
	}
	got, err := repo.GetVS(ctx, "~Common~test-vs")
	if err != nil || got.Destination != "/Common/192.0.2.1:443" {
		t.Fatalf("get: %+v, %v", got, err)
	}
	profiles, err := repo.VerifySSLInVS(ctx, "~Common~test-vs")
	if err != nil || profiles.Kind != "profiles" || profiles.SelfLink != "profile-list" || len(profiles.Items) != 2 || profiles.Items[0].Context != "clientside" {
		t.Fatalf("profiles: %+v, %v", profiles, err)
	}
}

func TestVirtualServerErrors(t *testing.T) {
	status := http.StatusNotFound
	body := `{"message":"virtual server not found"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	repo := NewF5Repository()
	ctx := context.Background()

	operations := []struct {
		name string
		call func() error
	}{
		{"create", func() error { _, err := repo.CreateVS(ctx, request.F5CreateVS{}); return err }},
		{"list", func() error { _, err := repo.ListVS(ctx); return err }},
		{"get", func() error { _, err := repo.GetVS(ctx, "~Common~missing"); return err }},
		{"profiles", func() error { _, err := repo.VerifySSLInVS(ctx, "~Common~missing"); return err }},
	}
	for _, operation := range operations {
		t.Run(operation.name, func(t *testing.T) {
			err := operation.call()
			if err == nil || !strings.Contains(err.Error(), "virtual server not found") {
				t.Fatalf("expected F5 error, got %v", err)
			}
		})
	}
	status, body = http.StatusBadGateway, "not json"
	if _, err := repo.ListVS(ctx); err == nil || !strings.Contains(err.Error(), "502") {
		t.Fatalf("expected status code error, got %v", err)
	}
	status, body = http.StatusOK, "not json"
	if _, err := repo.ListVS(ctx); err == nil || !strings.Contains(err.Error(), "unmarshal") {
		t.Fatalf("expected unmarshal error, got %v", err)
	}
}

func TestLoginResponseHandling(t *testing.T) {
	status := http.StatusOK
	body := `{"token":{"token":"new-token"}}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	oldToken := Token
	Token = "old-token"
	t.Cleanup(func() { Token = oldToken })

	repo := NewF5Repository()
	if err := repo.Login(context.Background()); err != nil || Token != "new-token" {
		t.Fatalf("login: token %q, error %v", Token, err)
	}
	status, body = http.StatusUnauthorized, `{"message":"invalid credentials"}`
	if err := repo.Login(context.Background()); err == nil || err.Error() != "login failed: invalid credentials" || Token != "new-token" {
		t.Fatalf("failed login: token %q, error %v", Token, err)
	}
	status, body = http.StatusOK, `invalid`
	if err := repo.Login(context.Background()); err == nil || !strings.Contains(err.Error(), "failed to unmarshal login response") || Token != "new-token" {
		t.Fatalf("invalid login response: token %q, error %v", Token, err)
	}
}

func TestUploadResponseHandling(t *testing.T) {
	var content bytes.Buffer
	writer := multipart.NewWriter(&content)
	part, err := writer.CreateFormFile("file", "test.crt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := part.Write([]byte("certificate")); err != nil {
		t.Fatal(err)
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	form, err := multipart.NewReader(&content, writer.Boundary()).ReadForm(1024)
	if err != nil {
		t.Fatal(err)
	}
	defer form.RemoveAll()

	status := http.StatusOK
	body := `{"temporaryFilePath":"/tmp/test.crt"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/mgmt/shared/file-transfer/uploads/test.crt" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	repo := NewF5Repository()
	file := form.File["file"][0]
	data, err := repo.Upload(context.Background(), file)
	if err != nil || data.TemporaryFilePath != "/tmp/test.crt" {
		t.Fatalf("upload: %+v, %v", data, err)
	}
	status, body = http.StatusBadRequest, `{}`
	if _, err := repo.Upload(context.Background(), file); err == nil || err.Error() != "upload failed with status code: 400" {
		t.Fatalf("upload error: %v", err)
	}
}

func TestCreateNodeResponseHandling(t *testing.T) {
	status := http.StatusOK
	body := `{"name":"test-node","address":"192.0.2.1"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	repo := NewF5Repository()
	data, err := repo.CreateNode(context.Background(), request.F5CreateNode{Name: "test-node"})
	if err != nil || data.Name != "test-node" || data.Address != "192.0.2.1" {
		t.Fatalf("create node: %+v, %v", data, err)
	}
	status, body = http.StatusConflict, `{"message":"already exists"}`
	if _, err := repo.CreateNode(context.Background(), request.F5CreateNode{}); err == nil || err.Error() != "create node failed: already exists" {
		t.Fatalf("create node error: %v", err)
	}
}

func TestPoolMemberAndSecurityPolicy(t *testing.T) {
	const memberJSON = `{"kind":"member","name":"node:443","partition":"Common","fullPath":"/Common/node:443","generation":123,"selfLink":"member-link","address":"192.0.2.1","connectionLimit":0,"dynamicRatio":1,"ephemeral":"false","fqdn":{"autopopulate":"disabled"},"inheritProfile":"enabled","logging":"disabled","monitor":"default","priorityGroup":0,"rateLimit":"disabled","ratio":1,"session":"user-enabled","state":"unchecked"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-F5-Auth-Token") != "test-token" || r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("missing F5 request headers")
		}
		switch {
		case r.Method == http.MethodGet && r.URL.EscapedPath() == "/mgmt/tm/ltm/pool/~Common~pool%2Fname/members/~Common~node:443":
			_, _ = w.Write([]byte(memberJSON))
		case r.Method == http.MethodGet && r.URL.EscapedPath() == "/mgmt/tm/ltm/pool/~Common~pool%2Fname/members":
			_, _ = w.Write([]byte(`{"kind":"members","selfLink":"list-link","items":[` + memberJSON + `]}`))
		case r.Method == http.MethodPost && r.URL.Path == "/mgmt/tm/asm/policies":
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("decode policy request: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			expected := map[string]any{
				"name": "policy", "applicationLanguage": "auto-detect", "enforcementMode": "blocking",
				"caseInsensitive": false, "protocolIndependent": false, "active": true, "policyBuilderEnabled": true,
			}
			if len(body) != len(expected) {
				t.Errorf("unexpected policy request fields: %v", body)
			}
			for key, value := range expected {
				if got, ok := body[key]; !ok || got != value {
					t.Errorf("policy request %s = %v, want %v", key, got, value)
				}
			}
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write([]byte(`{"id":"policy-id","name":"policy","active":true,"lastUpdateMicros":1794899852000000,"signatureRequirementReference":{"link":"signatures","isSubCollection":true},"templateReference":{"link":"template","title":"Fundamental"}}`))
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.EscapedPath())
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	oldToken := Token
	Token = "test-token"
	t.Cleanup(func() { Token = oldToken })
	repo := NewF5Repository()
	ctx := context.Background()

	member, err := repo.GetMemberPool(ctx, "~Common~pool/name", "~Common~node:443")
	if err != nil || member.Name != "node:443" || member.Address != "192.0.2.1" || member.Fqdn.Autopopulate != "disabled" || member.State != "unchecked" {
		t.Fatalf("get member: %+v, %v", member, err)
	}
	list, err := repo.ListMemberPool(ctx, "~Common~pool/name")
	if err != nil || list.Kind != "members" || list.SelfLink != "list-link" || len(list.Items) != 1 || list.Items[0] != member {
		t.Fatalf("list members: %+v, %v", list, err)
	}
	policy, err := repo.CreateSecurityPolicy(ctx, request.F5CreateSecurityPolicy{
		Name: "policy", ApplicationLanguage: "auto-detect", EnforcementMode: "blocking", Active: true, PolicyBuilderEnabled: true,
	})
	if err != nil || policy.ID != "policy-id" || !policy.Active || policy.LastUpdateMicros != 1794899852000000 ||
		!policy.SignatureRequirementReference.IsSubCollection || policy.SignatureRequirementReference.Link != "signatures" || policy.TemplateReference.Title != "Fundamental" {
		t.Fatalf("create policy: %+v, %v", policy, err)
	}
}

func TestPoolMemberAndSecurityPolicyErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"not found"}`))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	repo := NewF5Repository()
	ctx := context.Background()
	tests := []struct {
		want string
		call func() error
	}{
		{"get member pool failed: not found", func() error { _, err := repo.GetMemberPool(ctx, "pool", "member"); return err }},
		{"list member pool failed: not found", func() error { _, err := repo.ListMemberPool(ctx, "pool"); return err }},
		{"create security policy failed: not found", func() error { _, err := repo.CreateSecurityPolicy(ctx, request.F5CreateSecurityPolicy{}); return err }},
	}
	for _, tt := range tests {
		if err := tt.call(); err == nil || err.Error() != tt.want {
			t.Errorf("expected %q, got %v", tt.want, err)
		}
	}
}

func TestSecurityPolicyAndServerTechnologyLists(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.Header.Get("X-F5-Auth-Token") != "test-token" || r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected request method or headers: %s", r.Method)
		}
		switch r.URL.EscapedPath() {
		case "/mgmt/tm/asm/policies":
			_, _ = w.Write([]byte(`{"kind":"policies","selfLink":"policies-link","totalItems":2,"items":[{"name":"policy-1","id":"id-1","active":true,"lastUpdateMicros":1.794899924e+15,"virtualServers":["/Common/vs"],"manualVirtualServers":[],"subPath":"/Common","signatureRequirementReference":{"link":"signatures","isSubCollection":true}},{"name":"policy-2","lastUpdateMicros":1793801960000000,"virtualServers":[],"manualVirtualServers":["/Common/manual"]}]}`))
		case "/mgmt/tm/asm/server-technologies":
			_, _ = w.Write([]byte(`{"kind":"technologies","selfLink":"technology-link","totalItems":1,"items":[{"serverTechnologyDisplayName":"Flask","serverTechnologyName":"Flask","logoFileName":"flask.png","lastUpdateMicros":1784098296000000,"description":"framework","kind":"technology","selfLink":"flask-link","serverTechnologyReferences":[{"link":"python-link","serverTechnologyName":"Python"}],"id":"flask-id","isLearnable":true,"hasDeceptionResponses":false}]}`))
		case "/mgmt/tm/asm/policies/policy%2Fid/server-technologies":
			_, _ = w.Write([]byte(`{"kind":"policy-technologies","selfLink":"policy-link","totalItems":1,"items":[{"kind":"policy-technology","selfLink":"entry-link","id":"entry-id","lastUpdateMicros":1793096222000000,"serverTechnologyReference":{"link":"flask-link","serverTechnologyName":"Flask"}}]}`))
		default:
			t.Errorf("unexpected request path: %s", r.URL.EscapedPath())
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	oldToken := Token
	Token = "test-token"
	t.Cleanup(func() { Token = oldToken })
	repo := NewF5Repository()
	ctx := context.Background()

	policies, err := repo.ListSecurityPolicy(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if policies.Kind != "policies" || policies.SelfLink != "policies-link" || policies.TotalItems != 2 || len(policies.Items) != 2 {
		t.Fatalf("unexpected policy list: %+v", policies)
	}
	first := policies.Items[0]
	if first.Name != "policy-1" || first.ID != "id-1" || !first.Active || first.LastUpdateMicros.String() != "1.794899924e+15" ||
		first.SubPath != "/Common" || len(first.VirtualServers) != 1 || first.VirtualServers[0] != "/Common/vs" ||
		first.ManualVirtualServers == nil || !first.SignatureRequirementReference.IsSubCollection {
		t.Fatalf("unexpected policy: %+v", first)
	}
	if policies.Items[1].LastUpdateMicros.String() != "1793801960000000" || policies.Items[1].VirtualServers == nil || len(policies.Items[1].ManualVirtualServers) != 1 {
		t.Fatalf("unexpected second policy: %+v", policies.Items[1])
	}

	technologies, err := repo.ListServerTecnology(ctx)
	if err != nil || technologies.Kind != "technologies" || technologies.SelfLink != "technology-link" || technologies.TotalItems != 1 || len(technologies.Items) != 1 {
		t.Fatalf("technology list: %+v, %v", technologies, err)
	}
	technology := technologies.Items[0]
	if technology.ServerTechnologyDisplayName != "Flask" || technology.ID != "flask-id" || technology.LastUpdateMicros != 1784098296000000 ||
		!technology.IsLearnable || len(technology.ServerTechnologyReferences) != 1 || technology.ServerTechnologyReferences[0].ServerTechnologyName != "Python" {
		t.Fatalf("unexpected technology: %+v", technology)
	}
	policyTechnologies, err := repo.ListServerTecnologyInPolicy(ctx, "policy/id")
	if err != nil || policyTechnologies.Kind != "policy-technologies" || policyTechnologies.SelfLink != "policy-link" || policyTechnologies.TotalItems != 1 || len(policyTechnologies.Items) != 1 {
		t.Fatalf("policy technology list: %+v, %v", policyTechnologies, err)
	}
	entry := policyTechnologies.Items[0]
	if entry.ID != "entry-id" || entry.LastUpdateMicros != 1793096222000000 || entry.ServerTechnologyReference.Link != "flask-link" || entry.ServerTechnologyReference.ServerTechnologyName != "Flask" {
		t.Fatalf("unexpected policy technology: %+v", entry)
	}
}

func TestSecurityPolicyAndServerTechnologyListErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = w.Write([]byte(`{"message":"upstream unavailable"}`))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	repo := NewF5Repository()
	ctx := context.Background()
	tests := []struct {
		want string
		call func() error
	}{
		{"list security policy failed: upstream unavailable", func() error { _, err := repo.ListSecurityPolicy(ctx); return err }},
		{"list server tecnology failed: upstream unavailable", func() error { _, err := repo.ListServerTecnology(ctx); return err }},
		{"list server tecnology in policy failed: upstream unavailable", func() error { _, err := repo.ListServerTecnologyInPolicy(ctx, "id"); return err }},
	}
	for _, tt := range tests {
		if err := tt.call(); err == nil || err.Error() != tt.want {
			t.Errorf("expected %q, got %v", tt.want, err)
		}
	}
}

func TestPolicyAndDefenseOperations(t *testing.T) {
	var wantMethod, wantPath, wantBody string
	status := http.StatusOK
	body := `{}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != wantMethod || r.URL.EscapedPath() != wantPath {
			t.Errorf("request = %s %s, want %s %s", r.Method, r.URL.EscapedPath(), wantMethod, wantPath)
		}
		if r.Header.Get("Content-Type") != "application/json" || r.Header.Get("X-F5-Auth-Token") != "test-token" {
			t.Error("missing F5 request headers")
		}
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		} else if string(payload) != wantBody {
			t.Errorf("request body = %s, want %s", payload, wantBody)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	defer server.Close()
	t.Setenv("F5_URL", server.URL)
	oldToken := Token
	Token = "test-token"
	t.Cleanup(func() { Token = oldToken })
	repo := NewF5Repository()
	ctx := context.Background()
	const policyPath = "/mgmt/tm/asm/policies/policy%2Fid"
	tests := []struct {
		name, method, path, body, operation string
		call                                func() error
	}{
		{"assign policy", http.MethodPatch, policyPath, `{"virtualServers":["/Common/vs"]}`, "Assign security policy to vs",
			func() error {
				_, err := repo.AssignSecurityPolicyToVS(ctx, "policy/id", request.F5AssignSecurityPolicyToVS{VirtualServers: []string{"/Common/vs"}})
				return err
			}},
		{"assign technology", http.MethodPost, policyPath + "/server-technologies", `{"serverTechnologyReference":{"link":"technology-link"}}`, "Assign server tecnology to security policy",
			func() error {
				_, err := repo.AssignServerTecnologyToSecurityPolicy(ctx, "policy/id", request.F5AssignServerTecnologyToSecurityPolicy{ServerTechnologyReference: request.F5Link{Link: "technology-link"}})
				return err
			}},
		{"set learning mode", http.MethodPatch, policyPath + "/policy-builder", `{"learningMode":"automatic"}`, "set learning mode policy",
			func() error {
				_, err := repo.SetLearningModePolicy(ctx, "policy/id", request.F5SetLearningModePolicy{LearningMode: "automatic"})
				return err
			}},
		{"get learning mode", http.MethodGet, policyPath + "/policy-builder", "", "get learning mode policy",
			func() error { _, err := repo.GetLearningModePolicy(ctx, "policy/id"); return err }},
		{"get policy", http.MethodGet, policyPath, "", "get security policy",
			func() error { _, err := repo.GetSecurityPolicy(ctx, "policy/id"); return err }},
		{"list bot defense", http.MethodGet, "/mgmt/tm/security/bot-defense/profile", "", "get bot defense",
			func() error { _, err := repo.ListBotDefense(ctx); return err }},
		{"get bot defense", http.MethodGet, "/mgmt/tm/security/bot-defense/profile/~Common~bot%2Fname", "", "list bot defense",
			func() error { _, err := repo.GetBotDefense(ctx, "~Common~bot/name"); return err }},
		{"list log profile", http.MethodGet, "/mgmt/tm/security/log/profile", "", "get log profile",
			func() error { _, err := repo.ListLogProfile(ctx); return err }},
		{"get log profile", http.MethodGet, "/mgmt/tm/security/log/profile/~Common~log%2Fname", "", "get log profile",
			func() error { _, err := repo.GetLogProfile(ctx, "~Common~log/name"); return err }},
		{"attach profiles", http.MethodPatch, "/mgmt/tm/ltm/virtual/~Common~vs%2Fname", `{"securityLogProfiles":["/Common/log"],"botDefenseProfile":{"link":"bot-link"}}`, "attach bot defense and log profile to vs",
			func() error {
				_, err := repo.AttachBotDefenseandLogProfileToVS(ctx, "~Common~vs/name", request.F5AttachBotDefenseandLogProfileToVS{SecurityLogProfiles: []string{"/Common/log"}, BotDefenseProfile: request.F5Link{Link: "bot-link"}})
				return err
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantMethod, wantPath, wantBody = tt.method, tt.path, tt.body
			status, body = http.StatusOK, `{}`
			if err := tt.call(); err != nil {
				t.Fatalf("success response: %v", err)
			}
			status, body = http.StatusForbidden, `{"message":"denied"}`
			if err := tt.call(); err == nil || err.Error() != tt.operation+" failed: denied" {
				t.Fatalf("error response: %v", err)
			}
		})
	}
}
