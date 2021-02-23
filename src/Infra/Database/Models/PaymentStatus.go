package models

import "time"

// PaymentStatus .
type PaymentStatus struct {
	ID uint
	Name string
	Payments []Payment
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name used by User to `profiles`
func (PaymentStatus) TableName() string {
	return "payment_status"
}
