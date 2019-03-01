package db

import (
	"github.com/autom8ter/identify/users"
	"gopkg.in/authboss.v1"
)

var nextUserID int

type Database interface {
	Create(key string, attr authboss.Attributes) error
	Put(key string, attr authboss.Attributes) error
	Get(key string) (result interface{}, err error)
	PutOAuth(uid, provider string, attr authboss.Attributes) error
	GetOAuth(uid, provider string) (result interface{}, err error)
	AddToken(key, token string) error
	DelTokens(key string) error
	UseToken(givenKey, token string) error
	ConfirmUser(tok string) (result interface{}, err error)
	RecoverUser(rec string) (result interface{}, err error)
	GetUsers() map[string]users.User
	GetTokens() map[string][]string
}

type Db struct {
	Users           map[string]users.User
	Tokens          map[string][]string
	CreateFunc      func(key string, attr authboss.Attributes) error
	PutFunc         func(key string, attr authboss.Attributes) error
	GetFunc         func(key string) (result interface{}, err error)
	PutOAuthFunc    func(uid, provider string, attr authboss.Attributes) error
	GetOAuthFunc    func(uid, provider string) (result interface{}, err error)
	AddTokenFunc    func(key, token string) error
	DelTokensFunc   func(key string) error
	UseTokenFunc    func(givenKey, token string) error
	ConfirmUserFunc func(tok string) (result interface{}, err error)
	RecoverUserFunc func(rec string) (result interface{}, err error)
}

func (d *Db) GetTokens() map[string][]string {
	return d.Tokens
}

func (d *Db) GetUsers() map[string]users.User {
	return d.Users
}

func (d *Db) Create(key string, attr authboss.Attributes) error {
	return d.CreateFunc(key, attr)
}

func (d *Db) Put(key string, attr authboss.Attributes) error {
	return d.PutFunc(key, attr)
}

func (d *Db) Get(key string) (result interface{}, err error) {
	return d.GetFunc(key)
}

func (d *Db) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return d.PutOAuthFunc(uid, provider, attr)
}

func (d *Db) GetOAuth(uid, provider string) (result interface{}, err error) {
	return d.GetOAuthFunc(uid, provider)
}

func (d *Db) AddToken(key, token string) error {
	return d.AddTokenFunc(key, token)
}

func (d *Db) DelTokens(key string) error {
	return d.DelTokens(key)
}

func (d *Db) UseToken(givenKey, token string) error {
	return d.UseTokenFunc(givenKey, token)
}

func (d *Db) ConfirmUser(tok string) (result interface{}, err error) {
	return d.ConfirmUserFunc(tok)
}

func (d *Db) RecoverUser(rec string) (result interface{}, err error) {
	return d.RecoverUserFunc(rec)
}
