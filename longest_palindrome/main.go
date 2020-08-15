package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	test := longestPalindrome("teteeee")
	fmt.Printf("%v", test)
}

func longestPalindrome(s string) string {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			tmp := s[j : l-i+j]
			if isPalindrome(tmp) {
				return tmp
			}
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	l := len(s)
	if l == 1 {
		return true
	}

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

// 反转字符串的方式来判断一个字符串是否为回文,比较耗时
// 有几种字符串拼接方式,1. +; 2. Joins 3. Builder指定cap提示性能
func reverseString(s string) string {
	var b strings.Builder
	l := len(s)
	b.Grow(l)
	loop := l - 1
	for i := loop; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}
	return b.String()
}
