package main

import (
	"bytes"
	"time"
	"math/rand"
)

func generateKey() []byte {
	rand.Seed(time.Now().UnixNano())
	alphabet := []byte("abcdefghijklmnopqrstuvwxyz")
	size := len(alphabet)
	var key []byte
	for i := 0; i < size; i++ {
		letter := alphabet[rand.Intn(len(alphabet))]
		key = append(key, letter)
		index := bytes.IndexByte(alphabet, letter)
		alphabet = append(alphabet[:index], alphabet[index + 1:]...)
	}
	return key
}

func encrypt(text []byte) ([]byte, []byte) {
	key := generateKey()
	var result bytes.Buffer
	for _, char := range text {
		result.WriteString(string(key[char - 97]))
	}
	return result.Bytes(), key
}

func decrypt(text []byte, key []byte) []byte {
	var result bytes.Buffer
	for _, char := range text {
		result.WriteString(string(byte(bytes.IndexByte(key, char) + 97)))
	}
	return result.Bytes()
}