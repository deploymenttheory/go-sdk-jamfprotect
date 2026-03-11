package mocks

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"resty.dev/v3"
)

// NewMockResponse creates a *resty.Response for testing purposes.
// It sets up the RawResponse with the given status code, headers, and body.
func NewMockResponse(statusCode int, headers http.Header, body []byte) *resty.Response {
	if headers == nil {
		headers = make(http.Header)
	}
	if body == nil {
		body = []byte{}
	}

	statusText := http.StatusText(statusCode)
	if statusText == "" {
		statusText = fmt.Sprintf("%d", statusCode)
	}

	rawResp := &http.Response{
		StatusCode: statusCode,
		Status:     fmt.Sprintf("%d %s", statusCode, statusText),
		Header:     headers,
		Body:       io.NopCloser(bytes.NewReader(body)),
	}

	return &resty.Response{
		RawResponse: rawResp,
	}
}
