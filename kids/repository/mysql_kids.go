package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
	"github.com/sisimogangg/happyFeet.profileAI/kids"
	"github.com/sisimogangg/happyFeet.profileAI/models"
)

type mySQLKidsRepository struct {
	Conn *sql.DB
}

// NewMySQLKidsRepository returns a new instance of this repo
func NewMySQLKidsRepository(conn *sql.DB) kids.Repository {
	return &mySQLKidsRepository{Conn: conn}
}

func (m *mySQLKidsRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Kid, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Could not estable database connection")
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	result := make([]*models.Kid, 0)
	for rows.Next() {
		t := new(models.Kid)

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Surname,
			&t.Grade,
			&t.DateOfBirth,
			&t.EnrollmentDate,
		)

		if err != nil {
			return nil, errors.Wrap(err, "Could not serialize kids entity")
		}

		result = append(result, t)

	}

	return result, nil
}

func (m *mySQLKidsRepository) GetByProfileID(ctx context.Context, profileID int64) ([]*models.Kid, error) {
	query := ` SELECT id,name,surname, grade, date_of_birth, enrollment_date
	 			FROM kid
				WHERE profile_id=?
			  `

	kids, err := m.fetch(ctx, query, profileID)
	if err != nil {
		return nil, errors.Wrap(err, "Issues retrieving kids from database")
	}
	return kids, errors.Wrap(err, "Issues retrieving kids from database")
}
