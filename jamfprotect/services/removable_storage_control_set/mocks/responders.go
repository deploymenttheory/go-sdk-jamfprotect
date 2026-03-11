package mocks

import (
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	coremocks "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/mocks"
)

// USBControlSetMock provides mock responses for the RemovableStorageControlSet service GraphQL operations.
// All operations POST to the /app GraphQL endpoint and are distinguished by operation name
// in the request body.
type USBControlSetMock struct {
	*coremocks.GenericGraphQLMock
}

// NewUSBControlSetMock creates a new USBControlSetMock instance.
func NewUSBControlSetMock() *USBControlSetMock {
	return &USBControlSetMock{
		GenericGraphQLMock: coremocks.NewGenericGraphQLMock(coremocks.GenericGraphQLMockConfig{
			Name: "USBControlSetMock",
		}),
	}
}

// RegisterMocks registers all successful response mocks for USB control set operations.
func (m *USBControlSetMock) RegisterMocks() {
	m.RegisterCreateUSBControlSetMock()
	m.RegisterGetUSBControlSetMock()
	m.RegisterUpdateUSBControlSetMock()
	m.RegisterDeleteUSBControlSetMock()
	m.RegisterListUSBControlSetsMock()
	m.RegisterListUSBControlSetNamesMock()
}

// RegisterErrorMocks registers error response mocks.
func (m *USBControlSetMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterUnauthorizedErrorMock()
}

// RegisterCreateUSBControlSetMock registers a success mock for createUSBControlSet.
func (m *USBControlSetMock) RegisterCreateUSBControlSetMock() {
	m.Register(client.EndpointApp, "createUSBControlSet", 200, "create_usb_control_set_success.json")
}

// RegisterGetUSBControlSetMock registers a success mock for getUSBControlSet.
func (m *USBControlSetMock) RegisterGetUSBControlSetMock() {
	m.Register(client.EndpointApp, "getUSBControlSet", 200, "get_usb_control_set_success.json")
}

// RegisterUpdateUSBControlSetMock registers a success mock for updateUSBControlSet.
func (m *USBControlSetMock) RegisterUpdateUSBControlSetMock() {
	m.Register(client.EndpointApp, "updateUSBControlSet", 200, "update_usb_control_set_success.json")
}

// RegisterDeleteUSBControlSetMock registers a success mock for deleteUSBControlSet.
func (m *USBControlSetMock) RegisterDeleteUSBControlSetMock() {
	m.Register(client.EndpointApp, "deleteUSBControlSet", 200, "delete_usb_control_set_success.json")
}

// RegisterListUSBControlSetsMock registers a success mock for listUSBControlSets.
func (m *USBControlSetMock) RegisterListUSBControlSetsMock() {
	m.Register(client.EndpointApp, "listUSBControlSets", 200, "list_usb_control_sets_success.json")
}

// RegisterListUSBControlSetNamesMock registers a success mock for listUsbControlNames.
func (m *USBControlSetMock) RegisterListUSBControlSetNamesMock() {
	m.Register(client.EndpointApp, "listUsbControlNames", 200, "list_usb_control_set_names_success.json")
}

// RegisterNotFoundErrorMock registers a not-found error mock.
func (m *USBControlSetMock) RegisterNotFoundErrorMock() {
	m.RegisterError(client.EndpointApp, "getUSBControlSet", 200, "error_not_found.json", "graphql operation failed: USBControlSet not found")
}

// RegisterUnauthorizedErrorMock registers a 401 unauthorized error mock.
func (m *USBControlSetMock) RegisterUnauthorizedErrorMock() {
	m.RegisterError(client.EndpointApp, "getUSBControlSet", 401, "error_unauthorized.json", "Jamf Protect API error: unauthorized")
}
