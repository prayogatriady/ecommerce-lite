package table

import "time"

type Payment struct {
	PaymentID   int
	PaymentName string
	IsActive    string
	CreatedBy   string
	Created_at  time.Time
	Updated_at  time.Time
}
