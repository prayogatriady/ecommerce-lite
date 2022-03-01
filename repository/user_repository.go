package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/prayogatriady/ecommerce-lite/model/table"
)

type UserRepositoryInterface interface {
	FindAll() ([]table.UserDummy, error)
	SelectUsers(ctx context.Context, tx *sql.Tx) ([]table.User, error)

	// Insert(ctx context.Context, tx *sql.Tx, user table.User) table.User
	// Update(ctx context.Context, tx *sql.Tx, user table.User) table.User
	// Delete(ctx context.Context, tx *sql.Tx, user table.User)
	// FindByUserID(ctx context.Context, tx *sql.Tx, user table.User) (table.User, error)
	// FindAll(ctx context.Context, tx *sql.Tx) []table.User
	// FindByUserIDPassword(ctx context.Context, tx *sql.Tx, user table.User) (table.User, error)
}

type UserRepository struct {
}

// func NewUserRepository() UserRepository {
// 	return &UserRepositoryImpl{}
// }

// func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, user table.User) table.User {
// 	query := "insert into users (user_id, full_name, password, phone, email, isCustomer) VALUES (?,?,?,?)"
// 	_, err := tx.ExecContext(ctx, query, user.UserID, user.FullName, user.Password, user.Email, user.Email, user.IsCustomer)
// 	if err != nil {
// 		log.Printf("Insert: %s\n", err)
// 	}

// 	return user
// }

func (r *UserRepository) FindAll() ([]table.UserDummy, error) {
	users := []table.UserDummy{
		{UserID: "dobow", Password: "pass"},
		{UserID: "dobowski", Password: "pess"},
	}

	return users, nil
}

func (r *UserRepository) SelectUsers(ctx context.Context, tx *sql.Tx) ([]table.User, error) {
	query := `SELECT user_id, full_name, password, group_user, balance, phone, email, 
	 		isCustomer, isSeller, isShipper, created_at, updated_at FROM users`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[UserRepository][SelectUsers]: %s\n", err)
	}
	defer rows.Close()

	var users []table.User
	var user table.User
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.FullName, &user.Password, &user.GroupUser, &user.Balance, &user.Phone, &user.Email,
			&user.IsCustomer, &user.IsSeller, &user.IsShipper, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("[UserRepository][SelectUsers][Scan]: %s\n", err)
		}

		users = append(users, user)
	}

	return users, err
}
