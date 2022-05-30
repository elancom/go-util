package param

import (
	"fmt"
	"strconv"
	"testing"
)

func TestI(t *testing.T) {
	i, _ := strconv.Atoi("123456789012345678")
	fmt.Println(i)

	m := make(map[string]any)
	m["age"] = "100a00"
	p := NewParams(m)
	fmt.Println(p.Int64L("age"))
	fmt.Println(p.IntL("age"))
	i2, err := p.Int64("age")
	if err != nil {
		fmt.Println(err)
		fmt.Println(i2)
		return
	}
}
