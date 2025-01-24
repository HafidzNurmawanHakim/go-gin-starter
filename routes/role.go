package routes

import (
	"gin-template/controllers"
	"gin-template/internal/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func Role(r *gin.Engine) {
	protected := r.Group("/roles")
	protected.Use(middlewares.WithJwtAuth(os.Getenv("SECRET_KEY")))
	{
		protected.GET("", controllers.GetRoles)
	}
}