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
		// ‘b’ (-ddddp±ddd，二进制指数)
		// ‘e’ (-d.dddde±dd，十进制指数)
		// ‘E’ (-d.ddddE±dd，十进制指数)
		// ‘f’ (-ddd.dddd，没有指数)
		// ‘g’ (‘e’:大指数，‘f’:其它情况)
		// ‘G’ (‘E’:大指数，‘f’:其它情况)
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'g', 6, 64)
	default:
		return ""
	}
}

func IsBlank(s string) bool {
	return len(strings.Trim(s, "")) == 0
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

//func ToString[T int | int64 | uint64](i T) string {
//	return strconv.FormatInt(int64(i), 10)
//}

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

func Join(elems []string, sep string) string {
	return strings.Join(elems, sep)
}

func Trim(s string) string {
	return Trims(s, "")
}

func Trims(s string, cs string) string {
	return strings.Trim(s, cs)
}

func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// Replace
// old 如果为空则每隔1个字符串插入new字符串
// n 替换多次
func Replace(s, old, new string, n int) string {
	return strings.Replace(s, old, new, n)
}

func Contains(s, sub string) bool {
	return strings.Contains(s, sub)
}

func ContainsAny(s, chars string) bool {
	return strings.ContainsAny(s, chars)
}
