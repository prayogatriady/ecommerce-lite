package repository

import (
	"context"
	"log"

	"github.com/prayogatriady/ecommerce-lite/model/table"
)

type PaymentRepositoryInterface interface {

	// Top Up
	UpdateBalance(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error)
	InsertHistBalance(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error)

	// Master Payment
	InsertPayment(ctx context.Context, payment table.Payment) (table.Payment, error)
	SelectPayments(ctx context.Context) ([]table.Payment, error)
	UpdatePayment(ctx context.Context, payment table.Payment) (table.Payment, error)
	DeletePayment(ctx context.Context, paymentID int) error
}

type PaymentRepository struct {
}

func (p *PaymentRepository) UpdateBalance(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][UpdateBalance][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `UPDATE users SET balance = IFNULL(balance, 0) + ? WHERE user_id = ?`

	_, err = tx.ExecContext(ctx, query, topup.Amount, topup.UserID)
	if err != nil {
		log.Printf("[PaymentRepository][UpdateBalance][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[PaymentRepository][UpdateBalance][Commit]: %s\n", err)
	}

	return topup, err
}

func (p *PaymentRepository) InsertHistBalance(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][InsertHistBalance][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO hist_balance (user_id, payment_id, amount) VALUES (?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, topup.UserID, topup.PaymentID, topup.Amount)
	if err != nil {
		log.Printf("[PaymentRepository][InsertHistBalance][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[PaymentRepository][InsertHistBalance][Commit]: %s\n", err)
	}

	return topup, err
}

func (p *PaymentRepository) InsertPayment(ctx context.Context, payment table.Payment) (table.Payment, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][InsertPayment][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `INSERT INTO payments (payment_name, isActive, created_by) VALUES (?, ?, ?)`
	_, err = tx.ExecContext(ctx, query, payment.PaymentName, payment.IsActive, payment.CreatedBy)
	if err != nil {
		log.Printf("[PaymentRepository][InsertPayment][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[PaymentRepository][InsertPayment][Commit]: %s\n", err)
	}

	return payment, err
}

func (p *PaymentRepository) SelectPayments(ctx context.Context) ([]table.Payment, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][SelectPayments][Begin]: %s\n", err)
	}

	query := `SELECT payment_id, payment_name, isActive, created_by, updated_by, created_at, updated_at FROM payments`
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		log.Printf("[PaymentRepository][SelectPayments][QueryContext]: %s\n", err)
	}
	defer rows.Close()

	var payments []table.Payment
	var payment table.Payment
	for rows.Next() {
		err := rows.Scan(&payment.PaymentID, &payment.PaymentName, &payment.IsActive, &payment.CreatedBy, &payment.UpdatedBy, &payment.CreatedAt, &payment.UpdatedAt)
		if err != nil {
			log.Printf("[PaymentRepository][SelectPayments][Scan]: %s\n", err)
		}

		payments = append(payments, payment)
	}

	return payments, err
}

func (p *PaymentRepository) UpdatePayment(ctx context.Context, payment table.Payment) (table.Payment, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][UpdatePayment][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `UPDATE payments SET payment_name = ?, isActive = ?, updated_by = ? WHERE payment_id = ?`
	_, err = tx.ExecContext(ctx, query, payment.PaymentName, payment.IsActive, payment.UpdatedBy)
	if err != nil {
		log.Printf("[PaymentRepository][UpdatePayment][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[PaymentRepository][UpdatePayment][Commit]: %s\n", err)
	}

	return payment, err
}

func (p *PaymentRepository) DeletePayment(ctx context.Context, paymentID int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[PaymentRepository][DeletePayment][Begin]: %s\n", err)
	}
	defer tx.Rollback()

	query := `DELETE FROM payments WHERE payment_id = ?`

	_, err = tx.ExecContext(ctx, query, paymentID)
	if err != nil {
		log.Printf("[PaymentRepository][DeletePayment][ExecContext]: %s\n", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("[PaymentRepository][DeletePayment][Commit]: %s\n", err)
	}

	return err
}
