package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(test))
}

func findMin(nums []int) int {
	result := nums[0]

	if len(nums) == 1 {
		return result
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < 1 {
			result = nums[i]
		}
	}
	return result
}
