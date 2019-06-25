package models

import "time"

type Kid struct{
	Name string `json:"name"`
	Grade int32 `json:"grade"`
	DateStarted time.Time `json:"dateStarted"`
}