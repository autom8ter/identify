package driver

import (
	"time"
)

type AuthableUser struct {
	*User
	GetPasswordFunc func() (password string)
	PutPasswordFunc func(password string)
}

func NewAuthableUser(user *User, getPasswordFunc func() (password string), putPasswordFunc func(password string)) *AuthableUser {
	return &AuthableUser{User: user, GetPasswordFunc: getPasswordFunc, PutPasswordFunc: putPasswordFunc}
}

type ConfirmableUser struct {
	*User

	GetEmailFunc           func() (email string)
	GetConfirmedFunc       func() (confirmed bool)
	GetConfirmSelectorFunc func() (selector string)
	GetConfirmVerifierFunc func() (verifier string)

	PutEmailFunc           func(email string)
	PutConfirmedFunc       func(confirmed bool)
	PutConfirmSelectorFunc func(selector string)
	PutConfirmVerifierFunc func(verifier string)
}

func NewConfirmableUser(user *User, getEmailFunc func() (email string), getConfirmedFunc func() (confirmed bool), getConfirmSelectorFunc func() (selector string), getConfirmVerifierFunc func() (verifier string), putEmailFunc func(email string), putConfirmedFunc func(confirmed bool), putConfirmSelectorFunc func(selector string), putConfirmVerifierFunc func(verifier string)) *ConfirmableUser {
	return &ConfirmableUser{User: user, GetEmailFunc: getEmailFunc, GetConfirmedFunc: getConfirmedFunc, GetConfirmSelectorFunc: getConfirmSelectorFunc, GetConfirmVerifierFunc: getConfirmVerifierFunc, PutEmailFunc: putEmailFunc, PutConfirmedFunc: putConfirmedFunc, PutConfirmSelectorFunc: putConfirmSelectorFunc, PutConfirmVerifierFunc: putConfirmVerifierFunc}
}

type LockableUser struct {
	*User

	GetAttemptCountFunc func() (attempts int)
	GetLastAttemptFunc  func() (last time.Time)
	GetLockedFunc       func() (locked time.Time)

	PutAttemptCountFunc func(attempts int)
	PutLastAttemptFunc  func(last time.Time)
	PutLockedFunc       func(locked time.Time)
}

func NewLockableUser(user *User, getAttemptCountFunc func() (attempts int), getLastAttemptFunc func() (last time.Time), getLockedFunc func() (locked time.Time), putAttemptCountFunc func(attempts int), putLastAttemptFunc func(last time.Time), putLockedFunc func(locked time.Time)) *LockableUser {
	return &LockableUser{User: user, GetAttemptCountFunc: getAttemptCountFunc, GetLastAttemptFunc: getLastAttemptFunc, GetLockedFunc: getLockedFunc, PutAttemptCountFunc: putAttemptCountFunc, PutLastAttemptFunc: putLastAttemptFunc, PutLockedFunc: putLockedFunc}
}

type RecoverableUser struct {
	*AuthableUser

	GetEmailFunc           func() (email string)
	GetRecoverSelectorFunc func() (selector string)
	GetRecoverVerifierFunc func() (verifier string)
	GetRecoverExpiryFunc   func() (expiry time.Time)

	PutEmailFunc           func(email string)
	PutRecoverSelectorFunc func(selector string)
	PutRecoverVerifierFunc func(verifier string)
	PutRecoverExpiryFunc   func(expiry time.Time)
}

func NewRecoverableUser(authableUser *AuthableUser, getEmailFunc func() (email string), getRecoverSelectorFunc func() (selector string), getRecoverVerifierFunc func() (verifier string), getRecoverExpiryFunc func() (expiry time.Time), putEmailFunc func(email string), putRecoverSelectorFunc func(selector string), putRecoverVerifierFunc func(verifier string), putRecoverExpiryFunc func(expiry time.Time)) *RecoverableUser {
	return &RecoverableUser{AuthableUser: authableUser, GetEmailFunc: getEmailFunc, GetRecoverSelectorFunc: getRecoverSelectorFunc, GetRecoverVerifierFunc: getRecoverVerifierFunc, GetRecoverExpiryFunc: getRecoverExpiryFunc, PutEmailFunc: putEmailFunc, PutRecoverSelectorFunc: putRecoverSelectorFunc, PutRecoverVerifierFunc: putRecoverVerifierFunc, PutRecoverExpiryFunc: putRecoverExpiryFunc}
}

type ArbitraryUser struct {
	*User

	// GetArbitrary is used only to display the arbitrary data back to the user
	// when the form is reset.
	GetArbitraryFunc func() (arbitrary map[string]string)
	// PutArbitrary allows arbitrary fields defined by the authboss library
	// consumer to add fields to the user registration piece.
	PutArbitraryFunc func(arbitrary map[string]string)
}

type OAuth2User struct {
	*User
}

func (*OAuth2User) IsOAuth2User() bool {
	panic("implement me")
}

func (*OAuth2User) GetOAuth2UID() (uid string) {
	panic("implement me")
}

func (*OAuth2User) GetOAuth2Provider() (provider string) {
	panic("implement me")
}

func (*OAuth2User) GetOAuth2AccessToken() (token string) {
	panic("implement me")
}

func (*OAuth2User) GetOAuth2RefreshToken() (refreshToken string) {
	panic("implement me")
}

func (*OAuth2User) GetOAuth2Expiry() (expiry time.Time) {
	panic("implement me")
}

func (*OAuth2User) PutOAuth2UID(uid string) {
	panic("implement me")
}

func (*OAuth2User) PutOAuth2Provider(provider string) {
	panic("implement me")
}

func (*OAuth2User) PutOAuth2AccessToken(token string) {
	panic("implement me")
}

func (*OAuth2User) PutOAuth2RefreshToken(refreshToken string) {
	panic("implement me")
}

func (*OAuth2User) PutOAuth2Expiry(expiry time.Time) {
	panic("implement me")
}
