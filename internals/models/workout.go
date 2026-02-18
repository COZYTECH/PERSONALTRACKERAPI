package models

import (
	"time"
)

type Workout struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Duration  int       `json:"duration"` // in minutes
	Calories  int       `json:"calories"`
	CreatedAt time.Time `json:"created_at"`
}