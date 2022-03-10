package table

import "time"

type UserAddress struct {
	AddressID   int       `json:"address_id"`
	UserID      string    `json:"user_id"`
	AddressLine string    `json:"address_line"`
	PostalCode  string    `json:"postal_code"`
	City        string    `json:"city"`
	Phone       string    `json:"phone"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
