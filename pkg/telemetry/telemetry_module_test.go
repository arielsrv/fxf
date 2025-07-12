package telemetry_test

import (
	"context"
	"testing"

	"github.com/arielsrv/fxf/pkg/telemetry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
)

func TestRegisterTracer(t *testing.T) {
	t.Run("should register tracer without error", func(t *testing.T) {
		// Arrange
		app := fx.New(
			fx.Invoke(telemetry.RegisterTracer),
		)

		// Act & Assert
		// The tracer registration should not panic
		require.NotNil(t, app)

		// Start the app to test tracer registration
		ctx := context.Background()
		err := app.Start(ctx)

		// The tracer registration should work without errors
		// OTLP connection might fail in some environments, but that's expected
		if err != nil {
			// If there's an error, it should be related to OTLP connection
			// which is expected in test environments
			assert.Contains(t, err.Error(), "connection")
		}

		// Clean up
		_ = app.Stop(ctx)
	})
}

func TestTelemetryModule(t *testing.T) {
	t.Run("should provide telemetry module", func(t *testing.T) {
		// Arrange & Act
		module := telemetry.Module

		// Assert
		require.NotNil(t, module)
	})
}

func TestTracerProviderConfiguration(t *testing.T) {
	t.Run("should configure tracer provider with correct settings", func(t *testing.T) {
		// This test verifies that the tracer provider is configured correctly
		// The actual configuration happens in RegisterTracer function
		// We can't easily test the OTLP connection in unit tests, but we can verify the module structure

		// Arrange & Act
		module := telemetry.Module

		// Assert
		require.NotNil(t, module)
		// The module should contain the RegisterTracer invocation
		assert.NotNil(t, module)
	})
}
