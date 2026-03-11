package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	analyticset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_AnalyticSet_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("analytic_set")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating analytic set %s", name)
	req := &analyticset.CreateAnalyticSetRequest{
		Name:        name,
		Description: "SDK acceptance test analytic set",
		Types:       []string{},
		Analytics:   []string{},
	}
	created, _, err := acc.Client.AnalyticSet.CreateAnalyticSet(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.UUID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created analytic set UUID=%s", created.UUID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.AnalyticSet.DeleteAnalyticSet(ctx, created.UUID)
		acc.LogCleanupDeleteError(t, "AnalyticSet", created.UUID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing analytic sets")
	list, _, err := acc.Client.AnalyticSet.ListAnalyticSets(ctx)
	require.NoError(t, err)
	found := false
	for _, s := range list {
		if s.UUID == created.UUID {
			found = true
			break
		}
	}
	assert.True(t, found, "created analytic set should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting analytic set UUID=%s", created.UUID)
	got, _, err := acc.Client.AnalyticSet.GetAnalyticSet(ctx, created.UUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.UUID, got.UUID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating analytic set UUID=%s", created.UUID)
	updatedName := acc.UniqueName("analytic_set_updated")
	updateReq := &analyticset.UpdateAnalyticSetRequest{
		Name:        updatedName,
		Description: "Updated SDK acceptance test analytic set",
		Types:       []string{},
		Analytics:   []string{},
	}
	updated, _, err := acc.Client.AnalyticSet.UpdateAnalyticSet(ctx, created.UUID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated analytic set UUID=%s name=%s", created.UUID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting analytic set UUID=%s", created.UUID)
	_, err = acc.Client.AnalyticSet.DeleteAnalyticSet(ctx, created.UUID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted analytic set UUID=%s", created.UUID)
}

func TestAcceptance_AnalyticSet_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.AnalyticSet.CreateAnalyticSet(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty UUID for get", func(t *testing.T) {
		_, _, err := acc.Client.AnalyticSet.GetAnalyticSet(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty UUID for update", func(t *testing.T) {
		_, _, err := acc.Client.AnalyticSet.UpdateAnalyticSet(ctx, "", &analyticset.UpdateAnalyticSetRequest{
			Name:      "test",
			Analytics: []string{},
		})
		require.Error(t, err)
	})

	t.Run("empty UUID for delete", func(t *testing.T) {
		_, err := acc.Client.AnalyticSet.DeleteAnalyticSet(ctx, "")
		require.Error(t, err)
	})
}
