package number

import (
	"github.com/elancom/go-util/lang"
	"reflect"
	"strconv"
)

var IntCaseErr = lang.NewErr("case int err")
var IntCaseErrStr = lang.NewErr("case int err(s)")

func ToInt8L(v any, def ...int8) int8 {
	return intL[int8](v, def...)
}
func ToInt16L(v any, def ...int16) int16 {
	return intL[int16](v, def...)
}
func ToInt32L(v any, def ...int32) int32 {
	return intL[int32](v, def...)
}
func ToInt64L(v any, def ...int64) int64 {
	return intL[int64](v, def...)
}
func ToIntL(v any, def ...int) int {
	return intL[int](v, def...)
}

func intL[T int | int8 | int16 | int32 | int64](s any, def ...T) T {
	e, _ := ToInt(s, def...)
	return e
}

func ToInt[T int | int8 | int16 | int32 | int64](s any, def ...T) (T, error) {
	switch s.(type) {
	case T:
		return s.(T), nil
	case string:
		i, err := strconv.Atoi(s.(string))
		if err != nil {
			if len(def) > 0 {
				return def[0], IntCaseErrStr
			}
			return 0, IntCaseErrStr
		}
		return T(i), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return T(reflect.ValueOf(s).Int()), nil
	case float32, float64:
		return T(reflect.ValueOf(s).Float()), nil
	default:
		if len(def) > 0 {
			return def[0], IntCaseErr
		}
		return 0, IntCaseErr
	}
}
