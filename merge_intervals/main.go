package main

import "fmt"

func main() {
	var intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Printf("bubbleSort result %v\n", bubbleSort(intervals))
	fmt.Printf("merge result %v\n", merge(intervals))
}

func merge(intervals [][]int) [][]int {
	sortIntervals := bubbleSort(intervals)
	result := sortIntervals[0:1]
	fmt.Print(result, "\n")
	for i := 1; i < len(sortIntervals); i++ {
		if sortIntervals[i][0] <= result[len(result)-1][1] {
			if result[
			result[len(result)-1][1] = sortIntervals[i][1]
		} else {
			result = append(result, sortIntervals[i])
		}
	}
	return result

}

func bubbleSort(arr [][]int) [][]int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j][0] > arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
