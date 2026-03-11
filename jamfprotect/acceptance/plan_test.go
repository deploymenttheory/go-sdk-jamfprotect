package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	actionconfiguration "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createTestActionConfig creates a temporary action configuration for use in plan tests.
// The returned ID must be cleaned up by the caller via acc.Cleanup.
func createTestActionConfig(t *testing.T) string {
	t.Helper()

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("action_config_for_plan")
	req := &actionconfiguration.CreateActionConfigRequest{
		Name:        name,
		Description: "Ephemeral action config for plan acceptance test",
		AlertConfig: map[string]any{
			"data": map[string]any{},
		},
		Clients: []map[string]any{},
	}
	created, _, err := acc.Client.ActionConfig.CreateActionConfig(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.ActionConfig.DeleteActionConfig(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "ActionConfig (for plan)", created.ID, err)
	})

	return created.ID
}

func TestAcceptance_Plan_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	// Plans require an action configuration — create one for this test.
	actionConfigID := createTestActionConfig(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("plan")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating plan %s", name)
	req := &plan.CreatePlanRequest{
		Name:          name,
		Description:   "SDK acceptance test plan",
		ActionConfigs: actionConfigID,
		ExceptionSets: []string{},
		AnalyticSets:  []plan.AnalyticSetInput{},
		CommsConfig: plan.CommsConfigInput{
			Protocol: "mqtt",
			FQDN:     "",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{},
			InsightsSyncInterval: 0,
		},
		AutoUpdate: false,
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "",
		},
	}
	created, _, err := acc.Client.Plan.CreatePlan(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.ID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created plan ID=%s", created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.Plan.DeletePlan(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "Plan", created.ID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing plans")
	list, _, err := acc.Client.Plan.ListPlans(ctx)
	require.NoError(t, err)
	found := false
	for _, p := range list {
		if p.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created plan should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting plan ID=%s", created.ID)
	got, _, err := acc.Client.Plan.GetPlan(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating plan ID=%s", created.ID)
	updatedName := acc.UniqueName("plan_updated")
	updateReq := &plan.UpdatePlanRequest{
		Name:          updatedName,
		Description:   "Updated SDK acceptance test plan",
		ActionConfigs: actionConfigID,
		ExceptionSets: []string{},
		AnalyticSets:  []plan.AnalyticSetInput{},
		CommsConfig: plan.CommsConfigInput{
			Protocol: "mqtt",
			FQDN:     "",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{},
			InsightsSyncInterval: 0,
		},
		AutoUpdate: false,
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "",
		},
	}
	updated, _, err := acc.Client.Plan.UpdatePlan(ctx, created.ID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated plan ID=%s name=%s", created.ID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting plan ID=%s", created.ID)
	_, err = acc.Client.Plan.DeletePlan(ctx, created.ID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted plan ID=%s", created.ID)
}

func TestAcceptance_Plan_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.Plan.CreatePlan(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty ID for get", func(t *testing.T) {
		_, _, err := acc.Client.Plan.GetPlan(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty ID for delete", func(t *testing.T) {
		_, err := acc.Client.Plan.DeletePlan(ctx, "")
		require.Error(t, err)
	})
}
