package models

import (
	"time"
)

type User struct {
	ID        uint
	Password  string
	Email     string
	CreatedAt time.Time
}