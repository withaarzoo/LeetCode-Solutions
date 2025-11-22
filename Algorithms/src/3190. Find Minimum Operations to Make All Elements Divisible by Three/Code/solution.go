func minimumOperations(nums []int) int {
    operations := 0
    
    // Iterate over all numbers in the slice
    for _, x := range nums {
        // If x % 3 != 0, it takes exactly 1 operation to fix x
        if x%3 != 0 {
            operations++
        }
    }
    
    return operations
}
