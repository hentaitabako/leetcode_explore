package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(generate(25))
}

func generate(numRows int) [][]int {
	var result [][]int
	if numRows >= 1 {
		result = [][]int{{1}}
	}
	if numRows >= 2 {
		result = [][]int{{1}, {1, 1}}
	}
	for i := 2; i < numRows; i++ {
		tmp := generateOneLine(result[i-1])
		result = append(result, tmp)
	}
	return result
}

func generateOneLine(before []int) []int {
	var result = []int{1}
	for i := 0; i < len(before)-1; i++ {
		result = append(result, before[i]+before[i+1])
	}
	result = append(result, 1)
	return result
}
