package service

import (
	"context"
	"time"

	"github.com/sisimogangg/happyFeet.profileAI/personaldetails"

	"github.com/pkg/errors"
	"github.com/sisimogangg/happyFeet.profileAI/kids"

	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/profile"
	"golang.org/x/sync/errgroup"
)

type profileService struct {
	profileRepo            profile.Repository
	kidsRepo               kids.Repository
	personalDetailsService personaldetails.Servicer
	contextTimeout         time.Duration
}

// NewProfileService returns a new instance of profileService
func NewProfileService(p profile.Repository, k kids.Repository, pd personaldetails.Servicer, timeout time.Duration) profile.Servicer {
	return &profileService{
		profileRepo:            p,
		kidsRepo:               k,
		personalDetailsService: pd,
		contextTimeout:         timeout,
	}
}

func (p *profileService) fillPupilsDetails(c context.Context, profiles []*models.Profile) ([]*models.Profile, error) {
	type profileKids struct {
		kids     []*models.Kid
		parentID int64
	}

	g, ctx := errgroup.WithContext(c)

	mapProfileIDsToKids := map[int64]profileKids{}

	for _, profile := range profiles {
		mapProfileIDsToKids[profile.ID] = profileKids{}
	}

	chanKids := make(chan *profileKids)

	for profileID := range mapProfileIDsToKids {
		profileID := profileID
		g.Go(func() error {
			res, err := p.kidsRepo.GetByProfileID(ctx, profileID)
			if err != nil {
				return errors.Wrap(err, "Error retrieving kids from Repo")
			}

			response := profileKids{
				res,
				profileID,
			}

			chanKids <- &response
			return nil
		})
	}

	go func() {
		err := g.Wait()
		if err != nil {
			return
		}
		close(chanKids)
	}()

	for pkids := range chanKids {
		if pkids != nil {
			mapProfileIDsToKids[pkids.parentID] = *pkids
		}
	}

	if err := g.Wait(); err != nil {
		return nil, errors.Wrap(err, "Error in thread retrieving kids")
	}

	for _, p := range profiles {
		if k, ok := mapProfileIDsToKids[p.ID]; ok {
			p.Kids = k.kids
		}
	}

	return profiles, nil
}

func (p *profileService) Fetch(c context.Context, num int64) ([]*models.Profile, error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()

	listProfiles, err := p.profileRepo.Fetch(ctx, num)
	if err != nil {
		return nil, errors.Wrap(err, "Error retrieving profiles")
	}

	listProfiles, err = p.fillPupilsDetails(ctx, listProfiles)
	if err != nil {
		return nil, errors.Wrap(err, "Error filling kids details")
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
