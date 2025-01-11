package controllers

import (
	"gin-template/internal/models"
	"gin-template/lib/schema"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// paginator.PaginatedResult(models.DB.Model(&models.User{}))
	paginator := schema.NewPaginate(limit, page)
	metaPage, query := paginator.PaginatedResult(models.DB.Model(&models.User{}))
	query.Find(&users)

	meta := &schema.Meta{
		Status: 200,
		Message: "success",
		Page: metaPage.Page,
		Limit: metaPage.Limit,
		NextPage: metaPage.NextPage,
	}
	schema.NewResponse(c, users, meta )

}
