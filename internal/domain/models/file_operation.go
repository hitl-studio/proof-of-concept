package models

// FileOperation represents one operation over a project file.
type FileOperation interface {
	Kind() FileOperationKind
}

// FileOperationKind identifies the kind of a file operation.
type FileOperationKind string

const (
	FileOperationKindAdd    FileOperationKind = "add"
	FileOperationKindUpdate FileOperationKind = "update"
	FileOperationKindDelete FileOperationKind = "delete"
	FileOperationKindMove   FileOperationKind = "move"
)

// AddFileOperation creates a project file with the provided content.
type AddFileOperation struct {
	Path    string
	Content string
}

func (AddFileOperation) Kind() FileOperationKind {
	return FileOperationKindAdd
}

// UpdateFileOperation applies an ordered collection of changes to a project file.
type UpdateFileOperation struct {
	Path   string
	Chunks []FileUpdateChunk
}

func (UpdateFileOperation) Kind() FileOperationKind {
	return FileOperationKindUpdate
}

// DeleteFileOperation removes a project file.
type DeleteFileOperation struct {
	Path string
}

func (DeleteFileOperation) Kind() FileOperationKind {
	return FileOperationKindDelete
}

// MoveFileOperation moves a project file to another path.
type MoveFileOperation struct {
	SourcePath      string
	DestinationPath string
}

func (MoveFileOperation) Kind() FileOperationKind {
	return FileOperationKindMove
}

// FileUpdateChunk replaces one contiguous block of lines in a project file.
type FileUpdateChunk struct {
	Anchor      string
	OldLines    []string
	NewLines    []string
	AtEndOfFile bool
}
