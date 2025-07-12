package queries_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/arielsrv/fxf/internal/features/messages/queries"
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

func TestGetMessageByIDQueryHandler_Handle(t *testing.T) {
	t.Run("should get message successfully", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := queries.NewGetMessageByIDQueryHandler(mockRepo)

		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{
			ID: messageID,
		}

		expectedMessage := &models.Message{
			ID:   messageID,
			Text: "test message",
		}

		mockRepo.On("GetMessageByID", ctx, messageID).Return(expectedMessage, nil)

		// Act
		result, err := handler.Handle(ctx, query)

		// Assert
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, expectedMessage.ID, result.ID)
		assert.Equal(t, expectedMessage.Text, result.Text)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when message not found", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := queries.NewGetMessageByIDQueryHandler(mockRepo)

		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{
			ID: messageID,
		}

		expectedError := errors.New("message not found")
		mockRepo.On("GetMessageByID", ctx, messageID).Return(nil, expectedError)

		// Act
		result, err := handler.Handle(ctx, query)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		mockRepo := new(MockMessageRepository)
		handler := queries.NewGetMessageByIDQueryHandler(mockRepo)

		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{
			ID: messageID,
		}

		expectedError := errors.New("database error")
		mockRepo.On("GetMessageByID", ctx, messageID).Return(nil, expectedError)

		// Act
		result, err := handler.Handle(ctx, query)

		// Assert
		require.Error(t, err)
		require.Nil(t, result)
		assert.Equal(t, expectedError, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestNewGetMessageByIDQueryHandler(t *testing.T) {
	t.Run("should create handler with repository", func(t *testing.T) {
		// Arrange
		mockRepo := new(MockMessageRepository)

		// Act
		handler := queries.NewGetMessageByIDQueryHandler(mockRepo)

		// Assert
		require.NotNil(t, handler)
		assert.IsType(t, &queries.GetMessageByIDQueryHandler{}, handler)
	})
}
