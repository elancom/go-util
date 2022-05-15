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
