package models

import "time"

type PersonalDetails struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     Address   `json:"Address"`
}

const dateFormat = "2006-01-02"

// SetDate sets dateofbirth to the given date string
func (p *PersonalDetails) SetDateOfBirth(date string) error {
	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	p.DateOfBirth = t
	return nil
}
