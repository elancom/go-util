package collection

import (
	"fmt"
	"testing"
)

func TestMap_Int(t *testing.T) {
}

func TestClone(t *testing.T) {
	m := make(map[string]any)
	m["a"] = "22"
	m2 := Clone[string, any](m)
	m["a"] = "333"
	fmt.Println(m2)
}
