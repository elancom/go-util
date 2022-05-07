package sqln

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCat struct {
	Name     string `db:"name"`
	Password string `db:"password"`
}

func TestName(t *testing.T) {
	a := reflect.ValueOf(TestCat{})
	mapper, err := GetMappers().Mapper(a.Type())
	if err != nil {
		panic(err)
	}
	fmt.Println(mapper.Field("name"))
}
