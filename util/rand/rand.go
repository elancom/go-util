package rand

import (
	cryRand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func init() {
	var seed, _ = cryRand.Int(cryRand.Reader, big.NewInt(math.MaxInt64))
	var seed64 = seed.Int64()
	rand.Seed(seed64)
}

var numbers = "0123456789"
var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numLetters = numbers + letters

// RandomIntRange [min,max)
func RandomIntRange(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		panic("max > min")
	}
	n := max - min
	return min + rand.Intn(n)
}

func RandomUint64() uint64 {
	return rand.Uint64()
}

func RandomNumber(num int) string {
	return random(numbers, num)
}

func RandomLetter(num int) string {
	return random(letters, num)
}

func RandomStr(num int) string {
	return random(numLetters, num)
}

func random(s string, num int) string {
	rb := make([]byte, num)
	sLen := len(s)
	for i := 0; i < num; i++ {
		rb[i] = s[rand.Intn(sLen)]
	}
	return string(rb)
}
