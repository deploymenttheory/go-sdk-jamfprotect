package user_test

import (
	"context"
	"testing"

	user "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/user"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testUserID    = "user-id-1234"
	testUserEmail = "test@example.com"
)

func setupMockService(t *testing.T) (*user.Service, *mocks.UserMock) {
	t.Helper()
	mock := mocks.NewUserMock()
	return user.NewService(mock), mock
}

func TestUserService_CreateUser(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterCreateUserMock()

	req := &user.CreateUserRequest{
		Email:                 testUserEmail,
		RoleIDs:               []string{},
		GroupIDs:              []string{},
		ReceiveEmailAlert:     false,
		EmailAlertMinSeverity: "LOW",
	}

	result, _, err := service.CreateUser(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUserID, result.ID)
	assert.Equal(t, testUserEmail, result.Email)
	assert.Equal(t, "LOW", result.EmailAlertMinSeverity)
}

func TestUserService_GetUser(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetUserMock()

	result, _, err := service.GetUser(context.Background(), testUserID)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUserID, result.ID)
	assert.Equal(t, testUserEmail, result.Email)
}

func TestUserService_UpdateUser(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateUserMock()

	req := &user.UpdateUserRequest{
		RoleIDs:               []string{},
		GroupIDs:              []string{},
		ReceiveEmailAlert:     true,
		EmailAlertMinSeverity: "HIGH",
	}

	result, _, err := service.UpdateUser(context.Background(), testUserID, req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, testUserID, result.ID)
	assert.Equal(t, testUserEmail, result.Email)
	assert.True(t, result.ReceiveEmailAlert)
	assert.Equal(t, "HIGH", result.EmailAlertMinSeverity)
}

func TestUserService_DeleteUser(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterDeleteUserMock()

	_, err := service.DeleteUser(context.Background(), testUserID)

	require.NoError(t, err)
}

func TestUserService_ListUsers(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterListUsersMock()

	result, _, err := service.ListUsers(context.Background())

	require.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, testUserID, result[0].ID)
	assert.Equal(t, testUserEmail, result[0].Email)
}

func TestUserService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "CreateUser nil request",
			fn: func() error {
				_, _, err := service.CreateUser(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "CreateUser missing email",
			fn: func() error {
				_, _, err := service.CreateUser(context.Background(), &user.CreateUserRequest{
					EmailAlertMinSeverity: "LOW",
				})
				return err
			},
			wantErr: "email is required",
		},
		{
			name: "CreateUser missing emailAlertMinSeverity",
			fn: func() error {
				_, _, err := service.CreateUser(context.Background(), &user.CreateUserRequest{
					Email: testUserEmail,
				})
				return err
			},
			wantErr: "emailAlertMinSeverity is required",
		},
		{
			name: "GetUser empty id",
			fn: func() error {
				_, _, err := service.GetUser(context.Background(), "")
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateUser empty id",
			fn: func() error {
				_, _, err := service.UpdateUser(context.Background(), "", &user.UpdateUserRequest{})
				return err
			},
			wantErr: "id is required",
		},
		{
			name: "UpdateUser nil request",
			fn: func() error {
				_, _, err := service.UpdateUser(context.Background(), testUserID, nil)
				return err
			},
			wantErr: "request cannot be nil",
		},
		{
			name: "DeleteUser empty id",
			fn: func() error {
				_, err := service.DeleteUser(context.Background(), "")
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
