func stoneGameVIII(stones []int) int {
    // Convert the original array into prefix sums in place.
    for i := 1; i < len(stones); i++ {
        stones[i] += stones[i-1]
    }

    // Taking all stones gives the total prefix sum as the starting answer.
    best := stones[len(stones)-1]

    // Check every smaller valid prefix from right to left.
    for i := len(stones) - 2; i >= 1; i-- {
        // Keep the current best answer or take this prefix
        // and subtract the best result the opponent can achieve.
        best = max(best, stones[i]-best)
    }

    // Return the maximum score difference.
    return best
}

func max(a, b int) int {
    // Return the larger value needed for the DP transition.
    if a > b {
        return a
    }

    // Return b when it is greater than or equal to a.
    return b
}