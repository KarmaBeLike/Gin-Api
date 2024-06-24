package detector

import (
	"errors"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrDuplicate = errors.New("already exist")
)

type CustomeError struct {
	Err          string `json:"err"`
	DeveloperErr string `json:"developer_error`
	ClientErr    string `json:"client_error`
	Status       int    `json:"status"`
}

func (ce *CustomeError) Error() string {
	return ce.Err
}

func ErrorDetector(err error) (int, string) {
	var customError *CustomeError

	if errors.As(err, customError) {
	}

	return err.Status, err.Err
}
