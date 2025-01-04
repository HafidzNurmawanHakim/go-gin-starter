package routes

import (
	"gin-template/controllers"
	"gin-template/internal/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {

	proctected := r.Group("/users")
	proctected.Use(middlewares.WithJwtAuth(os.Getenv("SECRET_KEY")))
	{
		proctected.GET("", controllers.GetUsers)
	}
}
