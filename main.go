package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Counter is running")
	// to read files and access system content, we use "os" package
	// to read files we will use ReadFile from OS package
	// syntax: func ReadFile(name string) ([]byte, error)
	// since its exposed the func name would be starting with caps

	data, err := os.ReadFile("documents/words.txt")
	// data is slice of bytes, we need to type cast into ASCII string
	fmt.Println("data", string(data)) // in a slice of bytes
	fmt.Println(err)                  // if its nil, no errors, read is done
}
