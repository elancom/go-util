package crypto

import (
	"crypto/md5"
	"fmt"
	"github.com/elancom/go-util/str"
)

// Md6 摘要 小写
func Md6(s string) string {
	sum := md5.Sum([]byte(s))
	sum16 := fmt.Sprintf("%x", sum)
	return str.ToLower(sum16)
}

// Md6b 摘要 小写
func Md6b(b []byte) string {
	sum := md5.Sum(b)
	sum16 := fmt.Sprintf("%x", sum)
	return str.ToLower(sum16)
}
