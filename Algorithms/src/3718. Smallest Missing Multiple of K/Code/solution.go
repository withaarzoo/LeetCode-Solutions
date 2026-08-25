func missingMultiple(nums []int, k int) int {
    // Create a map that works like a hash set for fast membership checks.
    present := make(map[int]bool)

    // Store every value from nums in the map.
    for _, num := range nums {
        present[num] = true
    }

    // Start with the smallest positive multiple of k.
    multiple := k

    // Keep checking multiples while the current one exists.
    for present[multiple] {
        // Move to the next positive multiple of k.
        multiple += k
    }

    // Return the first multiple that is missing.
    return multiple
}