package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_TelemetryV2_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("telemetry_v2")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating telemetry v2 %s", name)
	req := &telemetry.CreateTelemetryV2Request{
		Name:               name,
		Description:        "SDK acceptance test telemetry v2",
		LogFiles:           []string{},
		LogFileCollection:  false,
		PerformanceMetrics: false,
		Events:             []string{},
		FileHashing:        false,
	}
	created, _, err := acc.Client.TelemetryV2.CreateTelemetryV2(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.ID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created telemetry v2 ID=%s", created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.TelemetryV2.DeleteTelemetryV2(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "TelemetryV2", created.ID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing telemetries v2")
	list, _, err := acc.Client.TelemetryV2.ListTelemetriesV2(ctx)
	require.NoError(t, err)
	found := false
	for _, tv := range list {
		if tv.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created telemetry v2 should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting telemetry v2 ID=%s", created.ID)
	got, _, err := acc.Client.TelemetryV2.GetTelemetryV2(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating telemetry v2 ID=%s", created.ID)
	updatedName := acc.UniqueName("telemetry_v2_updated")
	updateReq := &telemetry.UpdateTelemetryV2Request{
		Name:               updatedName,
		Description:        "Updated SDK acceptance test telemetry v2",
		LogFiles:           []string{},
		LogFileCollection:  false,
		PerformanceMetrics: true,
		Events:             []string{},
		FileHashing:        false,
	}
	updated, _, err := acc.Client.TelemetryV2.UpdateTelemetryV2(ctx, created.ID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated telemetry v2 ID=%s name=%s", created.ID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting telemetry v2 ID=%s", created.ID)
	_, err = acc.Client.TelemetryV2.DeleteTelemetryV2(ctx, created.ID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted telemetry v2 ID=%s", created.ID)
}

func TestAcceptance_TelemetryV2_list_combined(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	acc.LogTestStage(t, "List", "Listing combined telemetries (v1 and v2)")
	combined, _, err := acc.Client.TelemetryV2.ListTelemetriesCombined(ctx, false)
	require.NoError(t, err)
	require.NotNil(t, combined)
	acc.LogTestSuccess(t, "Listed combined telemetries: %d v1, %d v2", len(combined.Telemetries), len(combined.TelemetriesV2))
}

func TestAcceptance_TelemetryV2_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.TelemetryV2.CreateTelemetryV2(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty ID for get", func(t *testing.T) {
		_, _, err := acc.Client.TelemetryV2.GetTelemetryV2(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty ID for update", func(t *testing.T) {
		_, _, err := acc.Client.TelemetryV2.UpdateTelemetryV2(ctx, "", &telemetry.UpdateTelemetryV2Request{
			Name:     "test",
			LogFiles: []string{},
		})
		require.Error(t, err)
	})

	t.Run("empty ID for delete", func(t *testing.T) {
		_, err := acc.Client.TelemetryV2.DeleteTelemetryV2(ctx, "")
		require.Error(t, err)
	})
}
