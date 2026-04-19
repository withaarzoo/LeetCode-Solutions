func maxDistance(nums1 []int, nums2 []int) int {
    i, j := 0, 0
    ans := 0

    for i < len(nums1) && j < len(nums2) {

        // Ensure i <= j
        if i > j {
            j++
            continue
        }

        // Valid pair
        if nums1[i] <= nums2[j] {
            if j-i > ans {
                ans = j - i
            }
            j++ // Try for larger distance
        } else {
            // Invalid pair
            i++
        }
    }

    return ans
}