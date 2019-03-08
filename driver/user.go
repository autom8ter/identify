package driver

import (
	"time"
)

type AuthableUser struct {
	*User
	GetPasswordFunc func() (password string)
	PutPasswordFunc func(password string)
}

func (a *AuthableUser) GetPassword() (password string) {
	return a.GetPasswordFunc()
}

func (a *AuthableUser) PutPassword(password string) {
	a.PutPasswordFunc(password)
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

func (c *ConfirmableUser) GetEmail() (email string) {
	return c.GetEmailFunc()
}

func (c *ConfirmableUser) GetConfirmed() (confirmed bool) {
	return c.GetConfirmedFunc()

}

func (c *ConfirmableUser) GetConfirmSelector() (selector string) {
	return c.GetConfirmSelectorFunc()
}

func (c *ConfirmableUser) GetConfirmVerifier() (verifier string) {
	return c.GetConfirmVerifierFunc()
}

func (c *ConfirmableUser) PutEmail(email string) {
	c.PutEmailFunc(email)
}

func (c *ConfirmableUser) PutConfirmed(confirmed bool) {
	c.PutConfirmedFunc(confirmed)
}

func (c *ConfirmableUser) PutConfirmSelector(selector string) {
	c.PutConfirmSelectorFunc(selector)
}

func (c *ConfirmableUser) PutConfirmVerifier(verifier string) {
	c.PutConfirmVerifierFunc(verifier)
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

func (l *LockableUser) GetAttemptCount() (attempts int) {
	return l.GetAttemptCountFunc()
}

func (l *LockableUser) GetLastAttempt() (last time.Time) {
	return l.GetLastAttemptFunc()

}

func (l *LockableUser) GetLocked() (locked time.Time) {
	return l.GetLockedFunc()

}

func (l *LockableUser) PutAttemptCount(attempts int) {
	l.PutAttemptCountFunc(attempts)
}

func (l *LockableUser) PutLastAttempt(last time.Time) {
	l.PutLastAttemptFunc(last)
}

func (l *LockableUser) PutLocked(locked time.Time) {
	l.PutLockedFunc(locked)
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

func (r *RecoverableUser) GetEmail() (email string) {
	return r.GetEmailFunc()
}

func (r *RecoverableUser) GetRecoverSelector() (selector string) {
	return r.GetRecoverSelectorFunc()

}

func (r *RecoverableUser) GetRecoverVerifier() (verifier string) {
	return r.GetRecoverVerifierFunc()

}

func (r *RecoverableUser) GetRecoverExpiry() (expiry time.Time) {
	return r.GetRecoverExpiryFunc()
}

func (r *RecoverableUser) PutEmail(email string) {
	r.PutEmailFunc(email)
}

func (r *RecoverableUser) PutRecoverSelector(selector string) {
	r.PutRecoverSelectorFunc(selector)
}

func (r *RecoverableUser) PutRecoverVerifier(verifier string) {
	r.PutRecoverVerifierFunc(verifier)
}

func (r *RecoverableUser) PutRecoverExpiry(expiry time.Time) {
	r.PutRecoverExpiryFunc(expiry)
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

func (a *ArbitraryUser) GetArbitrary() (arbitrary map[string]string) {
	return a.GetArbitraryFunc()
}

func (a *ArbitraryUser) PutArbitrary(arbitrary map[string]string) {
	a.PutArbitraryFunc(arbitrary)
}

type OAuth2User struct {
	*User
	// IsOAuth2User checks to see if a user was registered in the site as an
	// oauth2 user.
	IsOAuth2UserFunc func() bool

	GetOAuth2UIDFunc          func() (uid string)
	GetOAuth2ProviderFunc     func() (provider string)
	GetOAuth2AccessTokenFunc  func() (token string)
	GetOAuth2RefreshTokenFunc func() (refreshToken string)
	GetOAuth2ExpiryFunc       func() (expiry time.Time)

	PutOAuth2UIDFunc          func(uid string)
	PutOAuth2ProviderFunc     func(provider string)
	PutOAuth2AccessTokenFunc  func(token string)
	PutOAuth2RefreshTokenFunc func(refreshToken string)
	PutOAuth2ExpiryFunc       func(expiry time.Time)
}

func NewOAuth2User(user *User, isOAuth2UserFunc func() bool, getOAuth2UIDFunc func() (uid string), getOAuth2ProviderFunc func() (provider string), getOAuth2AccessTokenFunc func() (token string), getOAuth2RefreshTokenFunc func() (refreshToken string), getOAuth2ExpiryFunc func() (expiry time.Time), putOAuth2UIDFunc func(uid string), putOAuth2ProviderFunc func(provider string), putOAuth2AccessTokenFunc func(token string), putOAuth2RefreshTokenFunc func(refreshToken string), putOAuth2ExpiryFunc func(expiry time.Time)) *OAuth2User {
	return &OAuth2User{User: user, IsOAuth2UserFunc: isOAuth2UserFunc, GetOAuth2UIDFunc: getOAuth2UIDFunc, GetOAuth2ProviderFunc: getOAuth2ProviderFunc, GetOAuth2AccessTokenFunc: getOAuth2AccessTokenFunc, GetOAuth2RefreshTokenFunc: getOAuth2RefreshTokenFunc, GetOAuth2ExpiryFunc: getOAuth2ExpiryFunc, PutOAuth2UIDFunc: putOAuth2UIDFunc, PutOAuth2ProviderFunc: putOAuth2ProviderFunc, PutOAuth2AccessTokenFunc: putOAuth2AccessTokenFunc, PutOAuth2RefreshTokenFunc: putOAuth2RefreshTokenFunc, PutOAuth2ExpiryFunc: putOAuth2ExpiryFunc}
}

func (a *OAuth2User) IsOAuth2User() bool {
	return a.IsOAuth2UserFunc()
}

func (a *OAuth2User) GetOAuth2UID() (uid string) {
	return a.GetOAuth2UIDFunc()
}

func (a *OAuth2User) GetOAuth2Provider() (provider string) {
	return a.GetOAuth2ProviderFunc()
}

func (a *OAuth2User) GetOAuth2AccessToken() (token string) {
	return a.GetOAuth2AccessTokenFunc()
}

func (a *OAuth2User) GetOAuth2RefreshToken() (refreshToken string) {
	return a.GetOAuth2RefreshTokenFunc()
}

func (a *OAuth2User) GetOAuth2Expiry() (expiry time.Time) {
	return a.GetOAuth2ExpiryFunc()
}

func (a *OAuth2User) PutOAuth2UID(uid string) {
	a.PutOAuth2UIDFunc(uid)
}

func (a *OAuth2User) PutOAuth2Provider(provider string) {
	a.PutOAuth2ProviderFunc(provider)
}

func (a *OAuth2User) PutOAuth2AccessToken(token string) {
	a.PutOAuth2AccessTokenFunc(token)
}

func (a *OAuth2User) PutOAuth2RefreshToken(refreshToken string) {
	a.PutOAuth2RefreshTokenFunc(refreshToken)
}

func (a *OAuth2User) PutOAuth2Expiry(expiry time.Time) {
	a.PutOAuth2ExpiryFunc(expiry)
}
