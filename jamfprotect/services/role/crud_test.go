package role_test

import (
	"context"
	"testing"

	role "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/role"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/role/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testRoleID = "role-id-1234"

func setupMockService(t *testing.T) (*role.Service, *mocks.RoleMock) {
	t.Helper()
	mock := mocks.NewRoleMock()
	return role.NewService(mock), mock
}

func TestRoleService_CreateRole(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateRoleMock()

	req := &role.CreateRoleRequest{
		Name:           "Test Role",
		ReadResources:  []string{"ANALYTIC"},
		WriteResources: []string{},
	}

	result, _, err := service.CreateRole(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testRoleID, result.ID)
	assert.Equal(t, "Test Role", result.Name)
	assert.Equal(t, []string{"ANALYTIC"}, result.Permissions.Read)
}

func TestRoleService_GetRole(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetRoleMock()

	result, _, err := service.GetRole(context.Background(), testRoleID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testRoleID, result.ID)
	assert.Equal(t, "Test Role", result.Name)
}

func TestRoleService_UpdateRole(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateRoleMock()

	req := &role.UpdateRoleRequest{
		Name:           "Updated Role",
		ReadResources:  []string{"ANALYTIC"},
		WriteResources: []string{},
	}

	result, _, err := service.UpdateRole(context.Background(), testRoleID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testRoleID, result.ID)
	assert.Equal(t, "Updated Role", result.Name)
}

func TestRoleService_DeleteRole(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteRoleMock()

	_, err := service.DeleteRole(context.Background(), testRoleID)

	require.NoError(t, err)
}

func TestRoleService_ListRoles(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListRolesMock()

	result, _, err := service.ListRoles(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testRoleID, result[0].ID)
	assert.Equal(t, "Test Role", result[0].Name)
}

func TestRoleService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateRole nil request",
			fn: func() error {
				_, _, err := service.CreateRole(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateRole missing name",
			fn: func() error {
				_, _, err := service.CreateRole(context.Background(), &role.CreateRoleRequest{
					ReadResources:  []string{},
					WriteResources: []string{},
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "CreateRole nil readResources",
			fn: func() error {
				_, _, err := service.CreateRole(context.Background(), &role.CreateRoleRequest{
					Name:           "Test",
					WriteResources: []string{},
				})
				return err
			},
			wantErr: "readResources is required",
		},
		{
			name: "CreateRole nil writeResources",
			fn: func() error {
				_, _, err := service.CreateRole(context.Background(), &role.CreateRoleRequest{
					Name:          "Test",
					ReadResources: []string{},
				})
				return err
			},
			wantErr: "writeResources is required",
		},
		{
			name: "GetRole empty id",
			fn: func() error {
				_, _, err := service.GetRole(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateRole empty id",
			fn: func() error {
				_, _, err := service.UpdateRole(context.Background(), "", &role.UpdateRoleRequest{Name: "test"})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateRole nil request",
			fn: func() error {
				_, _, err := service.UpdateRole(context.Background(), testRoleID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteRole empty id",
			fn: func() error {
				_, err := service.DeleteRole(context.Background(), "")
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
