package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("vim-go")
	fmt.Printf("%v\n", reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	splitWords := strings.Split(s, " ")
	result := ""
	for i := len(splitWords) - 1; i >= 0; i-- {
		if len(splitWords[i]) != 0 {
			if len(result) == 0 {
				result = result + splitWords[i]
			} else {
				result = result + " " + splitWords[i]
			}
		}
	}
	return result
}

func reverseWord(s string) string {
	result := ""
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		result = result + " " + string(s[i])
	}
	return result
}
