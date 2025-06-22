package queries

import (
	"context"
	"log"

	"github.com/arielsrv/fxf/internal/features/messages/repository"

	"github.com/google/uuid"
	"github.com/mehdihadeli/go-mediatr"
	"go.uber.org/fx"
)

// Module exports the query handler functionality.
var Module = fx.Options(
	fx.Provide(NewGetMessageByIDQueryHandler),
	fx.Invoke(registerGetMessageByIDQueryHandler),
)

// GetMessageByIDQuery is the query for retrieving a message by its ID.
type GetMessageByIDQuery struct {
	ID uuid.UUID
}

// GetMessageByIDQueryResponse is the response for GetMessageByIDQuery.
type GetMessageByIDQueryResponse struct {
	Text string    `json:"text"`
	ID   uuid.UUID `json:"id"`
}

// GetMessageByIDQueryHandler is the handler for GetMessageByIDQuery.
type GetMessageByIDQueryHandler struct {
	repo repository.IMessageRepository
}

// NewGetMessageByIDQueryHandler creates a new GetMessageByIDQueryHandler.
func NewGetMessageByIDQueryHandler(repo repository.IMessageRepository) *GetMessageByIDQueryHandler {
	return &GetMessageByIDQueryHandler{repo: repo}
}

// Handle handles the GetMessageByIDQuery.
func (h *GetMessageByIDQueryHandler) Handle(
	ctx context.Context,
	query *GetMessageByIDQuery,
) (*GetMessageByIDQueryResponse, error) {
	log.Printf("Handling GetMessageByIDQuery for ID: %s", query.ID)

	message, err := h.repo.GetMessageByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}

	return &GetMessageByIDQueryResponse{
		ID:   message.ID,
		Text: message.Text,
	}, nil
}

// registerGetMessageByIDQueryHandler registers the query handler with MediatR.
func registerGetMessageByIDQueryHandler(handler *GetMessageByIDQueryHandler) error {
	log.Println("Registering GetMessageByIDQueryHandler")
	return mediatr.RegisterRequestHandler[*GetMessageByIDQuery, *GetMessageByIDQueryResponse](handler)
}
