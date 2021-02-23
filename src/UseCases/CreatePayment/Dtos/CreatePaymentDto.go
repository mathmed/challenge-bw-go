package Dtos

type CreatePaymentDto struct {
	Card *CardDto
	User *UserDto
	Device *DeviceDto
}