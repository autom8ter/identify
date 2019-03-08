package api

import (
	"context"
	"github.com/volatiletech/authboss"
	"net/http"
	"time"
)

type User interface {
	GetPID() (pid string)
	PutPID(pid string)
}

type Validator interface {
	// Validate makes the type validate itself and return
	// a list of validation errors.
	Validate() []error
}

type UserValuer interface {
	Validator

	GetPID() string
	GetPassword() string
}

type UnderlyingResponseWriter interface {
	UnderlyingResponseWriter() http.ResponseWriter
}

type ServerStorer interface {
	// Load will look up the user based on the passed the PrimaryID
	Load(ctx context.Context, key string) (User, error)

	// Save persists the user in the database, this should never
	// create a user and instead return ErrUserNotFound if the user
	// does not exist.
	Save(ctx context.Context, user User) error
}

type Router interface {
	http.Handler

	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Delete(path string, handler http.Handler)
}

type RequestLogger interface {
	FromRequest(*http.Request) Logger
}

type Renderer interface {
	// Load the given templates, will most likely be called multiple times
	Load(names ...string) error

	// Render the given template
	Render(ctx context.Context, page string, data authboss.HTMLData) (output []byte, contentType string, err error)
}

type RememberingServerStorer interface {
	ServerStorer

	// AddRememberToken to a user
	AddRememberToken(ctx context.Context, pid, token string) error
	// DelRememberTokens removes all tokens for the given pid
	DelRememberTokens(ctx context.Context, pid string) error
	// UseRememberToken finds the pid-token pair and deletes it.
	// If the token could not be found return ErrTokenNotFound
	UseRememberToken(ctx context.Context, pid, token string) error
}

type RememberValuer interface {

	// GetShouldRemember is the checkbox or what have you that
	// tells the remember module if it should remember that user's
	// authentication or not.
	GetShouldRemember() bool
}

type RecoveringServerStorer interface {
	ServerStorer

	// LoadByRecoverSelector finds a user by his recover selector field
	// and should return ErrUserNotFound if that user cannot be found.
	LoadByRecoverSelector(ctx context.Context, selector string) (RecoverableUser, error)
}

type RecoverableUser interface {
	AuthableUser

	GetEmail() (email string)
	GetRecoverSelector() (selector string)
	GetRecoverVerifier() (verifier string)
	GetRecoverExpiry() (expiry time.Time)

	PutEmail(email string)
	PutRecoverSelector(selector string)
	PutRecoverVerifier(verifier string)
	PutRecoverExpiry(expiry time.Time)
}

type RecoverStartValuer interface {
	Validator

	GetPID() string
}

type RecoverMiddleValuer interface {
	Validator

	GetToken() string
}

type RecoverEndValuer interface {
	Validator

	GetPassword() string
	GetToken() string
}

type OAuth2User interface {
	User

	// IsOAuth2User checks to see if a user was registered in the site as an
	// oauth2 user.
	IsOAuth2User() bool

	GetOAuth2UID() (uid string)
	GetOAuth2Provider() (provider string)
	GetOAuth2AccessToken() (token string)
	GetOAuth2RefreshToken() (refreshToken string)
	GetOAuth2Expiry() (expiry time.Time)

	PutOAuth2UID(uid string)
	PutOAuth2Provider(provider string)
	PutOAuth2AccessToken(token string)
	PutOAuth2RefreshToken(refreshToken string)
	PutOAuth2Expiry(expiry time.Time)
}

