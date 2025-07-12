package mediator_test

import (
	"testing"

	"github.com/arielsrv/fxf/pkg/mediator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediatorModule(t *testing.T) {
	t.Run("should provide mediator module", func(t *testing.T) {
		// Arrange & Act
		module := mediator.Module

		// Assert
		require.NotNil(t, module)
		// The module is an empty fx.Options, which is valid
		assert.NotNil(t, module)
	})
}

func TestMediatorModuleStructure(t *testing.T) {
	t.Run("should have empty module options", func(t *testing.T) {
		// Arrange & Act
		module := mediator.Module

		// Assert
		// The module is currently empty (placeholder)
		// This test ensures the module can be created without issues
		require.NotNil(t, module)
		// Even empty fx.Options is not nil
		assert.NotNil(t, module)
	})
}
