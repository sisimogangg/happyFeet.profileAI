package profile

import (
	"context"

	"github.com/sisimogangg/happyFeet.profileAI/models"
)

// Repository represents profile's repository contract
type Repository interface {
	Fetch(ctx context.Context, num int64) (res []*models.Profile, err error)
	GetByUserID(ctx context.Context, string string) (*models.Profile, error)
	Update(ctx context.Context, p *models.Profile) error
	Create(ctx context.Context, p *models.Profile) error
	Delete(ctx context.Context, id string) error
}
