package http_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/arielsrv/fxf/internal/features/messages/delivery/http"
	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func TestMessageHandlers_CreateMessage(t *testing.T) {
	t.Run("should create message successfully", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		createMessageCmd := &dtos.CreateMessageCommand{
			Text: "test message",
		}

		expectedResponse := &dtos.CreateMessageCommandResponse{
			ID: uuid.New(),
		}

		// Use mock.AnythingOfType for context since Fiber passes its own context type
		mockService.On("CreateMessage", mock.AnythingOfType("*fasthttp.RequestCtx"), createMessageCmd).Return(expectedResponse, nil)

		http.RegisterRoutes(app, handlers)

		body, _ := json.Marshal(createMessageCmd)
		req := httptest.NewRequest("POST", "/messages", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.NotNil(t, responseBody["id"])

		mockService.AssertExpectations(t)
	})

	t.Run("should return bad request when body parsing fails", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		http.RegisterRoutes(app, handlers)

		req := httptest.NewRequest("POST", "/messages", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.Equal(t, "cannot parse request body", responseBody["error"])
	})

	t.Run("should return internal server error when service fails", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		createMessageCmd := &dtos.CreateMessageCommand{
			Text: "test message",
		}

		expectedError := errors.New("service error")
		mockService.On("CreateMessage", mock.AnythingOfType("*fasthttp.RequestCtx"), createMessageCmd).Return(nil, expectedError)

		http.RegisterRoutes(app, handlers)

		body, _ := json.Marshal(createMessageCmd)
		req := httptest.NewRequest("POST", "/messages", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.Equal(t, expectedError.Error(), responseBody["error"])

		mockService.AssertExpectations(t)
	})
}

func TestMessageHandlers_GetMessageByID(t *testing.T) {
	t.Run("should get message successfully", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{
			ID: messageID,
		}

		expectedResponse := &dtos.GetMessageByIDQueryResponse{
			ID:   messageID,
			Text: "test message",
		}

		mockService.On("GetMessageByID", mock.AnythingOfType("*fasthttp.RequestCtx"), query).Return(expectedResponse, nil)

		http.RegisterRoutes(app, handlers)

		req := httptest.NewRequest("GET", "/messages/"+messageID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.Equal(t, messageID.String(), responseBody["id"])
		assert.Equal(t, "test message", responseBody["text"])

		mockService.AssertExpectations(t)
	})

	t.Run("should return bad request when UUID is invalid", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		http.RegisterRoutes(app, handlers)

		req := httptest.NewRequest("GET", "/messages/invalid-uuid", nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.Equal(t, "invalid UUID format", responseBody["error"])
	})

	t.Run("should return not found when message not found", func(t *testing.T) {
		// Arrange
		app := fiber.New()
		mockService := new(MockMessageService)
		handlers := http.NewMessageHandlers(mockService)

		messageID := uuid.New()
		query := &dtos.GetMessageByIDQuery{
			ID: messageID,
		}

		expectedError := errors.New("message not found")
		mockService.On("GetMessageByID", mock.AnythingOfType("*fasthttp.RequestCtx"), query).Return(nil, expectedError)

		http.RegisterRoutes(app, handlers)

		req := httptest.NewRequest("GET", "/messages/"+messageID.String(), nil)

		// Act
		resp, err := app.Test(req)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

		var responseBody map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		require.NoError(t, err)
		assert.Equal(t, expectedError.Error(), responseBody["error"])

		mockService.AssertExpectations(t)
	})
}

func TestNewMessageHandlers(t *testing.T) {
	t.Run("should create handlers with service", func(t *testing.T) {
		// Arrange
		mockService := new(MockMessageService)

		// Act
		handlers := http.NewMessageHandlers(mockService)

		// Assert
		require.NotNil(t, handlers)
		assert.IsType(t, &http.MessageHandlers{}, handlers)
	})
}
