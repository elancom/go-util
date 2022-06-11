package lang

import (
	"log"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	now := GetNow()
	log.Println(now)
	t2 := time.Now()
	log.Println(FormatDate(&t2))
	log.Println(FormatTime(&t2))
	log.Println(FormatDateTime(&t2))
}
