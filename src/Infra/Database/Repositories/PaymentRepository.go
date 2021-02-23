package repositories

import (
	"time"

	database "github.com/mathmed/challenge-bw-go/src/Infra/Database"
	models "github.com/mathmed/challenge-bw-go/src/Infra/Database/Models"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// SavePayment .
func SavePayment(paymentData dtos.Card, gatewayValidator string, paymentStatus uint) {
	payment := models.Payment{
		Amount: paymentData.Amount,
		InitialCardNumber: paymentData.Number[0:4],
		CardOwner: paymentData.Owner,
		CardFlag: paymentData.Flag,
		GatewayValidator: gatewayValidator,
		Parcels: paymentData.Parcels,
		PaymentStatusID: paymentStatus,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	database.Instance.Create(&payment)
}