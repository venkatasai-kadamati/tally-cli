package main

import (
	"fmt"
	"os"

	"github.com/venkatasai-kadamati/tally-cli/algo"
)

func main() {
	fmt.Println("Counter is running")
	// to read files and access system content, we use "os" package
	// to read files we will use ReadFile from OS package
	// syntax: func ReadFile(name string) ([]byte, error)
	// since its exposed the func name would be starting with caps

	data, _ := os.ReadFile("documents/words.txt")
	// data is slice of bytes, we need to type cast into ASCII string

	// call algorithm for counting words
	totalCount := algo.CountWords(string(data))
	fmt.Println("total count of words:", totalCount)

}
