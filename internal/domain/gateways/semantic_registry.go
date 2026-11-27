package gateways

import (
	"context"

	"poc/internal/domain/models"
)

// SemanticRegistryGateway provides access to the origin semantic registry.
type SemanticRegistryGateway interface {
	// GetOrigin returns the complete registry from its origin.
	GetOrigin(ctx context.Context) (models.SemanticRegistry, error)
}
