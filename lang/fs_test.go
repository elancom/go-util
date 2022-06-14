package lang

import (
	"log"
	"testing"
)

func TestIsDir(t *testing.T) {
	log.Println(IsDir("c:\\Windows"))
	log.Println(IsDir("c:\\Windows\\111"))
	log.Println(IsFileExist("c:\\Windows"))
	log.Println(IsFileExist("c:\\Windows\\111"))
}
