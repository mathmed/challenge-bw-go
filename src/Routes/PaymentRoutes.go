package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mathmed/challenge-bw-go/src/UseCases/CreatePayment/Http"
)

func PaymentRoutes(route *gin.RouterGroup) {
	group := route.Group("/payment")
	group.POST("/", Http.CreatePaymentController)
}