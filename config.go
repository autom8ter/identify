package identify

import (
	"github.com/autom8ter/identify/driver"
	"github.com/volatiletech/authboss"
)

func NewConfig(opts ...driver.ConfigFunc) *authboss.Config {
	cfg := &authboss.Config{}
	for _, o := range opts {
		o(cfg)
	}
	return cfg
}
