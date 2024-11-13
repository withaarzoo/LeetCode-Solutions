import (
    "sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
    var count int64 = 0
    n := len(nums)
    
    for i := 0; i < n-1; i++ {
        minVal := lower - nums[i]
        maxVal := upper - nums[i]
        
        start := sort.Search(n-i-1, func(j int) bool { return nums[i+1+j] >= minVal }) + i + 1
        end := sort.Search(n-i-1, func(j int) bool { return nums[i+1+j] > maxVal }) + i + 1
        
        count += int64(end - start)
    }
    
    return count
}