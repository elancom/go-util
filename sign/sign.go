package sign

import (
	"crypto/md5"
	"fmt"
	"github.com/elancom/go-util/collection"
	"github.com/elancom/go-util/str"
	"sort"
	"strings"
)

func sign(key string, data map[string]string) string {
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

	// 连接秘钥
	b.WriteString("&key=")
	b.WriteString(key)
	s := b.String()

	// 摘要
	sum := md5.Sum([]byte(s))
	sum16 := fmt.Sprintf("%x", sum)

	return str.ToUpper(sum16)
}

func check(key string, data map[string]string, signs ...string) bool {
	var signStr string
	if len(signs) == 0 {
		signStr = data["sign"]
	} else {
		signStr = signs[0]
	}
	return signStr != "" && signStr == sign(key, data)
}
