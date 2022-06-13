package bytes

import (
	"bytes"
	"github.com/elancom/go-util/math"
)

// Join 连接数组
func Join(bss ...[]byte) []byte {
	return bytes.Join(bss, []byte(""))
}

func Joins(bss [][]byte, seq []byte) []byte {
	return bytes.Join(bss, seq)
}

func TrimUint8(b []byte, cut uint8) []byte {
	if len(b) == 0 {
		return b
	}
	b1 := b[0] == cut
	b2 := b[len(b)-1] == cut
	if b1 {
		if b2 {
			return b[1:math.Max(1, len(b)-1)]
		}
		return b[1:]
	} else if b2 {
		return b[:len(b)-1]
	}
	return b
}
