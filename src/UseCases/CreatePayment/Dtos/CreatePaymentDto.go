package dtos

// CreatePaymentDto .
type CreatePaymentDto struct {
	Card *CardDto
	User *UserDto
	Device *DeviceDto
}