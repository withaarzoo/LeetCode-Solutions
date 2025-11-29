package main

func minOperations(nums []int, k int) int {
    var sum int64 = 0
    
    // Calculate total sum of the array
    for _, x := range nums {
        sum += int64(x)
    }
    
    // Minimum operations equals sum % k
    remainder := int(sum % int64(k))
    
    return remainder
}
