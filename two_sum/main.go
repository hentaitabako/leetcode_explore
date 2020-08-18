package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := []int{2, 7, 11, 15}
	result := twoSum(test, 9)
	fmt.Printf("%v", result)
}

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	var result []int
	for i := 0; i < len(numbers); i++ {
		for j := l - 1; j > i; j-- {
			if numbers[i]+numbers[j] == target {
				result = append(result, i+1)
				result = append(result, j+1)
				return result
			}
		}
	}
	return result
}
