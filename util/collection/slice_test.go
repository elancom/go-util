package collection

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	nums := []int{1, 2, 3}
	find, _ := Find(nums, 0, func(t int) bool {
		return t%2 == 0
	})
	fmt.Println(find)
}
