package number

import (
	"testing"
)

func Test(t *testing.T) {
	var b = 1000
	if ToInt(b, 11) != b {
		t.Fail()
	}
}
