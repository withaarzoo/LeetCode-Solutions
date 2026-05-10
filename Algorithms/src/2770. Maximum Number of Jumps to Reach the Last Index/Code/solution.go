func maximumJumps(nums []int, target int) int {
    
    n := len(nums)

    // dp[i] = maximum jumps needed to reach index i
    dp := make([]int, n)

    // Initialize all values as -1
    for i := 0; i < n; i++ {
        dp[i] = -1
    }

    // Starting index
    dp[0] = 0

    // Try every current index
    for i := 0; i < n; i++ {

        // Skip unreachable indices
        if dp[i] == -1 {
            continue
        }

        // Try every next index
        for j := i + 1; j < n; j++ {

            // Difference between values
            diff := nums[j] - nums[i]

            // Check valid jump
            if diff >= -target && diff <= target {

                // Update maximum jumps
                if dp[i]+1 > dp[j] {
                    dp[j] = dp[i] + 1
                }
            }
        }
    }

    // Final answer
    return dp[n-1]
}