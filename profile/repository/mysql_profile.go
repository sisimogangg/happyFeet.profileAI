package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"log"
	"time"

	"github.com/sisimogangg/happyFeet.profileAI/models"
	"github.com/sisimogangg/happyFeet.profileAI/profile"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlProfileRepository struct {
	Conn *sql.DB
}

// NewMysqlProfileRepository Factory method
func NewMysqlProfileRepository(conn *sql.DB) profile.Repository {
	return &mysqlProfileRepository{conn}
}

func (m *mysqlProfileRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Profile, error) {
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

	result := make([]*models.Profile, 0)
	for rows.Next() {
		t := new(models.Profile)
		sgb := 0
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&sgb,
		)

		if err != nil {
			return nil, err
		}

		if sgb == 0 {
			t.SGBMember = false
		} else {
			t.SGBMember = true
		}

		result = append(result, t)

	}

	return result, nil
}

// Fetch fetches all the current profiles
func (m *mysqlProfileRepository) Fetch(ctx context.Context, num int64) ([]*models.Profile, error) {
	query := `SELECT id,user_id,sgb_member FROM profile LIMIT ? `

	p, err := m.fetch(ctx, query, num)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (m *mysqlProfileRepository) GetByID(ctx context.Context, id string) (res *models.Profile, err error) {
	return nil, nil
}

func (m *mysqlProfileRepository) Create(ctx context.Context, p *models.Profile) error {
	return nil
}

func (m *mysqlProfileRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (m *mysqlProfileRepository) Update(ctx context.Context, p *models.Profile) error {
	return nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
