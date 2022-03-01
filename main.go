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
	users, _ := serv.FindUsers(ctx)

	bytes, _ := json.Marshal(users)
	fmt.Println(string(bytes))
}
