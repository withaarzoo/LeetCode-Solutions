func backtrack(nums []int, index int, currentOR int, maxOR int, count *int) {
    if currentOR == maxOR {
        *count++
    }
    
    for i := index; i < len(nums); i++ {
        backtrack(nums, i+1, currentOR|nums[i], maxOR, count)
    }
}

func countMaxOrSubsets(nums []int) int {
    maxOR := 0
    
    // Step 1: Compute the maximum OR
    for _, num := range nums {
        maxOR |= num
    }
    
    count := 0
    // Step 2: Backtrack to count the subsets
    backtrack(nums, 0, 0, maxOR, &count)
    
    return count
}
