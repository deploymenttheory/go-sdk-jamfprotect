package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"resty.dev/v3"
)

// GraphQLRequest represents a GraphQL request payload.
type GraphQLRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

// NewRequest returns a GraphQLRequestBuilder for this transport.
// The service layer uses it to construct the full request — query, variables,
// target, headers — before calling Post to execute it.
// Auth, retry, and concurrency limiting are applied by the transport.
func (t *Transport) NewRequest(ctx context.Context) *GraphQLRequestBuilder {
	return &GraphQLRequestBuilder{
		ctx:      ctx,
		headers:  make(map[string]string),
		executor: t,
	}
}

// executeGraphQL implements graphQLExecutor for Transport.
func (t *Transport) executeGraphQL(ctx context.Context, path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
	if path == "" {
		return nil, fmt.Errorf("%w: path is required", ErrInvalidInput)
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	payload := GraphQLRequest{Query: query, Variables: variables}
	var gqlResp GraphQLResponse

	clientResp, err := t.Post(ctx, path, payload, headers, &gqlResp)
	if err != nil {
		return clientResp, err
	}

	if err := MapGraphQLErrors(gqlResp.Errors); err != nil {
		return clientResp, err
	}

	if target == nil || len(gqlResp.Data) == 0 {
		return clientResp, nil
	}

	if err := json.Unmarshal(gqlResp.Data, target); err != nil {
		return clientResp, fmt.Errorf("decoding graphql response: %w", err)
	}

	return clientResp, nil
}
