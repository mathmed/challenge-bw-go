package routes

import (
	"github.com/gin-gonic/gin"
	http "github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Http"
)

// PaymentRoutes .
func PaymentRoutes(route *gin.RouterGroup) {
	group := route.Group("/payment")
	group.POST("/", http.CreatePaymentController)
}