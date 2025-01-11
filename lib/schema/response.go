package schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{}  `json:"data"`
	Meta *Meta `json:"meta"`
}

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	NextPage int 	`json:"nextPage,omitempty"`
	Page	int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
}

func NewResponse(c *gin.Context, data any,  meta *Meta) {
	if meta.Status == 0 {
		meta.Status = http.StatusOK
	}
	c.JSON(meta.Status, Response{
		Data: data,
		Meta: meta,
	})
}
