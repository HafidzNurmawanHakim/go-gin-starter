package controllers

import (
	"gin-template/internal/models"
	"strconv"
	"time"

	"gin-template/lib/schema"
	"gin-template/lib/utils"

	"github.com/dgrijalva/jwt-go"
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

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			Subject: strconv.Itoa(existingUser.ID),
			ExpiresAt: expirationTime.Unix(),
		},
		Email: existingUser.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		meta = &schema.Meta{
			Status: 500,
			Message: "Could not generate token",
		}
		schema.NewResponse(c, nil, meta)
		return
	}

	meta = &schema.Meta{
		Status: 200,
		Message: "success",
	}

	schema.NewResponse(c, gin.H{"user": existingUser, "token" : tokenString}, meta)

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

	models.DB.Create(&user)
	meta := &schema.Meta{
		Status: 201,
		Message: "success",
	}
	schema.NewResponse(c, gin.H{"userId" : user.ID}, meta)

}
