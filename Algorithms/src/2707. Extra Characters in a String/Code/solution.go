// Function to find the minimum number of extra characters
// required to split the string `s` into valid words from the dictionary.
func minExtraChar(s string, dictionary []string) int {
    
    // Step 1: Initialize a map for quick lookup of dictionary words.
    // This allows O(1) lookup time for any word in the dictionary.
    dict := make(map[string]bool)
    for _, word := range dictionary {
        dict[word] = true // Add each word in the dictionary to the map.
    }

    // Step 2: Define the length of the string `s`.
    n := len(s)

    // Step 3: Initialize a DP (Dynamic Programming) array `dp` of size `n+1`.
    // The DP array will store the minimum extra characters needed up to index `i`.
    dp := make([]int, n+1)

    // Step 4: Set a large initial value in the DP array to represent the worst case.
    // Initially, we assume all characters are extra, so we set each `dp[i]` to `n` (maximum possible).
    for i := 0; i <= n; i++ {
        dp[i] = n
    }

    // Step 5: The base case.
    // dp[0] represents no characters, hence 0 extra characters are needed for an empty string.
    dp[0] = 0

    // Step 6: Iterate through each index `i` of the string `s` (from 1 to n).
    // `i` represents the current position in the string we are processing.
    for i := 1; i <= n; i++ {

        // Step 7: For each `i`, we consider all substrings that end at `i`.
        // This involves checking every possible starting index `j` (0 to i-1).
        for j := 0; j < i; j++ {
            // Extract the substring `s[j:i]` (substring from index `j` to `i`).
            sub := s[j:i]

            // Step 8: Check if this substring exists in the dictionary.
            if dict[sub] {
                // If the substring is a valid word, update `dp[i]` to the minimum between
                // its current value and `dp[j]`, meaning we can split the string at `j`.
                dp[i] = min(dp[i], dp[j])
            }
        }

        // Step 9: In case no valid split was found, consider the current character
        // `s[i-1]` as an extra character and update `dp[i]` accordingly.
        dp[i] = min(dp[i], dp[i-1]+1)
    }

    // Step 10: The result is stored in `dp[n]`, which represents the minimum extra
    // characters needed for the entire string `s`.
    return dp[n]
}

// Helper function to return the minimum of two integers `a` and `b`.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
