package repository

import (
	"context"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) FindAll() ([]table.UserDummy, error) {
	args := r.Mock.Called()
	// users := []table.UserDummy{
	// 	{UserID: "mock", Password: "*****"},
	// }

	users := args.Get(0).([]table.UserDummy)
	return users, args.Error(1)
}

func (r *UserRepositoryMock) SelectUsers(ctx context.Context) ([]table.User, error) {
	args := r.Mock.Called(ctx)

	// users := args.Get(0).([]table.User)
	// return users, nil

	if args.Get(0) == nil {
		return nil, nil
	} else {
		users := args.Get(0).([]table.User)
		return users, nil
	}
}

func (r *UserRepositoryMock) SelectUserByUserID(ctx context.Context, userID string) (table.User, error) {
	args := r.Mock.Called(ctx, userID)

	var user table.User
	if args.Get(0) == nil {
		return user, nil
	} else {
		user = args.Get(0).(table.User)
		return user, nil
	}
}

func (r *UserRepositoryMock) SelectByUserIDPassword(ctx context.Context, userID string, password string) (table.User, error) {
	args := r.Mock.Called(ctx, userID, password)

	var user table.User
	if args.Get(0) == nil {
		return user, nil
	} else {
		user = args.Get(0).(table.User)
		return user, nil
	}
}

func (r *UserRepositoryMock) InsertUser(ctx context.Context, user table.User) (table.User, error) {
	args := r.Mock.Called(ctx, user)

	if args.Get(0) == nil {
		return user, nil
	} else {
		user = args.Get(0).(table.User)
		return user, nil
	}
}

func (r *UserRepositoryMock) UpdateUser(ctx context.Context, user table.User) (table.User, error) {
	args := r.Mock.Called(ctx, user)

	if args.Get(0) == nil {
		return user, nil
	} else {
		user = args.Get(0).(table.User)
		return user, nil
	}
}

func (r *UserRepositoryMock) DeleteUser(ctx context.Context, userID string) error {
	r.Mock.Called(ctx, userID)

	return nil
}
