package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	removablestoragecontrolset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/removable_storage_control_set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_USBControlSet_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("usb_control_set")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating USB control set %s", name)
	req := &removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:                 name,
		Description:          "SDK acceptance test USB control set",
		DefaultMountAction:   "ALLOW",
		DefaultMessageAction: "",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}
	created, _, err := acc.Client.USBControlSet.CreateUSBControlSet(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.ID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created USB control set ID=%s", created.ID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.USBControlSet.DeleteUSBControlSet(ctx, created.ID)
		acc.LogCleanupDeleteError(t, "USBControlSet", created.ID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing USB control sets")
	list, _, err := acc.Client.USBControlSet.ListUSBControlSets(ctx)
	require.NoError(t, err)
	found := false
	for _, u := range list {
		if u.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "created USB control set should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting USB control set ID=%s", created.ID)
	got, _, err := acc.Client.USBControlSet.GetUSBControlSet(ctx, created.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.ID, got.ID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating USB control set ID=%s", created.ID)
	updatedName := acc.UniqueName("usb_control_set_updated")
	updateReq := &removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:                 updatedName,
		Description:          "Updated SDK acceptance test USB control set",
		DefaultMountAction:   "ALLOW",
		DefaultMessageAction: "",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}
	updated, _, err := acc.Client.USBControlSet.UpdateUSBControlSet(ctx, created.ID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated USB control set ID=%s name=%s", created.ID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting USB control set ID=%s", created.ID)
	_, err = acc.Client.USBControlSet.DeleteUSBControlSet(ctx, created.ID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted USB control set ID=%s", created.ID)
}

func TestAcceptance_USBControlSet_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.USBControlSet.CreateUSBControlSet(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty ID for get", func(t *testing.T) {
		_, _, err := acc.Client.USBControlSet.GetUSBControlSet(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty ID for update", func(t *testing.T) {
		_, _, err := acc.Client.USBControlSet.UpdateUSBControlSet(ctx, "", &removablestoragecontrolset.UpdateUSBControlSetRequest{
			Name:               "test",
			DefaultMountAction: "ALLOW",
			Rules:              []removablestoragecontrolset.USBControlRuleInput{},
		})
		require.Error(t, err)
	})

	t.Run("empty ID for delete", func(t *testing.T) {
		_, err := acc.Client.USBControlSet.DeleteUSBControlSet(ctx, "")
		require.Error(t, err)
	})
}
