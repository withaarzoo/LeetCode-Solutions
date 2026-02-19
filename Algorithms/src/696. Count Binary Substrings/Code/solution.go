func countBinarySubstrings(s string) int {
    
    prevGroup := 0
    currGroup := 1
    result := 0
    
    for i := 1; i < len(s); i++ {
        
        if s[i] == s[i-1] {
            currGroup++
        } else {
            if prevGroup < currGroup {
                result += prevGroup
            } else {
                result += currGroup
            }
            
            prevGroup = currGroup
            currGroup = 1
        }
    }
    
    if prevGroup < currGroup {
        result += prevGroup
    } else {
        result += currGroup
    }
    
    return result
}
