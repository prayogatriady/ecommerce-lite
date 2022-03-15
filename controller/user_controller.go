package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/service"
)

type UserControllerInterface interface {
	FindAllUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	Signup(c *gin.Context)
	Profile(c *gin.Context)
	EditProfile(c *gin.Context)
	DeleteUser(c *gin.Context)

	CreateFeed(c *gin.Context)
}

type UserController struct {
	Service service.UserServiceInterface
}

func (uc *UserController) Signup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userRequest table.User

	if err := c.BindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	result, err := uc.Service.FindUserByUserID(ctx, userRequest.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "An error occured while checking for the email"},
		)
		return
	}

	if result.UserID != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": "Email already used"},
		)
		return
	}

	if _, err := uc.Service.SignUp(ctx, userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": "User created",
	})
}