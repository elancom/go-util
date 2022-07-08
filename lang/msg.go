package lang

const OK = 200
const Err = 400

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
	code := 0
	if c, ok := d["code"].(int); ok {
		code = c
	} else if c1, ok1 := d["code"].(float64); ok1 {
		code = int(c1)
	} else {
		return NewErr("msg code required int type")
	}
	var msg0 string
	if m, ok := d["msg"].(string); ok {
		msg0 = m
	}

	data := d["data"]

	return NewMsg(code, msg0, data)
}

func CastMsgToErr(err error) (*Msg, error) {
	msg, ok := err.(*Msg)
	if !ok {
		return nil, err
	}
	return msg, nil
}

// Msg 一等结构，安全输出
type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func (m *Msg) Error() string {
	return m.Msg
}

func (m *Msg) IsOk() bool {
	return m.Code == 200
}

func (m *Msg) IsErr() bool {
	return !m.IsOk()
}
