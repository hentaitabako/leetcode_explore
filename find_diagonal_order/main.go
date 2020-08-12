package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	test := [][]int{{1, 2, 3}}

	result := findDiagonalOrder(test)
	fmt.Printf("%v", result)
}

func findDiagonalOrder(matrix [][]int) []int {
	// 获取矩阵维数
	m, n := len(matrix), len(matrix[0])

	// 对角线循环次数
	loop := m + n - 1

	var result []int

	// k是当前坐标和
	for k := 0; k < loop; k++ {

		// 开始位置
		i, j := 0, 0

		// 确定内循环次数
		innerLoop := k + 1

		for l := 0; l < innerLoop; l++ {
			//j = k - i
			if k%2 == 0 {
				i = k - j
			} else {
				j = k - i
			}
			if j >= 0 && j < n && i >= 0 && i < m {
				result = append(result, matrix[i][j])
			}
			// i++ or j++
			//	i++
			if k%2 == 0 {
				j++
			} else {
				i++
			}
		}
	}
	return result
}
