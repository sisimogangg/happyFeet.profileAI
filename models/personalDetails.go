package models

import "time"
import u "happyFeet.profileAI/utils"

type PersonalDetails struct {
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address     Address   `json:"Address"`
}

const dateFormat = "2006-01-02"

func (p *PersonalDetails) Validate() (map[string]interface{}, bool) {
	if p.Name == "" {
		return u.Message(false, "No Name"), false
	}

	if p.DateOfBirth.IsZero() {
		return u.Message(false, "Provide Date Of Birth"), false
	}

	if resp, ok := p.Address.Validate(); !ok {
		return resp, ok
	}

	return u.Message(false, "Requirements Passed"), true
}

// SetDate sets dateofbirth to the given date string
func (p *PersonalDetails) SetDateOfBirth(date string) error {
	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return err
	}

	p.DateOfBirth = t
	return nil
}
