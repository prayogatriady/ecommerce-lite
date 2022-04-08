package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	FindUsers(ctx context.Context) ([]table.User, error)
	FindUserByUserID(ctx context.Context, userID string) (table.User, error)
	FindUserProfile(ctx context.Context, userID string) (table.User, error)
	Login(ctx context.Context, userID string, password string) (table.User, error)
	SignUp(ctx context.Context, user table.User) (table.User, error)
	EditProfile(ctx context.Context, user table.User) (table.User, error)
	RemoveUser(ctx context.Context, userID string) error

	AddAddress(ctx context.Context, address table.UserAddress) (table.User, error)
}

type UserService struct {
	Repository repository.UserRepositoryInterface
}

func (s *UserService) FindUsers(ctx context.Context) ([]table.User, error) {

	user, err := s.Repository.SelectUsers(ctx)
	if err != nil {
		log.Printf("[UserService][FindUsers][SelectUsers]: %s\n", err)
		return user, errors.New("An error occured while finding users")
	}

	return user, nil
}

func (s *UserService) FindUserByUserID(ctx context.Context, userID string) (table.User, error) {

	user, err := s.Repository.SelectUserByUserID(ctx, userID)
	if err != nil {
		log.Printf("[UserService][FindUserByUserID][SelectUserByUserID]: %s\n", err)
		return user, errors.New("An error occured while checking users")
	}
	if user.UserID != "" {
		return user, errors.New("UserID already used")
	}

	return user, nil
}

func (s *UserService) FindUserProfile(ctx context.Context, userID string) (table.User, error) {

	user, err := s.Repository.SelectUserByUserID(ctx, userID)
	if err != nil {
		log.Printf("[UserService][FindUserProfile][SelectUserByUserID]: %s\n", err)
		return user, errors.New("An error occured while getting your profile")
	}
	if user.UserID == "" {
		return user, errors.New("UserID is not found")
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, userID string, password string) (table.User, error) {

	userFound, err := s.Repository.SelectUserByUserID(ctx, userID)
	if err != nil {
		log.Printf("[UserService][Login][SelectUserByUserID]: %s\n", err)
	}

	var user table.User
	// compare between hash password from db and inputted password from user
	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(password))
	if err != nil {
		log.Printf("[UserService][Login][CompareHashAndPassword]: %s\n", err)
	}

	user, err = s.Repository.SelectByUserIDPassword(ctx, userID, userFound.Password)
	if err != nil {
		log.Printf("[UserService][Login][SelectByUserIDPassword]: %s\n", err)
	}

	return user, err
}

func (s *UserService) SignUp(ctx context.Context, user table.User) (table.User, error) {

	user.FullName = strings.ToUpper(user.FullName)

	// encrypting the password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		log.Printf("[UserService][SignUp][GenerateFromPassword]: %s\n", err)
		return user, errors.New("An error occured while checking password")
	}

	user.Password = string(hashPassword)

	user, err = s.Repository.InsertUser(ctx, user)
	if err != nil {
		log.Printf("[UserService][SignUp][InsertUser]: %s\n", err)
		return user, errors.New("An error occured while signing up")
	}

	return user, nil
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

func (s *UserService) AddAddress(ctx context.Context, address table.UserAddress) (table.User, error) {

	userFound, err := s.Repository.SelectUserByUserID(ctx, address.UserID)
	if err != nil {
		log.Printf("[UserService][AddAddress][SelectUserByUserID]: %s\n", err)
	}

	address, err = s.Repository.InsertAddress(ctx, address)
	if err != nil {
		log.Printf("[UserService][AddAddress][InsertAddress]: %s\n", err)
	}

	userFound.AddressDetail = append(userFound.AddressDetail, address)

	return userFound, err
}
