package table

import "time"

type User struct {
	UserID        string        `json:"user_id"`
	FullName      string        `json:"full_name"`
	Password      string        `json:"password"`
	GroupUser     string        `json:"group_user"`
	Balance       int           `json:"balance"`
	Phone         string        `json:"phone"`
	Email         string        `json:"email"`
	IsCustomer    string        `json:"is_customer"`
	IsSeller      string        `json:"is_seller"`
	IsShipper     string        `json:"is_shipper"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	AddressDetail []UserAddress `json:"address_detail"`
}
