func twoSum(nums []int, target int) []int {

	for i, j := range nums {
		for l, k := range nums[i+1:] {
			if (j + k) == target {
				return []int{i, l + i + 1}
			}
		}
	}
	return []int{10, 10}
}
