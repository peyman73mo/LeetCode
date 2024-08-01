package main

import "fmt"

func plusOne(digits []int) []int {

	carry_over := 1
	var tmp int
	for i := len(digits) - 1; i >= 0; i-- {
		tmp = digits[i]
		digits[i] = (digits[i] + carry_over) % 10
		// fmt.Println(digits[i])
		carry_over = (tmp + carry_over) / 10
		// fmt.Println(carry_over)
		if i == 0 && carry_over > 0 {
			digits = append([]int{carry_over}, digits...)
		}
	}
	return digits
}

func main() {
	// nums := []int{9, 9, 9}
	nums := []int{1, 2, 3}
	// nums := []int{4, 3, 2, 1}
	fmt.Println(plusOne(nums))
}
