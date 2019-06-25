package models

import "time"

type PersonalDetails struct {
	Name string `json:"name"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Address Address `json:"Address"`
}