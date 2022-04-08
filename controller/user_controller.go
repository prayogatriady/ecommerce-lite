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

func (uc *UserController) Profile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userID := c.Param("userid")

	var user table.User
	user, err := uc.Service.FindUserProfile(ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "400 - BAD REQUEST",
			"message": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": user,
	})
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
			"message": err.Error()},
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
		"message": result,
	})
}

func (uc *UserController) FindAllUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []table.User
	users, err := uc.Service.FindUsers(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "500 - INTERNAL SERVER ERROR",
			"message": err.Error()},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200 - OK",
		"message": users,
	})
}
