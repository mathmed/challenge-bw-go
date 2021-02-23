package models

import (
	"time"
)

// Device .
type Device struct {
	ID uint
	IP string
	Device string
	Platform string
	Payments []Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}
