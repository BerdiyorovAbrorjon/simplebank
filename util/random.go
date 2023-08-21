package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomFullName() string {
	return fmt.Sprintf("%s %s", RandomString(5), RandomString(5))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10))
}

func RandomUsername() string {
	return RandomString(int(RandomNumber(5, 10)))
}

func RandomMoney() int64 {
	return RandomNumber(0, 1000)
}

func RandomCurrency() string {
	c := []string{USD, UZS, EUR}
	n := len(c)
	return c[rand.Intn(n)]
}
