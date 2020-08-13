package main

import "fmt"

func main() {
	//test := []string{"flower", "flow", "flight"}
	var test []string
	prefix := longestCommonPrefix(test)
	fmt.Printf("%v", prefix)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	loop := findShortestLength(strs)
	commonPrefix := ""
	for i := 0; i < loop; i++ {
		prefix := strs[0][0 : i+1]
		for j := 1; j < len(strs); j++ {
			if prefix == strs[j][0:i+1] {
				if j == len(strs)-1 {
					commonPrefix = prefix
				}
				continue
			} else {
				break
			}
		}
	}

	return commonPrefix
}

func findShortestLength(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	result := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < result {
			result = len(strs[i])
		}
	}
	return result
}
