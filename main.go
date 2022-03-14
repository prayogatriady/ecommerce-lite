package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

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

	// user := table.User{
	// 	UserID:     "admin1",
	// 	FullName:   "admin1",
	// 	Password:   "admin1",
	// 	GroupUser:  "ADMIN",
	// 	Balance:    0,
	// 	Phone:      "081234567892",
	// 	Email:      "admin1@gmail.com",
	// 	IsCustomer: "N",
	// 	IsSeller:   "N",
	// 	IsShipper:  "N",
	// }

	// result, _ := serv.SignUp(ctx, user)

	// payments := []table.Payment{
	// 	{PaymentName: "BNI", IsActive: "Y", CreatedBy: "admin1"},
	// 	{PaymentName: "Permata", IsActive: "Y", CreatedBy: "admin1"},
	// 	{PaymentName: "Mandiri", IsActive: "Y", CreatedBy: "admin1"},
	// 	{PaymentName: "BCA", IsActive: "Y", CreatedBy: "admin1"},
	// }

	// for _, i := range payments {
	// 	result, _ := serv.AddPayment(ctx, i, "ADMIN")

	// 	bytes, _ := json.Marshal(result)
	// 	fmt.Println(string(bytes))
	// }

	result, _ := serv.FindUserByUserID(ctx, "admin3")
	bytes, _ := json.Marshal(result)
	fmt.Println(string(bytes))
	fmt.Println(result.UserID)
}
