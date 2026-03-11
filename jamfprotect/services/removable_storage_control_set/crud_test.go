package removablestoragecontrolset_test

import (
	"context"
	"testing"

	removablestoragecontrolset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/removable_storage_control_set"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/removable_storage_control_set/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*removablestoragecontrolset.Service, *mocks.USBControlSetMock) {
	t.Helper()
	mock := mocks.NewUSBControlSetMock()
	return removablestoragecontrolset.NewService(mock), mock
}

func TestUSBControlSetService_CreateUSBControlSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateUSBControlSetMock()

	req := &removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:                 "Test USB Control Set",
		Description:          "A test USB control set",
		DefaultMountAction:   "ReadOnly",
		DefaultMessageAction: "NOTIFY",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}

	result, _, err := service.CreateUSBControlSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test USB Control Set", result.Name)
}

func TestUSBControlSetService_GetUSBControlSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetUSBControlSetMock()

	result, _, err := service.GetUSBControlSet(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test USB Control Set", result.Name)
}

func TestUSBControlSetService_UpdateUSBControlSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateUSBControlSetMock()

	req := &removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:                 "Updated USB Control Set",
		Description:          "An updated USB control set",
		DefaultMountAction:   "ReadOnly",
		DefaultMessageAction: "NOTIFY",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}

	result, _, err := service.UpdateUSBControlSet(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated USB Control Set", result.Name)
}

func TestUSBControlSetService_DeleteUSBControlSet(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteUSBControlSetMock()

	_, err := service.DeleteUSBControlSet(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestUSBControlSetService_ListUSBControlSets(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListUSBControlSetsMock()

	result, _, err := service.ListUSBControlSets(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
}

func TestUSBControlSetService_ListUSBControlSetNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListUSBControlSetNamesMock()

	result, _, err := service.ListUSBControlSetNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test USB Control Set", result[0])
}

func TestUSBControlSetService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateUSBControlSet nil request",
			fn: func() error {
				_, _, err := service.CreateUSBControlSet(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateUSBControlSet missing name",
			fn: func() error {
				_, _, err := service.CreateUSBControlSet(context.Background(), &removablestoragecontrolset.CreateUSBControlSetRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetUSBControlSet empty id",
			fn: func() error {
				_, _, err := service.GetUSBControlSet(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "DeleteUSBControlSet empty id",
			fn: func() error {
				_, err := service.DeleteUSBControlSet(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fn()
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}
