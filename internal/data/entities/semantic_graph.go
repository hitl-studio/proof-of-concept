package entities

// Node represents one typed element of the persisted graph.
type Node struct {
	ID   string
	Kind string
}

// Attribute represents one value owned by a graph node.
type Attribute struct {
	ID     string
	NodeID string
	Key    string
	Value  string
}

// Relation represents one directed connection between graph nodes.
type Relation struct {
	ID            string
	Kind          string
	OriginID      string
	DestinationID string
	Position      int
}
