package http

import (
	"github.com/arielsrv/fxf/internal/features/messages/commands"
	"github.com/arielsrv/fxf/internal/features/messages/queries"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

// Module exports the HTTP handlers functionality.
var Module = fx.Options(
	fx.Provide(NewMessageHandlers),
	fx.Invoke(RegisterRoutes),
)

// MessageHandlers contains the handlers for message-related routes.
type MessageHandlers struct {
	// fx.In // This was causing the issue. Removed as it's not needed.
}

// NewMessageHandlers creates new message handlers.
func NewMessageHandlers() *MessageHandlers {
	return &MessageHandlers{}
}

// RegisterRoutes registers the message routes to the Fiber app.
func RegisterRoutes(app *fiber.App, handlers *MessageHandlers) {
	app.Post("/messages", handlers.CreateMessage)
	app.Get("/messages/:id", handlers.GetMessageByID)
}

// CreateMessage handles the creation of a new message.
func (h *MessageHandlers) CreateMessage(c *fiber.Ctx) error {
	cmd := new(commands.CreateMessageCommand)
	if err := c.BodyParser(cmd); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse request body",
		})
	}

	result, err := mediatr.Send[*commands.CreateMessageCommand, *commands.CreateMessageCommandResponse](
		c.Context(),
		cmd,
	)
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

	query := &queries.GetMessageByIDQuery{ID: id}

	result, err := mediatr.Send[*queries.GetMessageByIDQuery, *queries.GetMessageByIDQueryResponse](c.Context(), query)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
