package usecases

import (
	"context"

	"poc/internal/domain/models"
)

// InitializeProjectUseCase prepares a project and its current working state.
type InitializeProjectUseCase interface {
	// Initialize prepares and returns the current project.
	Initialize(ctx context.Context) (models.Project, error)
}
