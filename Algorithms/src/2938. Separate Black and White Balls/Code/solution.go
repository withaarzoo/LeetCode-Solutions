func minimumSteps(s string) int64 {
    var ans int64 = 0
    var blackCount int = 0 // Tracks the number of black balls (1s)

    // Traverse through the string
    for i := 0; i < len(s); i++ {
        if s[i] == '0' {
            // White ball encountered, add the number of black balls on its left
            ans += int64(blackCount)
        } else {
            // Black ball encountered, increment the black ball count
            blackCount++
        }
    }

    return ans
}
