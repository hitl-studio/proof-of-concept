package repositories

import (
	"context"

	"poc/internal/domain/models"
)

// SemanticGraphRepository provides persistent access to the semantic graph.
type SemanticGraphRepository interface {
	// GetGraph returns the current semantic graph.
	GetGraph(ctx context.Context) (models.SemanticGraph, error)

	// SaveGraph saves the complete semantic graph.
	SaveGraph(ctx context.Context, graph models.SemanticGraph) error

	// DeleteGraph deletes the complete semantic graph.
	DeleteGraph(ctx context.Context) error

	// GetNode returns one semantic node by its identifier.
	GetNode(ctx context.Context, nodeID string) (models.SemanticNode, error)

	// SaveNode saves one semantic node.
	SaveNode(ctx context.Context, node models.SemanticNode) error

	// SaveNodes atomically saves a collection of semantic nodes.
	SaveNodes(ctx context.Context, nodes []models.SemanticNode) error

	// DeleteNode deletes one semantic node.
	DeleteNode(ctx context.Context, nodeID string) error

	// DeleteNodes atomically deletes the semantic nodes identified by nodeIDs.
	DeleteNodes(ctx context.Context, nodeIDs []string) error
}
