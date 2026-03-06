func checkOnesSegment(s string) bool {
    // Check if substring "01" exists
    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' && s[i+1] == '1' {
            return false
        }
    }
    return true
}