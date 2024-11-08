func getMaximumXor(nums []int, maximumBit int) []int {
    n := len(nums)
    answer := make([]int, n)
    XORed := 0
    
    // Calculate the cumulative XOR of the entire nums array
    for _, num := range nums {
        XORed ^= num
    }
    
    // max_k is 2^maximumBit - 1
    max_k := (1 << maximumBit) - 1
    
    // Process each query in reverse
    for i := 0; i < n; i++ {
        // Calculate the k that maximizes XOR
        answer[i] = XORed ^ max_k
        
        // Update XORed by removing the effect of the last element
        XORed ^= nums[n - 1 - i]
    }
    
    return answer
}
