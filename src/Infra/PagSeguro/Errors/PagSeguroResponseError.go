package errors

// PagSeguroResponseError .
type PagSeguroResponseError struct {
	s string
}

func (e *PagSeguroResponseError) Error() string {
	return "Error to receive to PagSeguro response"
}