package controllers

import (
	"gin-template/internal/models"
	"net/http"
	"time"

	"gin-template/lib/schema"
	"gin-template/lib/utils"

	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")


func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		meta := &schema.Meta{
			Status: 500,
			Message: err.Error(),
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)
	meta := &schema.Meta{
		Status: 404,
		Message: "User does not exist",
	}
	if existingUser.ID == 0 {
		schema.NewResponse(c, nil, meta)
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)
	if !errHash {
		meta = &schema.Meta{
			Status: 400,
			Message: "Invalid password",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	tokenExpTime := time.Now().Add(5 * time.Minute)
	refreshExpTime := time.Now().Add(7 * 24 * time.Hour)

	token, err := utils.GenerateToken(&existingUser, string(jwtKey), tokenExpTime)

	if err != nil {
		meta := &schema.Meta{
			Status: http.StatusInternalServerError,
			Message: "Could not generate token",
		}
		schema.NewResponse(c, nil, meta)
		return
	}
	refreshToken, err := utils.GenerateToken(&existingUser, string(jwtKey), refreshExpTime)

	if err != nil {
		meta = &schema.Meta{
			Status: 500,
			Message: "Could not generate refresh token",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	meta = &schema.Meta{
		Status: 200,
		Message: "success",
	}

	schema.NewResponse(c, gin.H{"user": existingUser, "token" : gin.H{"access_token": token, "refresh_token" : refreshToken}}, meta)

}


func SignUp (c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		meta := &schema.Meta{
			Status: 500,
			Message: err.Error(),
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		meta := &schema.Meta{
			Status: 422,
			Message: "User already exist",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		meta := &schema.Meta{
			Status: 500,
			Message: "Could not generate password hash",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	if user.RoleID == 0 {
		var defaultRole models.Role
		if err := models.DB.Where("name = ?", "staff").First(&defaultRole).Error; err != nil {
			meta := &schema.Meta{
				Status: 500,
				Message: err.Error(),
			}
			schema.NewResponse(c, nil, meta)
			return 
		}
		user.RoleID = defaultRole.ID
	}

	models.DB.Create(&user)
	meta := &schema.Meta{
		Status: 201,
		Message: "success",
	}
	schema.NewResponse(c, gin.H{"userId" : user.ID}, meta)

}


func RefreshToken(c *gin.Context) {
	var user models.User
	refresh_token, err := c.Cookie("refresh_token")

	if err != nil {
		meta := &schema.Meta{
			Status: http.StatusUnauthorized,
			Message: "Refresh Token is required!",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	userId, err := utils.ExtraxtIdFromToken(refresh_token, string(jwtKey))

	if err != nil {
		meta := &schema.Meta{
			Status: http.StatusUnauthorized,
			Message: err.Error(),
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	models.DB.Where("ID = ?", userId).First(&user)

	if user.ID == 0 {
		meta := &schema.Meta{
			Status: 404,
			Message: "User not found",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	tokenExpTime := time.Now().Add(5 * time.Minute)
	refreshExpTime := time.Now().Add(7 * 24 * time.Hour)


	token, err := utils.GenerateToken(&user, string(jwtKey), tokenExpTime)

	if err != nil {
		meta := &schema.Meta{
			Status: http.StatusInternalServerError,
			Message: "Could not generate token",
		}
		schema.NewResponse(c, nil, meta)
		return
	}
	refreshToken, err := utils.GenerateToken(&user, string(jwtKey), refreshExpTime)

	if err != nil {
		meta := &schema.Meta{
			Status: http.StatusInternalServerError,
			Message: "Could not generate refresh token",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	meta := &schema.Meta{
		Status: 200,
		Message: "success",
	}

	schema.NewResponse(c, gin.H{ "token": gin.H{ "token" : token, "access_token" : refreshToken}}, meta )

}
