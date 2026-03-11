package custompreventlist_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*custompreventlist.Service, *mocks.PreventListMock) {
	t.Helper()
	mock := mocks.NewPreventListMock()
	return custompreventlist.NewService(mock), mock
}

func TestPreventListService_CreatePreventList(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreatePreventListMock()

	req := &custompreventlist.CreatePreventListRequest{
		Name:        "Test Prevent List",
		Description: "A test prevent list",
		Type:        "FILEHASH",
	}

	result, _, err := service.CreatePreventList(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Prevent List", result.Name)
}

func TestPreventListService_GetPreventList(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetPreventListMock()

	result, _, err := service.GetPreventList(context.Background(), "test-id-1234")

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Test Prevent List", result.Name)
}

func TestPreventListService_UpdatePreventList(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdatePreventListMock()

	req := &custompreventlist.UpdatePreventListRequest{
		Name:        "Updated Prevent List",
		Description: "An updated prevent list",
		Type:        "FILEHASH",
	}

	result, _, err := service.UpdatePreventList(context.Background(), "test-id-1234", req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "test-id-1234", result.ID)
	assert.Equal(t, "Updated Prevent List", result.Name)
}

func TestPreventListService_DeletePreventList(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeletePreventListMock()

	_, err := service.DeletePreventList(context.Background(), "test-id-1234")

	require.NoError(t, err)
}

func TestPreventListService_ListPreventLists(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListPreventListsMock()

	result, _, err := service.ListPreventLists(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-id-1234", result[0].ID)
	assert.Equal(t, "Test Prevent List", result[0].Name)
}

func TestPreventListService_ListPreventListNames(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListPreventListNamesMock()

	result, _, err := service.ListPreventListNames(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, "Test Prevent List", result[0])
}

func TestPreventListService_ListPreventLists_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listPreventLists", 200, "list_prevent_lists_empty.json")

	result, _, err := service.ListPreventLists(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestPreventListService_ListPreventListNames_EmptyResult(t *testing.T) {
	service, mock := setupMockService(t)
	mock.Register("/graphql", "listPreventListNames", 200, "list_prevent_list_names_empty.json")

	result, _, err := service.ListPreventListNames(context.Background())

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}

func TestPreventListService_CreatePreventList_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "createPreventList", 500, "", "Internal Server Error")

	req := &custompreventlist.CreatePreventListRequest{
		Name: "test",
		Type: "FILEHASH",
	}

	result, _, err := service.CreatePreventList(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to create prevent list")
}

func TestPreventListService_GetPreventList_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "getPreventList", 500, "", "Internal Server Error")

	result, _, err := service.GetPreventList(context.Background(), "test-id")

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to get prevent list")
}

func TestPreventListService_UpdatePreventList_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "updatePreventList", 500, "", "Internal Server Error")

	req := &custompreventlist.UpdatePreventListRequest{
		Name: "test",
		Type: "FILEHASH",
	}

	result, _, err := service.UpdatePreventList(context.Background(), "test-id", req)

	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to update prevent list")
}

func TestPreventListService_DeletePreventList_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "deletePreventList", 500, "", "Internal Server Error")

	_, err := service.DeletePreventList(context.Background(), "test-id")

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to delete prevent list")
}

func TestPreventListService_ListPreventLists_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "listPreventLists", 500, "", "Internal Server Error")

	result, _, err := service.ListPreventLists(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestPreventListService_ListPreventListNames_Error(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterError("/app", "listPreventListNames", 500, "", "Internal Server Error")

	result, _, err := service.ListPreventListNames(context.Background())

	require.Error(t, err)
	assert.Nil(t, result)
}

func TestPreventListService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreatePreventList nil request",
			fn: func() error {
				_, _, err := service.CreatePreventList(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreatePreventList missing name",
			fn: func() error {
				_, _, err := service.CreatePreventList(context.Background(), &custompreventlist.CreatePreventListRequest{
					Type: "FILEHASH",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "GetPreventList empty id",
			fn: func() error {
				_, _, err := service.GetPreventList(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdatePreventList empty id",
			fn: func() error {
				_, _, err := service.UpdatePreventList(context.Background(), "", &custompreventlist.UpdatePreventListRequest{
					Name: "test",
					Type: "FILEHASH",
				})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdatePreventList nil request",
			fn: func() error {
				_, _, err := service.UpdatePreventList(context.Background(), "test-id", nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "UpdatePreventList missing name",
			fn: func() error {
				_, _, err := service.UpdatePreventList(context.Background(), "test-id", &custompreventlist.UpdatePreventListRequest{
					Type: "FILEHASH",
				})
				return err
			},
			wantErr: "name is required",
		},
		{
			name: "UpdatePreventList missing type",
			fn: func() error {
				_, _, err := service.UpdatePreventList(context.Background(), "test-id", &custompreventlist.UpdatePreventListRequest{
					Name: "test",
				})
				return err
			},
			wantErr: "type is required",
		},
		{
			name: "DeletePreventList empty id",
			fn: func() error {
				_, err := service.DeletePreventList(context.Background(), "")
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

func TestPreventListService_Validators(t *testing.T) {
	assert.NoError(t, custompreventlist.ValidatePreventListID("test-id"))
	assert.NoError(t, custompreventlist.ValidatePreventListID(""))

	assert.NoError(t, custompreventlist.ValidatePreventListType("TEAMID"))
	assert.NoError(t, custompreventlist.ValidatePreventListType("FILEHASH"))
	assert.NoError(t, custompreventlist.ValidatePreventListType("CDHASH"))
	assert.NoError(t, custompreventlist.ValidatePreventListType("SIGNINGID"))
	assert.Error(t, custompreventlist.ValidatePreventListType("INVALID"))

	assert.NoError(t, custompreventlist.ValidateCreatePreventListRequest(&custompreventlist.CreatePreventListRequest{Type: "TEAMID"}))
	assert.NoError(t, custompreventlist.ValidateCreatePreventListRequest(nil))
	assert.Error(t, custompreventlist.ValidateCreatePreventListRequest(&custompreventlist.CreatePreventListRequest{Type: "INVALID"}))

	assert.NoError(t, custompreventlist.ValidateUpdatePreventListRequest(&custompreventlist.UpdatePreventListRequest{Type: "FILEHASH"}))
	assert.NoError(t, custompreventlist.ValidateUpdatePreventListRequest(nil))
	assert.Error(t, custompreventlist.ValidateUpdatePreventListRequest(&custompreventlist.UpdatePreventListRequest{Type: "INVALID"}))
}
