package crypto

import (
	"fmt"
	"github.com/elancom/go-util/rand"
	"testing"
)

func TestMd6(t *testing.T) {
	fmt.Println(Md6(rand.RandomNumber(125)))
	fmt.Println(Md6(rand.RandomNumber(125)))
	fmt.Println(Md6(rand.RandomNumber(125)))
}
