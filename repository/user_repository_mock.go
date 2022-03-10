package repository

import (
	"context"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
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

func (r *UserRepositoryMock) InsertAddress(ctx context.Context, address table.UserAddress) (table.UserAddress, error) {
	args := r.Mock.Called(ctx, address)

	if args.Get(0) == nil {
		return address, nil
	} else {
		address = args.Get(0).(table.UserAddress)
		return address, nil
	}
}
