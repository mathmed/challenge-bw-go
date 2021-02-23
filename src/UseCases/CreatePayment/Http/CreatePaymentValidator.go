package Http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

func CreatePaymentValidator(payload Dtos.CreatePaymentDto) bool {

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