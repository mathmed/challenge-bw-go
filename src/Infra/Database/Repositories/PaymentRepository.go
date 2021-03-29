package repositories

import (
	"time"

	database "github.com/mathmed/challenge-bw-go/src/Infra/Database"
	models "github.com/mathmed/challenge-bw-go/src/Infra/Database/Models"
	create_payment_dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
	payment_details_dtos "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails/Dtos"
)

// SavePayment .
func SavePayment(
	paymentData create_payment_dtos.Card,
	gatewayValidator string,
	paymentStatus uint,
	userID uint,
	deviceID uint,
) (uint) {
	payment := models.Payment{
		Amount: paymentData.Amount,
		InitialCardNumber: paymentData.Number[0:4],
		CardOwner: paymentData.Owner,
		CardFlag: paymentData.Flag,
		GatewayValidator: gatewayValidator,
		Parcels: paymentData.Parcels,
		PaymentStatusID: paymentStatus,
		UserID: userID,
		DeviceID: deviceID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	database.Instance.Create(&payment)

	return payment.ID
}

// GetPayment .
func GetPayment(paymentID uint)(*payment_details_dtos.PaymentDetailsResponse){
	payment := &payment_details_dtos.PaymentDetailsResponse{}
	database.Instance.Model(&models.Payment{}).Select(
		`
		payments.id,payments.amount,
		payments.initial_card_number,
		payments.card_flag,
		payments.gateway_validator,
		payments.parcels,
		users.name,
		users.email,
		payment_status.name as status
		`,
	).Joins(`
		LEFT JOIN users ON payments.user_id = users.id
		LEFT JOIN payment_status ON payments.payment_status_id = payment_status.id
	`).Where(paymentID).Scan(payment)
	return payment
}