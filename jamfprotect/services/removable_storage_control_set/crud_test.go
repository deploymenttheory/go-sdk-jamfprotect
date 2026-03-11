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

func TestUSBControlSetService_ListUSBControlSets_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "listUSBControlSets", 200, "list_usb_control_sets_empty.json")

	result, _, err := service.ListUSBControlSets(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestUSBControlSetService_ListUSBControlSetNames_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/app", "listUsbControlNames", 200, "list_usb_control_set_names_empty.json")

	result, _, err := service.ListUSBControlSetNames(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestUSBControlSetService_CreateUSBControlSet_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "createUSBControlSet", 500, "", "Internal Server Error")

	req := &removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:               "test",
		DefaultMountAction: "Prevented",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					Vendors:     []string{"0x1234"},
					MountAction: "ReadWrite",
				},
			},
		},
	}

	result, _, err := service.CreateUSBControlSet(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create USB control set")
}

func TestUSBControlSetService_GetUSBControlSet_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "getUSBControlSet", 500, "", "Internal Server Error")

	result, _, err := service.GetUSBControlSet(context.Background(), "test-id")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get USB control set")
}

func TestUSBControlSetService_UpdateUSBControlSet_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "updateUSBControlSet", 500, "", "Internal Server Error")

	req := &removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:               "test",
		DefaultMountAction: "Prevented",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					Vendors:     []string{"0x1234"},
					MountAction: "ReadWrite",
				},
			},
		},
	}

	result, _, err := service.UpdateUSBControlSet(context.Background(), "test-id", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update USB control set")
}

func TestUSBControlSetService_DeleteUSBControlSet_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "deleteUSBControlSet", 500, "", "Internal Server Error")

	_, err := service.DeleteUSBControlSet(context.Background(), "test-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete USB control set")
}

