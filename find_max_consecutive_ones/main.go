package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := []int{1, 0, 0, 1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(test))

}

func findMaxConsecutiveOnes(nums []int) int {
	nums = append(nums, 0)
	slow := 0
	result := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] < 1 {
			if tmp := fast - slow; tmp > result {
				result = tmp
			}
			slow = fast + 1
		}
	}
	return result
}
