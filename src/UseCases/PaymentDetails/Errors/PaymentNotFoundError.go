package errors

// PaymentNotFoundError .
type PaymentNotFoundError struct {
	s string
}

func (e *PaymentNotFoundError) Error() string {
	return "Payment not found"
}