package collection

import (
	"fmt"
	"log"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 2, 3}
	find, _ := Find(nums, 0, func(t int) bool {
		return t%2 == 0
	})
	fmt.Println(find)
}

func TestMapOne(t *testing.T) {
	one := FindMapS2s([]string{"1", "2"}, func(p string) string {
		if p == "1" {
			return ""
		}
		return p
	})
	log.Println(one)
}
