package createpayment

import (
	"fmt"

	pagseguro "github.com/mathmed/challenge-bw-go/src/Infra/PagSeguro"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// UseCase .
func UseCase(payment dtos.CreatePayment)  {

	hash := pagseguro.CreatePaymentPagSeguro(payment)
	fmt.Printf(hash)

}
