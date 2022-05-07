package sqln

import (
	"reflect"
	"sync"
)

var mapperLock = sync.Mutex{}
var mappers = Mappers{
	cache: map[reflect.Type]*Mapper{},
}

type Mappers struct {
	cache map[reflect.Type]*Mapper
}

func GetMappers() *Mappers {
	return &mappers
}

func (m Mappers) Mapper(t reflect.Type) (*Mapper, error) {
	mapper := m.cache[t]
	if mapper != nil {
		return mapper, nil
	}
	mapperLock.Lock()
	defer mapperLock.Unlock()
	mapper = m.parseMapper(t)
	m.cache[t] = mapper
	return mapper, nil
}

func (m Mappers) parseMapper(t reflect.Type) *Mapper {
	mapper := Mapper{fields: make(map[string]*Field, t.NumField())}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		column := field.Tag.Get("db")
		if len(column) == 0 {
			column = field.Name
		}
		mapper.Field(column, Field{Name: field.Name, Column: column})
	}
	return &mapper
}

type Mapper struct {
	fields map[string]*Field
}

func (e Mapper) Field(column string, field ...Field) *Field {
	if len(field) != 0 {
		if column != field[0].Column {
			panic("field name error")
		}
		e.fields[column] = &field[0]
	}
	return e.fields[column]
}

type Field struct {
	Name   string
	Column string
}
