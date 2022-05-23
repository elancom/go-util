package collection

import (
	"fmt"
	"testing"
)

func TestMap_Int(t *testing.T) {
	var a = make(map[string]any)
	newMap1 := Params(a)
	newMap2 := Params(a)
	newMap1["0"] = "111"
	fmt.Println(newMap1)
	fmt.Println(newMap2["0"])
}

func TestClone(t *testing.T) {
	m := make(map[string]any)
	m["a"] = "22"
	m2 := clone[string, any](m)
	m["a"] = "333"
	fmt.Println(m2)
}
