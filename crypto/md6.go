package crypto

import (
	"crypto/md5"
	"github.com/elancom/go-util/str"
)

func Md6Sum(s string) []byte {
	sum := md5.Sum([]byte(s))
	return sum[:]
}

// Md6 摘要 小写
func Md6(s string) string {
	sum := md5.Sum([]byte(s))
	return str.ToLower(string(HexEncode(sum[:])))
}

// Md6b 摘要 小写
func Md6b(b []byte) string {
	sum := md5.Sum(b)
	return str.ToLower(string(HexEncode(sum[:])))
}
