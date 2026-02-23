func hasAllCodes(s string, k int) bool {
    n := len(s)
    
    if n < k {
        return false
    }
    
    total := 1 << k
    if n-k+1 < total {
        return false
    }
    
    seen := make([]bool, total)
    mask := total - 1
    
    curr := 0
    count := 0
    
    // First window
    for i := 0; i < k; i++ {
        curr = (curr << 1) | int(s[i]-'0')
    }
    
    if !seen[curr] {
        seen[curr] = true
        count++
    }
    
    // Sliding window
    for i := k; i < n; i++ {
        curr = ((curr << 1) & mask) | int(s[i]-'0')
        
        if !seen[curr] {
            seen[curr] = true
            count++
            if count == total {
                return true
            }
        }
    }
    
    return count == total
}