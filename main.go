package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prayogatriady/ecommerce-lite/controller"
	"github.com/prayogatriady/ecommerce-lite/repository"
	"github.com/prayogatriady/ecommerce-lite/service"
)

var PORT string = "3003"

func main() {

	userRepo := repository.UserRepository{}
	userServ := service.UserService{Repository: &userRepo}
	userCont := controller.UserController{Service: &userServ}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/users", userCont.FindAllUser)
	router.POST("/signup", userCont.Signup)
	router.GET("/profile/:userid", userCont.Profile)
	router.PUT("/edit/:userid", userCont.EditProfile)

	log.Printf("Server running on port %s.. \n", PORT)
	if err := router.Run(":" + PORT); err != nil {
		log.Fatalf("Error when running port %s \n", PORT)
	}
}
