package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
	"github.com/prayogatriady/ecommerce-lite/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo := repository.UserRepository{}
	serv := service.UserService{Repository: &repo}
	// users, _ := serv.FindUsers(ctx)

	// user, _ := serv.FindUserByUserID(ctx, "dobow")

	user := table.User{
		UserID:     "admin1",
		FullName:   "admin1",
		Password:   "admin1",
		GroupUser:  "ADMIN",
		Balance:    0,
		Phone:      "081234567890",
		Email:      "admin1@gmail.com",
		IsCustomer: "Y",
		IsSeller:   "N",
		IsShipper:  "N",
	}

	result, _ := serv.SignUp(ctx, user)

	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
}
