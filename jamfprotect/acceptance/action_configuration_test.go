package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	actionconfiguration "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// minimalAlertConfig returns a minimal valid alertConfig map for action configuration requests.
func minimalAlertConfig() map[string]any {
	return map[string]any{
		"data": map[string]any{},
	}
}

func TestAcceptance_ActionConfig_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("action_config")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating action configuration %s", name)
	req := &actionconfiguration.CreateActionConfigRequest{
		Name:        name,
		Description: "SDK acceptance test action configuration",
		AlertConfig: minimalAlertConfig(),
		Clients:     []map[string]any{},
	}
	created, _, err := acc.Client.ActionConfig.CreateActionConfig(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.ID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created action configuration ID=%s", created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.ActionConfig.DeleteActionConfig(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "ActionConfig", created.ID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing action configurations")
	list, _, err := acc.Client.ActionConfig.ListActionConfigs(ctx)
	require.NoError(t, err)
	found := false
	for _, ac := range list {
		if ac.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created action configuration should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting action configuration ID=%s", created.ID)
	got, _, err := acc.Client.ActionConfig.GetActionConfig(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating action configuration ID=%s", created.ID)
	updatedName := acc.UniqueName("action_config_updated")
	updateReq := &actionconfiguration.UpdateActionConfigRequest{
		Name:        updatedName,
		Description: "Updated SDK acceptance test action configuration",
		AlertConfig: minimalAlertConfig(),
		Clients:     []map[string]any{},
	}
	updated, _, err := acc.Client.ActionConfig.UpdateActionConfig(ctx, created.ID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated action configuration ID=%s name=%s", created.ID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting action configuration ID=%s", created.ID)
	_, err = acc.Client.ActionConfig.DeleteActionConfig(ctx, created.ID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted action configuration ID=%s", created.ID)
}

func TestAcceptance_ActionConfig_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.ActionConfig.CreateActionConfig(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty ID for get", func(t *testing.T) {
		_, _, err := acc.Client.ActionConfig.GetActionConfig(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty ID for update", func(t *testing.T) {
		_, _, err := acc.Client.ActionConfig.UpdateActionConfig(ctx, "", &actionconfiguration.UpdateActionConfigRequest{
			Name:        "test",
			AlertConfig: minimalAlertConfig(),
		})
		require.Error(t, err)
	})

	t.Run("empty ID for delete", func(t *testing.T) {
		_, err := acc.Client.ActionConfig.DeleteActionConfig(ctx, "")
		require.Error(t, err)
	})
}
