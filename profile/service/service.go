package service

import (
	"context"
	"time"

	"github.com/sisimogangg/happyFeet.profileAI/kids"

	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/profile"
	"golang.org/x/sync/errgroup"
)

type profileService struct {
	profileRepo    profile.Repository
	kidsRepo       kids.Repository
	contextTimeout time.Duration
}

// NewProfileService returns a new instance of profileService
func NewProfileService(p profile.Repository, k kids.Repository, timeout time.Duration) profile.Servicer {
	return &profileService{profileRepo: p, kidsRepo: k, contextTimeout: timeout}
}

func (p *profileService) fillPupilsDetails(c context.Context, profiles []*models.Profile) ([]*models.Profile, error) {
	g, ctx := errgroup.WithContext(c)

	mapKids := map[int64]models.Kid{}

	for _, profile := range profiles {
		mapKids[profile.ID] = models.Kid{}
	}

	chanKids := make(chan []*models.Kid)

	for kid := range mapKids {
		g.Go(func() error {
			res, err := p.kidsRepo.GetByProfileID(ctx, kid)
			if err != nil {
				return err
			}
			chanKids <- res
			return nil
		})
	}
	return nil, nil
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
