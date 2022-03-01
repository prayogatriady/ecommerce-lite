package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prayogatriady/ecommerce-lite/database"
	"github.com/prayogatriady/ecommerce-lite/repository"
	"github.com/prayogatriady/ecommerce-lite/service"
)

func main() {
	db := database.NewDB()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	repo := repository.UserRepository{}
	serv := service.UserService{Repository: &repo, DB: db}
	users, _ := serv.FindUsers(ctx)

	bytes, _ := json.Marshal(users)
	fmt.Println(string(bytes))
}
