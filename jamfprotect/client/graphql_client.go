package client

import (
	"context"

	"go.uber.org/zap"
)

// GraphQLClient is the interface service implementations depend on.
// Transport satisfies it; tests supply GenericGraphQLMock.
type GraphQLClient interface {
	// NewRequest returns a GraphQLRequestBuilder that the service layer uses
	// to construct a complete GraphQL request — query, variables, target, headers —
	// before executing it via Post. Auth, retry, and concurrency limiting are
	// applied by the transport at execution time.
	NewRequest(ctx context.Context) *GraphQLRequestBuilder

	// GetLogger returns the configured zap logger instance.
	GetLogger() *zap.Logger
}
