package bytes

import (
	"bytes"
	"log"
	"testing"
)

func TestJoin(t *testing.T) {
	bytes1 := []byte("123456")
	bytes2 := []byte("789")
	join := Join(bytes1, bytes2)
	log.Println(string(join))
}

func TestName(t *testing.T) {
	log.Println(string(bytes.Trim([]byte("\"123\""), "\"")))
	log.Println(string(TrimUint8([]byte("\"123\""), 34)))
	log.Println(string(TrimUint8([]byte("\"123"), 34)))
	log.Println(string(TrimUint8([]byte("123\""), 34)))
	log.Println(string(TrimUint8([]byte("\""), 34)))
	log.Println(string(bytes.Trim([]byte("\""), "\"")))
	log.Println(string(bytes.Trim([]byte("\"\""), "\"")))
}
