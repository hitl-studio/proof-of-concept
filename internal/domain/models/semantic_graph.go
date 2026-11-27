package models

// SemanticGraph represents the shared semantic structure of a project change.
type SemanticGraph struct {
	RootID string
	Nodes  []SemanticNode
}

// SemanticNode represents one typed part of a semantic graph.
type SemanticNode struct {
	ID          string
	Type        string
	Changes     []Change
	ChildrenIDs []string
	States      []SemanticNodeStateType
}

// Change represents one project change associated with a semantic node.
type Change struct {
	ID        string
	Revision  uint64
	Operation FileOperation
	Error     string
}

// SemanticNodeStateType identifies a state associated with a semantic node.
type SemanticNodeStateType string

const (
	SemanticNodeStateSelected   SemanticNodeStateType = "selected"
	SemanticNodeStateCheckedOut SemanticNodeStateType = "checked_out"
	SemanticNodeStateDetached   SemanticNodeStateType = "detached"
	SemanticNodeStateFrozen     SemanticNodeStateType = "frozen"
	SemanticNodeStateFocused    SemanticNodeStateType = "focused"
	SemanticNodeStateExpanded   SemanticNodeStateType = "expanded"
)
