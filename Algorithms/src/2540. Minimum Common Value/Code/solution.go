func getCommon(nums1 []int, nums2 []int) int {
    
    // Pointer for nums1
    i := 0

    // Pointer for nums2
    j := 0

    // Traverse both arrays
    for i < len(nums1) && j < len(nums2) {

        // If both values are equal,
        // return the answer
        if nums1[i] == nums2[j] {
            return nums1[i]
        }

        // Move the pointer with smaller value
        if nums1[i] < nums2[j] {
            i++
        } else {
            j++
        }
    }

    // No common value found
    return -1
}