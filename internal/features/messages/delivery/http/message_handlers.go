package http

import (
	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

// Module exports the HTTP handlers functionality.
var Module = fx.Options(
	fx.Provide(NewMessageHandlers),
	fx.Invoke(RegisterRoutes),
)

// MessageHandlers contains the handlers for message-related routes.
type MessageHandlers struct {
	service interfaces.IMessageService
}

// NewMessageHandlers creates new message handlers.
func NewMessageHandlers(service interfaces.IMessageService) *MessageHandlers {
	return &MessageHandlers{service: service}
}

// RegisterRoutes registers the message routes to the Fiber app.
func RegisterRoutes(app *fiber.App, handlers *MessageHandlers) {
	app.Post("/messages", handlers.CreateMessage)
	app.Get("/messages/:id", handlers.GetMessageByID)
}

// CreateMessage handles the creation of a new message.
func (h *MessageHandlers) CreateMessage(c *fiber.Ctx) error {
	cmd := new(dtos.CreateMessageCommand)
	if err := c.BodyParser(cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	result, err := h.service.CreateMessage(c.Context(), cmd)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetMessageByID handles retrieving a message by its ID.
func (h *MessageHandlers) GetMessageByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid UUID format",
		})
	}

	query := &dtos.GetMessageByIDQuery{ID: id}

	result, err := h.service.GetMessageByID(c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
