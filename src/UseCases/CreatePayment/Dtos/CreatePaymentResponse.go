package dtos

// CreatePaymentResponse .
type CreatePaymentResponse struct {
	PaymentID uint
	GatewayHash string
	Amount float64
	Parcels int
}