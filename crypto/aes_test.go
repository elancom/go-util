package crypto

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	bytes := []byte("12345678901234561234567890123456")
	encrypt, err := ECBEncrypt([]byte("123456"), bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))
	decrypt, _ := ECBDecrypt(encrypt, bytes)
	fmt.Println(string(decrypt))
}
