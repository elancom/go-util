package t

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	ch := make(chan any)

	// 增加关闭
	go close(ch)

	for {
		select {
		case v, ok := <-ch:
			fmt.Printf("v: %v, ok: %v\n", v, ok)
		}
	}
}
