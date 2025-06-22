package commands

import (
	"context"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/features/messages/models"
	"github.com/arielsrv/fxf/internal/features/messages/repository"
	"github.com/arielsrv/fxf/internal/interfaces"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

// Module exports the command handler functionality.
var Module = fx.Options(
	fx.Provide(NewCreateMessageCommandHandler),
	fx.Invoke(registerCreateMessageCommandHandler),
)

// CreateMessageCommandHandler is the handler for CreateMessageCommand.
type CreateMessageCommandHandler struct {
	repo repository.IMessageRepository
}

// NewCreateMessageCommandHandler creates a new CreateMessageCommandHandler.
func NewCreateMessageCommandHandler(repo repository.IMessageRepository) interfaces.ICreateMessageCommandHandler {
	return &CreateMessageCommandHandler{repo: repo}
}

// Handle handles the CreateMessageCommand.
func (h *CreateMessageCommandHandler) Handle(
	ctx context.Context,
	cmd *dtos.CreateMessageCommand,
) (*dtos.CreateMessageCommandResponse, error) {
	message := &models.Message{
		Text: cmd.Text,
	}

	createdMessage, err := h.repo.CreateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateMessageCommandResponse{ID: createdMessage.ID}, nil
}

// registerCreateMessageCommandHandler registers the command handler with MediatR.
func registerCreateMessageCommandHandler(handler interfaces.ICreateMessageCommandHandler) error {
	return mediatr.RegisterRequestHandler[*dtos.CreateMessageCommand, *dtos.CreateMessageCommandResponse](handler)
}
