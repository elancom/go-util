package json

import (
	"encoding/json"
	"github.com/elancom/go-util/lang"
	"log"
)

func ToJson(v any) (string, error) {
	marshal, err := json.Marshal(v)
	if err != nil {
		return "", lang.NewErr("json deserialize err")
	}
	return string(marshal), nil
}

func ToJsonStr(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Println("json deserialize err")
		return ""
	}
	return string(marshal)
}

func ToObj[T any](s string, v T) (T, error) {
	err := json.Unmarshal([]byte(s), v)
	if err != nil {
		return v, lang.NewErr("json serialize err")
	}
	return v, nil
}
