package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	unifiedloggingfilter "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_UnifiedLoggingFilter_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("ulfilter")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating unified logging filter %s", name)
	req := &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
		Name:        name,
		Description: "SDK acceptance test unified logging filter",
		Tags:        []string{},
		Filter:      "subsystem == 'com.example.test'",
		Enabled:     true,
	}
	created, _, err := acc.Client.UnifiedLoggingFilter.CreateUnifiedLoggingFilter(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.UUID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created unified logging filter UUID=%s", created.UUID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.UnifiedLoggingFilter.DeleteUnifiedLoggingFilter(ctx, created.UUID)
		acc.LogCleanupDeleteError(t, "UnifiedLoggingFilter", created.UUID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing unified logging filters")
	list, _, err := acc.Client.UnifiedLoggingFilter.ListUnifiedLoggingFilters(ctx)
	require.NoError(t, err)
	found := false
	for _, f := range list {
		if f.UUID == created.UUID {
			found = true
			break
		}
	}
	assert.True(t, found, "created unified logging filter should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting unified logging filter UUID=%s", created.UUID)
	got, _, err := acc.Client.UnifiedLoggingFilter.GetUnifiedLoggingFilter(ctx, created.UUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.UUID, got.UUID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating unified logging filter UUID=%s", created.UUID)
	updatedName := acc.UniqueName("ulfilter_updated")
	updateReq := &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
		Name:        updatedName,
		Description: "Updated SDK acceptance test unified logging filter",
		Tags:        []string{},
		Filter:      "subsystem == 'com.example.updated'",
		Enabled:     false,
	}
	updated, _, err := acc.Client.UnifiedLoggingFilter.UpdateUnifiedLoggingFilter(ctx, created.UUID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated unified logging filter UUID=%s name=%s", created.UUID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting unified logging filter UUID=%s", created.UUID)
	_, err = acc.Client.UnifiedLoggingFilter.DeleteUnifiedLoggingFilter(ctx, created.UUID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted unified logging filter UUID=%s", created.UUID)
}

func TestAcceptance_UnifiedLoggingFilter_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.UnifiedLoggingFilter.CreateUnifiedLoggingFilter(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty UUID for get", func(t *testing.T) {
		_, _, err := acc.Client.UnifiedLoggingFilter.GetUnifiedLoggingFilter(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty UUID for update", func(t *testing.T) {
		_, _, err := acc.Client.UnifiedLoggingFilter.UpdateUnifiedLoggingFilter(ctx, "", &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
			Name:   "test",
			Filter: "subsystem == 'com.example.test'",
		})
		require.Error(t, err)
	})

	t.Run("empty UUID for delete", func(t *testing.T) {
		_, err := acc.Client.UnifiedLoggingFilter.DeleteUnifiedLoggingFilter(ctx, "")
		require.Error(t, err)
	})
}
