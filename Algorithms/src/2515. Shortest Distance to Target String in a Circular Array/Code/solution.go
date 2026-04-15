func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    ans := int(^uint(0) >> 1) // Maximum possible int value

    // Traverse all indices
    for i := 0; i < n; i++ {
        // If current word matches target
        if words[i] == target {
            diff := i - startIndex
            if diff < 0 {
                diff = -diff
            }

            // Circular distance
            circularDist := n - diff

            // Minimum distance for this index
            current := diff
            if circularDist < current {
                current = circularDist
            }

            // Update answer
            if current < ans {
                ans = current
            }
        }
    }

    // If target was never found
    if ans == int(^uint(0)>>1) {
        return -1
    }

    return ans
}