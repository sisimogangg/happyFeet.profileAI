package repository

import (
	"context"
	"database/sql"

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

// GetByProfileID returns user personaldetails given profile ID
func (m *mySQLPersonalDetailsRepo) GetByProfileID(ctx context.Context, profileID int64) (*models.PersonalDetails, error) {
	return nil, nil
}
