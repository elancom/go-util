package msg

import (
	"github.com/elancom/go-util/number"
	"github.com/elancom/go-util/str"
)

const OK = 200
const Err = 400

type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (m *Msg) Error() string {
	return m.Msg
}

func NewOk(data ...any) *Msg {
	d := any(nil)
	if len(data) > 0 {
		d = data[0]
	}
	return NewMsg(200, "", d)
}

func NewErr(msg string) *Msg {
	return NewMsg(400, msg, nil)
}

func NewMsg(code int, msg string, data ...any) *Msg {
	d := any(nil)
	if len(data) > 0 {
		d = data[0]
	}
	return &Msg{Code: code, Msg: msg, Data: d}
}

func NewFrom(d map[string]any) *Msg {
	code := number.ToInt(d["code"], 0)
	msg0 := str.String(d["msg"])
	data := d["data"]
	return NewMsg(code, msg0, data)
}

func (m Msg) IsOk() bool {
	return m.Code == 200
}

func (m Msg) IsErr() bool {
	return !m.IsOk()
}
