package main

import "fmt"

func main() {
	arr1 := []int{1, 3}
	arr2 := []int{2, 7}
	fmt.Println(findMedianSortedArrays(arr1, arr2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	arr := make([]int, m+n)
	i, j := 0, 0
	if m == 0 && n != 0 {
		arr = nums2
		return median(arr)
	} else if n == 0 && m != 0 {
		arr = nums1
		return median(arr)
	} else if n == 0 && m == 0 {
		return 0
	}
	for index, _ := range arr {

		if i == m {
			arr = append(arr[0:index], nums2[j:]...)
			break
		}
		if j == n {
			arr = append(arr[0:index], nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			fmt.Println(nums1[i])
			arr[index] = nums1[i]
			i += 1
		} else {
			fmt.Println(nums2[j])
			arr[index] = nums2[j]
			j += 1
		}

	}
	fmt.Println(arr)
	return median(arr)
}

func median(arr []int) float64 {
	if len(arr)%2 == 0 {
		return float64(arr[(len(arr)/2)-1]+arr[(len(arr)/2)]) / float64(2)
	} else {
		return float64((arr[len(arr)/2]))
	}
}
