package str

import (
	"fmt"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(String(123.12))
}

func TestRepeat(t *testing.T) {
	s := "0"
	log.Println(Repeat(s, 10))
}

func TestReplace(t *testing.T) {
	s := "aabbccddbbee"
	log.Println(Replace(s, "bb", "xx", 2))
	log.Println(Replace(s, "", "xx", 10))
}
