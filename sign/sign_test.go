package sign

import (
	"fmt"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	m1 := map[string]string{
		"name": "1234",
	}
	fmt.Println(Sign(m1, "123456"))
}

func TestCheck(t *testing.T) {
	m1 := map[string]string{
		"name": "1234",
	}
	fmt.Println("Check:", Check("123456", m1, "D8263F84D0352191D9C19E2D36C1F0B8"))
}

func TestSignStr(t *testing.T) {
	str := Str("123456000", "123456")
	log.Println(str)
}

func TestCheckStr(t *testing.T) {
	str := Str("123456000", "123456")
	log.Println(str)
	log.Println(CheckStr("123456000", "123456", Str("123456000", "123456")))
}
