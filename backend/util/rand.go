package util

import (
	"math/rand"
	"time"
)

var seededGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var stringCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateGameCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = stringCharset[seededGenerator.Intn(len(stringCharset))]
	}
	return string(b)
}

func GeneratePlayerID() string {
	b := make([]byte, 64)
	for i := range b {
		b[i] = stringCharset[seededGenerator.Intn(len(stringCharset))]
	}
	return string(b)
}

func RollDie() int {
	return seededGenerator.Intn(6) + 1
}
