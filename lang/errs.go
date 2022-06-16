package lang

import "errors"

var NotFound = errors.New("not found")

func isNotFound(err error) bool {
	return err == NotFound
}
