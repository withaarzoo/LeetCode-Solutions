func canBeValid(s string, locked string) bool {
    if len(s)%2 != 0 {
        return false // Odd length can't be balanced
    }
    
    open, flexible := 0, 0
    // Left-to-right pass
    for i := 0; i < len(s); i++ {
        if locked[i] == '1' {
            if s[i] == '(' {
                open++
            } else {
                open--
            }
        } else {
            flexible++
        }
        if open+flexible < 0 {
            return false
        }
    }
    
    open, flexible = 0, 0
    // Right-to-left pass
    for i := len(s) - 1; i >= 0; i-- {
        if locked[i] == '1' {
            if s[i] == ')' {
                open++
            } else {
                open--
            }
        } else {
            flexible++
        }
        if open+flexible < 0 {
            return false
        }
    }
    
    return true
}
