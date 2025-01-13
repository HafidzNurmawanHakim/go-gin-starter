package routes

import (
	"gin-template/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.SignUp)
	r.POST("/refresh-token", controllers.RefreshToken)
}
