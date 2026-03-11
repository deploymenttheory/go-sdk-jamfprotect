package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Analytic_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("analytic")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating analytic %s", name)
	req := &analytic.CreateAnalyticRequest{
		Name:            name,
		InputType:       "GPFSEvent",
		Filter:          "process.name == 'test'",
		Description:     "SDK acceptance test analytic",
		Level:           3,
		Severity:        "MEDIUM",
		Tags:            []string{},
		Categories:      []string{},
		AnalyticActions: []analytic.AnalyticActionInput{},
		Context:         []analytic.AnalyticContextInput{},
		SnapshotFiles:   []string{},
	}
	created, _, err := acc.Client.Analytic.CreateAnalytic(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.UUID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created analytic UUID=%s", created.UUID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.Analytic.DeleteAnalytic(ctx, created.UUID)
		acc.LogCleanupDeleteError(t, "Analytic", created.UUID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing analytics")
	list, _, err := acc.Client.Analytic.ListAnalytics(ctx)
	require.NoError(t, err)
	found := false
	for _, a := range list {
		if a.UUID == created.UUID {
			found = true
			break
		}
	}
	assert.True(t, found, "created analytic should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting analytic UUID=%s", created.UUID)
	got, _, err := acc.Client.Analytic.GetAnalytic(ctx, created.UUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.UUID, got.UUID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating analytic UUID=%s", created.UUID)
	updatedName := acc.UniqueName("analytic_updated")
	updateReq := &analytic.UpdateAnalyticRequest{
		Name:            updatedName,
		InputType:       "GPFSEvent",
		Filter:          "process.name == 'updated'",
		Description:     "Updated SDK acceptance test analytic",
		Level:           4,
		Tags:            []string{},
		Categories:      []string{},
		AnalyticActions: []analytic.AnalyticActionInput{},
		Context:         []analytic.AnalyticContextInput{},
		SnapshotFiles:   []string{},
	}
	updated, _, err := acc.Client.Analytic.UpdateAnalytic(ctx, created.UUID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated analytic UUID=%s name=%s", created.UUID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting analytic UUID=%s", created.UUID)
	_, err = acc.Client.Analytic.DeleteAnalytic(ctx, created.UUID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted analytic UUID=%s", created.UUID)
}

func TestAcceptance_Analytic_list_lite(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	acc.LogTestStage(t, "List", "Listing analytics lite")
	list, _, err := acc.Client.Analytic.ListAnalyticsLite(ctx)
	require.NoError(t, err)
	assert.NotNil(t, list)
	acc.LogTestSuccess(t, "Listed %d analytics (lite)", len(list))
}

func TestAcceptance_Analytic_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.Analytic.CreateAnalytic(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty UUID for get", func(t *testing.T) {
		_, _, err := acc.Client.Analytic.GetAnalytic(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty UUID for update", func(t *testing.T) {
		_, _, err := acc.Client.Analytic.UpdateAnalytic(ctx, "", &analytic.UpdateAnalyticRequest{
			Name:      "test",
			InputType: "GPFSEvent",
			Filter:    "process.name == 'test'",
		})
		require.Error(t, err)
	})

	t.Run("empty UUID for delete", func(t *testing.T) {
		_, err := acc.Client.Analytic.DeleteAnalytic(ctx, "")
		require.Error(t, err)
	})
}
