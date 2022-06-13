package sign

import (
	"github.com/elancom/go-util/collection"
	"github.com/elancom/go-util/crypto"
	"github.com/elancom/go-util/str"
	"sort"
	"strings"
)

func Sign(data map[string]string, key string) string {
	data = collection.Clone(data)

	// 移除关键字
	k1, k2 := "sign", "_sign"
	if data[k1] != "" || data[k2] != "" {
		delete(data, k1)
		delete(data, k2)
	}

	// 过滤空
	for k := range data {
		if str.IsBlank(data[k]) {
			delete(data, k)
		}
	}

	// 排序
	keys := collection.Keys(data)
	sort.Strings(keys)

	// 连接字符串
	b := strings.Builder{}
	for _, k := range keys {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(data[k])
	}
	s := b.String()

	// 摘要
	return Str(s, key)
}

func Str(data string, key string) string {
	bs := []byte(data + "&key=" + key)
	sign := crypto.Md6b(bs)
	return sign
}

func Check(key string, data map[string]string, signs ...string) bool {
	var signStr string
	if len(signs) == 0 {
		signStr = data["sign"]
	} else {
		signStr = signs[0]
	}
	return signStr != "" && signStr == Sign(data, key)
}

func CheckStr(data string, key string, sign string) bool {
	return Str(data, key) == sign
}
