package oauth

import (
	"context"
	"fmt"
	"github.com/autom8ter/identify/util"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

/*
const (
    // AuthStyleAutoDetect means to auto-detect which authentication
    // style the provider wants by trying both ways and caching
    // the successful way for the future.
    AuthStyleAutoDetect AuthStyle = 0

    // AuthStyleInParams sends the "client_id" and "client_secret"
    // in the POST body as application/x-www-form-urlencoded parameters.
    AuthStyleInParams AuthStyle = 1

    // AuthStyleInHeader sends the client_id and client_password
    // using HTTP Basic Authorization. This is an optional style
    // described in the OAuth2 RFC 6749 section 2.3.1.
    AuthStyleInHeader AuthStyle = 2
)
*/

var ValidStyles = []string{"AutoDetect", "auto", "autodetect", "Auto", "InParams", "inparams", "params", "Params", "InHeader", "inheader", "Header", "header"}

type OAuth func(cfg *oauth2.Config) *oauth2.Config

func NewOauthFromEnv() OAuth {
	return func(cfg *oauth2.Config) *oauth2.Config {
		var style oauth2.AuthStyle

		switch os.Getenv("OAUTH_STYLE") {
		case "AutoDetect", "auto", "autodetect", "Auto":
			style = oauth2.AuthStyleAutoDetect
		case "InParams", "inparams", "params", "Params":
			style = oauth2.AuthStyleInParams
		case "InHeader", "inheader", "Header", "header":
			style = oauth2.AuthStyleInHeader
		default:
			log.Printf("Unsupported Authentication Style. Please use one of: %s", ValidStyles)

		}
		scopes, err := util.ReadAsCSV(os.Getenv("OAUTH_SCOPES"))
		if err != nil {
			log.Println(err.Error())
		}
		return &oauth2.Config{
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			Endpoint: oauth2.Endpoint{
				AuthURL:   os.Getenv("OAUTH_ENDPOINT_URL"),
				TokenURL:  os.Getenv("OAUTH_ENDPOINT_TOKEN_URL"),
				AuthStyle: style,
			},
			RedirectURL: os.Getenv("OAUTH_REDIRECT_URL"),
			Scopes:      scopes,
		}
	}
}

func (c OAuth) GetClient(ctx context.Context, code string) *http.Client {
	config := c(&oauth2.Config{})
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	return config.Client(ctx, tok)
}
