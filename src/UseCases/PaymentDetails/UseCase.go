package paymentdetails

import (
	repositories "github.com/mathmed/challenge-bw-go/src/Infra/Database/Repositories"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails/Dtos"
	errors "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails/Errors"
)

// UseCase .
func UseCase(paymentID uint) (*dtos.PaymentDetailsResponse, error) {
	paymentData := repositories.GetPayment(paymentID) 
	if paymentData.ID == 0 {
		return nil, &errors.PaymentNotFoundError{}
	}
	return paymentData, nil
}
