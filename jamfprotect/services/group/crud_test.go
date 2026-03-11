package group_test

import (
	"context"
	"testing"

	group "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/group"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/group/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testGroupID = "group-id-1234"

func setupMockService(t *testing.T) (*group.Service, *mocks.GroupMock) {
	t.Helper()
	mock := mocks.NewGroupMock()
	return group.NewService(mock), mock
}

func TestGroupService_CreateGroup(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateGroupMock()

	req := &group.CreateGroupRequest{
		Name:        "Test Group",
		AccessGroup: false,
		RoleIDs:     []string{},
	}

	result, _, err := service.CreateGroup(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testGroupID, result.ID)
	assert.Equal(t, "Test Group", result.Name)
}

func TestGroupService_GetGroup(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetGroupMock()

	result, _, err := service.GetGroup(context.Background(), testGroupID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testGroupID, result.ID)
	assert.Equal(t, "Test Group", result.Name)
}

func TestGroupService_UpdateGroup(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateGroupMock()

	req := &group.UpdateGroupRequest{
		Name:        "Updated Group",
		AccessGroup: false,
		RoleIDs:     []string{},
	}

	result, _, err := service.UpdateGroup(context.Background(), testGroupID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testGroupID, result.ID)
	assert.Equal(t, "Updated Group", result.Name)
}

func TestGroupService_DeleteGroup(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteGroupMock()

	_, err := service.DeleteGroup(context.Background(), testGroupID)

	require.NoError(t, err)
}

func TestGroupService_ListGroups(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListGroupsMock()

	result, _, err := service.ListGroups(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testGroupID, result[0].ID)
	assert.Equal(t, "Test Group", result[0].Name)
}

func TestGroupService_ListGroups_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listGroups", 200, "list_groups_empty.json")

	result, _, err := service.ListGroups(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestGroupService_CreateGroup_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "createGroup", 500, "", "Internal Server Error")

	req := &group.CreateGroupRequest{
		Name:    "test",
		RoleIDs: []string{"role-1"},
	}

	result, _, err := service.CreateGroup(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create group")
}

func TestGroupService_GetGroup_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "getGroup", 500, "", "Internal Server Error")

	result, _, err := service.GetGroup(context.Background(), "test-id")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get group")
}

func TestGroupService_UpdateGroup_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "updateGroup", 500, "", "Internal Server Error")

	req := &group.UpdateGroupRequest{
		Name:    "test",
		RoleIDs: []string{"role-1"},
	}

	result, _, err := service.UpdateGroup(context.Background(), "test-id", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update group")
}

func TestGroupService_DeleteGroup_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "deleteGroup", 500, "", "Internal Server Error")

	_, err := service.DeleteGroup(context.Background(), "test-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete group")
}

func TestGroupService_ListGroups_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/graphql", "listGroups", 500, "", "Internal Server Error")

	result, _, err := service.ListGroups(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestGroupService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateGroup nil request",
			fn: func() error {
				_, _, err := service.CreateGroup(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateGroup missing name",
			fn: func() error {
				_, _, err := service.CreateGroup(context.Background(), &group.CreateGroupRequest{})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetGroup empty id",
			fn: func() error {
				_, _, err := service.GetGroup(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateGroup empty id",
			fn: func() error {
				_, _, err := service.UpdateGroup(context.Background(), "", &group.UpdateGroupRequest{Name: "test"})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateGroup nil request",
			fn: func() error {
				_, _, err := service.UpdateGroup(context.Background(), testGroupID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteGroup empty id",
			fn: func() error {
				_, err := service.DeleteGroup(context.Background(), "")
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
