package errors

// ConnectToPagseguroError .
type ConnectToPagseguroError struct {
	s string
}

func (e *ConnectToPagseguroError) Error() string {
	return "Error to connect to PagSeguro API"
}