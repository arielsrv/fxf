package models_test

import (
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessage(t *testing.T) {
	t.Run("should create message with valid data", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := "test message"

		// Act
		message := &models.Message{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, message)
		assert.Equal(t, id, message.ID)
		assert.Equal(t, text, message.Text)
	})

	t.Run("should create message with empty text", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := ""

		// Act
		message := &models.Message{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, message)
		assert.Equal(t, id, message.ID)
		assert.Equal(t, text, message.Text)
	})

	t.Run("should create message with nil ID", func(t *testing.T) {
		// Arrange
		id := uuid.Nil
		text := "test message"

		// Act
		message := &models.Message{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, message)
		assert.Equal(t, id, message.ID)
		assert.Equal(t, text, message.Text)
	})

	t.Run("should create message with long text", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := "This is a very long message that contains multiple words and should be handled properly by the system"

		// Act
		message := &models.Message{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, message)
		assert.Equal(t, id, message.ID)
		assert.Equal(t, text, message.Text)
	})
}

func TestMessageStructure(t *testing.T) {
	t.Run("should have correct field types", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := "test message"

		// Act
		message := &models.Message{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, message)
		assert.IsType(t, uuid.UUID{}, message.ID)
		assert.IsType(t, "", message.Text)
	})
}
