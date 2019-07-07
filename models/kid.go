package models

import "time"

// Kid stores users kids details
type Kid struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Grade          string    `json:"grade"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	EnrollmentDate time.Time `json:"enrollmentDate"`
}
