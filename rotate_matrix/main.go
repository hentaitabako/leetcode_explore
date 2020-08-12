package main

import "fmt"

func main() {
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate_matrix(matrix)
}

func rotate_matrix(matrix [][]int) {
	fmt.Printf("%v", matrix)
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if i < j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}

			if j > l/2 {
				matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
			}
		}
	}
	fmt.Printf("%v", matrix)
}
