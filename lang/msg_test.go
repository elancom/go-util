package lang

import (
	"log"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	newOk := NewOk()
	var ok any = NewOk()

	msg := ok.(*Msg)
	log.Println(msg.Code)

	switch ok.(type) {
	case *Msg:
		log.Println("..................")
	}

	if reflect.ValueOf(ok).Type() == reflect.ValueOf(newOk).Type() {
		log.Println("ok", reflect.ValueOf(ok).Type())
	} else {
		log.Println("no")
	}
}
