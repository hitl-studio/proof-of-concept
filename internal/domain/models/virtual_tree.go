package models

// VirtualTree represents the virtual file state produced by applying prepared
// changes to the current project without materializing them on disk.
type VirtualTree struct {
	Files []VirtualFile
}

// VirtualFile represents one affected file in a virtual tree.
type VirtualFile struct {
	Path         string
	PreviousPath string
	Content      string
	Status       VirtualFileStatus
	Diff         VirtualFileDiff
}

// VirtualFileStatus identifies how a virtual file differs from the project on disk.
type VirtualFileStatus string

const (
	VirtualFileStatusAdded    VirtualFileStatus = "added"
	VirtualFileStatusModified VirtualFileStatus = "modified"
	VirtualFileStatusDeleted  VirtualFileStatus = "deleted"
	VirtualFileStatusMoved    VirtualFileStatus = "moved"
)

// VirtualFileDiff describes the full line-by-line difference between a virtual
// file and its corresponding state in the project on disk.
type VirtualFileDiff struct {
	Lines []VirtualFileDiffLine
}

// VirtualFileDiffLine represents one line in a full-file diff.
// A missing line number on either side is represented by zero.
type VirtualFileDiffLine struct {
	Kind               VirtualFileDiffLineKind
	Content            string
	PreviousLineNumber int
	ResultLineNumber   int
}

// VirtualFileDiffLineKind identifies how a line differs between the project on
// disk and the virtual file.
type VirtualFileDiffLineKind string

const (
	VirtualFileDiffLineUnchanged VirtualFileDiffLineKind = "unchanged"
	VirtualFileDiffLineAdded     VirtualFileDiffLineKind = "added"
	VirtualFileDiffLineRemoved   VirtualFileDiffLineKind = "removed"
)
