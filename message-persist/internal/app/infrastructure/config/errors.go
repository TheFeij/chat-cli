package config

import (
	"errors"
	"fmt"
)

const (
	errReadingConfigMsg       = "error reading config file"
	errUnmarshallingConfigMsg = "error unmarshalling config file"
)

var (
	ErrReadingConfig       = &configError{msg: errReadingConfigMsg}
	ErrUnmarshallingConfig = &configError{msg: errUnmarshallingConfigMsg}
	ErrEnvFilePathNotSet   = fmt.Errorf("environment variable %s is not set", configFilePathKey)
	ErrEnvFileNameNotSet   = fmt.Errorf("environment variable %s is not set", configFileNameKey)
	ErrEnvFileTypeNotSet   = fmt.Errorf("environment variable %s is not set", configFileTypeKey)
)

func newErrReadingConfig(err error) error {
	return &configError{msg: errReadingConfigMsg, err: err}
}

func newErrUnmarshallingConfig(err error) error {
	return &configError{msg: errUnmarshallingConfigMsg, err: err}
}

type configError struct {
	msg string
	err error
}

func (e *configError) Error() string {
	if e.err == nil {
		return e.msg
	}

	return fmt.Sprintf("%s: %v", e.msg, e.err)
}

func (e *configError) Unwrap() error { return e.err }

func (e *configError) Is(target error) bool {
	var t *configError

	ok := errors.As(target, &t)

	return ok && e.msg == t.msg
}
