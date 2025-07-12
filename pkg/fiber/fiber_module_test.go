package fiber_test

import (
	"testing"

	fiberpkg "github.com/arielsrv/fxf/pkg/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFiberServer(t *testing.T) {
	t.Run("should create fiber server with correct configuration", func(t *testing.T) {
		// Act
		app := fiberpkg.NewFiberServer()

		// Assert
		require.NotNil(t, app)
		assert.IsType(t, &fiber.App{}, app)

		// Check that the app is properly configured
		config := app.Config()
		assert.True(t, config.EnablePrintRoutes)
		assert.True(t, config.DisableStartupMessage)
	})
}
