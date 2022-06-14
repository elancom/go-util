package lang

import (
	"log"
	"testing"
)

func TestIsDir(t *testing.T) {
	log.Println(IsDir("c:\\Windows"))
	log.Println(IsDir("c:\\Windows\\111"))
	log.Println(IsExist("c:\\Windows"))
	log.Println(IsExist("c:\\Windows\\111"))
}
