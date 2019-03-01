package users

import (
	"time"
)

type User struct {
	ID           int    `json:"id"`
	UID          string `json:"uid"`
	FullName     string `json:"full_name"`
	Phone        string `json:"phone"`
	Location     string `json:"location"`
	Subscription string `json:"subscription"`
	CreatedAt    string `json:"created_at"`
	LastLogin    string `json:"last_login"`
	Suspended    bool   `json:"suspended"`
	// Auth
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`

	// Confirm
	ConfirmToken string `json:"confirmed_token"`
	Confirmed    bool   `json:"confirmed"`

	// Lock
	AttemptNumber int64     `json:"attempt_number"`
	AttemptTime   time.Time `json:"attempt_time"`
	Locked        time.Time `json:"locked"`

	// Recover
	RecoverToken       string    `json:"recover_token"`
	RecoverTokenExpiry time.Time `json:"recover_token_expiry"`

	Data interface{}
}
