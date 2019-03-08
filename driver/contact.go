package driver

import (
	"context"
	"github.com/volatiletech/authboss"
)

type Mailer func(context.Context, authboss.Email) error

func (m Mailer) Send(ctx context.Context, e authboss.Email) error {
	return m(ctx, e)
}
