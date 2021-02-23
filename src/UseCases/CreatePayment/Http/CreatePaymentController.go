package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	createpayment "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment"
	dtos "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Dtos"
)

// CreatePaymentController .
func CreatePaymentController(c *gin.Context) {

	var createPaymentDto dtos.CreatePaymentDto

	if err := c.BindJSON(&createPaymentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if CreatePaymentValidator(createPaymentDto) == false {
		return
	}

	createpayment.CreatePaymentUseCase(createPaymentDto)

}