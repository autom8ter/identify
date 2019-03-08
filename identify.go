package identify

import (
	"github.com/autom8ter/identify/options"
	"github.com/volatiletech/authboss/defaults"
	"github.com/volatiletech/authboss"
)


func New(opts ...options.Option) *authboss.Authboss {
	ab := authboss.New()
	for _, o := range opts {
		o(ab)
	}
	if err := ab.Init(); err != nil {
		panic(err)
	}
	return ab
}

func NewWithDefaults(readJson, usUsername bool, opts ...options.Option) *authboss.Authboss {
	ab := authboss.New()
	for _, o := range opts {
		o(ab)
	}
	// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
	defaults.SetCore(&ab.Config, readJson, usUsername)
	if err := ab.Init(); err != nil {
		panic(err)
	}
	return ab
}