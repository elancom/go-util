package lang

import (
	"log"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	now := GetNow()
	log.Println(now)
	log.Println(time.Now().Format("2006-01-02 15:04:05"))
}
