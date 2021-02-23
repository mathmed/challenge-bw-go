package dtos

// User .
type User struct {
	Name string `json:"name" binding:"required" validate:"required"`
	Email string `json:"email" binding:"required" validate:"required,email"`
	Document string `json:"document" binding:"required" validate:"required,numeric"`
	BirthDate string `json:"birthDate" binding:"required" validate:"required"`
	Phone string `json:"phone" binding:"required" validate:"required"`
	Address *Address
}