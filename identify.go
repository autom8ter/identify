package identify

import (
	"github.com/autom8ter/identify/db"
	"github.com/autom8ter/identify/saml"
	"net/http"
)

type Identifier struct {
	Db          db.Database
	OAuthClient *http.Client
	SamlReq     *saml.AuthorizationRequest
	Handler     http.HandlerFunc
}

func NewIdentifier(db db.Database, OAuthClient *http.Client, samlReq *saml.AuthorizationRequest, handler http.HandlerFunc) *Identifier {
	return &Identifier{Db: db, OAuthClient: OAuthClient, SamlReq: samlReq, Handler: handler}
}
