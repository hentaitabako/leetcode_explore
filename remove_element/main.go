package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := []int{0, 1, 2, 2, 3, 0, 4, 2}
	length := removeElement(test, 2)
	fmt.Printf("%v \n", test)
	fmt.Printf("%v \n", length)
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
