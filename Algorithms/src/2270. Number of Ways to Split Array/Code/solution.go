func waysToSplitArray(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }
    
    prefixSum := 0
    count := 0
    
    for i := 0; i < len(nums)-1; i++ {
        prefixSum += nums[i]
        rightSum := totalSum - prefixSum
        if prefixSum >= rightSum {
            count++
        }
    }
    
    return count
}
