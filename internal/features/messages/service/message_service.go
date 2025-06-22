package service

import (
	"context"

	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/interfaces"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMessageService),
)

type MessageService struct{}

func NewMessageService() interfaces.IMessageService {
	return &MessageService{}
}

func (s *MessageService) CreateMessage(
	ctx context.Context,
	cmd *dtos.CreateMessageCommand,
) (*dtos.CreateMessageCommandResponse, error) {
	return mediatr.Send[*dtos.CreateMessageCommand, *dtos.CreateMessageCommandResponse](ctx, cmd)
}

func (s *MessageService) GetMessageByID(
	ctx context.Context,
	query *dtos.GetMessageByIDQuery,
) (*dtos.GetMessageByIDQueryResponse, error) {
	return mediatr.Send[*dtos.GetMessageByIDQuery, *dtos.GetMessageByIDQueryResponse](ctx, query)
}
