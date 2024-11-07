func largestCombination(candidates []int) int {
    var bitCount [31]int // Array to count '1's at each bit position
    
    // Count '1's in each bit position across all numbers
    for _, num := range candidates {
        for i := 0; i < 31; i++ {
            if num & (1 << i) != 0 {
                bitCount[i]++
            }
        }
    }
    
    // Find the maximum count in any bit position
    maxCombinationSize := 0
    for _, count := range bitCount {
        if count > maxCombinationSize {
            maxCombinationSize = count
        }
    }
    
    return maxCombinationSize
}
