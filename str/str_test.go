package str

import (
	"fmt"
	"log"
	"regexp"
	"testing"
)

func Test1(t *testing.T) {
	s := "123,123,1     23   123  981、72，31312。3213"
	compile, err := regexp.Compile("[,。，、|\n]|\\s+")
	log.Println(err)
	allString := compile.ReplaceAllString(s, " ")
	log.Println(allString)
}

func TestName(t *testing.T) {
	fmt.Println(String(123.12))
}

func TestRepeat(t *testing.T) {
	s := "0"
	log.Println(Repeat(s, 10))
}

func TestReplace(t *testing.T) {
	s := "aabbccddbbee"
	log.Println(Replace(s, "bb", "xx", 2))
	log.Println(Replace(s, "", "xx", 10))
}

func TestEqualFold(t *testing.T) {
	log.Println(EqualFold("AAA", "aaA"))
}

func TestContains(t *testing.T) {
	log.Println(Contains("aaabbbccc", "bbb"))
	log.Println(Contains("aaabbbccc", "bb"))
	log.Println(Contains("aaabbbccc", "bbbb"))
	log.Println(ContainsAny("aaabbbccc", "dc"))
}
