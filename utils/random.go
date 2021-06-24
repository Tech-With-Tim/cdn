package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//RandomInt Generates a random number between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // +1 if max - min = 0
}

//RandomString Generates  a Random String of N characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(charset)

	for i := 0; i < n; i++ {
		c := charset[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func StrToBinary(s string, base int) []byte {

	var b []byte

	for _, c := range s {
		b = strconv.AppendInt(b, int64(c), base)
	}

	return b
}
