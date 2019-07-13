package service

import (
	"context"
	"time"

	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/personaldetails"
)

type personalService struct {
	repo           personaldetails.Repository
	contextTimeout time.Duration
	// Address repo here
}

// NewPersonalDetailsService returns a new instance of this service
func NewPersonalDetailsService(repo personaldetails.Repository, timeout time.Duration) personaldetails.Servicer {
	return &personalService{repo: repo, contextTimeout: timeout}
}

func (p *personalService) GetByProfileID(ctx context.Context, profileID int64) (*models.PersonalDetails, error) {

	// Get personal details from the repo

	// Get Address from Address repo

	return nil, nil
}
