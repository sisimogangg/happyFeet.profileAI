package profile

import (
	"context"

	"github.com/sisimogangg/happyFeet.profileAI/models"
)

// Servicer represents the profile service
type Servicer interface {
	Fetch(ctx context.Context, cursor string, num int64) (*[]models.Profile, string, error)
	GetByID(ctx context.Context, id string) (*models.Profile, error)
	Update(ctx context.Context, p *models.Profile) error
	Create(ctx context.Context, p *models.Profile) error
	Delete(ctx context.Context, id string) error
}
