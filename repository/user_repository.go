package repository

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prayogatriady/ecommerce-lite/database"
	"github.com/prayogatriady/ecommerce-lite/model/table"
)

type UserRepositoryInterface interface {
	FindAll() ([]table.UserDummy, error)
	SelectUsers(ctx context.Context) ([]table.User, error)
	SelectUserByUserID(ctx context.Context, userID string) (table.User, error)
	InsertUser(ctx context.Context, user table.User) (table.User, error)

	// Insert(ctx context.Context, user table.User) table.User
	// Update(ctx context.Context, user table.User) table.User
	// Delete(ctx context.Context, user table.User)
	// FindByUserID(ctx context.Context, user table.User) (table.User, error)
	// FindAll(ctx context.Context, tx *sql.Tx) []table.User
	// FindByUserIDPassword(ctx context.Context, user table.User) (table.User, error)
}

type UserRepository struct {
}

var db *sql.DB = database.NewDB()

// func NewUserRepository() UserRepository {
// 	return &UserRepositoryImpl{}
// }

// func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, user table.User) table.User {
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

func (r *UserRepository) SelectUsers(ctx context.Context) ([]table.User, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][SelectUsers][Begin]: %s\n", err)
	}

	query := `SELECT user_id, full_name, password, group_user, balance, phone, email, 
	 		isCustomer, isSeller, isShipper, created_at, updated_at FROM users`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[UserRepository][SelectUsers][QueryContext]: %s\n", err)
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

func (r *UserRepository) SelectUserByUserID(ctx context.Context, userID string) (table.User, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][SelectUserByUserID][Begin]: %s\n", err)
	}

	query := `SELECT user_id, full_name, password, group_user, balance, phone, email, 
	 		isCustomer, isSeller, isShipper, created_at, updated_at FROM users WHERE user_id = ?`
	rows, err := tx.QueryContext(ctx, query, userID)
	if err != nil {
		log.Printf("[UserRepository][SelectUserByUserID][QueryContext]: %s\n", err)
	}
	defer rows.Close()

	var user table.User
	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.FullName, &user.Password, &user.GroupUser, &user.Balance, &user.Phone, &user.Email,
			&user.IsCustomer, &user.IsSeller, &user.IsShipper, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("[UserRepository][SelectUserByUserID][Scan]: %s\n", err)
		}
	}

	return user, err
}

func (r *UserRepository) InsertUser(ctx context.Context, user table.User) (table.User, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][InsertUser][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO users (user_id, full_name, password, group_user, balance, phone, email, 
	 		isCustomer, isSeller, isShipper) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, user.UserID, user.FullName, user.Password, user.GroupUser, user.Balance,
		user.Phone, user.Email, user.IsCustomer, user.IsSeller, user.IsShipper)
	if err != nil {
		log.Printf("[UserRepository][InsertUser][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[UserRepository][InsertUser][Commit]: %s\n", err)
	}

	return user, err
}
