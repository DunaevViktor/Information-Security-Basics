package main

import (
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	way := os.Args[1]
	inputPath := os.Args[2]
	outputPath := os.Args[3]
	keyPath := os.Args[4]

	input, err := ioutil.ReadFile(inputPath)
	check(err)

	var result, key []byte

	switch way {
	case "encrypt": 
		result, key = encrypt(input)
		err = ioutil.WriteFile(outputPath, result, 0644)
		check(err)
		err = ioutil.WriteFile(keyPath, key, 0644)
		check(err)
	case "decrypt":
		key, err = ioutil.ReadFile(keyPath)
		check(err)
		result = decrypt(input, key)
		err = ioutil.WriteFile(outputPath, result, 0644)
		check(err)
	default:
		panic("wrong way")
	}

}
