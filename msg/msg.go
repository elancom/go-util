package msg

import "github.com/elancom/go-util/lang"

// 快捷

type Msg = lang.Msg

func NewErr(msg string) *Msg {
	return lang.NewErr(msg)
}

func NewOk(data ...any) *Msg {
	return lang.NewOk(data)
}

func NewMsg(code int, msg string, data ...any) *Msg {
	return lang.NewMsg(code, msg, data...)
}
