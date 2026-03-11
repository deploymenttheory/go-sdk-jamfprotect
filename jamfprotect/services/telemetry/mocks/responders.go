package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// TelemetryMock provides mock responses for the Telemetry service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type TelemetryMock struct {
	*coremocks.GenericGraphQLMock
}

// NewTelemetryMock creates a new TelemetryMock instance.
func NewTelemetryMock() *TelemetryMock {
	return &TelemetryMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "TelemetryMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for telemetry operations.
func (m *TelemetryMock) RegisterMocks() {
	m.RegisterCreateTelemetryV2Mock()
	m.RegisterGetTelemetryV2Mock()
	m.RegisterUpdateTelemetryV2Mock()
	m.RegisterDeleteTelemetryV2Mock()
	m.RegisterListTelemetriesV2Mock()
	m.RegisterListTelemetriesCombinedMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *TelemetryMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateTelemetryV2Mock registers a success mock for createTelemetryV2.
func (m *TelemetryMock) RegisterCreateTelemetryV2Mock() {
	m.Register(client.EndpointApp, "createTelemetryV2", 200, "create_telemetry_v2_success.json")
}

// RegisterGetTelemetryV2Mock registers a success mock for getTelemetryV2.
func (m *TelemetryMock) RegisterGetTelemetryV2Mock() {
	m.Register(client.EndpointApp, "getTelemetryV2", 200, "get_telemetry_v2_success.json")
}

// RegisterUpdateTelemetryV2Mock registers a success mock for updateTelemetryV2.
func (m *TelemetryMock) RegisterUpdateTelemetryV2Mock() {
	m.Register(client.EndpointApp, "updateTelemetryV2", 200, "update_telemetry_v2_success.json")
}

// RegisterDeleteTelemetryV2Mock registers a success mock for deleteTelemetryV2.
func (m *TelemetryMock) RegisterDeleteTelemetryV2Mock() {
	m.Register(client.EndpointApp, "deleteTelemetryV2", 200, "delete_telemetry_v2_success.json")
}

// RegisterListTelemetriesV2Mock registers a success mock for listTelemetriesV2.
func (m *TelemetryMock) RegisterListTelemetriesV2Mock() {
	m.Register(client.EndpointApp, "listTelemetriesV2", 200, "list_telemetries_v2_success.json")
}

// RegisterListTelemetriesCombinedMock registers a success mock for listTelemetriesCombined.
func (m *TelemetryMock) RegisterListTelemetriesCombinedMock() {
	m.Register(client.EndpointApp, "listTelemetriesCombined", 200, "list_telemetries_combined_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *TelemetryMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getTelemetryV2", 200, "error_not_found.json", "graphql operation failed: Telemetry not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *TelemetryMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getTelemetryV2", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
