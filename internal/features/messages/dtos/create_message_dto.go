package dtos

import "github.com/google/uuid"

// CreateMessageCommand is the command for creating a new message.
type CreateMessageCommand struct {
	Text string `json:"text"`
}

// CreateMessageCommandResponse is the response for CreateMessageCommand.
type CreateMessageCommandResponse struct {
	ID uuid.UUID `json:"id"`
}
