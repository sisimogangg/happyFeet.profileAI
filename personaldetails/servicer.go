package personaldetails

import (
	"context"

	"github.com/sisimogangg/happyFeet.profileAI/models"
)

// Servicer services all personal details queries
type Servicer interface {
	GetByProfileID(ctx context.Context, profileID int64) (*models.PersonalDetails, error)
}
