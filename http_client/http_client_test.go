package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

type mockRoundTripper struct{}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	bodyBytes, _ := io.ReadAll(req.Body)
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var r Req
	json.Unmarshal(bodyBytes, &r)

	respBody := Resp{
		Status:  200,
		Message: "OK",
		Payload: map[string]string{
			"name_received": r.Name,
		},
	}

	respJSON, _ := json.Marshal(respBody)

	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBuffer(respJSON)),
		Header:     make(http.Header),
	}, nil
}

func TestRun(t *testing.T) {
	client := &http.Client{
		Transport: &mockRoundTripper{},
	}

	Run(client, "http://example.com")
}
