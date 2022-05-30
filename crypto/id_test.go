package crypto

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	fmt.Println(NewUUID())
	fmt.Println(NewUUID())
	fmt.Println(NewId32())
	fmt.Println(NewId32())
	fmt.Println(NewId64())
	fmt.Println(NewId64())
}
