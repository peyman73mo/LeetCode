func intersect(nums1 []int, nums2 []int) []int {
    
    output := []int{}
    if len(nums1) >= len(nums2){
        for _, item := range nums2 {
            for index, element := range nums1 {
                if item == element {
                    output = append(output, item)
                    nums1[index] = -1
                    break
                }
            }
        }
    }else {
        for _, item := range nums1 {
            for index, element := range nums2 {
                if item == element {
                    output = append(output, item)
                    nums2[index] = -1
                    break
                }
            }
        }
    }
    return output
    
}