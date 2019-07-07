package service

import (
	"context"
	"time"

	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/profile"
)

type profileService struct {
	profileRepo    profile.Repository
	contextTimeout time.Duration
}

// NewProfileService returns a new instance of profileService
func NewProfileService(p profile.Repository, timeout time.Duration) profile.Servicer {
	return &profileService{profileRepo: p, contextTimeout: timeout}
}

func (p *profileService) Fetch(c context.Context, num int64) ([]*models.Profile, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	listProfiles, err := p.profileRepo.Fetch(ctx, num)
	if err != nil {
		return nil, err
	}

	return listProfiles, nil
}

func (p *profileService) GetByID(ctx context.Context, id string) (*models.Profile, error) {
	return nil, nil
}

func (p *profileService) Update(ctx context.Context, profile *models.Profile) error {
	return nil
}
func (p *profileService) Create(ctx context.Context, profile *models.Profile) error {
	return nil
}
func (p *profileService) Delete(ctx context.Context, id string) error {
	return nil
}
