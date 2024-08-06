package main

import "fmt"

func setZeroes(matrix [][]int) {

	m := len(matrix)
	n := len(matrix[0])
	finds := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				// fmt.Println("find zero at ", i, j)
				finds = append(finds, []int{i, j})
			}
		}
	}
	// fmt.Println(finds)
	for _, find := range finds {

		setRowsZero(matrix, find[0])
		setColsZero(matrix, find[1])
	}
}

func setRowsZero(matrix [][]int, row int) {
	for i := 0; i < len(matrix[row]); i++ {
		// fmt.Println("set row to zero", row, i)
		matrix[row][i] = 0
	}
}
func setColsZero(matrix [][]int, col int) {
	for i := 0; i < len(matrix); i++ {
		// fmt.Println("set col to zero", i, col)
		matrix[i][col] = 0
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
