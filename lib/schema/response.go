package schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{}  `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewResponse(c *gin.Context, data any,  message string, status int) {
	if status == 0 {
		status = http.StatusOK
	}
	c.JSON(status, Response{
		Data: data,
		Meta: Meta{
			Status: status,
			Message: message,
		},
	})
}
