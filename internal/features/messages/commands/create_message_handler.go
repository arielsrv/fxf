package commands

import (
	"context"
	"log"

	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/arielsrv/fxf/internal/features/messages/repository"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

// Module exports the command handler functionality.
var Module = fx.Options(
	fx.Provide(NewCreateMessageCommandHandler),
	fx.Invoke(registerCreateMessageCommandHandler),
)

// CreateMessageCommand is the command for creating a new message.
type CreateMessageCommand struct {
	Text string `json:"text"`
}

// CreateMessageCommandResponse is the response for CreateMessageCommand.
type CreateMessageCommandResponse struct {
	ID uuid.UUID `json:"id"`
}

// CreateMessageCommandHandler is the handler for CreateMessageCommand.
type CreateMessageCommandHandler struct {
	repo repository.IMessageRepository
}

// NewCreateMessageCommandHandler creates a new CreateMessageCommandHandler.
func NewCreateMessageCommandHandler(repo repository.IMessageRepository) *CreateMessageCommandHandler {
	return &CreateMessageCommandHandler{repo: repo}
}

// Handle handles the CreateMessageCommand.
func (h *CreateMessageCommandHandler) Handle(
	ctx context.Context,
	cmd *CreateMessageCommand,
) (*CreateMessageCommandResponse, error) {
	message := &models.Message{
		Text: cmd.Text,
	}

	createdMessage, err := h.repo.CreateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return &CreateMessageCommandResponse{ID: createdMessage.ID}, nil
}

// registerCreateMessageCommandHandler registers the command handler with MediatR.
func registerCreateMessageCommandHandler(handler *CreateMessageCommandHandler) error {
	log.Println("Registering CreateMessageCommandHandler")
	return mediatr.RegisterRequestHandler[*CreateMessageCommand, *CreateMessageCommandResponse](handler)
}
