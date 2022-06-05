package msg

import "github.com/elancom/go-util/lang"

// 快捷

const OK = lang.OK
const Err = lang.Err

func NewErr(msg string) *lang.Msg {
	return lang.NewErr(msg)
}

func NewOk(data ...any) *lang.Msg {
	return lang.NewOk(data...)
}

func NewMsg(code int, msg string, data ...any) *lang.Msg {
	return lang.NewMsg(code, msg, data...)
}
