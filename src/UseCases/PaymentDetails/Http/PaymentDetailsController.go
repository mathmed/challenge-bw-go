package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	paymentdetails "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails"
)

// PaymentDetailsController .
func PaymentDetailsController(c *gin.Context) {
	
	paymentID, err := strconv.ParseUint(c.Param("id"), 10, 64) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{"message": err.Error(), "code": 400},
		})
		return
	}

	paymentData, err := paymentdetails.UseCase(uint(paymentID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{"message": err.Error(), "code": 400},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{"message": "ok", "code": 200},
		"data": gin.H{
			"payment_id" : paymentData.ID,
			"gateway_hash" : paymentData.GatewayValidator,
			"amount" : paymentData.Amount,
			"parcels" : paymentData.Parcels,
			"status" : paymentData.Status,
			"card" : gin.H{
				"initial_number": paymentData.InitialCardNumber,
				"flag": paymentData.CardFlag,
			},
			"user" : gin.H{
				"name": paymentData.Name,
				"email": paymentData.Email,
			},
		},
	})

	return
}