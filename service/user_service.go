package service

import (
	"context"
	"database/sql"
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
	DB         *sql.DB
}

func (s *UserService) GetAll() ([]table.UserDummy, error) {
	users, _ := s.Repository.FindAll()
	for i := range users {
		users[i].Password = "*****"
	}
	return users, nil
}

func (s *UserService) FindUsers(ctx context.Context) ([]table.User, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Printf("[UserService][FindUsers][Begin]: %s\n", err)
	}

	user, err := s.Repository.SelectUsers(ctx, tx)
	if err != nil {
		log.Printf("[UserService][FindUsers][SelectUsers]: %s\n", err)
	}

	return user, err
}
