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

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}

func TestService_GetUser(t *testing.T) {

	userRepository.Mock.On("FindAll").Return([]table.UserDummy{}, nil)

	service := UserService{Repository: userRepository}
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

	userRepository.Mock.On("SelectUsers", ctx).Return(users, nil)

	service := UserService{Repository: userRepository}
	result, err := service.FindUsers(ctx)
	if err != nil {
		log.Printf("[TestService_SelectUsers][FindUsers]: %s\n", err)
	}

	fmt.Println(users)
	fmt.Println(result)
	assert.Equal(t, users[0].UserID, result[0].UserID)
}

func TestService_SelectUserByUserID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := table.User{
		UserID:   "dobow",
		Password: "dobow",
		Email:    "dobow@gmail.com",
	}

	userRepository.Mock.On("SelectUserByUserID", ctx, "dobow").Return(user, nil)

	service := UserService{Repository: userRepository}
	result, err := service.FindUserByUserID(ctx, "dobow")
	if err != nil {
		log.Printf("[TestService_SelectUserByUserID][FindUserByUserID]: %s\n", err)
	}

	fmt.Println(user)
	fmt.Println(result)
	assert.Equal(t, user.UserID, result.UserID)
}

func TestService_SelectByUserIDPassword(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := table.User{
		UserID:   "dobow",
		Password: "dobow",
		Email:    "dobow@gmail.com",
	}

	userRepository.Mock.On("SelectByUserIDPassword", ctx, "dobow", "dobow").Return(user, nil)

	service := UserService{Repository: userRepository}
	_, err := service.Login(ctx, "dobow", "dobow")
	if err != nil {
		log.Printf("[TestService_SelectByUserIDPassword][Login]: %s\n", err)
	}
}

func TestService_InsertUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := table.User{
		UserID:   "dobow",
		Password: "dobow",
		Email:    "dobow@gmail.com",
	}

	userRepository.Mock.On("InsertUser", ctx, user).Return(user, nil)

	service := UserService{Repository: userRepository}
	result, err := service.SignUp(ctx, user)
	if err != nil {
		log.Printf("[TestService_InsertUser][SignUp]: %s\n", err)
	}

	fmt.Println(user)
	fmt.Println(result)
	assert.Equal(t, user.UserID, result.UserID)
}

func TestService_UpdateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := table.User{
		UserID:   "dobow",
		Password: "dobow",
		Email:    "dobow@gmail.com",
	}

	userRepository.Mock.On("UpdateUser", ctx, user).Return(user, nil)

	service := UserService{Repository: userRepository}
	result, err := service.EditProfile(ctx, user)
	if err != nil {
		log.Printf("[TestService_UpdateUser][EditProfile]: %s\n", err)
	}

	fmt.Println(user)
	fmt.Println(result)
	assert.Equal(t, user.UserID, result.UserID)
}

func TestService_DeleteUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userID := "dobow"

	userRepository.Mock.On("DeleteUser", ctx, userID).Return(nil)

	service := UserService{Repository: userRepository}
	err := service.RemoveUser(ctx, userID)
	if err != nil {
		log.Printf("[TestService_DeleteUser][RemoveUser]: %s\n", err)
	}

	assert.NoError(t, err)
}
