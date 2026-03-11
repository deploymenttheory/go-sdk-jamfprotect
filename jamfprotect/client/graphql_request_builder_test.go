package client

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

func TestGraphQLRequestBuilder_SetQuery(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx:     context.Background(),
		headers: make(map[string]string),
	}

	result := builder.SetQuery("query { test }")
	assert.Equal(t, "query { test }", builder.query)
	assert.Equal(t, builder, result)
}

func TestGraphQLRequestBuilder_SetVariables(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx:     context.Background(),
		headers: make(map[string]string),
	}

	vars := map[string]any{"id": "123", "name": "test"}
	result := builder.SetVariables(vars)
	assert.Equal(t, vars, builder.variables)
	assert.Equal(t, builder, result)
}

func TestGraphQLRequestBuilder_SetTarget(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx:     context.Background(),
		headers: make(map[string]string),
	}

	var target map[string]any
	result := builder.SetTarget(&target)
	assert.Equal(t, &target, builder.target)
	assert.Equal(t, builder, result)
}

func TestGraphQLRequestBuilder_AddHeader(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx:     context.Background(),
		headers: make(map[string]string),
	}

	result := builder.AddHeader("X-Custom", "value")
	assert.Equal(t, "value", builder.headers["X-Custom"])
	assert.Equal(t, builder, result)
}

func TestGraphQLRequestBuilder_AddHeader_EmptyValue(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx:     context.Background(),
		headers: make(map[string]string),
	}

	builder.AddHeader("X-Custom", "")
	_, exists := builder.headers["X-Custom"]
	assert.False(t, exists)
}

func TestGraphQLRequestBuilder_AddHeader_NilMap(t *testing.T) {
	builder := &GraphQLRequestBuilder{
		ctx: context.Background(),
	}

	result := builder.AddHeader("X-Custom", "value")
	assert.NotNil(t, builder.headers)
	assert.Equal(t, "value", builder.headers["X-Custom"])
	assert.Equal(t, builder, result)
}

func TestGraphQLRequestBuilder_Post(t *testing.T) {
	called := false
	fn := func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
		called = true
		assert.Equal(t, "/app", path)
		assert.Equal(t, "query { test }", query)
		assert.Equal(t, "123", variables["id"])
		assert.Equal(t, "value", headers["X-Custom"])
		return &resty.Response{}, nil
	}

	builder := NewMockGraphQLRequestBuilder(context.Background(), fn)
	builder.query = "query { test }"
	builder.variables = map[string]any{"id": "123"}
	builder.headers["X-Custom"] = "value"

	_, err := builder.Post("/app")
	require.NoError(t, err)
	assert.True(t, called)
}

func TestGraphQLRequestBuilder_ChainedCalls(t *testing.T) {
	fn := func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
		return &resty.Response{}, nil
	}

	builder := NewMockGraphQLRequestBuilder(context.Background(), fn)

	var target map[string]any
	_, err := builder.
		SetQuery("query { test }").
		SetVariables(map[string]any{"id": "123"}).
		SetTarget(&target).
		AddHeader("X-Custom", "value").
		Post("/app")

	require.NoError(t, err)
	assert.Equal(t, "query { test }", builder.query)
	assert.Equal(t, "123", builder.variables["id"])
	assert.Equal(t, &target, builder.target)
	assert.Equal(t, "value", builder.headers["X-Custom"])
}

func TestNewMockGraphQLRequestBuilder(t *testing.T) {
	ctx := context.Background()
	called := false

	fn := func(path, query string, variables map[string]any, target any, headers map[string]string) (*resty.Response, error) {
		called = true
		return &resty.Response{}, nil
	}

	builder := NewMockGraphQLRequestBuilder(ctx, fn)
	assert.NotNil(t, builder)
	assert.Equal(t, ctx, builder.ctx)
	assert.NotNil(t, builder.headers)

	_, err := builder.Post("/test")
	require.NoError(t, err)
	assert.True(t, called)
}
