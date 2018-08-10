package main

import (
	"fmt"
	"os"
)

func isPalindrom(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	res := isPalindrom(os.Args[1])
	fmt.Println(res)
}
