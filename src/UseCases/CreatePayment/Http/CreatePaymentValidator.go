package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// CreatePaymentValidator .
func CreatePaymentValidator(payload dtos.CreatePaymentDto) bool {

	var validate *validator.Validate
	validate = validator.New()
	err := validate.Struct(payload)

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return false
		}
	}
	
	return true
}