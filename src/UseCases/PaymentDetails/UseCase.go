package paymentdetails

import (
	repositories "github.com/mathmed/challenge-bw-go/src/Infra/Database/Repositories"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails/Dtos"
)

// UseCase .
func UseCase(paymentID uint) (*dtos.PaymentDetailsResponse, error) {
	return repositories.GetPayment(paymentID), nil
}
