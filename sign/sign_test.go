package sign

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	m1 := map[string]string{
		"name": "1234",
	}
	fmt.Println(sign("123456", m1))
}

func TestCheck(t *testing.T) {
	m1 := map[string]string{
		"name": "1234",
	}
	fmt.Println("check:", check("123456", m1, "D8263F84D0352191D9C19E2D36C1F0B8"))
}