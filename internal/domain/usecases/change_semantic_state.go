package usecases

import "context"

// ChangeSemanticNodeStateUseCase changes the interactive state of semantic nodes.
type ChangeSemanticNodeStateUseCase interface {
	// Select marks a node as selected.
	Select(ctx context.Context, nodeID string) error

	// Checkout makes a node the current agent context.
	Checkout(ctx context.Context, nodeID string) error

	// Focus makes a node the local root of the graph presentation.
	Focus(ctx context.Context, nodeID string) error

	// Unfocus removes a node as the local root of the graph presentation.
	Unfocus(ctx context.Context, nodeID string) error

	// Expand reveals the children of a node.
	Expand(ctx context.Context, nodeID string) error

	// Collapse hides the children of a node.
	Collapse(ctx context.Context, nodeID string) error

	// Detach removes a node and its descendants from the current line of work.
	Detach(ctx context.Context, nodeID string) error

	// Attach returns a detached node and its descendants to the current line of work.
	Attach(ctx context.Context, nodeID string) error

	// Freeze prevents changes to a node and its descendants.
	Freeze(ctx context.Context, nodeID string) error

	// Unfreeze permits changes to a frozen node and its descendants.
	Unfreeze(ctx context.Context, nodeID string) error
}
