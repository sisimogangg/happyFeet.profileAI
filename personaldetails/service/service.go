package service

import (
	"github.com/sisimogangg/happyFeet.profileAI/personaldetails"
)

type personalService struct {
	repo *personaldetails.Repository
}

// NewPersonalDetailsService returns a new instance of this service
func NewPersonalDetailsService(repo *personaldetails.Repository) personaldetails.Servicer {
	return &personalService{repo: repo}
}
