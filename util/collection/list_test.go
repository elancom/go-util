package collection

import (
	"container/list"
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ls := List{list.New()}
	ls.PushBack(100)
	fmt.Println(ls.Random())
}
