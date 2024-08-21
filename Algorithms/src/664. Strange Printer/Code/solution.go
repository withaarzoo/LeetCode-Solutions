package main

// strangePrinter calculates the minimum number of turns required to print the string `s`.
func strangePrinter(s string) int {
    n := len(s) // Length of the string
    // Create a 2D slice `dp` where dp[i][j] represents the minimum turns needed to print substring s[i:j+1]
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n) // Initialize each row of the 2D slice
    }

    // Fill the DP table by considering substrings of various lengths
    for i := n - 1; i >= 0; i-- { // Start from the end of the string and move backwards
        dp[i][i] = 1 // A single character substring requires exactly 1 turn to print

        for j := i + 1; j < n; j++ { // For every substring starting at `i` and ending at `j`
            // Initial assumption: Print the character `s[j]` separately after printing s[i:j]
            dp[i][j] = dp[i][j-1] + 1

            // Check if there's a character within s[i:j] that matches `s[j]`
            for k := i; k < j; k++ {
                if s[k] == s[j] {
                    // If a match is found, the turns to print s[i:j] could be optimized
                    // by combining the printing of s[k] and s[j]
                    dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j-1])
                }
            }
        }
    }

    // Return the minimum turns needed to print the entire string s[0:n]
    return dp[0][n-1]
}

// min returns the smaller of two integers `a` and `b`
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
