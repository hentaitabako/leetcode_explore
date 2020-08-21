package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	s := []byte{'s', 'b', 'a'}
	fmt.Printf("%v", s)
	reverseString(s)
	fmt.Printf("%v", s)
}

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}
