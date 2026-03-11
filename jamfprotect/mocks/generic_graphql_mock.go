package mocks

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// Ensure GenericGraphQLMock implements client.GraphQLClient at compile time.
var _ client.GraphQLClient = (*GenericGraphQLMock)(nil)

// registeredGraphQLResponse holds a pre-canned response for a single GraphQL operation.
type registeredGraphQLResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string // if non-empty, return this error without parsing body
}

// GenericGraphQLMock is a reusable test double implementing client.GraphQLClient.
// Responses are keyed by "path:operationName" where operationName is the exact GraphQL
// operation name extracted from the query string (e.g. "createAnalytic", "listAnalytics").
type GenericGraphQLMock struct {
	name       string
	responses  map[string]registeredGraphQLResponse
	logger     *zap.Logger
	fixtureDir string
}

// GenericGraphQLMockConfig configures a GenericGraphQLMock instance.
type GenericGraphQLMockConfig struct {
	// Name is used in error messages (e.g., "AnalyticMock"). Defaults to "GenericGraphQLMock".
	Name string
	// FixtureDir is the directory containing fixture JSON files.
	// Defaults to the "mocks" subdirectory relative to the calling service package.
	FixtureDir string
}

// opNameRegex extracts the operation name from a GraphQL query or mutation.
var opNameRegex = regexp.MustCompile(`(?:query|mutation)\s+(\w+)`)

// extractOperationName returns the GraphQL operation name from a query string.
// Returns empty string if no operation name is found (anonymous query).
func extractOperationName(query string) string {
	if matches := opNameRegex.FindStringSubmatch(query); len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// NewGenericGraphQLMock creates a new GenericGraphQLMock with the given configuration.
func NewGenericGraphQLMock(config GenericGraphQLMockConfig) *GenericGraphQLMock {
	if config.Name == "" {
		config.Name = "GenericGraphQLMock"
	}
	if config.FixtureDir == "" {
		// Default to "mocks" directory relative to the calling service package.
		// Walk up the call stack to find the first frame outside the mocks package.
		for i := 1; i < 10; i++ {
			_, filename, _, ok := runtime.Caller(i)
			if !ok {
				break
			}
			dir := filepath.Dir(filename)
			if filepath.Base(dir) == "mocks" {
				continue
			}
			config.FixtureDir = filepath.Join(dir, "mocks")
			break
		}
	}
	return &GenericGraphQLMock{
		name:       config.Name,
		responses:  make(map[string]registeredGraphQLResponse),
		logger:     zap.NewNop(),
		fixtureDir: config.FixtureDir,
	}
}

// Register registers a mock success response for the given GraphQL path and operation name.
//
//   - path is the endpoint path (e.g., client.EndpointApp or client.EndpointGraphQL).
//   - operationName is the exact GraphQL operation name (e.g., "createAnalytic").
//   - statusCode is the HTTP status code to return.
//   - fixture is the filename of a JSON fixture in the fixture directory ("" for empty body).
func (m *GenericGraphQLMock) Register(path, operationName string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := m.loadFixture(fixture)
		if err != nil {
			panic(fmt.Sprintf("%s: failed to load fixture %q: %v", m.name, fixture, err))
		}
		body = data
	}
	m.responses[path+":"+operationName] = registeredGraphQLResponse{
		statusCode: statusCode,
		rawBody:    body,
	}
}

// RegisterError registers a mock error response.
//
//   - errMsg is returned as the error. If empty, a default message is generated from statusCode.
func (m *GenericGraphQLMock) RegisterError(path, operationName string, statusCode int, fixture string, errMsg string) {
	var body []byte
	if fixture != "" {
		data, err := m.loadFixture(fixture)
		if err != nil {
			panic(fmt.Sprintf("%s: failed to load error fixture %q: %v", m.name, fixture, err))
		}
		body = data
	}
	if errMsg == "" {
		errMsg = fmt.Sprintf("%s: error response %d for operation %s", m.name, statusCode, operationName)
	}
	m.responses[path+":"+operationName] = registeredGraphQLResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// NewRequest implements client.GraphQLClient.
// It returns a GraphQLRequestBuilder backed by this mock's dispatch logic.
func (m *GenericGraphQLMock) NewRequest(ctx context.Context) *client.GraphQLRequestBuilder {
	return client.NewMockGraphQLRequestBuilder(ctx, func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
		return m.dispatch(path, query, target)
	})
}

// dispatch extracts the operation name from query, looks up the registered response,
// builds a *resty.Response, decodes fixture data into target, and returns.
func (m *GenericGraphQLMock) dispatch(path, query string, target any) (*resty.Response, error) {
	opName := extractOperationName(query)
	key := path + ":" + opName

	r, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("%s: no response registered for path=%q operationName=%q", m.name, path, opName)
	}

	mockResp := newMockRestyResponse(r.statusCode, r.rawBody)

	// If an explicit error message is set, return it without parsing body.
	if r.errMsg != "" {
		return mockResp, fmt.Errorf("%s", r.errMsg)
	}

	// Nothing to decode.
	if target == nil || len(r.rawBody) == 0 {
		return mockResp, nil
	}

	// Decode the full GraphQL response envelope.
	var gqlResp struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.Unmarshal(r.rawBody, &gqlResp); err != nil {
		return mockResp, fmt.Errorf("%s: unmarshal graphql envelope: %w", m.name, err)
	}

	// Surface GraphQL-layer errors.
	if len(gqlResp.Errors) > 0 {
		msgs := make([]string, 0, len(gqlResp.Errors))
		for _, e := range gqlResp.Errors {
			if e.Message != "" {
				msgs = append(msgs, e.Message)
			}
		}
		if len(msgs) > 0 {
			return mockResp, fmt.Errorf("graphql operation failed: %s", strings.Join(msgs, "; "))
		}
	}

	// Unmarshal the data field into the caller-supplied target.
	if len(gqlResp.Data) > 0 {
		if err := json.Unmarshal(gqlResp.Data, target); err != nil {
			return mockResp, fmt.Errorf("%s: unmarshal graphql data: %w", m.name, err)
		}
	}

	return mockResp, nil
}

// GetLogger implements client.GraphQLClient.
func (m *GenericGraphQLMock) GetLogger() *zap.Logger {
	return m.logger
}

// loadFixture reads a fixture file from the configured fixture directory.
func (m *GenericGraphQLMock) loadFixture(filename string) ([]byte, error) {
	path := filepath.Join(m.fixtureDir, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

// newMockRestyResponse constructs a minimal *resty.Response for use in tests.
func newMockRestyResponse(statusCode int, body []byte) *resty.Response {
	if body == nil {
		body = []byte{}
	}

	headers := http.Header{"Content-Type": []string{"application/json"}}
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
