package number

import (
	"reflect"
	"strconv"
)

func ToInt8(v any, def int8) int8 {
	return toInt[int8](v, def)
}
func ToInt16(v any, def int16) int16 {
	return toInt[int16](v, def)
}
func ToInt32(v any, def int32) int32 {
	return toInt[int32](v, def)
}
func ToInt64(v any, def int64) int64 {
	return toInt[int64](v, def)
}
func ToInt(v any, def int) int {
	return toInt[int](v, def)
}

func toInt[T int | int8 | int16 | int32 | int64](s any, def T) T {
	switch s.(type) {
	case T:
		return s.(T)
	case string:
		i, err := strconv.Atoi(s.(string))
		if err != nil {
			return def
		}
		return T(i)
	case int, int8, int16, int32, int64:
		return T(reflect.ValueOf(s).Int())
	case uint, uint8, uint16, uint32, uint64:
		return T(reflect.ValueOf(s).Int())
	case float32, float64:
		return T(reflect.ValueOf(s).Float())
	default:
		return def
	}
}
