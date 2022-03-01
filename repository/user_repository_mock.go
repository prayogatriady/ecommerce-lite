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
