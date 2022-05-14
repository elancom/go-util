package crypto

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/elancom/go-util/rand"
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	bytes := []byte("1234567890123456")
	encrypt, err := ECBEncrypt([]byte("123456"), bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt, _ := ECBDecrypt(encrypt, bytes)
	fmt.Println(string(decrypt))
}

func TestName2(t *testing.T) {
	bytes := []byte("1234567890123456")
	user := map[string]any{
		"id":        "123456",
		"username":  "张三",
		"key":       rand.RandomStr(16),
		"timestamp": time.Now().UnixMilli(),
	}
	marshal, _ := json.Marshal(&user)
	encrypt, err := ECBEncrypt(marshal, bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt, _ := ECBDecrypt(encrypt, bytes)
	fmt.Println(string(decrypt))
}
