package lang

import (
	"log"
	"testing"
)

func TestGetNow(t *testing.T) {
	now := GetNow()
	log.Println(now)
}
