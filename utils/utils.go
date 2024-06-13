package utils

import "fmt"

func AppendError(existError error, newError error) error {
	if existError == nil {
		return newError
	}
	return fmt.Errorf("%v: %w", existError, newError)

}
