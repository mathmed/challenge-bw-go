package dtos

// PaymentDetailsResponse .
type PaymentDetailsResponse struct {
	ID uint
	GatewayValidator string
	Amount float64
	Parcels int
	Status string
	InitialCardNumber string
	CardFlag string
	Name string
	Email string
}