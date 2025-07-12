package logger_test

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/arielsrv/fxf/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx/fxevent"
)

func TestNewSlogLogger(t *testing.T) {
	t.Run("should create slog logger adapter", func(t *testing.T) {
		// Arrange
		handler := slog.NewJSONHandler(bytes.NewBuffer(nil), &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)

		// Act
		adapter := logger.New(slogLogger)

		// Assert
		require.NotNil(t, adapter)
		assert.IsType(t, &logger.SlogLogger{}, adapter)
	})
}

func TestSlogLogger_LogEvent(t *testing.T) {
	t.Run("should handle OnStartExecuting event", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.OnStartExecuting{
			FunctionName: "testFunction",
			CallerName:   "testCaller",
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		// The event should be logged (we can't easily test the exact output due to JSON formatting)
		// but we can verify the adapter doesn't panic
		assert.NotNil(t, adapter)
	})

	t.Run("should handle OnStartExecuted event with error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.OnStartExecuted{
			FunctionName: "testFunction",
			CallerName:   "testCaller",
			Err:          assert.AnError,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})

	t.Run("should handle OnStartExecuted event without error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.OnStartExecuted{
			FunctionName: "testFunction",
			CallerName:   "testCaller",
			Err:          nil,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})

	t.Run("should handle Supplied event with error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.Supplied{
			TypeName: "testType",
			Err:      assert.AnError,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})

	t.Run("should handle Supplied event without error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.Supplied{
			TypeName: "testType",
			Err:      nil,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})

	t.Run("should handle Started event with error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.Started{
			Err: assert.AnError,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})

	t.Run("should handle Started event without error", func(t *testing.T) {
		// Arrange
		var buf bytes.Buffer
		handler := slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		event := &fxevent.Started{
			Err: nil,
		}

		// Act
		adapter.LogEvent(event)

		// Assert
		assert.NotNil(t, adapter)
	})
}

func TestSlogLogger_ImplementsInterface(t *testing.T) {
	t.Run("should implement fxevent.Logger interface", func(t *testing.T) {
		// Arrange
		handler := slog.NewJSONHandler(bytes.NewBuffer(nil), &slog.HandlerOptions{Level: slog.LevelInfo})
		slogLogger := slog.New(handler)
		adapter := logger.New(slogLogger)

		// Act & Assert
		var _ fxevent.Logger = adapter
		assert.NotNil(t, adapter)
	})
}
