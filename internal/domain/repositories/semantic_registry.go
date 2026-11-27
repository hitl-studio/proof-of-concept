package repositories

import (
	"context"

	"poc/internal/domain/models"
)

// SemanticRegistryRepository provides read-only access to the semantic registry.
type SemanticRegistryRepository interface {
	// GetRegistry returns the complete semantic registry.
	GetRegistry(ctx context.Context) (models.SemanticRegistry, error)

	// GetContext returns one semantic context by its semantic path.
	GetContext(ctx context.Context, path string) (models.SemanticContext, error)
}
