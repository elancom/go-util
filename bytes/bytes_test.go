package bytes

import (
	"log"
	"testing"
)

func TestJoin(t *testing.T) {
	bytes1 := []byte("123456")
	bytes2 := []byte("789")
	join := Join(bytes1, bytes2)
	log.Println(string(join))
}
