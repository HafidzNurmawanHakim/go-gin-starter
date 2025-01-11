package middlewares

import (
	"gin-template/lib/schema"
	"gin-template/lib/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)


func WithJwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("access_token")
		if err != nil {
			meta := &schema.Meta{
				Status: http.StatusUnauthorized,
				Message: "Invalid token",
			}
			schema.NewResponse(c, nil, meta )
			c.Abort()
			return
		}
		userId, err := utils.ExtraxtIdFromToken(cookie, secret)

		if err != nil {
			meta := &schema.Meta{
				Status: http.StatusUnauthorized,
				Message: err.Error(),
			}
			schema.NewResponse(c, nil, meta)
			c.Abort()
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
