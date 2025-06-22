package dtos

import "github.com/google/uuid"

// GetMessageByIDQuery is the query for retrieving a message by its ID.
type GetMessageByIDQuery struct {
	ID uuid.UUID
}

// GetMessageByIDQueryResponse is the response for GetMessageByIDQuery.
type GetMessageByIDQueryResponse struct {
	Text string    `json:"text"`
	ID   uuid.UUID `json:"id"`
}
