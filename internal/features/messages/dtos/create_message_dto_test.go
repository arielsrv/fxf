package dtos_test

import (
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateMessageCommand(t *testing.T) {
	t.Run("should create command with valid text", func(t *testing.T) {
		// Arrange
		text := "test message"

		// Act
		cmd := &dtos.CreateMessageCommand{
			Text: text,
		}

		// Assert
		require.NotNil(t, cmd)
		assert.Equal(t, text, cmd.Text)
	})

	t.Run("should create command with empty text", func(t *testing.T) {
		// Arrange
		text := ""

		// Act
		cmd := &dtos.CreateMessageCommand{
			Text: text,
		}

		// Assert
		require.NotNil(t, cmd)
		assert.Equal(t, text, cmd.Text)
	})

	t.Run("should create command with long text", func(t *testing.T) {
		// Arrange
		text := "This is a very long message that contains multiple words and should be handled properly by the system"

		// Act
		cmd := &dtos.CreateMessageCommand{
			Text: text,
		}

		// Assert
		require.NotNil(t, cmd)
		assert.Equal(t, text, cmd.Text)
	})
}

func TestCreateMessageCommandResponse(t *testing.T) {
	t.Run("should create response with valid ID", func(t *testing.T) {
		// Arrange
		id := uuid.New()

		// Act
		resp := &dtos.CreateMessageCommandResponse{
			ID: id,
		}

		// Assert
		require.NotNil(t, resp)
		assert.Equal(t, id, resp.ID)
	})

	t.Run("should create response with nil ID", func(t *testing.T) {
		// Arrange
		id := uuid.Nil

		// Act
		resp := &dtos.CreateMessageCommandResponse{
			ID: id,
		}

		// Assert
		require.NotNil(t, resp)
		assert.Equal(t, id, resp.ID)
	})
}
