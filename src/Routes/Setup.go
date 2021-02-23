package routes

import (
	"github.com/gin-gonic/gin"
)

// Setup .
func Setup(route *gin.Engine) {
	api := route.Group("/api")
	PaymentRoutes(api)
}