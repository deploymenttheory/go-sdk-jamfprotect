package dataforwarding_test

import (
	"context"
	"testing"

	dataforwarding "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_forwarding"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/data_forwarding/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*dataforwarding.Service, *mocks.DataForwardingMock) {
	t.Helper()
	mock := mocks.NewDataForwardingMock()
	return dataforwarding.NewService(mock), mock
}

func TestDataForwardingService_GetDataForwarding(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetDataForwardingMock()

	result, _, err := service.GetDataForwarding(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "org-uuid-1234", result.UUID)
	assert.True(t, result.Forward.S3.Enabled)
	assert.Equal(t, "my-bucket", result.Forward.S3.Bucket)
}

func TestDataForwardingService_UpdateDataForwarding(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterUpdateDataForwardingMock()

	req := &dataforwarding.UpdateDataForwardingRequest{
		S3: dataforwarding.ForwardS3Input{
			Bucket:  "updated-bucket",
			Enabled: true,
		},
	}

	result, _, err := service.UpdateDataForwarding(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "updated-bucket", result.Forward.S3.Bucket)
}

func TestDataForwardingService_ValidationErrors(t *testing.T) {
	service, _ := setupMockService(t)

	tests := []struct {
		name    string
		fn      func() error
		wantErr string
	}{
		{
			name: "UpdateDataForwarding nil request",
			fn: func() error {
				_, _, err := service.UpdateDataForwarding(context.Background(), nil)
				return err
			},
			wantErr: "request cannot be nil",
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
