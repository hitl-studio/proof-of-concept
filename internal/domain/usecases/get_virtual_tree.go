package usecases

import (
	"context"

	"poc/internal/domain/models"
)

// GetVirtualTreeUseCase provides virtual trees for semantic nodes.
type GetVirtualTreeUseCase interface {
	// GetForNode returns a virtual tree formed from the direct changes of a node.
	GetForNode(ctx context.Context, nodeID string) (models.VirtualTree, error)

	// GetForSubtree returns a virtual tree formed from a node and its effective subtree.
	GetForSubtree(ctx context.Context, nodeID string) (models.VirtualTree, error)
}
