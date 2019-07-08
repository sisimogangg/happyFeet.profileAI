package personaldetails

import (
	"context"

	"github.com/sisimogangg/happyFeet.profileAI/models"
)

// Repository represents the Kids repository
type Repository interface {
	GetByProfileID(ctx context.Context, profileID int64) (*models.PersonalDetails, error)
}
