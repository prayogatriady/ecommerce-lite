package table

import "time"

type Payment struct {
	PaymentID   int
	PaymentName string
	IsActive    string
	CreatedBy   string
	UpdatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
