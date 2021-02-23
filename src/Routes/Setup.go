package Routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(route *gin.Engine) {
	api := route.Group("/api")
	PaymentRoutes(api)
}