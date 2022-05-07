package str

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(String(123))
	fmt.Println(StringEx(123))
}
