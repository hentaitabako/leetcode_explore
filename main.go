package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(moveZeroes(nums))
}

func moveZeroes(nums []int) {
	l := len(nums)

	nonZeroNum := removeElement(nums, 0)

	for i := nonZeroNum; i < l; i++ {
		nums[i] = 0
	}
}

func removeElement(nums []int, target int) int {
	slow := 0
	n := len(nums)
	for fast := 0; fast < n; fast++ {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow += 1
		}
	}
	nums = nums[:slow]
	return slow
}
