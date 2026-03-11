package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
	"resty.dev/v3"
)

// Post executes a POST request with JSON body. Auth is applied automatically via middleware.
func (t *Transport) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	req := t.client.R().
		SetContext(ctx).
		SetResult(result)

	if body != nil {
		req.SetBody(body)
	}

	t.applyHeaders(req, headers)

	return t.executeRequest(req, "POST", path)
}

// executeRequest is a centralized request executor that handles error processing.
func (t *Transport) executeRequest(req *resty.Request, method, path string) (*resty.Response, error) {
	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// Total retry deadline
	if t.totalRetryDuration > 0 {
		if _, hasDeadline := ctx.Deadline(); !hasDeadline {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, t.totalRetryDuration)
			defer cancel()
			req.SetContext(ctx)
		}
	}

	// Concurrency semaphore
	if t.sem != nil {
		select {
		case t.sem <- struct{}{}:
			defer func() { <-t.sem }()
		case <-ctx.Done():
			return nil, fmt.Errorf("concurrency limit: %w", ctx.Err())
		}
	}

	t.logger.Debug("Executing API request",
		zap.String("method", method),
		zap.String("path", path))

	var resp *resty.Response
	var err error

	switch method {
	case "GET":
		resp, err = req.Get(path)
	case "POST":
		resp, err = req.Post(path)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		t.logger.Error("Request failed",
			zap.String("method", method),
			zap.String("path", path),
			zap.Error(err))
		return resp, fmt.Errorf("request failed: %w", err)
	}

	if err := t.validateResponse(resp, method, path); err != nil {
		return resp, err
	}

	if resp.IsError() {
		return resp, ParseErrorResponse(
			[]byte(resp.String()),
			resp.StatusCode(),
			resp.Status(),
			method,
			path,
			t.logger,
		)
	}

	t.logger.Debug("Request completed successfully",
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status_code", resp.StatusCode()))

	// Mandatory fixed delay
	if t.requestDelay > 0 {
		time.Sleep(t.requestDelay)
	}

	return resp, nil
}
