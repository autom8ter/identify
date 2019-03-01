package mem

import (
	"github.com/autom8ter/identify/db"
	"github.com/autom8ter/identify/users"
	"gopkg.in/authboss.v1"
)

var nextUserID int

func NewDb() db.Database {
	s := &db.Db{
		Users: map[string]users.User{
			"testperson@gmail.com": users.User{
				ID:        1,
				UID:       "TestPerson",
				Password:  "$2a$10$XtW/BrS5HeYIuOCXYe8DFuInetDMdaarMUJEOg/VA/JAIDgw3l4aG", // pass = 1234
				Email:     "TestPerson@gmail.com",
				Confirmed: true,
			},
		},
		Tokens: make(map[string][]string),
	}
	s.CreateFunc = func(key string, attr authboss.Attributes) error {
		var user users.User
		if err := attr.Bind(&user, true); err != nil {
			return err
		}

		user.ID = nextUserID
		nextUserID++

		s.Users[key] = user
		return nil
	}
	s.PutFunc = func(key string, attr authboss.Attributes) error {
		return s.CreateFunc(key, attr)
	}

	s.GetFunc = func(key string) (result interface{}, err error) {
		user, ok := s.Users[key]
		if !ok {
			return nil, authboss.ErrUserNotFound
		}

		return &user, nil
	}
	s.PutOAuthFunc = func(uid, provider string, attr authboss.Attributes) error {
		return s.CreateFunc(uid+provider, attr)
	}
	s.GetOAuthFunc = func(uid, provider string) (result interface{}, err error) {
		user, ok := s.Users[uid+provider]
		if !ok {
			return nil, authboss.ErrUserNotFound
		}

		return &user, nil
	}
	s.AddTokenFunc = func(key, token string) error {
		s.Tokens[key] = append(s.Tokens[key], token)
		return nil
	}
	s.DelTokensFunc = func(key string) error {
		delete(s.Tokens, key)
		return nil
	}
	s.UseTokenFunc = func(givenKey, token string) error {
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
	s.ConfirmUserFunc = func(tok string) (result interface{}, err error) {

		for _, u := range s.Users {
			if u.ConfirmToken == tok {
				return &u, nil
			}
		}

		return nil, authboss.ErrUserNotFound
	}
	s.RecoverUserFunc = func(rec string) (result interface{}, err error) {
		for _, u := range s.Users {
			if u.RecoverToken == rec {
				return &u, nil
			}
		}

		return nil, authboss.ErrUserNotFound
	}
	return s
}
