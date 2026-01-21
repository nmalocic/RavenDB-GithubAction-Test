package ravendbtest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTTLSetup(t *testing.T) {
	t.Run("Should initialize store with TTL settings", func(t *testing.T) {
		wrapper := RavenDB_Wrapper{}
		err := wrapper.Init()
		require.NoError(t, err)
		require.NotNil(t, wrapper.documentStore)
	})
}
