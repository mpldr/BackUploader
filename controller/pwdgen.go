package controller

import (
	"math/rand"
	"time"
)

var (
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenPwd(length int) string {
	if length < 0 {
		length = 0
	}

	result := ""

	for i := 0; i < length; i++ {
		result += string(Alphabet[rand.Intn(len(Alphabet))])
	}

	return result
}
