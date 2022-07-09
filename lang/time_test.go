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

func Test1(t *testing.T) {
	parse, err := ToDateTime("2022-07-09 18:16:07")
	log.Println(err)
	log.Println(parse)
	log.Println(ToDateTimeWith(time.Now().UnixMilli()))
}
