package controller

import (
	"math/rand"
	"time"
)

var (
	Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz01234567890"
	Length   = 16
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenPwd(length ...int) string {
	leng := 5
	if len(length) > 0 {
		leng = length[0]
	}

	result := ""

	for i := 0; i < leng; i++ {
		result += string(Alphabet[rand.Intn(len(Alphabet))])
	}

	return result
}
