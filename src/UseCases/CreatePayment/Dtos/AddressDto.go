package dtos

// AddressDto .
type AddressDto struct {
	Street string `json:"street" binding:"required" validate:"required"`
	Neighborhood string `json:"neighborhood" binding:"required" validate:"required"`
	Number string `json:"number" binding:"required" validate:"required"`
	City string `json:"city" binding:"required" validate:"required"`
	AdministrativeArea string `json:"administrativeArea" binding:"required" validate:"required"`
	PostalCode string `json:"postalCode" binding:"required" validate:"required"`
	Complement string `json:"complement" binding:"required" validate:"required"`
}