package errors

import (
	"errors"
	"fmt"
)

const (
	ErrorCodeBadRequest   = "bad request"
	ErrorCodeNotFound     = "not found"
	ErrorCodeUnauthorized = "unauthorized"
	ErrTooManyRequests    = "too many requests"
	ErrInternalServer     = "internal server error"
)

var (
	ErrPhoneAlreadyExists = errors.New("phone already registered")
	ErrPhoneIsInvalid     = errors.New("phone is invalid")
	ErrPasswordNotMatches = errors.New("password does not match")
	ErrValidationPassword = errors.New("validation password failed")
	ErrValidationEmail    = errors.New("validation email failed")
)

var (
	ErrIncorrectPhoneOrPassword = errors.New("incorrect phone or password")
	ErrPhoneNotFound            = errors.New("phone not found")
	ErrRefreshTokenNotFound     = errors.New("refresh token not found")
	ErrCreateRefreshToken       = errors.New("create refresh token failed")
	ErrGenerateAccessToken      = errors.New("error generating access token")
	ErrUpdateRefreshToken       = errors.New("error updating refresh token")
	ErrGenerateRefreshToken     = errors.New("error generating refresh token")
	ErrGetRefreshToken          = errors.New("get refresh token failed")
	ErrDeletingUser             = errors.New("user has already been deleted")
)

func New(msg string) error {
	return errors.New(msg)
}

func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, args...))
}

func NewF(format string, args ...interface{}) error {
	return fmt.Errorf(format, args...)
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
