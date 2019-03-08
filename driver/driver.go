package driver

import (
	"context"
	"github.com/autom8ter/identify/driver/api"
	"github.com/volatiletech/authboss"
	"net/http"
)

type UnderlyingResponseWriter func() http.ResponseWriter

func (u UnderlyingResponseWriter) UnderlyingResponseWriter() http.ResponseWriter {
	return u()
}

type HTTPResponder func(w http.ResponseWriter, r *http.Request, code int, templateName string, data authboss.HTMLData) error

func (h HTTPResponder) Respond(w http.ResponseWriter, r *http.Request, code int, templateName string, data authboss.HTMLData) error {
	return h.Respond(w, r, code, templateName, data)
}

type HTTPRedirector func(w http.ResponseWriter, r *http.Request, ro authboss.RedirectOptions) error

func (h HTTPRedirector) Redirect(w http.ResponseWriter, r *http.Request, ro authboss.RedirectOptions) error {
	return h(w, r, ro)
}

type BodyReader func(page string, r *http.Request) (api.Validator, error)

func (b BodyReader) Read(page string, r *http.Request) (api.Validator, error) {
	return b(page, r)
}

type ClientState func(key string) (string, bool)

func (c ClientState) Get(key string) (string, bool) {
	return c(key)
}

type Moduler func(*authboss.Authboss) error

func (m Moduler) Init(a *authboss.Authboss) error {
	return m(a)
}

type Renderer struct {

	// Load the given templates, will most likely be called multiple times
	LoadFunc func(names ...string) error

	// Render the given template
	RenderFunc func(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error)
}

func NewRenderer(loadFunc func(names ...string) error, renderFunc func(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error)) *Renderer {
	return &Renderer{LoadFunc: loadFunc, RenderFunc: renderFunc}
}

type ClientStateReadWriter struct {

	// ReadState should return a map like structure allowing it to look up
	// any values in the current session, or any cookie in the request
	ReadStateFunc func(*http.Request) (ClientState, error)
	// WriteState can sometimes be called with a nil ClientState in the event
	// that no ClientState was read in from LoadClientState
	WriteStateFunc func(http.ResponseWriter, ClientState, []authboss.ClientStateEvent) error
}

func NewClientStateReadWriter(readStateFunc func(*http.Request) (ClientState, error), writeStateFunc func(http.ResponseWriter, ClientState, []authboss.ClientStateEvent) error) *ClientStateReadWriter {
	return &ClientStateReadWriter{ReadStateFunc: readStateFunc, WriteStateFunc: writeStateFunc}
}
