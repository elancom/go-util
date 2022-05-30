package crypto

import (
	"crypto/md5"
	"fmt"
)

func md6(s string) string {
	sum := md5.Sum([]byte(s))
	sum16 := fmt.Sprintf("%x", sum)
	return sum16
}
