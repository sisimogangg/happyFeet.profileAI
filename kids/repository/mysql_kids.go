package repository

import (
	"context"
	"database/sql"
	"log"

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
		return nil, err
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
			//			&t.DateOfBirth,
			//			&t.EnrollmentDate,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, t)

	}

	return result, nil
}

func (m *mySQLKidsRepository) GetByProfileID(ctx context.Context, profileID int64) ([]*models.Kid, error) {
	query := ` SELECT id,name,surname, grade
	 			FROM kid
				WHERE profile_id=?
			  `

	kids, err := m.fetch(ctx, query, profileID)
	if err != nil {
		return nil, err
	}
	return kids, nil
}
