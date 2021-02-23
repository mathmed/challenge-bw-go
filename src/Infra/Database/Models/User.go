package models

import (
	"time"
)

// User .
type User struct {
	ID uint
	Name string
	Email string
	Document string
	BirthDate string `gorm:"type:date"`
	Phone string
	Street string
	Number string
	Neighborhood string
	City string
	Area string
	PostalCode string
	Complement string
	CreatedAt time.Time
	UpdatedAt time.Time
	Payments []Payment
}


