package models

// SemanticRegistry contains the semantic contexts available in a project.
type SemanticRegistry struct {
	Context []SemanticContext
}

// SemanticContext describes one reusable named context identified by its semantic path.
type SemanticContext struct {
	Path        string
	Description string
}
