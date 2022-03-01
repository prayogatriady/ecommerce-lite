package table

import "time"

type BalanceHist struct {
	UserID     string
	PaymentID  int
	Amount     int
	Created_at time.Time
}
