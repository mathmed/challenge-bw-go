package createpayment

import (
	"errors"

	repositories "github.com/mathmed/challenge-bw-go/src/Infra/Database/Repositories"
	pagseguro "github.com/mathmed/challenge-bw-go/src/Infra/PagSeguro"
	custom_errors "github.com/mathmed/challenge-bw-go/src/Infra/PagSeguro/Errors"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
	enums "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Enums"
)

// UseCase .
func UseCase(payment dtos.CreatePayment) (*dtos.CreatePaymentResponse, error) {

	var paymentStatus uint
	hash, err := pagseguro.CreatePaymentPagSeguro(payment)
	
	if err != nil {
		if (!errors.Is(err, &custom_errors.PaymentDeclinedError{})) {
			return nil, err
		}
		paymentStatus = enums.DECLINED;
	}

	paymentStatus = enums.RESERVED;
	deviceID := repositories.SaveDevice(*payment.Device)
	userID := repositories.SaveUser(*payment.User)
	paymentID := repositories.SavePayment(*payment.Card, hash, paymentStatus, userID, deviceID)
	return &dtos.CreatePaymentResponse{
		PaymentID: paymentID,
		GatewayHash: hash,
		Amount: payment.Card.Amount,
		Parcels: payment.Card.Parcels,
	}, nil
}
