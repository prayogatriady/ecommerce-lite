package service

import (
	"context"
	"log"
	"strings"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
)

type UserServiceInterface interface {
	GetAll() ([]table.UserDummy, error)
	FindUsers(ctx context.Context) ([]table.User, error)
	FindUserByUserID(ctx context.Context, userID string) (table.User, error)
	SignUp(ctx context.Context, user table.User) (table.User, error)
	EditProfile(ctx context.Context, user table.User) (table.User, error)
	RemoveUser(ctx context.Context, userID string) error
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

func (s *UserService) FindUserByUserID(ctx context.Context, userID string) (table.User, error) {

	user, err := s.Repository.SelectUserByUserID(ctx, userID)
	if err != nil {
		log.Printf("[UserService][FindUserByUserID][SelectUserByUserID]: %s\n", err)
	}

	return user, err
}

func (s *UserService) SignUp(ctx context.Context, user table.User) (table.User, error) {

	user.FullName = strings.ToUpper(user.FullName)

	user, err := s.Repository.InsertUser(ctx, user)
	if err != nil {
		log.Printf("[UserService][SignUp][InsertUser]: %s\n", err)
	}

	return user, err
}

func (s *UserService) EditProfile(ctx context.Context, user table.User) (table.User, error) {
	user.FullName = strings.ToUpper(user.FullName)

	user, err := s.Repository.UpdateUser(ctx, user)
	if err != nil {
		log.Printf("[UserService][EditProfile][UpdateUser]: %s\n", err)
	}

	return user, err
}

func (s *UserService) RemoveUser(ctx context.Context, userID string) error {
	err := s.Repository.DeleteUser(ctx, userID)
	if err != nil {
		log.Printf("[UserService][RemoveUser][DeleteUser]: %s\n", err)
	}

	return err
}
