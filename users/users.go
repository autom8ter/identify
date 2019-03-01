package users

import (
	"gopkg.in/authboss.v1"
	"time"
)

type User struct {
	ID           int `json:"id"`
	UID     string `json:"uid"`
	FullName     string `json:"full_name"`
	Phone	string `json:"phone"`
	Location string `json:"location"`
	Subscription string `json:"subscription"`
	CreatedAt    string `json:"created_at"`
	LastLogin    string `json:"last_login"`
	Suspended    bool `json:"suspended"`
	// Auth
	Email    string `json:"email"`
	Password string `json:"password"`
	Roles    []string `json:"roles"`

	// Confirm
	ConfirmToken string `json:"confirmed_token"`
	Confirmed    bool `json:"confirmed"`

	// Lock
	AttemptNumber int64 `json:"attempt_number"`
	AttemptTime   time.Time `json:"attempt_time"`
	Locked        time.Time `json:"locked"`

	// Recover
	RecoverToken       string `json:"recover_token"`
	RecoverTokenExpiry time.Time `json:"recover_token_expiry"`

	Data interface{}
}

var nextUserID int

type MemStorer struct {
	Users  map[string]User
	Tokens map[string][]string
}

func NewMemStorer() *MemStorer {
	return &MemStorer{
		Users: map[string]User{
			"testperson@gmail.com": User{
				ID:        1,
				UID:  "TestPerson",
				Password:  "$2a$10$XtW/BrS5HeYIuOCXYe8DFuInetDMdaarMUJEOg/VA/JAIDgw3l4aG", // pass = 1234
				Email:     "TestPerson@gmail.com",
				Confirmed: true,
			},
		},
		Tokens: make(map[string][]string),
	}
}

func (s MemStorer) Create(key string, attr authboss.Attributes) error {
	var user User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	user.ID = nextUserID
	nextUserID++

	s.Users[key] = user
	return nil
}

func (s MemStorer) Put(key string, attr authboss.Attributes) error {

	return s.Create(key, attr)
}

func (s MemStorer) Get(key string) (result interface{}, err error) {
	user, ok := s.Users[key]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MemStorer) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return s.Create(uid+provider, attr)
}

func (s MemStorer) GetOAuth(uid, provider string) (result interface{}, err error) {
	user, ok := s.Users[uid+provider]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

func (s MemStorer) AddToken(key, token string) error {
	s.Tokens[key] = append(s.Tokens[key], token)
	return nil
}

func (s MemStorer) DelTokens(key string) error {
	delete(s.Tokens, key)
	return nil
}

func (s MemStorer) UseToken(givenKey, token string) error {
	toks, ok := s.Tokens[givenKey]
	if !ok {
		return authboss.ErrTokenNotFound
	}

	for i, tok := range toks {
		if tok == token {
			toks[i], toks[len(toks)-1] = toks[len(toks)-1], toks[i]
			s.Tokens[givenKey] = toks[:len(toks)-1]
			return nil
		}
	}

	return authboss.ErrTokenNotFound
}

func (s MemStorer) ConfirmUser(tok string) (result interface{}, err error) {

	for _, u := range s.Users {
		if u.ConfirmToken == tok {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}

func (s MemStorer) RecoverUser(rec string) (result interface{}, err error) {
	for _, u := range s.Users {
		if u.RecoverToken == rec {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}
