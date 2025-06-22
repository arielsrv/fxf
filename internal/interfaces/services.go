package interfaces

import (
	"context"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
)

// IMessageService defines the application service for messages.
type IMessageService interface {
	CreateMessage(ctx context.Context, cmd *dtos.CreateMessageCommand) (*dtos.CreateMessageCommandResponse, error)
	GetMessageByID(ctx context.Context, query *dtos.GetMessageByIDQuery) (*dtos.GetMessageByIDQueryResponse, error)
}
