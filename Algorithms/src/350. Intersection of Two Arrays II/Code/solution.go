func intersect(nums1 []int, nums2 []int) []int {
    // Step 1: Create a map to count occurrences of each number in nums1.
    countMap := make(map[int]int)
    var result []int

    // Step 2: Populate countMap with counts of each number in nums1.
    for _, num := range nums1 {
        countMap[num]++
    }

    // Step 3: Iterate through nums2 to find common elements.
    for _, num := range nums2 {
        // Step 4: If num exists in countMap and has remaining counts, add it to result.
        if countMap[num] > 0 {
            result = append(result, num)
            // Step 5: Decrease the count of num in countMap.
            countMap[num]--
        }
    }

    // Step 6: Return the result containing the intersection of nums1 and nums2.
    return result
}
