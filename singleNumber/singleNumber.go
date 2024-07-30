package main

func singleNumber(nums []int) int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp ^= nums[i]
	}
	return tmp

}

func main() {
	// nums := []int{2, 2, 3, 2}
	// nums := []int{4, 1, 2, 1, 2}
	// nums := []int{1}
	nums := []int{2, 2, 1}
	println(singleNumber(nums))
}
