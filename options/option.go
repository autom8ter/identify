package options

import (
	"github.com/volatiletech/authboss"
	"github.com/volatiletech/authboss/defaults"
)

type Option func(a *authboss.Authboss)

func Empty() Option {
	return func(a *authboss.Authboss) {
	}
}
func WithDefaults(readJson, useUserName bool) Option {
	return func(a *authboss.Authboss) {
		// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
		defaults.SetCore(&a.Config, readJson, useUserName)
	}
}

func WithViewRenderer(renderer authboss.Renderer) Option {
	return func(a *authboss.Authboss) {
		a.Config.Core.ViewRenderer = renderer
	}
}

func WithMailRenderer(renderer authboss.Renderer) Option {
	return func(a *authboss.Authboss) {
		a.Config.Core.MailRenderer = renderer
	}
}

func WithRootUrl(rootUrl string) Option {
	return func(a *authboss.Authboss) {
		a.Config.Paths.RootURL = rootUrl
	}
}

func WithMount(mountPath string) Option {
	return func(a *authboss.Authboss) {
		a.Config.Paths.Mount = mountPath
	}
}

func WithStorage(strg authboss.ServerStorer) Option {
	return func(a *authboss.Authboss) {
		a.Config.Storage.Server = strg
	}
}

func WithSessionState(staterw authboss.ClientStateReadWriter) Option {
	return func(a *authboss.Authboss) {
		a.Config.Storage.SessionState = staterw
	}
}

func WithCookieState(staterw authboss.ClientStateReadWriter) Option {
	return func(a *authboss.Authboss) {
		a.Config.Storage.CookieState = staterw
	}
}