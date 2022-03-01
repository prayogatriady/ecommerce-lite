package table

import "time"

type UserAddress struct {
	AddressID   int
	UserID      string
	AddressLine string
	PostalCode  string
	City        string
	Phone       string
	Created_at  time.Time
	Updated_at  time.Time
}
