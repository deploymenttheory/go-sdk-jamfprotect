package acceptance_test

import (
	"testing"

	acc "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/acceptance"
	exceptionset "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_ExceptionSet_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	name := acc.UniqueName("exception_set")

	// --- Create ---
	acc.LogTestStage(t, "Create", "Creating exception set %s", name)
	req := &exceptionset.CreateExceptionSetRequest{
		Name:         name,
		Description:  "SDK acceptance test exception set",
		Exceptions:   []exceptionset.ExceptionInput{},
		EsExceptions: []exceptionset.EsExceptionInput{},
	}
	created, _, err := acc.Client.ExceptionSet.CreateExceptionSet(ctx, req)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.NotEmpty(t, created.UUID)
	assert.Equal(t, name, created.Name)
	acc.LogTestSuccess(t, "Created exception set UUID=%s", created.UUID)

	acc.Cleanup(t, func() {
		ctx, cancel := acc.NewContext()
		defer cancel()
		_, err := acc.Client.ExceptionSet.DeleteExceptionSet(ctx, created.UUID)
		acc.LogCleanupDeleteError(t, "ExceptionSet", created.UUID, err)
	})

	// --- List ---
	acc.LogTestStage(t, "List", "Listing exception sets")
	list, _, err := acc.Client.ExceptionSet.ListExceptionSets(ctx)
	require.NoError(t, err)
	found := false
	for _, es := range list {
		if es.UUID == created.UUID {
			found = true
			break
		}
	}
	assert.True(t, found, "created exception set should appear in list")

	// --- Get ---
	acc.LogTestStage(t, "Get", "Getting exception set UUID=%s", created.UUID)
	got, _, err := acc.Client.ExceptionSet.GetExceptionSet(ctx, created.UUID)
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, created.UUID, got.UUID)
	assert.Equal(t, name, got.Name)

	// --- Update ---
	acc.LogTestStage(t, "Update", "Updating exception set UUID=%s", created.UUID)
	updatedName := acc.UniqueName("exception_set_updated")
	updateReq := &exceptionset.UpdateExceptionSetRequest{
		Name:         updatedName,
		Description:  "Updated SDK acceptance test exception set",
		Exceptions:   []exceptionset.ExceptionInput{},
		EsExceptions: []exceptionset.EsExceptionInput{},
	}
	updated, _, err := acc.Client.ExceptionSet.UpdateExceptionSet(ctx, created.UUID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, updatedName, updated.Name)
	acc.LogTestSuccess(t, "Updated exception set UUID=%s name=%s", created.UUID, updatedName)

	// --- Delete ---
	acc.LogTestStage(t, "Delete", "Deleting exception set UUID=%s", created.UUID)
	_, err = acc.Client.ExceptionSet.DeleteExceptionSet(ctx, created.UUID)
	require.NoError(t, err)
	acc.LogTestSuccess(t, "Deleted exception set UUID=%s", created.UUID)
}

func TestAcceptance_ExceptionSet_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	ctx, cancel := acc.NewContext()
	defer cancel()

	t.Run("nil request", func(t *testing.T) {
		_, _, err := acc.Client.ExceptionSet.CreateExceptionSet(ctx, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "request cannot be nil")
	})

	t.Run("empty UUID for get", func(t *testing.T) {
		_, _, err := acc.Client.ExceptionSet.GetExceptionSet(ctx, "")
		require.Error(t, err)
	})

	t.Run("empty UUID for update", func(t *testing.T) {
		_, _, err := acc.Client.ExceptionSet.UpdateExceptionSet(ctx, "", &exceptionset.UpdateExceptionSetRequest{
			Name: "test",
		})
		require.Error(t, err)
	})

	t.Run("empty UUID for delete", func(t *testing.T) {
		_, err := acc.Client.ExceptionSet.DeleteExceptionSet(ctx, "")
		require.Error(t, err)
	})
}
