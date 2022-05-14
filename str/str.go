package str

import (
	"reflect"
	"strconv"
	"strings"
)

func String(v any) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int, int8, int16, int32, int64:
		return strconv.Itoa(int(reflect.ValueOf(v).Int()))
	case uint, uint8, uint16, uint32, uint64:
		return strconv.Itoa(int(reflect.ValueOf(v).Int()))
	case float32, float64:
		return strconv.Itoa(int(reflect.ValueOf(v).Int()))
	default:
		return ""
	}
}

func StringEx(v any) string {
	switch v.(type) {
	case string:
		return v.(string)
	default:
		return reflect.ValueOf(v).String()
	}
}

func IsBlank(s string) bool {
	return len(strings.Trim(s, "")) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func ToString[T int | int64 | uint64](i T) string {
	return strconv.FormatInt(int64(i), 10)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Equal(a, b string) bool {
	return strings.Compare(a, b) == 0
}

func EqualFold(a, b string) bool {
	return strings.EqualFold(a, b)
}

func Index(s, substr string) int {
	return strings.Index(s, substr)
}

func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

func Split(s, sep string) []string {
	return strings.Split(s, sep)
}
