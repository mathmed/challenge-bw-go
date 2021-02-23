package models

import (
	"time"
)

// Payment .
type Payment struct {
	ID uint
	Amount float64
	InitialCardNumber string
	CardOwner string
	CardFlag string
	GatewayValidator string
	Parcels int
	PaymentStatusID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
