package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 5, 4, 8, 3, 7}))
}

// this gives answer but has o(n^2) which is not accepted ( Time Limit Exceeded )
func maxArea_first_try(height []int) int {

	tmp := 0
	for i := 0; i < len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i == j || height[i] > height[j] {
				continue
			} else {
				x_axis := func(i, j int) int {
					if i-j < 0 {
						return -1 * (i - j)
					} else {
						return i - j
					}
				}(i, j)
				area := x_axis * height[i]
				if area > tmp {
					tmp = area
				}
			}
		}
	}

	return tmp
}

// secound try
func maxArea(height []int) int {

	pointer_1 := 0
	pointer_2 := len(height) - 1
	tmp := 0
	y_axis := 0

	for pointer_1 != pointer_2 {
		x_axis := pointer_2 - pointer_1
		if height[pointer_1] > height[pointer_2] {
			y_axis = height[pointer_2]
			pointer_2 -= 1
		} else {
			y_axis = height[pointer_1]
			pointer_1 += 1
		}
		area := x_axis * y_axis
		if tmp < area {
			tmp = area
		}

	}
	return tmp
}
