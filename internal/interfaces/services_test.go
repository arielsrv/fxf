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
type MockMessageService struct{}

func (m *MockMessageService) CreateMessage(ctx context.Context, cmd *dtos.CreateMessageCommand) (*dtos.CreateMessageCommandResponse, error) {
	return &dtos.CreateMessageCommandResponse{ID: uuid.New()}, nil
}

func (m *MockMessageService) GetMessageByID(ctx context.Context, query *dtos.GetMessageByIDQuery) (*dtos.GetMessageByIDQueryResponse, error) {
	return &dtos.GetMessageByIDQueryResponse{ID: query.ID, Text: "test"}, nil
}

func TestIMessageService(t *testing.T) {
	t.Run("should implement interface correctly", func(t *testing.T) {
		// Arrange
		service := &MockMessageService{}

		// Act & Assert
		var _ interfaces.IMessageService = service
		require.NotNil(t, service)
	})

	t.Run("should create message", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		service := &MockMessageService{}
		cmd := &dtos.CreateMessageCommand{Text: "test"}

		// Act
		result, err := service.CreateMessage(ctx, cmd)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.NotEqual(t, uuid.Nil, result.ID)
	})

	t.Run("should get message by ID", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		service := &MockMessageService{}
		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{ID: messageID}

		// Act
		result, err := service.GetMessageByID(ctx, query)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, messageID, result.ID)
		assert.Equal(t, "test", result.Text)
	})
}
