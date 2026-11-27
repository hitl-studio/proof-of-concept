package models

// Project contains the source directory and its current working state.
type Project struct {
	RootPath      string
	Conversation  Conversation
	Registry      SemanticRegistry
	SemanticGraph SemanticGraph
	VirtualTree   VirtualTree
}
