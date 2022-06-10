package collection

import (
	"errors"
)

var NotFound = errors.New("not found")

func Find[T any](slice []T, def T, predicate func(T) bool) (T, error) {
	for _, v := range slice {
		if predicate(v) {
			return v, nil
		}
	}
	return def, NotFound
}

func FindMapS2s(ps []string, gen func(idx int, it string) string) string {
	for idx, it := range ps {
		r := gen(idx, it)
		if r != "" {
			return r
		}
	}
	return ""
}
