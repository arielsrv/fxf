package repository_test

import (
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/repository"

	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInMemoryMessageRepository(t *testing.T) {
	ctx := t.Context()
	repo := repository.NewInMemoryMessageRepository()

	t.Run("should create and get a message successfully", func(t *testing.T) {
		// Arrange
		createMsg := &models.Message{
			Text: "hello world",
		}

		// Act
		createdMsg, err := repo.CreateMessage(ctx, createMsg)
		require.NoError(t, err)
		require.NotNil(t, createdMsg)
		assert.NotEqual(t, uuid.Nil, createdMsg.ID)

		retrievedMsg, err := repo.GetMessageByID(ctx, createdMsg.ID)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, retrievedMsg)
		assert.Equal(t, createdMsg.ID, retrievedMsg.ID)
		assert.Equal(t, "hello world", retrievedMsg.Text)
	})

	t.Run("should return an error when message not found", func(t *testing.T) {
		// Arrange
		nonExistentID := uuid.New()

		// Act
		retrievedMsg, err := repo.GetMessageByID(ctx, nonExistentID)

		// Assert
		require.Error(t, err)
		require.Nil(t, retrievedMsg)
		assert.Contains(t, err.Error(), "not found")
	})
}
