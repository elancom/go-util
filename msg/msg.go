package msg

import "go-util/collection"

type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
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

func NewFrom(m map[string]any) *Msg {
	mp := collection.Params(m)
	code := mp.Int("code", 0)
	msg0 := mp.String("msg")
	data := mp["data"]
	return NewMsg(code, msg0, data)
}

func (m Msg) IsOk() bool {
	return m.Code == 200
}

func (m Msg) IsErr() bool {
	return !m.IsOk()
}
