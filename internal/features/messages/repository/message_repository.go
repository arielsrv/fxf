package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/arielsrv/fxf/internal/features/messages/models"

	"github.com/google/uuid"
	"go.uber.org/fx"
)

// Module exports the repository functionality.
var Module = fx.Options(
	fx.Provide(NewInMemoryMessageRepository),
)

// IMessageRepository defines the interface for message repository.
type IMessageRepository interface {
	CreateMessage(ctx context.Context, message *models.Message) (*models.Message, error)
	GetMessageByID(ctx context.Context, id uuid.UUID) (*models.Message, error)
}

// InMemoryMessageRepository is an in-memory implementation of IMessageRepository.
type InMemoryMessageRepository struct {
	messages map[uuid.UUID]*models.Message
	mu       sync.RWMutex
}

// NewInMemoryMessageRepository creates a new InMemoryMessageRepository.
// The fx.In is not strictly necessary here but shows how dependencies would be injected.
func NewInMemoryMessageRepository() IMessageRepository {
	return &InMemoryMessageRepository{
		messages: make(map[uuid.UUID]*models.Message),
	}
}

// CreateMessage adds a new message to the store.
func (r *InMemoryMessageRepository) CreateMessage(
	ctx context.Context,
	message *models.Message,
) (*models.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if message.ID == uuid.Nil {
		message.ID = uuid.New()
	}

	r.messages[message.ID] = message
	return message, nil
}

// GetMessageByID retrieves a message from the store by its ID.
func (r *InMemoryMessageRepository) GetMessageByID(ctx context.Context, id uuid.UUID) (*models.Message, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	message, ok := r.messages[id]
	if !ok {
		return nil, fmt.Errorf("message with ID %s not found", id)
	}
	return message, nil
}
