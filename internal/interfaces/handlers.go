package interfaces

import (
	"context"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
)

// ICreateMessageCommandHandler defines the interface for the create message command handler.
type ICreateMessageCommandHandler interface {
	Handle(ctx context.Context, cmd *dtos.CreateMessageCommand) (*dtos.CreateMessageCommandResponse, error)
}

// IGetMessageByIDQueryHandler defines the interface for the get message by id query handler.
type IGetMessageByIDQueryHandler interface {
	Handle(ctx context.Context, query *dtos.GetMessageByIDQuery) (*dtos.GetMessageByIDQueryResponse, error)
}
