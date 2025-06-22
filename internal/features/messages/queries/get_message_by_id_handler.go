package queries

import (
	"context"
	"github.com/arielsrv/fxf/internal/features/messages/dtos"
	"github.com/arielsrv/fxf/internal/features/messages/repository"
	"github.com/arielsrv/fxf/internal/interfaces"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

// Module exports the query handler functionality.
var Module = fx.Options(
	fx.Provide(NewGetMessageByIDQueryHandler),
	fx.Invoke(registerGetMessageByIDQueryHandler),
)

// GetMessageByIDQueryHandler is the handler for GetMessageByIDQuery.
type GetMessageByIDQueryHandler struct {
	repo repository.IMessageRepository
}

// NewGetMessageByIDQueryHandler creates a new GetMessageByIDQueryHandler.
func NewGetMessageByIDQueryHandler(repo repository.IMessageRepository) interfaces.IGetMessageByIDQueryHandler {
	return &GetMessageByIDQueryHandler{repo: repo}
}

// Handle handles the GetMessageByIDQuery.
func (h *GetMessageByIDQueryHandler) Handle(
	ctx context.Context,
	query *dtos.GetMessageByIDQuery,
) (*dtos.GetMessageByIDQueryResponse, error) {
	message, err := h.repo.GetMessageByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return &dtos.GetMessageByIDQueryResponse{
		ID:   message.ID,
		Text: message.Text,
	}, nil
}

// registerGetMessageByIDQueryHandler registers the query handler with MediatR.
func registerGetMessageByIDQueryHandler(handler interfaces.IGetMessageByIDQueryHandler) error {
	return mediatr.RegisterRequestHandler[*dtos.GetMessageByIDQuery, *dtos.GetMessageByIDQueryResponse](handler)
}
