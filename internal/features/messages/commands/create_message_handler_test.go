package commands_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/commands"
	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock repository for testing
type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) CreateMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	args := m.Called(ctx, message)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Message), args.Error(1)
}

func (m *MockMessageRepository) GetMessageByID(ctx context.Context, id uuid.UUID) (*models.Message, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Message), args.Error(1)
}

func TestCreateMessageCommandHandler_Handle(t *testing.T) {
	t.Run("should create message successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := commands.NewCreateMessageCommandHandler(mockRepo)

		cmd := &dtos.CreateMessageCommand{
			Text: "test message",
		}

		expectedMessage := &models.Message{
			ID:   uuid.New(),
			Text: "test message",
		}

		mockRepo.On("CreateMessage", ctx, mock.MatchedBy(func(msg *models.Message) bool {
			return msg.Text == "test message"
		})).Return(expectedMessage, nil)

		// Act
		result, err := handler.Handle(ctx, cmd)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, expectedMessage.ID, result.ID)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := commands.NewCreateMessageCommandHandler(mockRepo)

		cmd := &dtos.CreateMessageCommand{
			Text: "test message",
		}

		expectedError := errors.New("database error")
		mockRepo.On("CreateMessage", ctx, mock.AnythingOfType("*models.Message")).Return(nil, expectedError)

		// Act
		result, err := handler.Handle(ctx, cmd)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should handle empty text", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := commands.NewCreateMessageCommandHandler(mockRepo)

		cmd := &dtos.CreateMessageCommand{
			Text: "",
		}

		expectedMessage := &models.Message{
			ID:   uuid.New(),
			Text: "",
		}

		mockRepo.On("CreateMessage", ctx, mock.MatchedBy(func(msg *models.Message) bool {
			return msg.Text == ""
		})).Return(expectedMessage, nil)

		// Act
		result, err := handler.Handle(ctx, cmd)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, expectedMessage.ID, result.ID)
		mockRepo.AssertExpectations(t)
	})
}

func TestNewCreateMessageCommandHandler(t *testing.T) {
	t.Run("should create handler with repository", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockMessageRepository)

		// Act
		handler := commands.NewCreateMessageCommandHandler(mockRepo)

		// Assert
		require.NotNil(t, handler)
		assert.IsType(t, &commands.CreateMessageCommandHandler{}, handler)
	})
}
