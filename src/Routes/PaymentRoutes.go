package routes

import (
	"github.com/gin-gonic/gin"
	http_create_payment "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Http"
	http_payment_details "github.com/mathmed/challenge-bw-go/src/UseCases/PaymentDetails/Http"
)

// PaymentRoutes .
func PaymentRoutes(route *gin.RouterGroup) {
	group := route.Group("/payment")

	group.POST("/", http_create_payment.CreatePaymentController)
	group.GET("/:id", http_payment_details.PaymentDetailsController)
}