package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("vim-go")
	fmt.Printf("%q\n", strings.Split("hello world!        example", " "))
	r := reverseWord("hello")
	fmt.Printf("%v", r)
}

func reverseWords(s string) string {
	return "sb"
}

func reverseWord(s string) string {
	result := ""
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		result = result + string(s[i])
	}
	return result
}
