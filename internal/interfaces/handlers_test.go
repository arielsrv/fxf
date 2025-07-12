package interfaces_test

import (
	"context"
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/interfaces"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mock implementations for testing
type MockCreateMessageCommandHandler struct{}

func (m *MockCreateMessageCommandHandler) Handle(ctx context.Context, cmd *dtos.CreateMessageCommand) (*dtos.CreateMessageCommandResponse, error) {
	return &dtos.CreateMessageCommandResponse{ID: uuid.New()}, nil
}

type MockGetMessageByIDQueryHandler struct{}

func (m *MockGetMessageByIDQueryHandler) Handle(ctx context.Context, query *dtos.GetMessageByIDQuery) (*dtos.GetMessageByIDQueryResponse, error) {
	return &dtos.GetMessageByIDQueryResponse{ID: query.ID, Text: "test"}, nil
}

func TestICreateMessageCommandHandler(t *testing.T) {
	t.Run("should implement interface correctly", func(t *testing.T) {
		// Arrange
		handler := &MockCreateMessageCommandHandler{}

		// Act & Assert
		var _ interfaces.ICreateMessageCommandHandler = handler
		require.NotNil(t, handler)
	})

	t.Run("should handle create message command", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		handler := &MockCreateMessageCommandHandler{}
		cmd := &dtos.CreateMessageCommand{Text: "test"}

		// Act
		result, err := handler.Handle(ctx, cmd)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEqual(t, uuid.Nil, result.ID)
	})
}

func TestIGetMessageByIDQueryHandler(t *testing.T) {
	t.Run("should implement interface correctly", func(t *testing.T) {
		// Arrange
		handler := &MockGetMessageByIDQueryHandler{}

		// Act & Assert
		var _ interfaces.IGetMessageByIDQueryHandler = handler
		require.NotNil(t, handler)
	})

	t.Run("should handle get message by ID query", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		handler := &MockGetMessageByIDQueryHandler{}
		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{ID: messageID}

		// Act
		result, err := handler.Handle(ctx, query)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, messageID, result.ID)
		assert.Equal(t, "test", result.Text)
	})
}
