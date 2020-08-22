package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("hello world"))
}

func reverseWords(s string) string {
	splitWords := strings.Split(s, " ")
	result := ""
	for i := 0; i < len(splitWords); i++ {
		if i == 0 {
			result = result + reverseWord(splitWords[i])
		} else {
			result = result + " " + reverseWord(splitWords[i])
		}
	}
	return result
}

func reverseWord(s string) string {
	ss := []rune(s)

	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[len(ss)-1-i] = ss[len(ss)-1-i], ss[i]
	}
	return string(ss)
}
