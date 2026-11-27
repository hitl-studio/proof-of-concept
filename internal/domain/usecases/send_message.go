package usecases

import (
	"context"

	"poc/internal/domain/models"
)

// SendMessageUseCase sends a user message and observes the resulting messages.
type SendMessageUseCase interface {
	// Send returns a stream of current message states for one conversation turn.
	Send(ctx context.Context, content string) (<-chan models.Message, error)
}
