func plusOne(digits []int) []int {
    // Traverse from last digit
    for i := len(digits) - 1; i >= 0; i-- {
        digits[i]++
        
        if digits[i] < 10 { // No carry
            return digits
        }
        
        digits[i] = 0 // Carry forward
    }
    
    // All digits were 9
    result := make([]int, len(digits)+1)
    result[0] = 1
    return result
}
