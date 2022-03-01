package service

import (
	"context"
	"log"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
)

type UserServiceInterface interface {
	GetAll() ([]table.UserDummy, error)
	FindUsers(ctx context.Context) ([]table.User, error)
}

type UserService struct {
	Repository repository.UserRepositoryInterface
}

func (s *UserService) GetAll() ([]table.UserDummy, error) {
	users, _ := s.Repository.FindAll()
	for i := range users {
		users[i].Password = "*****"
	}
	return users, nil
}

func (s *UserService) FindUsers(ctx context.Context) ([]table.User, error) {

	user, err := s.Repository.SelectUsers(ctx)
	if err != nil {
		log.Printf("[UserService][FindUsers][SelectUsers]: %s\n", err)
	}

	return user, err
}
