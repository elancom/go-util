package http

import (
	"bytes"
	"encoding/json"
	"github.com/elancom/go-util/lang"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
)

func toMsg(s string) *lang.Msg {
	m := make(map[string]any)
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return lang.NewErr("解析结果失败")
	}
	return lang.NewFrom(m)
}

func GetMsg(url string, data ...map[string]string) *lang.Msg {
	get := Get(url, data...)
	if get.IsErr() {
		return get
	}
	return toMsg(get.Data.(string))
}

func Get(url string, data ...map[string]string) *lang.Msg {
	if data != nil && len(data) > 0 {
		bs := bytes.NewBufferString(url)
		bs.WriteString("?")
		for k, v := range data[0] {
			bs.WriteString(url2.QueryEscape(k))
			bs.WriteString("=")
			bs.WriteString(url2.QueryEscape(v))
		}
		url = bs.String()
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println("[GET]", err)
		return lang.NewErr("request err")
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[GET]read", err)
		return lang.NewErr("request err")
	}
	return lang.NewOk(string(bs))
}

func PostMsg(url string, data ...map[string]any) *lang.Msg {
	post := Post(url, data...)
	if post.IsErr() {
		return post
	}
	return toMsg(post.Data.(string))
}

func Post(url string, data ...map[string]any) *lang.Msg {
	if len(data) == 0 {
		data = append(data, make(map[string]any, 0))
	}

	jb, err := json.Marshal(data[0])
	if err != nil {
		return lang.NewErr("data to json error")
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jb))
	if err != nil {
		log.Println("[POST]", err)
		return lang.NewErr("request err")
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[POST]read", err)
		return lang.NewErr("request err")
	}

	return lang.NewOk(string(bs))
}

func PostFormMsg(url string, data ...map[string]string) *lang.Msg {
	form := PostForm(url, data...)
	if form.IsErr() {
		return form
	}
	return toMsg(form.Data.(string))
}

func PostForm(url string, data ...map[string]string) *lang.Msg {
	if len(data) == 0 {
		data = append(data, make(map[string]string, 0))
	}

	vs := url2.Values{}
	for k, v := range data[0] {
		vs.Set(k, v)
	}

	resp, err := http.PostForm(url, vs)
	if err != nil {
		log.Println("[PostForm]", err)
		return lang.NewErr("request err")
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[PostForm]read", err)
		return lang.NewErr("request err")
	}

	return lang.NewOk(string(bs))
}
