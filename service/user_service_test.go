package service

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetUser(t *testing.T) {
	repository := repository.UserRepositoryMock{}
	repository.Mock.On("FindAll").Return([]table.UserDummy{}, nil)

	service := UserService{Repository: &repository}
	users, _ := service.GetAll()
	for i := range users {
		fmt.Println(users[i].Password)
		assert.Equal(t, users[i].Password, "*****", "user password must be encrypted")
	}
	fmt.Println(users)
}

func TestService_SelectUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	users := []table.User{
		{UserID: "dobow", Password: "dobow", Email: "dobow@gmail.com"},
		{UserID: "jisoo", Password: "jisoo", Email: "jisoo@gmail.com"},
	}

	repository := &repository.UserRepositoryMock{Mock: mock.Mock{}}
	repository.Mock.On("SelectUsers", ctx, repository.Mock).Return(users, nil)

	service := UserService{Repository: repository}
	users, err := service.FindUsers(ctx)
	if err != nil {
		log.Printf("[TestService_SelectUsers][FindUsers]: %s\n", err)
	}
	fmt.Println(users)
}
