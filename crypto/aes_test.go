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
	encrypt, err := AesEcbEncrypt([]byte("123456"), bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt, _ := AesEcbDecrypt(encrypt, bytes)
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

	// rune (characters in Go are represented using `rune` data type)
	marshal, _ := json.Marshal(&user)
	marshal = []byte("{\"timestamp\":\"100000\",\"random\":\"100000\"}")
	bytes = []byte{49, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	b := byte(int(' '))
	bytes = []byte{49, b, b, b, b, b, b, b, b, b, b, b, b, b, b, b}
	encrypt, err := AesEcbEncrypt(marshal, bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt, _ := AesEcbDecrypt(encrypt, bytes)
	fmt.Println(string(decrypt))
}

func TestName3(t *testing.T) {
	d := "4ZP31mjs8Gk+7DrhAXp9LELpeHjmpIt+Jqpm+6K9AGXOxZ/CAtbhalKQGphVdsqbc2mgrGWVwGzxPlz7vCiAB6jIV8RBEIbkv6QG0B5b8VE="
	decodeString, err2 := base64.StdEncoding.DecodeString(d)
	if err2 != nil {
		t.Error(err2)
	}
	for i := 0; i < 100; i++ {
		key := Md6("login-password-key:1:11" + rand.RandomStr(10))
		decrypt, _ := AesEcbDecrypt(decodeString, []byte(key))
		log.Println(string(decrypt))
	}
}

func TestName4(t *testing.T) {

}
