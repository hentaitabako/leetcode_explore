package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func arrayPairSum(nums []int) int {
	quickSort(nums, 0, len(nums))
	result := 0
	for i := 0; i < len(nums); {
		result += nums[i]
		i = i + 2
	}
	return result
}

func quickSort(source []int, l, u int) {
	if l < u {
		m := partition(source, l, u)
		quickSort(source, l, m-1)
		quickSort(source, m, u)
	}
}

func partition(source []int, l, u int) int { //划分
	var (
		pivot = source[l]
		left  = l
		right = l + 1
	)
	for ; right < u; right++ {
		if source[right] <= pivot {
			left++
			source[left], source[right] = source[right], source[left]
		}
	}
	source[l], source[left] = source[left], source[l]
	return left + 1
}
