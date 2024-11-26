package routes

import (
	"elogistics-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/orders", controllers.CreateOrder)
}
