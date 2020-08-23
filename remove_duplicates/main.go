package main

import "fmt"

func main() {
	test := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(test))
}

func removeDuplicates(nums []int) int {
	slow := 0
	n := len(nums)
	if n == 0 {
		return n
	}
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[slow] {
			nums[slow+1] = nums[fast]
			slow += 1
		}
	}
	slow += 1
	nums = nums[:slow]
	return slow
}
