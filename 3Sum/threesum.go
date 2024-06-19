package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		left_pointer := i + 1
		right_pointer := len(nums) - 1

		for left_pointer < right_pointer {
			sum := nums[left_pointer] + nums[right_pointer]

			if sum == target {
				if !contains(result, []int{nums[i], nums[left_pointer], nums[right_pointer]}) {
					result = append(result, []int{nums[i], nums[left_pointer], nums[right_pointer]})
				}
				left_pointer++
				right_pointer--
			} else if sum < target {
				left_pointer++
			} else {
				right_pointer--
			}

		}
	}

	return result
}

func contains(arr [][]int, item []int) bool {
	for _, a := range arr {
		if a[0] == item[0] && a[1] == item[1] && a[2] == item[2] {
			return true
		}
	}
	return false
}
