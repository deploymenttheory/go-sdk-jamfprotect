# Jamf Protect API SDK for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/deploymenttheory/go-api-sdk-jamfprotect.svg)](https://pkg.go.dev/github.com/deploymenttheory/go-api-sdk-jamfprotect)
[![Go Report Card](https://goreportcard.com/badge/github.com/deploymenttheory/go-api-sdk-jamfprotect)](https://goreportcard.com/report/github.com/deploymenttheory/go-api-sdk-jamfprotect)
[![codecov](https://codecov.io/gh/deploymenttheory/go-sdk-jamfprotect/graph/badge.svg)](https://codecov.io/gh/deploymenttheory/go-sdk-jamfprotect)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A comprehensive Go SDK for interacting with the Jamf Protect GraphQL API. This SDK provides a type-safe, idiomatic Go interface for managing Jamf Protect resources including plans, analytics, action configurations, and more.

## Features

- **GraphQL-First Design**: Built specifically for GraphQL APIs with proper query/mutation support
- **OAuth2 Authentication**: Secure client credentials flow with automatic token refresh
- **Type-Safe Operations**: Strongly-typed Go structs for all API resources
- **Automatic Pagination**: Built-in handling of paginated responses
- **Comprehensive Error Handling**: Detailed GraphQL error parsing and reporting
- **Logging Support**: Integrated with zap for structured logging
- **Context Support**: Full context.Context support for cancellation and timeouts
- **Production Ready**: Used in production Terraform providers

## Installation

```bash
go get github.com/deploymenttheory/go-api-sdk-jamfprotect
```

## Quick Start

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
    "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
)

func main() {
    // Create client with credentials
    client, err := jamfprotect.NewClient(
        "your-client-id",
        "your-client-secret",
    )
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    
    // List all plans
    plans, err := client.Plans.ListPlans(ctx)
    if err != nil {
        log.Fatal(err)
    }
    
    for _, plan := range plans {
        log.Printf("Plan: %s (ID: %s)", plan.Name, plan.ID)
    }
}
```

### Using Environment Variables

```go
// Set environment variables:
// export JAMFPROTECT_CLIENT_ID="your-client-id"
// export JAMFPROTECT_CLIENT_SECRET="your-client-secret"
// export JAMFPROTECT_BASE_URL="https://apis.jamfprotect.cloud" (optional)

client, err := jamfprotect.NewClientFromEnv()
if err != nil {
    log.Fatal(err)
}
```

### With Custom Configuration

```go
client, err := jamfprotect.NewClient(
    "your-client-id",
    "your-client-secret",
    client.WithBaseURL("https://custom.jamfprotect.cloud"),
    client.WithTimeout(60 * time.Second),
    client.WithDebug(),
)
```

## Services

### Plans

Manage Jamf Protect security plans:

```go
// Create a plan
logLevel := "INFO"
request := &plans.CreatePlanRequest{
    Name:          "Production Security Plan",
    Description:   "Security configuration for production systems",
    LogLevel:      &logLevel,
    ActionConfigs: "action-config-id",
    AutoUpdate:    true,
    CommsConfig: plans.CommsConfigInput{
        FQDN:     "protect.example.com",
        Protocol: "HTTPS",
    },
    InfoSync: plans.InfoSyncInput{
        Attrs:                []string{"hostname", "osVersion"},
        InsightsSyncInterval: 3600,
    },
    SignaturesFeedConfig: plans.SignaturesFeedConfigInput{
        Mode: "AUTO",
    },
}

plan, err := client.Plans.CreatePlan(ctx, request)

// Get a plan
plan, err := client.Plans.GetPlan(ctx, "plan-id")

// Update a plan
plan, err := client.Plans.UpdatePlan(ctx, "plan-id", updateRequest)

// Delete a plan
err := client.Plans.DeletePlan(ctx, "plan-id")

// List all plans (with automatic pagination)
plans, err := client.Plans.ListPlans(ctx)
```

## Configuration Options

The SDK supports various configuration options:

### Base URL
```go
client.WithBaseURL("https://custom.jamfprotect.cloud")
```

### Timeout
```go
client.WithTimeout(90 * time.Second)
```

### Custom HTTP Client
```go
httpClient := &http.Client{
    Timeout: 60 * time.Second,
}
client.WithHTTPClient(httpClient)
```

### Logging
```go
// With zap logger
logger, _ := zap.NewProduction()
client.WithLogger(logger)

// Debug mode
client.WithDebug()
```

### User Agent
```go
// Custom user agent
client.WithUserAgent("MyApp/1.0.0")

// Append to default user agent
client.WithCustomAgent("MyApp/1.0.0")
```

## Examples

Comprehensive examples for each service are available in the [examples](./examples) directory:

### Plans
- [Create Plan](./examples/jamfprotect/plans/CreatePlan/main.go)
- [Get Plan](./examples/jamfprotect/plans/GetPlan/main.go)
- [Update Plan](./examples/jamfprotect/plans/UpdatePlan/main.go)
- [Delete Plan](./examples/jamfprotect/plans/DeletePlan/main.go)
- [List Plans](./examples/jamfprotect/plans/ListPlans/main.go)

## Error Handling

The SDK provides comprehensive error handling with specific error types:

```go
plan, err := client.Plans.GetPlan(ctx, "invalid-id")
if err != nil {
    if errors.Is(err, client.ErrNotFound) {
        log.Println("Plan not found")
    } else if errors.Is(err, client.ErrAuthentication) {
        log.Println("Authentication failed")
    } else if errors.Is(err, client.ErrGraphQL) {
        log.Printf("GraphQL error: %v", err)
    } else {
        log.Printf("Unexpected error: %v", err)
    }
}
```

## GraphQL API

This SDK is built specifically for the Jamf Protect GraphQL API. The API uses:

- **OAuth2 Client Credentials Flow** for authentication
- **GraphQL** for all data operations
- **Two Endpoints**:
  - `/app` - Full API access (recommended)
  - `/graphql` - Limited schema endpoint

## Architecture

The SDK follows a layered architecture:

```
jamfprotect/
├── client/          # HTTP transport, auth, GraphQL handling
├── services/        # Service-specific operations
│   └── plans/      # Plans service
└── new.go          # Main client wrapper
```

### Key Components

- **Transport Layer**: Handles HTTP communication, OAuth2 authentication, and GraphQL request/response processing
- **Service Layer**: Provides domain-specific operations (Plans, Analytics, etc.)
- **Client Wrapper**: Aggregates all services into a unified interface

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: [Full API Documentation](https://pkg.go.dev/github.com/deploymenttheory/go-api-sdk-jamfprotect)
- **Issues**: [GitHub Issues](https://github.com/deploymenttheory/go-api-sdk-jamfprotect/issues)
- **Jamf Protect API**: [Official API Documentation](https://learn.jamf.com/bundle/jamf-protect-documentation/page/API_Documentation.html)

## Related Projects

- [terraform-provider-jamfprotect](https://github.com/smithjw/terraform-provider-jamfprotect) - Terraform provider using this SDK
- [go-api-sdk-jamfpro](https://github.com/deploymenttheory/go-api-sdk-jamfpro) - Jamf Pro API SDK
- [go-api-sdk-virustotal](https://github.com/deploymenttheory/go-api-sdk-virustotal) - VirusTotal API SDK

## Acknowledgments

Built with inspiration from the patterns established in the VirusTotal SDK and the smithjw Terraform provider for Jamf Protect.
