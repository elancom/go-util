package http

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	get := Get("https://baidu.com", map[string]string{"id": "123"})
	log.Println(get.Data)
}

func TestPost(t *testing.T) {
	marshal, _ := json.Marshal(Post("http://120.79.214.77:10020/prod-api/login/login", map[string]any{"username": ""}))
	log.Println(string(marshal))
	msg := PostMsg("http://120.79.214.77:10020/prod-api/login/login", map[string]any{"username": ""})
	bytes, _ := json.Marshal(msg)
	log.Println(string(bytes))
}

func TestToMsg(t *testing.T) {
	m := make(map[string]any)
	m["code"] = 200
	m["msg"] = "aaa"
	marshal, _ := json.Marshal(m)
	msg := toMsg(string(marshal))
	log.Println(*msg)
}
