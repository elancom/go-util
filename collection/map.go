package collection

import (
	"go-util/number"
	"go-util/str"
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

func (m Map) StringEx(key string) string {
	return str.StringEx(m[key])
}
