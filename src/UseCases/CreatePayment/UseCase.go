package createpayment

import (
	"errors"

	repositories "github.com/mathmed/challenge-bw-go/src/Infra/Database/Repositories"
	pagseguro "github.com/mathmed/challenge-bw-go/src/Infra/PagSeguro"
	custom_errors "github.com/mathmed/challenge-bw-go/src/Infra/PagSeguro/Errors"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// UseCase .
func UseCase(payment dtos.CreatePayment) error {

	var paymentStatus uint
	hash, err := pagseguro.CreatePaymentPagSeguro(payment)
	
	if err != nil {
		if (!errors.Is(err, &custom_errors.PaymentDeclinedError{})) {
			return err
		}
		paymentStatus = 2
	}
	paymentStatus = 1
	deviceID := repositories.SaveDevice(*payment.Device)
	userID := repositories.SaveUser(*payment.User)
	repositories.SavePayment(*payment.Card, hash, paymentStatus, userID, deviceID)
	return nil
}
