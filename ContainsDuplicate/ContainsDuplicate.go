package main

import (
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false

}

func main() {
	// nums := []int{1, 2, 3, 1}
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	println(containsDuplicate(nums))
}