func TestUSBControlSetService_ListUSBControlSets_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "listUSBControlSets", 500, "", "Internal Server Error")

	result, _, err := service.ListUSBControlSets(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestUSBControlSetService_ListUSBControlSetNames_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "listUSBControlSetNames", 500, "", "Internal Server Error")

	result, _, err := service.ListUSBControlSetNames(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
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
			name: "CreateUSBControlSet missing defaultMountAction",
			fn: func() error {
				_, _, err := service.CreateUSBControlSet(context.Background(), &removablestoragecontrolset.CreateUSBControlSetRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "defaultMountAction is required",
		},
		{
			name: "CreateUSBControlSet nil rules",
			fn: func() error {
				_, _, err := service.CreateUSBControlSet(context.Background(), &removablestoragecontrolset.CreateUSBControlSetRequest{
					Name:               "test",
					DefaultMountAction: "ReadOnly",
				})
				return err
			},
			wantErr: "rules is required",
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
			name: "UpdateUSBControlSet empty id",
			fn: func() error {
				_, _, err := service.UpdateUSBControlSet(context.Background(), "", &removablestoragecontrolset.UpdateUSBControlSetRequest{
					Name:               "test",
					DefaultMountAction: "ReadOnly",
					Rules:              []removablestoragecontrolset.USBControlRuleInput{},
				})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateUSBControlSet nil request",
			fn: func() error {
				_, _, err := service.UpdateUSBControlSet(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdateUSBControlSet missing name",
			fn: func() error {
				_, _, err := service.UpdateUSBControlSet(context.Background(), "test-id", &removablestoragecontrolset.UpdateUSBControlSetRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "UpdateUSBControlSet missing defaultMountAction",
			fn: func() error {
				_, _, err := service.UpdateUSBControlSet(context.Background(), "test-id", &removablestoragecontrolset.UpdateUSBControlSetRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "defaultMountAction is required",
		},
		{
			name: "UpdateUSBControlSet nil rules",
			fn: func() error {
				_, _, err := service.UpdateUSBControlSet(context.Background(), "test-id", &removablestoragecontrolset.UpdateUSBControlSetRequest{
					Name:               "test",
					DefaultMountAction: "ReadOnly",
				})
				return err
			},
			wantErr: "rules is required",
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

func TestUSBControlSetService_CreateWithRules(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateUSBControlSetMock()

	req := &removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:                 "Test with Rules",
		Description:          "USB control set with rules",
		DefaultMountAction:   "ReadOnly",
		DefaultMessageAction: "NOTIFY",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "Prevented",
					Vendors:     []string{"1234"},
				},
			},
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "ReadWrite",
					Vendors:     []string{"abcd"},
				},
			},
		},
	}

	result, _, err := service.CreateUSBControlSet(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestUSBControlSetService_UpdateWithRules(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateUSBControlSetMock()

	req := &removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:                 "Updated with Rules",
		Description:          "Updated USB control set",
		DefaultMountAction:   "Prevented",
		DefaultMessageAction: "ALERT",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "Prevented",
					Vendors:     []string{"9999"},
				},
			},
		},
	}

	result, _, err := service.UpdateUSBControlSet(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestUSBControlSetService_Validators(t *testing.T) {
	assert.NoError(t, removablestoragecontrolset.ValidateUSBControlSetID("test-id"))
	assert.NoError(t, removablestoragecontrolset.ValidateUSBControlSetID(""))

	assert.NoError(t, removablestoragecontrolset.ValidateDefaultMountAction("ReadOnly"))
	assert.NoError(t, removablestoragecontrolset.ValidateDefaultMountAction("ReadWrite"))
	assert.NoError(t, removablestoragecontrolset.ValidateDefaultMountAction("Prevented"))
	assert.Error(t, removablestoragecontrolset.ValidateDefaultMountAction("Block"))

	assert.NoError(t, removablestoragecontrolset.ValidateRuleMountAction("ReadOnly"))
	assert.NoError(t, removablestoragecontrolset.ValidateRuleMountAction("ReadWrite"))
	assert.NoError(t, removablestoragecontrolset.ValidateRuleMountAction("Prevented"))
	assert.Error(t, removablestoragecontrolset.ValidateRuleMountAction("INVALID"))

	assert.NoError(t, removablestoragecontrolset.ValidateCreateUSBControlSetRequest(&removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:               "test",
		DefaultMountAction: "ReadOnly",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "Prevented",
					Vendors:     []string{"1234"},
				},
			},
			{
				Type: "serial",
				SerialRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "ReadWrite",
					Serials:     []string{"ABC123"},
				},
			},
			{
				Type: "product",
				ProductRule: &removablestoragecontrolset.USBControlProductRuleDetails{
					MountAction: "ReadOnly",
					Products: []removablestoragecontrolset.USBControlProductPair{
						{Vendor: "1234", Product: "5678"},
					},
				},
			},
			{
				Type: "encryption",
				EncryptionRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "ReadWrite",
				},
			},
		},
	}))
	assert.NoError(t, removablestoragecontrolset.ValidateCreateUSBControlSetRequest(nil))
	assert.Error(t, removablestoragecontrolset.ValidateCreateUSBControlSetRequest(&removablestoragecontrolset.CreateUSBControlSetRequest{
		DefaultMountAction: "INVALID",
	}))

	assert.NoError(t, removablestoragecontrolset.ValidateUpdateUSBControlSetRequest(&removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:               "test",
		DefaultMountAction: "ReadWrite",
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "Prevented",
					Vendors:     []string{"1234"},
				},
			},
		},
	}))
	assert.NoError(t, removablestoragecontrolset.ValidateUpdateUSBControlSetRequest(nil))
	assert.Error(t, removablestoragecontrolset.ValidateUpdateUSBControlSetRequest(&removablestoragecontrolset.UpdateUSBControlSetRequest{
		Rules: []removablestoragecontrolset.USBControlRuleInput{
			{
				Type: "vendor",
				VendorRule: &removablestoragecontrolset.USBControlRuleDetails{
					MountAction: "INVALID",
					Vendors:     []string{"1234"},
				},
			},
		},
	}))
}
