package collection

import (
	"github.com/elancom/go-util/number"
	"github.com/elancom/go-util/str"
)

type Map map[string]any

func Params(m map[string]any) Map {
	return m
}

func (m Map) Int(key string, def int) int {
	return number.ToInt(m[key], def)
}

func (m Map) String(key string) string {
	return str.String(m[key])
}

func Clone[K string | int, V any](m map[K]V) map[K]V {
	nm := make(map[K]V, len(m))
	for k, v := range m {
		nm[k] = v
	}
	return nm
}
