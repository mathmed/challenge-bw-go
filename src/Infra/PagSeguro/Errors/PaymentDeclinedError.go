package errors

// PaymentDeclinedError .
type PaymentDeclinedError struct {
	s string
}

func (e *PaymentDeclinedError) Error() string {
	return "Payment declined - PagSeguro"
}