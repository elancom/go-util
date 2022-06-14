package crypto

import (
	"log"
	"testing"
)

func TestHexEncodeToString(t *testing.T) {
	s := "idyxUU3dq4KJPIqylw"
	md6 := Md6(s)
	toString := HexEncodeToString([]byte(md6))
	log.Println(toString)
}

func TestHexEncode(t *testing.T) {
	src := "123456"
	log.Println(src)
	s := HexEncodeToString([]byte(src))
	log.Println(s)
	decodeString, err := HexDecodeString(s)
	if err != nil {
		panic(err)
	}
	log.Println(string(decodeString))
	if src != string(decodeString) {
		t.Error(string(decodeString))
	}
}

func TestHexDecode(t *testing.T) {
	decode, err := HexDecode([]byte("aaa"))
	if err != nil {
		panic(err)
	}
	log.Println(string(decode))
}

func TestHexDecodeString(t *testing.T) {
	decode, err := HexDecodeString("0b0b0b")
	if err != nil {
		t.Error(err)
	}
	for _, b := range decode {
		if b != 11 {
			t.Error("err")
		}
	}
	log.Println(string(decode))
}
