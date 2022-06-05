package msg

import (
	"github.com/elancom/go-util/lang"
	"log"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	newOk := lang.NewOk()
	var ok any = NewOk()

	msg := ok.(*lang.Msg)
	log.Println(msg.Code)

	switch ok.(type) {
	case *lang.Msg:
		log.Println("..................")
	}

	if reflect.ValueOf(ok).Type() == reflect.ValueOf(newOk).Type() {
		log.Println("ok", reflect.ValueOf(ok).Type())
	} else {
		log.Println("no")
	}
}
