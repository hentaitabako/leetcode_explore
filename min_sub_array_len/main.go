package main

import "fmt"

func main() {
	test := []int{2, 3, 1, 2, 4, 3}

	fmt.Println(minSubArrayLen(15, test))
}

func minSubArrayLen(s int, nums []int) int {
	slow := 0
	l := len(nums)
	result := l
	sum := 0

	for fast := 0; fast < l; {
		sum += nums[fast]
		if sum >= s {
			if tmp2 := fast - slow; tmp2 < result {
				result = tmp2 + 1
			}
			slow++
			fast = slow
			sum = 0
		} else {
			if slow == 0 && fast == l-1 {
				return 0
			}
			fast++
		}

	}

	return result

}

func sumArray(nums []int, start int, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum = sum + nums[i]
	}

	return sum
}
