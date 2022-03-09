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
	SelectByUserIDPassword(ctx context.Context, userID string, password string) (table.User, error)
	InsertUser(ctx context.Context, user table.User) (table.User, error)
	UpdateUser(ctx context.Context, user table.User) (table.User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type UserRepository struct {
}

var db *sql.DB = database.NewDB()

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

func (r *UserRepository) SelectByUserIDPassword(ctx context.Context, userID string, password string) (table.User, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][SelectByUserIDPassword][Begin]: %s\n", err)
	}

	query := `SELECT user_id, full_name, password, group_user, balance, phone, email, 
	 		isCustomer, isSeller, isShipper, created_at, updated_at FROM users WHERE user_id = ? AND password = ?`
	rows, err := tx.QueryContext(ctx, query, userID, password)
	if err != nil {
		log.Printf("[UserRepository][SelectByUserIDPassword][QueryContext]: %s\n", err)
	}
	defer rows.Close()

	var user table.User
	if rows.Next() {
		err := rows.Scan(&user.UserID, &user.FullName, &user.Password, &user.GroupUser, &user.Balance, &user.Phone, &user.Email,
			&user.IsCustomer, &user.IsSeller, &user.IsShipper, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Printf("[UserRepository][SelectByUserIDPassword][Scan]: %s\n", err)
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

func (r *UserRepository) UpdateUser(ctx context.Context, user table.User) (table.User, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][UpdateUser][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `UPDATE users SET full_name = ?, password = ?, group_user = ?, balance = ?, phone = ?, email = ?, 
			isCustomer = ?, isSeller = ?, isShipper = ? WHERE user_id = ?`

	_, err = tx.ExecContext(ctx, query, user.FullName, user.Password, user.GroupUser, user.Balance,
		user.Phone, user.Email, user.IsCustomer, user.IsSeller, user.IsShipper, user.UserID)
	if err != nil {
		log.Printf("[UserRepository][UpdateUser][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[UserRepository][UpdateUser][Commit]: %s\n", err)
	}

	return user, err
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[UserRepository][DeleteUser][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `DELETE FROM users WHERE user_id = ?`

	_, err = tx.ExecContext(ctx, query, userID)
	if err != nil {
		log.Printf("[UserRepository][DeleteUser][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[UserRepository][DeleteUser][Commit]: %s\n", err)
	}

	return err
}
