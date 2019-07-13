package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/personaldetails"
)

type mySQLPersonalDetailsRepo struct {
	Conn *sql.DB
}

//NewMySQLPersonalDetailsRepo returns a new instance of profile repo
func NewMySQLPersonalDetailsRepo(Conn *sql.DB) personaldetails.Repository {
	return &mySQLPersonalDetailsRepo{Conn: Conn}
}

func (m *mySQLPersonalDetailsRepo) fetch(ctx context.Context, query string, args ...interface{}) (*models.PersonalDetails, error) {
	row, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Could establish database connection")
	}

	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal("Unable to closed database connection")
		}
	}()

	d := new(models.PersonalDetails)
	// Will need to read the address from another repository
	if row.Next() {
		if err := row.Scan(&d.Address, &d.DateOfBirth, &d.Name); err != nil {
			return nil, errors.Wrap(err, "Error closing the database connection")
		}
	}

	return d, nil
}

// GetByProfileID returns user personaldetails given profile ID
func (m *mySQLPersonalDetailsRepo) GetByProfileID(ctx context.Context, profileID int64) (*models.PersonalDetails, error) {
	return nil, nil
}
