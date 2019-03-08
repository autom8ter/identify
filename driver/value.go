package driver

type UserValuer struct {
	*Validator
	GetPIDFunc      func() string
	GetPasswordFunc func() string
}

func NewUserValuer(validator *Validator, getPIDFunc func() string, getPasswordFunc func() string) *UserValuer {
	return &UserValuer{Validator: validator, GetPIDFunc: getPIDFunc, GetPasswordFunc: getPasswordFunc}
}

func (u *UserValuer) GetPID() string {
	return u.GetPIDFunc()
}

func (u *UserValuer) GetPassword() string {
	return u.GetPasswordFunc()
}

type ArbitraryValuer struct {
	*Validator
	GetValuesFunc func() map[string]string
}

func NewArbitraryValuer(validator *Validator, getValuesFunc func() map[string]string) *ArbitraryValuer {
	return &ArbitraryValuer{Validator: validator, GetValuesFunc: getValuesFunc}
}

func (a *ArbitraryValuer) GetValues() map[string]string {
	return a.GetValuesFunc()
}

type ConfirmValuer struct {
	*Validator
	GetTokenFunc func() string
}

func (c *ConfirmValuer) GetToken() string {
	return c.GetTokenFunc()
}

func NewConfirmValuer(validator *Validator, getTokenFunc func() string) *ConfirmValuer {
	return &ConfirmValuer{Validator: validator, GetTokenFunc: getTokenFunc}
}

type RecoverStartValuer struct {
	*Validator

	GetPIDFunc func() string
}

func (r *RecoverStartValuer) GetPID() string {
	return r.GetPIDFunc()
}

func NewRecoverStartValuer(validator *Validator, getPIDFunc func() string) *RecoverStartValuer {
	return &RecoverStartValuer{Validator: validator, GetPIDFunc: getPIDFunc}
}

type RecoverMiddleValuer struct {
	*Validator

	GetTokenFunc func() string
}

func (r *RecoverMiddleValuer) GetToken() string {
	return r.GetTokenFunc()
}

func NewRecoverMiddleValuer(validator *Validator, getTokenFunc func() string) *RecoverMiddleValuer {
	return &RecoverMiddleValuer{Validator: validator, GetTokenFunc: getTokenFunc}
}

type RecoverEndValuer struct {
	*Validator

	GetPasswordFunc func() string
	GetTokenFunc    func() string
}

func (r *RecoverEndValuer) GetPassword() string {
	return r.GetPasswordFunc()
}

func (r RecoverEndValuer) GetToken() string {
	return r.GetTokenFunc()
}

func NewRecoverEndValuer(validator *Validator, getPasswordFunc func() string, getTokenFunc func() string) *RecoverEndValuer {
	return &RecoverEndValuer{Validator: validator, GetPasswordFunc: getPasswordFunc, GetTokenFunc: getTokenFunc}
}

type RememberValuer struct {

	// GetShouldRemember is the checkbox or what have you that
	// tells the remember module if it should remember that user's
	// authentication or not.
	GetShouldRememberFunc func() bool
}

func (r *RememberValuer) GetShouldRemember() bool {
	return r.GetShouldRememberFunc()
}

func NewRememberValuer(getShouldRememberFunc func() bool) *RememberValuer {
	return &RememberValuer{GetShouldRememberFunc: getShouldRememberFunc}
}
