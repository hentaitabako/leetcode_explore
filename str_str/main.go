package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	haystack := "hello"
	needle := "ll"
	fmt.Printf("%v", strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {

	m, n := len(haystack), len(needle)

	if haystack == needle || n == 0 {
		return 0
	}

	if n > m {
		return -1
	}

	positionInH := 0
	for i := 0; i < m-n+1; i++ {
		for j := 0; j < n; j++ {
			if needle[j] != haystack[positionInH+j] {
				positionInH++
				break
			} else {
				if j == n-1 {
					return positionInH
				}
			}
		}
	}

	return -1
}
