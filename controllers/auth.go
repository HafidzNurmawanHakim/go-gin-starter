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
		schema.NewResponse(c, nil, err.Error(), 500)
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		schema.NewResponse(c, nil, "User does not exist", 404)
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		schema.NewResponse(c, nil, "Invalid password", 400)
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
		schema.NewResponse(c, nil, "Could not generate token", 500)
		return
	}

	schema.NewResponse(c, gin.H{"user": existingUser, "token" : tokenString}, "success", 200)

}


func SignUp (c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		schema.NewResponse(c, nil, err.Error(), 500)
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		schema.NewResponse(c, nil, "User already exist", 422)
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		schema.NewResponse(c, nil, "Could not generate password hash", 500)
		return
	}

	models.DB.Create(&user)
	schema.NewResponse(c, gin.H{"userId" : user.ID}, "success", 201)

}
