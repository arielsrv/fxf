package dtos_test

import (
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetMessageByIDQuery(t *testing.T) {
	t.Run("should create query with valid ID", func(t *testing.T) {
		// Arrange
		id := uuid.New()

		// Act
		query := &dtos.GetMessageByIDQuery{
			ID: id,
		}

		// Assert
		require.NotNil(t, query)
		assert.Equal(t, id, query.ID)
	})

	t.Run("should create query with nil ID", func(t *testing.T) {
		// Arrange
		id := uuid.Nil

		// Act
		query := &dtos.GetMessageByIDQuery{
			ID: id,
		}

		// Assert
		require.NotNil(t, query)
		assert.Equal(t, id, query.ID)
	})
}

func TestGetMessageByIDQueryResponse(t *testing.T) {
	t.Run("should create response with valid data", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := "test message"

		// Act
		resp := &dtos.GetMessageByIDQueryResponse{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, resp)
		assert.Equal(t, id, resp.ID)
		assert.Equal(t, text, resp.Text)
	})

	t.Run("should create response with empty text", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := ""

		// Act
		resp := &dtos.GetMessageByIDQueryResponse{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, resp)
		assert.Equal(t, id, resp.ID)
		assert.Equal(t, text, resp.Text)
	})

	t.Run("should create response with long text", func(t *testing.T) {
		// Arrange
		id := uuid.New()
		text := "This is a very long message that contains multiple words and should be handled properly by the system"

		// Act
		resp := &dtos.GetMessageByIDQueryResponse{
			ID:   id,
			Text: text,
		}

		// Assert
		require.NotNil(t, resp)
		assert.Equal(t, id, resp.ID)
		assert.Equal(t, text, resp.Text)
	})
}
