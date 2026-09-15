func maxPalindromes(s string, k int) int {
    n := len(s)
    // isPal[i][j] == true  iff  s[i..j] is a palindrome
    isPal := make([][]bool, n)
    for i := range isPal {
        isPal[i] = make([]bool, n)
    }
    
    // every single character is a palindrome
    for i := 0; i < n; i++ {
        isPal[i][i] = true
    }
    
    // every pair of identical characters is a palindrome
    for i := 0; i+1 < n; i++ {
        if s[i] == s[i+1] {
            isPal[i][i+1] = true
        }
    }
    
    // longer lengths: ends equal and inner part already known to be palindrome
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            if s[i] == s[j] && isPal[i+1][j-1] {
                isPal[i][j] = true
            }
        }
    }
    
    // dp[i] = maximum number of valid pieces inside s[0..i-1]
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] // skip the last character
        // try every possible start of a piece that ends at i-1
        for j := 0; j <= i-k; j++ {
            if isPal[j][i-1] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }
    return dp[n]
}