package main

import "fmt"

func main() {
	fmt.Println("vim-go")

	testList := [][]int{{0, 1, 2, 0}, {3, 5, 4, 2}, {1, 3, 1, 5}}

	setZeros(testList)
}

func setZeros(matrix [][]int) {
	var cols, rows []int

	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	noRepeatRows := removeRepeat(rows)
	noRepeatCols := removeRepeat(cols)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if findIndex(noRepeatRows, i) != -1 || findIndex(noRepeatCols, j) != -1 {
				matrix[i][j] = 0
			}
		}
	}
}

func removeRepeat(list []int) []int {
	temMap := make(map[int]int)

	var result []int
	for _, value := range list {
		temMap[value] = value
	}

	for key := range temMap {
		result = append(result, key)
	}
	return result
}

func findIndex(list []int, value int) int {
	for index, i := range list {
		if i == value {
			return index
		}
	}

	return -1
}
