package identify

import (
	"github.com/autom8ter/identify/options"
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