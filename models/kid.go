package models

import "time"

// Kid stores users kids details
type Kid struct {
	Name        string    `json:"name"`
	Grade       int32     `json:"grade"`
	DateStarted time.Time `json:"dateStarted"`
}
