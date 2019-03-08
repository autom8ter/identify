package driver

type User struct {
	Get func() (pid string)
	Put func(pid string)
}

func NewUser(get func() (pid string), put func(pid string)) *User {
	return &User{Get: get, Put: put}
}

func (u *User) GetPID() (pid string) {
	return u.Get()
}

func (u *User) PutPID(pid string) {
	u.Put(pid)
}

type Validator struct {
	ValidatorFunc func() []error
}

func NewValidator(validatorFunc func() []error) *Validator {
	return &Validator{ValidatorFunc: validatorFunc}
}

func (v *Validator) Validate() []error {
	return v.ValidatorFunc()
}
