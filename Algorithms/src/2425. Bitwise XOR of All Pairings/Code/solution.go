func xorAllNums(nums1 []int, nums2 []int) int {
    xor1, xor2 := 0, 0

    // XOR all elements in nums1
    for _, num := range nums1 {
        xor1 ^= num
    }

    // XOR all elements in nums2
    for _, num := range nums2 {
        xor2 ^= num
    }

    // If nums1 has odd length, include xor2
    // If nums2 has odd length, include xor1
    if len(nums1)%2 == 1 {
        xor2 ^= xor1
    }
    if len(nums2)%2 == 1 {
        xor2 ^= xor1
    }
    
    return xor2
}
