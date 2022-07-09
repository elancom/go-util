package lang

import (
	"encoding/json"
	"log"
	"testing"
)

func TestPage(t *testing.T) {
	p := new(Page)
	p.SetPage(2)
	p.SetRows(10)
	marshal, _ := json.Marshal(p)
	log.Println(string(marshal))
	log.Println(p.GetPageIndex())
	log.Println(p.GetFirst())
	p.SetPage(3)
	p.SetMaxRows(8)
	marshal, _ = json.Marshal(p)
	log.Println(string(marshal))
	log.Println(p.GetPageIndex())
	log.Println(p.GetFirst())
	log.Println(*NewPage(2, 25))
}
