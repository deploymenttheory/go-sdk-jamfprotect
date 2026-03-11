package client

import (
	"context"

	"resty.dev/v3"
)

// graphQLExecutor is the internal execution backend for a GraphQLRequestBuilder.
// Transport implements it directly; tests supply a mock via NewMockGraphQLRequestBuilder.
type graphQLExecutor interface {
	executeGraphQL(ctx context.Context, path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error)
}

// GraphQLRequestBuilder constructs a single GraphQL request. The service layer
// (serialization) owns the full request shape — query, variables, target, headers —
// before handing the completed request to the executor (transport) which handles
// auth, retry, concurrency limiting, and throttling.
//
// Usage:
//
//	resp, err := s.client.NewRequest(ctx).
//	    SetQuery(createAnalyticMutation).
//	    SetVariables(vars).
//	    SetTarget(&result).
//	    Post(client.EndpointApp)
type GraphQLRequestBuilder struct {
	ctx       context.Context
	query     string
	variables map[string]any
	headers   map[string]string
	target    any
	executor  graphQLExecutor
}

// SetQuery sets the GraphQL query or mutation string.
func (b *GraphQLRequestBuilder) SetQuery(q string) *GraphQLRequestBuilder {
	b.query = q
	return b
}

// SetVariables sets the GraphQL variables map.
func (b *GraphQLRequestBuilder) SetVariables(v map[string]any) *GraphQLRequestBuilder {
	b.variables = v
	return b
}

// SetTarget sets the pointer for JSON unmarshaling of a successful response.
func (b *GraphQLRequestBuilder) SetTarget(t any) *GraphQLRequestBuilder {
	b.target = t
	return b
}

// AddHeader adds a request-level header. Empty values are ignored.
func (b *GraphQLRequestBuilder) AddHeader(key, value string) *GraphQLRequestBuilder {
	if b.headers == nil {
		b.headers = make(map[string]string)
	}
	if value != "" {
		b.headers[key] = value
	}
	return b
}

// Post executes the GraphQL request as POST against path.
func (b *GraphQLRequestBuilder) Post(path string) (*resty.Response, error) {
	return b.executor.executeGraphQL(b.ctx, path, b.query, b.variables, b.target, b.headers)
}

// mockGraphQLExecutor backs a GraphQLRequestBuilder in tests, routing execution
// through a caller-supplied dispatch function instead of a real Transport.
type mockGraphQLExecutor struct {
	fn func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error)
}

func (m *mockGraphQLExecutor) executeGraphQL(ctx context.Context, path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
	return m.fn(path, query, variables, target, headers)
}

// NewMockGraphQLRequestBuilder returns a GraphQLRequestBuilder suitable for unit tests.
// The fn callback receives all GraphQL parameters and returns a pre-programmed response.
func NewMockGraphQLRequestBuilder(ctx context.Context, fn func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error)) *GraphQLRequestBuilder {
	return &GraphQLRequestBuilder{
		ctx:      ctx,
		headers:  make(map[string]string),
		executor: &mockGraphQLExecutor{fn: fn},
	}
}
