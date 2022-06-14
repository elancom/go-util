package crypto

import (
	"encoding/hex"
)

// 字节<=>16进制字符串

func HexEncode(b []byte) []byte {
	bs := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(bs, b)
	return bs
}

func HexEncodeToString(b []byte) string {
	return string(HexEncode(b))
}

func HexDecode(src []byte) ([]byte, error) {
	bs := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(bs, src)
	return bs, err
}

func HexDecodeString(src string) ([]byte, error) {
	return HexDecode([]byte(src))
}
