func rotatedDigits(n int) int {
    count := 0 // total good numbers
    
    for i := 1; i <= n; i++ {
        num := i
        isValid := true   // assume valid
        hasChange := false // check if it changes
        
        for num > 0 {
            digit := num % 10 // extract last digit
            
            // invalid digits
            if digit == 3 || digit == 4 || digit == 7 {
                isValid = false
                break
            }
            
            // digits that change
            if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
                hasChange = true
            }
            
            num /= 10 // remove last digit
        }
        
        if isValid && hasChange {
            count++
        }
    }
    
    return count
}