package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	custompreventlist "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_PreventList_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("prevent_list")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating prevent list %s", name)
	req := &custompreventlist.CreatePreventListRequest{
		Name:        name,
		Description: "SDK acceptance test prevent list",
		Type:        "SHA256",
		Tags:        []string{},
		List:        []string{},
	}
	created, _, err := acc.Client.PreventList.CreatePreventList(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.ID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created prevent list ID=%s", created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.PreventList.DeletePreventList(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "PreventList", created.ID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing prevent lists")
	list, _, err := acc.Client.PreventList.ListPreventLists(ctx)
	require.NoError(t, err)
	found := false
	for _, pl := range list {
		if pl.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created prevent list should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting prevent list ID=%s", created.ID)
	got, _, err := acc.Client.PreventList.GetPreventList(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating prevent list ID=%s", created.ID)
	updatedName := acc.UniqueName("prevent_list_updated")
	updateReq := &custompreventlist.UpdatePreventListRequest{
		Name:        updatedName,
		Description: "Updated SDK acceptance test prevent list",
		Type:        "SHA256",
		Tags:        []string{},
		List:        []string{},
	}
	updated, _, err := acc.Client.PreventList.UpdatePreventList(ctx, created.ID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated prevent list ID=%s name=%s", created.ID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting prevent list ID=%s", created.ID)
	_, err = acc.Client.PreventList.DeletePreventList(ctx, created.ID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted prevent list ID=%s", created.ID)
}

func TestAcceptance_PreventList_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.PreventList.CreatePreventList(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty ID for get", func(t *testing.T) {
		_, _, err := acc.Client.PreventList.GetPreventList(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty ID for update", func(t *testing.T) {
		_, _, err := acc.Client.PreventList.UpdatePreventList(ctx, "", &custompreventlist.UpdatePreventListRequest{
			Name: "test",
			Type: "SHA256",
		})
		require.Error(t, err)
	})

	t.Run("empty ID for delete", func(t *testing.T) {
		_, err := acc.Client.PreventList.DeletePreventList(ctx, "")
		require.Error(t, err)
	})
}
