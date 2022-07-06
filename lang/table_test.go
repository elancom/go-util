package lang

import (
	"log"
	"testing"
)

func TestTable(t *testing.T) {
	ss := make([]string, 0, 100)
	ss = append(ss, "1")
	ss = append(ss, "1")
	ss = append(ss, "1")
	ss = append(ss, "1")
	table := NewTable(ss, int64(len(ss)))
	ss = ss[:len(ss)-1]
	log.Println(len(table.List))
	log.Println(len(ss))
}
