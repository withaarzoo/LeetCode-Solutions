func minPartitions(n string) int {
    maxDigit := 0 // Store maximum digit
    
    for i := 0; i < len(n); i++ {
        digit := int(n[i] - '0') // Convert byte to int
        
        if digit > maxDigit {
            maxDigit = digit
        }
        
        // If 9 found, break early
        if maxDigit == 9 {
            break
        }
    }
    
    return maxDigit
}