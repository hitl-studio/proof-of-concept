package usecases

import (
	"context"
)

// MaterializeChangesUseCase transfers prepared changes from the semantic graph
// to the current project on a best-effort basis.
type MaterializeChangesUseCase interface {
	// MaterializeNode materializes the direct changes of a node.
	MaterializeNode(ctx context.Context, nodeID string) error

	// MaterializeSubtree materializes the changes of a node and its effective
	// subtree without including detached areas.
	MaterializeSubtree(ctx context.Context, nodeID string) error
}
