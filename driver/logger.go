package driver

import (
	"context"
	"github.com/autom8ter/identify/driver/api"
	"net/http"
)

type Logger struct {
	InfoFunc  func(string)
	ErrorFunc func(string)
}

func NewLogger(infoFunc func(string), errorFunc func(string)) *Logger {
	return &Logger{InfoFunc: infoFunc, ErrorFunc: errorFunc}
}

func (l *Logger) Info(s string) {
	l.InfoFunc(s)
}

func (l *Logger) Error(s string) {
	l.ErrorFunc(s)
}

type ContextLogger func(context.Context) api.Logger

func (f ContextLogger) FromContext(ctx context.Context) api.Logger {
	return f(ctx)
}

type ErrorHandler func(func(w http.ResponseWriter, r *http.Request) error) http.Handler

func (e ErrorHandler) Wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return e(fn)
}

type FieldError struct {
	error
	NameFunc func() string
	ErrFunc  func() error
}

func NewFieldError(error error, nameFunc func() string, errFunc func() error) *FieldError {
	return &FieldError{error: error, NameFunc: nameFunc, ErrFunc: errFunc}
}

func (f *FieldError) Name() string {
	return f.NameFunc()
}

func (f *FieldError) Err() error {
	return f.ErrFunc()
}
