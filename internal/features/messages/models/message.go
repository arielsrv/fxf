package models

import "github.com/google/uuid"

type Message struct {
	Text string
	ID   uuid.UUID
}
