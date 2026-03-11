package downloads_test

import (
	"context"
	"testing"

	downloads "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/downloads"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/downloads/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*downloads.Service, *mocks.DownloadsMock) {
	t.Helper()
	mock := mocks.NewDownloadsMock()
	return downloads.NewService(mock), mock
}

func TestDownloadsService_GetOrganizationDownloads(t *testing.T) {
	service, mock := setupMockService(t)
	mock.RegisterGetOrganizationDownloadsMock()

	result, _, err := service.GetOrganizationDownloads(context.Background())

	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "pppc-content", result.PPPC)
	assert.Equal(t, "rootca-content", result.RootCA)
	assert.Equal(t, "csr-content", result.CSR)
	assert.Equal(t, "installer-uuid-1234", result.InstallerUUID)
	require.NotNil(t, result.VanillaPackage)
	assert.Equal(t, "5.0.0", result.VanillaPackage.Version)
	assert.Equal(t, "ws-auth-token", result.WebsocketAuth)
	assert.Equal(t, "tamper-profile-content", result.TamperPreventionProfile)
}
