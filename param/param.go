package param

import (
	"github.com/elancom/go-util/number"
	"github.com/elancom/go-util/str"
)

func NewParams(m map[string]string) *Params {
	p := new(Params)
	p.src = m
	return p
}

type Params struct {
	src map[string]string
}

func (p *Params) Get(key string, def ...string) string {
	s := p.src[key]
	if s == "" && len(def) > 0 {
		return def[0]
	}
	return s
}

func (p *Params) Trim(key string, def ...string) string {
	s := str.Trim(p.src[key])
	if len(s) == 0 && len(def) > 0 {
		return def[0]
	}
	return s
}

func (p *Params) Int(key string, def ...int) (int, error) {
	s := p.src[key]
	return number.ToInt(s, def...)
}

func (p *Params) IntL(key string, def ...int) int {
	s := p.src[key]
	return number.ToIntL(s, def...)
}

func (p *Params) Int64(key string, def ...int64) (int64, error) {
	s := p.src[key]
	return number.ToInt(s, def...)
}

func (p *Params) Int64L(key string, def ...int64) int64 {
	s := p.src[key]
	return number.ToInt64L(s, def...)
}
