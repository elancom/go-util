package rand

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func RandLow(n int) []byte {
	if n <= 0 {
		return []byte{}
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return []byte{}
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return b
}

func TestName(t *testing.T) {
	min := 0
	max := 2
	for i := 0; i < 10; i++ {
		fmt.Println(RandomIntRange(min, max+1))
	}
	//fmt.Printf(string(RandLow(10)))

	//fmt.Println(RandomStr(10))
	//fmt.Println(RandomNumber(10))
	//fmt.Println(RandomLetter(10))

	//1、Int
	//n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	//if err == nil {
	//	fmt.Println("rand.Int：", n, n.BitLen())
	//}
	////2、Prime
	//p, err := rand.Prime(rand.Reader, 5)
	//if err == nil {
	//	fmt.Println("rand.Prime：", p)
	//}
	////3、Read
	//b := make([]byte, 32)
	//m, err := rand.Read(b)
	//if err == nil {
	//	fmt.Println("rand.Read：", b[:m])
	//	fmt.Println("rand.Read：", base64.URLEncoding.EncodeToString(b))
	//}

	//var numbers = "0123456789"
	//var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//var numLetters = numbers + letters
	//b := make([]byte, 10)
	//read, _ := rand.Read(b) // []byte(numLetters)
	//fmt.Println(read)

	//b := make([]byte, 10)
	//// rand.Seed(time.Now().UnixNano())
	//if _, err := rand.Read(b[:]); err != nil {
	//	fmt.Println("err")
	//}
	//fmt.Println(b)
}
