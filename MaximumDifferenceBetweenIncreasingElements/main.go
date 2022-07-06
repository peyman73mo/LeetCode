package main

import "fmt"

func main() {
	fmt.Println(maximumDifference([]int{999, 997, 980, 976, 948, 940, 938, 928, 924, 917, 907, 907, 881, 878, 864, 862, 859, 857, 848, 840, 824, 824, 824, 805, 802, 798, 788, 777, 775, 766, 755, 748, 735, 732, 727, 705, 700, 697, 693, 679, 676, 644, 634, 624, 599, 596, 588, 583, 562, 558, 553, 539, 537, 536, 509, 491, 485, 483, 454, 449, 438, 425, 403, 368, 345, 327, 287, 285, 270, 263, 255, 248, 235, 234, 224, 221, 201, 189, 187, 183, 179, 168, 155, 153, 150, 144, 107, 102, 102, 87, 80, 57, 55, 49, 48, 45, 26, 26, 23, 15}))
}

func maximumDifference(nums []int) int {
	i_num := nums[0]
	// j_num := -1
	tmp := -1
	for i_index := 0; i_index < len(nums)-1; i_index++ {
		if nums[i_index] <= i_num {
			i_num = nums[i_index]
			for j_index := i_index + 1; j_index < len(nums); j_index++ {
				if nums[j_index]-nums[i_index] > tmp && nums[j_index] != nums[i_index] {
					tmp = nums[j_index] - nums[i_index]
					// i_num = i_index
					// j_num = j_index
				}
			}
		}
	}

	return tmp
}
