package options

import (
	"fmt"
	"github.com/autom8ter/identify/db"
	"github.com/spf13/viper"
	"github.com/volatiletech/authboss"
	"github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss-renderer"
	"github.com/volatiletech/authboss/defaults"
	aboauth "github.com/volatiletech/authboss/oauth2"
	"github.com/volatiletech/authboss/otp/twofactor"
	"github.com/volatiletech/authboss/otp/twofactor/totp2fa"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"regexp"
)

type Option func(a *authboss.Authboss)

func Empty() Option {
	return func(a *authboss.Authboss) {
	}
}

func Init(v *viper.Viper, cookie abclientstate.CookieStorer, session abclientstate.SessionStorer) Option {
	v.SetDefault("root-url", "http://localhost:8080")
	v.SetDefault("read-json", true)
	v.SetDefault("use-username", true)
	v.SetDefault("issuer", "Autom8ter")
	v.SetDefault("render.override", "views")
	v.SetDefault("render.mount", "/auth")
	return func(a *authboss.Authboss) {
		rootUrl := v.GetString("root-url")
		readJson := v.GetBool("read-json")
		useUserName := v.GetBool("use-username")
		issuer := v.GetString("issuer")
		clientId := v.GetString("oauth.client-id")
		clientSecret := v.GetString("oauth.client-secret")
		renderMount := v.GetString("email.mount")
		renderOverride := v.GetString("email.override")
		a.Config.Paths.RootURL = rootUrl
		if readJson {
			a.Config.Core.ViewRenderer = defaults.JSONRenderer{}
		} else {
			a.Config.Core.ViewRenderer = abrenderer.NewHTML(renderMount, renderOverride)
		}
		a.Config.Modules.TwoFactorEmailAuthRequired = true

		a.Config.Storage.Server = db.NewMemStorer()
		a.Config.Storage.SessionState = session
		a.Config.Storage.CookieState = cookie
		// The preserve fields are things we don't want to
		// lose when we're doing user registration (prevents having
		// to type them again)
		a.Config.Modules.RegisterPreserveFields = []string{"email", "name"}
		// Set up defaults for basically everything besides the ViewRenderer/MailRenderer in the HTTP stack
		// TOTP2FAIssuer is the name of the issuer we use for totp 2fa
		a.Config.Modules.TOTP2FAIssuer = issuer
		a.Config.Core.MailRenderer = abrenderer.NewEmail(renderMount, renderOverride)

		defaults.SetCore(&a.Config, readJson, useUserName)

		// Here we initialize the bodyreader as something customized in order to accept a name
		// parameter for our user as well as the standard e-mail and password.
		//
		// We also change the validation for these fields
		// to be something less secure so that we can use test data easier.
		emailRule := defaults.Rules{
			FieldName: "email", Required: true,
			MatchError: "Must be a valid e-mail address",
			MustMatch:  regexp.MustCompile(`.*@.*\.[a-z]{1,}`),
		}
		passwordRule := defaults.Rules{
			FieldName: "password", Required: true,
			MinLength: 4,
		}
		nameRule := defaults.Rules{
			FieldName: "name", Required: true,
			MinLength: 2,
		}

		a.Config.Core.BodyReader = defaults.HTTPBodyReader{
			ReadJSON: readJson,
			Rulesets: map[string][]defaults.Rules{
				"register":    {emailRule, passwordRule, nameRule},
				"recover_end": {passwordRule},
			},
			Confirms: map[string][]string{
				"register":    {"password", authboss.ConfirmPrefix + "password"},
				"recover_end": {"password", authboss.ConfirmPrefix + "password"},
			},
			Whitelist: map[string][]string{
				"register": []string{"email", "name", "password"},
			},
		}

		// Set up 2fa
		twofaRecovery := &twofactor.Recovery{Authboss: a}
		if err := twofaRecovery.Setup(); err != nil {
			panic(err)
		}

		totp := &totp2fa.TOTP{Authboss: a}
		if err := totp.Setup(); err != nil {
			panic(err)
		}
		if clientId != "" && clientSecret != "" {
			a.Config.Modules.OAuth2Providers = map[string]authboss.OAuth2Provider{
				"google": authboss.OAuth2Provider{
					OAuth2Config: &oauth2.Config{
						ClientID:     clientId,
						ClientSecret: clientSecret,
						Scopes:       []string{`profile`, `email`},
						Endpoint:     google.Endpoint,
					},
					FindUserDetails: aboauth.GoogleUserDetails,
				},
			}
		} else {
			fmt.Println("oauth.client-id and oauth.client-secret not found in config, disabling oauth...")
		}
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
