package json

import (
	"encoding/json"
	"github.com/elancom/go-util/lang"
)

func ToJson(v any) (string, error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "", lang.NewErr("json convert err")
	}
	return string(marshal), nil
}

func ToObj(s string, v any) error {
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		return lang.NewErr("json convert err")
	}
	return nil
}
