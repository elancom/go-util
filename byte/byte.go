package byte

import "bytes"

// Join 连接数组
func Join(bss ...[]byte) []byte {
	return bytes.Join(bss, []byte(""))
}

func Joins(bss [][]byte, seq []byte) []byte {
	return bytes.Join(bss, seq)
}
