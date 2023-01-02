package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var letterMap map[rune]int

func init() {
	letterMap = createLetterMap()
}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	sum := 0
	for i := range b {
		index := rand.Intn(len(letterRunes))
		b[i] = letterRunes[index]
		sum += index
	}
	return string(b) + strconv.Itoa(sum)
}

func VerifyUrlId(shortUrl string) bool {
	if len(shortUrl) < 6 {
		return false
	}
	urlId := shortUrl[0:6]
	checkSum := shortUrl[6:]

	sum := 0
	for _, v := range urlId {
		sum += letterMap[v]
	}
	return strconv.Itoa(sum) == checkSum
}

func createLetterMap() map[rune]int {
	lm := make(map[rune]int)
	for i, v := range letterRunes {
		lm[v] = i
	}
	return lm
}
