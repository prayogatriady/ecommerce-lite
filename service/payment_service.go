package service

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/prayogatriady/ecommerce-lite/model/table"
	"github.com/prayogatriady/ecommerce-lite/repository"
)

type PaymentServiceInterface interface {
	Topup(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error)

	AddPayment(ctx context.Context, payment table.Payment, groupUser string) (table.Payment, error)
	FindPayments(ctx context.Context, groupUser string) ([]table.Payment, error)
	EditPayment(ctx context.Context, payment table.Payment, groupUser string) (table.Payment, error)
	DeactivatePayment(ctx context.Context, paymentID int, groupUser string) error
}

type PaymentService struct {
	Repository repository.PaymentRepositoryInterface
}

func (s *PaymentService) Topup(ctx context.Context, topup table.BalanceHist) (table.BalanceHist, error) {

	_, err := s.Repository.UpdateBalance(ctx, topup)
	if err != nil {
		log.Printf("[PaymentService][Topup][UpdateBalance]: %s\n", err)
	}

	_, err = s.Repository.InsertHistBalance(ctx, topup)
	if err != nil {
		log.Printf("[PaymentService][Topup][InsertHistBalance]: %s\n", err)
	}

	return topup, err
}

func (s *PaymentService) AddPayment(ctx context.Context, payment table.Payment, groupUser string) (table.Payment, error) {

	if groupUser != "ADMIN" {
		return payment, errors.New("ADMIN REQUIRED")
	}

	payment.PaymentName = strings.ToUpper(payment.PaymentName)

	payment, err := s.Repository.InsertPayment(ctx, payment)
	if err != nil {
		log.Printf("[PaymentService][AddPayment][InsertPayment]: %s\n", err)
	}

	return payment, err
}

func (s *PaymentService) FindPayments(ctx context.Context, groupUser string) ([]table.Payment, error) {

	var payments []table.Payment

	if groupUser != "ADMIN" {
		return payments, errors.New("ADMIN REQUIRED")
	}

	payments, err := s.Repository.SelectPayments(ctx)
	if err != nil {
		log.Printf("[PaymentService][FindPayments][SelectPayments]: %s\n", err)
	}

	return payments, err
}

func (s *PaymentService) EditPayment(ctx context.Context, payment table.Payment, groupUser string) (table.Payment, error) {

	if groupUser != "ADMIN" {
		return payment, errors.New("ADMIN REQUIRED")
	}

	payment, err := s.Repository.UpdatePayment(ctx, payment)
	if err != nil {
		log.Printf("[PaymentService][EditPayment][UpdatePayment]: %s\n", err)
	}

	return payment, err
}

func (s *PaymentService) DeactivatePayment(ctx context.Context, paymentID int, groupUser string) error {

	if groupUser != "ADMIN" {
		return errors.New("ADMIN REQUIRED")
	}

	err := s.Repository.DeletePayment(ctx, paymentID)
	if err != nil {
		log.Printf("[PaymentService][DeactivatePayment][DeletePayment]: %s\n", err)
	}

	return err
}
