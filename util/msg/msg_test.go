package msg

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ok := NewOk()
	fmt.Println(ok)
}
