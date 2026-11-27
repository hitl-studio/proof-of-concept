package usecases

import (
	"context"

	"poc/internal/domain/models"
)

// GetSemanticGraphUseCase provides the current project semantic graph and its updates.
type GetSemanticGraphUseCase interface {
	// Get returns a stream containing the current graph followed by subsequent states.
	Get(ctx context.Context) (<-chan models.SemanticGraph, error)
}
