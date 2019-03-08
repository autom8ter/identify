package driver

import (
	"context"
	"github.com/autom8ter/identify/driver/api"
)

type ServerStorer struct {
	LoadFunc func(ctx context.Context, key string) (api.User, error)

	// Save persists the user in the database, this should never
	// create a user and instead return ErrUserNotFound if the user
	// does not exist.
	SaveFunc func(ctx context.Context, user api.User) error
}

func NewServerStorer(loadFunc func(ctx context.Context, key string) (api.User, error), saveFunc func(ctx context.Context, user api.User) error) *ServerStorer {
	return &ServerStorer{LoadFunc: loadFunc, SaveFunc: saveFunc}
}

func (s *ServerStorer) Load(ctx context.Context, key string) (api.User, error) {
	return s.LoadFunc(ctx, key)
}

func (s *ServerStorer) Save(ctx context.Context, user api.User) error {
	return s.SaveFunc(ctx, user)
}

type ConfirmingServerStorer struct {
	*ServerStorer

	// LoadByConfirmSelector finds a user by his confirm selector field
	// and should return ErrUserNotFound if that user cannot be found.
	LoadByConfirmSelectorFunc func(ctx context.Context, selector string) (api.ConfirmableUser, error)
}

func NewConfirmingServerStorer(serverStorer *ServerStorer, loadByConfirmSelectorFunc func(ctx context.Context, selector string) (api.ConfirmableUser, error)) *ConfirmingServerStorer {
	return &ConfirmingServerStorer{ServerStorer: serverStorer, LoadByConfirmSelectorFunc: loadByConfirmSelectorFunc}
}

type CreatingServerStorer struct {
	*ServerStorer

	// New creates a blank user, it is not yet persisted in the database
	// but is just for storing data
	NewFunc func(ctx context.Context) api.User
	// Create the user in storage, it should not overwrite a user
	// and should return ErrUserFound if it currently exists.
	CreateFunc func(ctx context.Context, user api.User) error
}

func NewCreatingServerStorer(serverStorer *ServerStorer, newFunc func(ctx context.Context) api.User, createFunc func(ctx context.Context, user api.User) error) *CreatingServerStorer {
	return &CreatingServerStorer{ServerStorer: serverStorer, NewFunc: newFunc, CreateFunc: createFunc}
}

type OAuth2ServerStorer struct {
	*ServerStorer

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
	NewFromOAuth2Func func(ctx context.Context, provider string, details map[string]string) (api.OAuth2User, error)

	// SaveOAuth2 has different semantics from the typical ServerStorer.Save,
	// in this case we want to insert a user if they do not exist.
	// The difference must be made clear because in the non-oauth2 case,
	// we know exactly when we want to Create vs Update. However since we're
	// simply trying to persist a user that may have been in our database,
	// but if not should already be (since you can think of the operation as
	// a caching of what's on the oauth2 provider's servers).
	SaveOAuth2Func func(ctx context.Context, user api.OAuth2User) error
}

func NewOAuth2ServerStorer(serverStorer *ServerStorer, newFromOAuth2Func func(ctx context.Context, provider string, details map[string]string) (api.OAuth2User, error), saveOAuth2Func func(ctx context.Context, user api.OAuth2User) error) *OAuth2ServerStorer {
	return &OAuth2ServerStorer{ServerStorer: serverStorer, NewFromOAuth2Func: newFromOAuth2Func, SaveOAuth2Func: saveOAuth2Func}
}

type RecoveringServerStorer struct {
	*ServerStorer

	// LoadByRecoverSelector finds a user by his recover selector field
	// and should return ErrUserNotFound if that user cannot be found.
	LoadByRecoverSelectorFunc func(ctx context.Context, selector string) (api.RecoverableUser, error)
}

func NewRecoveringServerStorer(serverStorer *ServerStorer, loadByRecoverSelectorFunc func(ctx context.Context, selector string) (api.RecoverableUser, error)) *RecoveringServerStorer {
	return &RecoveringServerStorer{ServerStorer: serverStorer, LoadByRecoverSelectorFunc: loadByRecoverSelectorFunc}
}

type RememberingServerStorer struct {
	*ServerStorer

	// AddRememberToken to a user
	AddRememberTokenFunc func(ctx context.Context, pid, token string) error
	// DelRememberTokens removes all tokens for the given pid
	DelRememberTokensFunc func(ctx context.Context, pid string) error
	// UseRememberToken finds the pid-token pair and deletes it.
	// If the token could not be found return ErrTokenNotFound
	UseRememberTokenFunc func(ctx context.Context, pid, token string) error
}

func NewRememberingServerStorer(serverStorer *ServerStorer, addRememberTokenFunc func(ctx context.Context, pid, token string) error, delRememberTokensFunc func(ctx context.Context, pid string) error, useRememberTokenFunc func(ctx context.Context, pid, token string) error) *RememberingServerStorer {
	return &RememberingServerStorer{ServerStorer: serverStorer, AddRememberTokenFunc: addRememberTokenFunc, DelRememberTokensFunc: delRememberTokensFunc, UseRememberTokenFunc: useRememberTokenFunc}
}
