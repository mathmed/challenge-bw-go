package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createpayment "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// CreatePaymentController .
func CreatePaymentController(c *gin.Context) {

	var createPaymentDto dtos.CreatePayment

	if err := c.BindJSON(&createPaymentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{"message": err.Error(), "code": 400},
		})
		return
	}

	if CreatePaymentValidator(createPaymentDto) == false {
		return
	}

	paymentData, err := createpayment.UseCase(createPaymentDto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{"message": err.Error(), "code": 400},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{"message": "ok", "code": 200},
		"data": gin.H{
			"payment_id" : paymentData.PaymentID,
			"gateway_hash" : paymentData.GatewayHash,
			"amount" : paymentData.Amount,
			"parcels" : paymentData.Parcels,
		},
	})

	return
}