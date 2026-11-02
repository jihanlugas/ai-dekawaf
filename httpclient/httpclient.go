package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client *http.Client

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Body    interface{}
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func do(ctx context.Context, req Request) (*Response, error) {
	var body io.Reader

	if req.Body != nil {
		payload, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}

		body = bytes.NewBuffer(payload)
		fmt.Printf("\n--- HTTP Request ---\n%s %s\nBody:\n%s\n", req.Method, req.URL, formatBody(payload))
	} else {
		fmt.Printf("\n--- HTTP Request ---\n%s %s\nBody: (empty)\n", req.Method, req.URL)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		req.Method,
		req.URL,
		body,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	if req.Body != nil && httpReq.Header.Get("Content-Type") == "" {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}
	fmt.Printf("--- HTTP Response ---\nStatus: %s\nBody:\n%s\n\n", httpResp.Status, formatBody(respBody))

	return &Response{
		StatusCode: httpResp.StatusCode,
		Headers:    httpResp.Header,
		Body:       respBody,
	}, nil
}

func formatBody(body []byte) string {
	if len(body) == 0 {
		return "(empty)"
	}

	var prettyBody bytes.Buffer
	if err := json.Indent(&prettyBody, body, "", "  "); err == nil {
		return prettyBody.String()
	}

	return string(body)
}

func Get(ctx context.Context, url string, headers map[string]string) (*Response, error) {
	return do(ctx, Request{
		Method:  http.MethodGet,
		URL:     url,
		Headers: headers,
	})
}

func Post(ctx context.Context, url string, headers map[string]string, payload interface{}) (*Response, error) {
	return do(ctx, Request{
		Method:  http.MethodPost,
		URL:     url,
		Headers: headers,
		Body:    payload,
	})
}

func Put(ctx context.Context, url string, headers map[string]string, payload interface{}) (*Response, error) {
	return do(ctx, Request{
		Method:  http.MethodPut,
		URL:     url,
		Headers: headers,
		Body:    payload,
	})
}

func Delete(ctx context.Context, url string, headers map[string]string) (*Response, error) {
	return do(ctx, Request{
		Method:  http.MethodDelete,
		URL:     url,
		Headers: headers,
	})
}

func Patch(ctx context.Context, url string, headers map[string]string, payload interface{}) (*Response, error) {
	return do(ctx, Request{
		Method:  http.MethodPatch,
		URL:     url,
		Headers: headers,
		Body:    payload,
	})
}
