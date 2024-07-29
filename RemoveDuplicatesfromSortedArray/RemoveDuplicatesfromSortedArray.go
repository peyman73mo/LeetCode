package main

import "fmt"

func removeDuplicates(nums []int) int {
	var tmp int = -100000
	var count int = 0
	var pntr int = -1

	for _, num := range nums {
		// fmt.Println("num: ", num, "tmp: ", tmp, "count: ", count)
		if num != tmp {
			pntr++
			nums[pntr] = num
			tmp = num
			count++
		}

	}
	// fmt.Println("final round: :", "tmp: ", tmp, "count: ", count)
	return count
}

func main() {
	// nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums), nums)
}
