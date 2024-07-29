package main

import "fmt"

func rotate(nums []int, k int) {

	rtrn_arr := make([]int, len(nums))

	for index, _ := range nums {
		if index+k < len(nums) {
			rtrn_arr[index+k] = nums[index]
		} else {
			rtrn_arr[((index + k) % len(nums))] = nums[index]
		}

	}

	// replace the original array with the rotated array
	for index, _ := range nums {
		nums[index] = rtrn_arr[index]
	}
	// fmt.Println(rtrn_arr)

}

func main() {
	// sample := []int{1, 2, 3, 4, 5, 6, 7}
	sample := []int{-1, -100, 3, 99}
	rotate(sample, 2)
	fmt.Println(sample)
}
