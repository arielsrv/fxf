package service_test

import (
	"context"
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/features/messages/service"
	"github.com/arielsrv/fxf/internal/interfaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Mock service for testing
type MockMessageService struct {
	mock.Mock
}

func (m *MockMessageService) CreateMessage(ctx context.Context, cmd *dtos.CreateMessageCommand) (*dtos.CreateMessageCommandResponse, error) {
	args := m.Called(ctx, cmd)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.CreateMessageCommandResponse), args.Error(1)
}

func (m *MockMessageService) GetMessageByID(ctx context.Context, query *dtos.GetMessageByIDQuery) (*dtos.GetMessageByIDQueryResponse, error) {
	args := m.Called(ctx, query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dtos.GetMessageByIDQueryResponse), args.Error(1)
}

func TestMessageService_CreateMessage(t *testing.T) {
	t.Run("should create message service successfully", func(t *testing.T) {
		// Act
		msgService := service.NewMessageService()

		// Assert
		require.NotNil(t, msgService)
		assert.IsType(t, &service.MessageService{}, msgService)
	})
}

func TestMessageService_GetMessageByID(t *testing.T) {
	t.Run("should create message service successfully", func(t *testing.T) {
		// Act
		msgService := service.NewMessageService()

		// Assert
		require.NotNil(t, msgService)
		assert.IsType(t, &service.MessageService{}, msgService)
	})
}

func TestNewMessageService(t *testing.T) {
	t.Run("should create message service", func(t *testing.T) {
		// Act
		msgService := service.NewMessageService()

		// Assert
		require.NotNil(t, msgService)
		assert.IsType(t, &service.MessageService{}, msgService)
		assert.Implements(t, (*interfaces.IMessageService)(nil), msgService)
	})
}
