package usecases

import (
	"context"

	"poc/internal/domain/models"
)

// GetConversationUseCase provides the current project conversation.
type GetConversationUseCase interface {
	// Get returns the current conversation.
	Get(ctx context.Context) (models.Conversation, error)
}
