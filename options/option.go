package options

import (
	"github.com/volatiletech/authboss"
)

type Option func(a *authboss.Authboss)

func EmptyOption() Option {
	return func(a *authboss.Authboss) {

	}
}