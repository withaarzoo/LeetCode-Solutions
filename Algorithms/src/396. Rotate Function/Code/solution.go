func maxRotateFunction(nums []int) int {
    n := len(nums)
    
    var sum int64 = 0 // total sum
    var F int64 = 0   // F(0)
    
    // Step 1: compute sum and F(0)
    for i := 0; i < n; i++ {
        sum += int64(nums[i])
        F += int64(i) * int64(nums[i])
    }
    
    result := F
    
    // Step 2: compute next rotations
    for k := 1; k < n; k++ {
        F = F + sum - int64(n)*int64(nums[n-k])
        
        if F > result {
            result = F
        }
    }
    
    return int(result)
}