type OAuth2ServerStorer interface {
	ServerStorer

	// NewFromOAuth2 should return an OAuth2User from a set
	// of details returned from OAuth2Provider.FindUserDetails
	// A more in-depth explanation is that once we've got an access token
	// for the service in question (say a service that rhymes with book)
	// the FindUserDetails function does an http request to a known endpoint
	// that provides details about the user, those details are captured in a
	// generic way as map[string]string and passed into this function to be
	// turned into a real user.
	//
	// It's possible that the user exists in the database already, and so
	// an attempt should be made to look that user up using the details.
	// Any details that have changed should be updated. Do not save the user
	// since that will be done later by OAuth2ServerStorer.SaveOAuth2()
	NewFromOAuth2(ctx context.Context, provider string, details map[string]string) (OAuth2User, error)

	// SaveOAuth2 has different semantics from the typical ServerStorer.Save,
	// in this case we want to insert a user if they do not exist.
	// The difference must be made clear because in the non-oauth2 case,
	// we know exactly when we want to Create vs Update. However since we're
	// simply trying to persist a user that may have been in our database,
	// but if not should already be (since you can think of the operation as
	// a caching of what's on the oauth2 provider's servers).
	SaveOAuth2(ctx context.Context, user OAuth2User) error
}

type Mailer interface {
	Send(context.Context, authboss.Email) error
}

type Moduler interface {
	// Init the module
	Init(*authboss.Authboss) error
}

type Logger interface {
	Info(string)
	Error(string)
}

type LockableUser interface {
	User

	GetAttemptCount() (attempts int)
	GetLastAttempt() (last time.Time)
	GetLocked() (locked time.Time)

	PutAttemptCount(attempts int)
	PutLastAttempt(last time.Time)
	PutLocked(locked time.Time)
}

type HTTPResponder interface {
	Respond(w http.ResponseWriter, r *http.Request, code int, templateName string, data authboss.HTMLData) error
}

type HTTPRedirector interface {
	Redirect(w http.ResponseWriter, r *http.Request, ro authboss.RedirectOptions) error
}

type FieldError interface {
	error
	Name() string
	Err() error
}

type ErrorHandler interface {
	Wrap(func(w http.ResponseWriter, r *http.Request) error) http.Handler
}

type CreatingServerStorer interface {
	ServerStorer

	// New creates a blank user, it is not yet persisted in the database
	// but is just for storing data
	New(ctx context.Context) User
	// Create the user in storage, it should not overwrite a user
	// and should return ErrUserFound if it currently exists.
	Create(ctx context.Context, user User) error
}

type ContextLogger interface {
	FromContext(context.Context) Logger
}

type ConfirmingServerStorer interface {
	ServerStorer

	// LoadByConfirmSelector finds a user by his confirm selector field
	// and should return ErrUserNotFound if that user cannot be found.
	LoadByConfirmSelector(ctx context.Context, selector string) (ConfirmableUser, error)
}

type ConfirmableUser interface {
	User

	GetEmail() (email string)
	GetConfirmed() (confirmed bool)
	GetConfirmSelector() (selector string)
	GetConfirmVerifier() (verifier string)

	PutEmail(email string)
	PutConfirmed(confirmed bool)
	PutConfirmSelector(selector string)
	PutConfirmVerifier(verifier string)
}

type ConfirmValuer interface {
	Validator

	GetToken() string
}

type ClientStateReadWriter interface {
	// ReadState should return a map like structure allowing it to look up
	// any values in the current session, or any cookie in the request
	ReadState(*http.Request) (ClientState, error)
	// WriteState can sometimes be called with a nil ClientState in the event
	// that no ClientState was read in from LoadClientState
	WriteState(http.ResponseWriter, ClientState, []authboss.ClientStateEvent) error
}

type ClientState interface {
	Get(key string) (string, bool)
}

type BodyReader interface {
	Read(page string, r *http.Request) (Validator, error)
}

type AuthableUser interface {
	User

	GetPassword() (password string)
	PutPassword(password string)
}

type ArbitraryValuer interface {
	Validator

	GetValues() map[string]string
}

type ArbitraryUser interface {
	User

	// GetArbitrary is used only to display the arbitrary data back to the user
	// when the form is reset.
	GetArbitrary() (arbitrary map[string]string)
	// PutArbitrary allows arbitrary fields defined by the authboss library
	// consumer to add fields to the user registration piece.
	PutArbitrary(arbitrary map[string]string)
}
