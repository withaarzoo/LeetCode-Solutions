func findThePrefixCommonArray(A []int, B []int) []int {
    
    n := len(A)
    
    // Frequency array to track appearances
    freq := make([]int, n+1)
    
    // Result array
    ans := make([]int, n)
    
    // Stores count of common elements
    common := 0
    
    for i := 0; i < n; i++ {
        
        // Add current element from A
        freq[A[i]]++
        
        // If frequency becomes 2,
        // number appeared in both arrays
        if freq[A[i]] == 2 {
            common++
        }
        
        // Add current element from B
        freq[B[i]]++
        
        // Same check for B
        if freq[B[i]] == 2 {
            common++
        }
        
        // Store current prefix answer
        ans[i] = common
    }
    
    return ans
}