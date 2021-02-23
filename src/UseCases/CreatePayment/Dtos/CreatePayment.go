package dtos

// CreatePayment .
type CreatePayment struct {
	Card *Card
	User *User
	Device *Device
}