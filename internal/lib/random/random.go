package random

import (
	//"crypto/rand"
	"math/rand"
	"time"
)

// Генерирует случайную строку заданной длины из букв и цифр. (не безопасно, т.к. используется rand.Intn)
func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}
	return string(b)
}

// Безопасный вариант с библиотекой crypto/rand
// func NewRandomCryptoString(size int) (string, error) {
// 	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
// 		"abcdefghijklmnopqrstuvwxyz" +
// 		"0123456789")
// 	b := make([]rune, size)
// 	for i := range b {
// 		// Чтобы заработало, импортируй "crypto/rand"
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
// 		if err != nil {
// 			return "", err
// 		}
// 		b[i] = chars[num.Int64()]
// 	}
// 	return string(b), nil
// }
