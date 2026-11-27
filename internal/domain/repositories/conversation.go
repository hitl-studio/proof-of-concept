package repositories

import (
	"context"

	"poc/internal/domain/models"
)

// ConversationRepository provides persistent access to the project conversation.
type ConversationRepository interface {
	// GetConversation returns the current conversation.
	GetConversation(ctx context.Context) (models.Conversation, error)

	// SaveMessage appends a new message or updates an existing message while
	// preserving its position in the conversation.
	SaveMessage(ctx context.Context, message models.Message) error

	// ClearConversation removes all messages from the conversation.
	ClearConversation(ctx context.Context) error
}
