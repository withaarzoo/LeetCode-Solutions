func canReach(s string, minJump int, maxJump int) bool {

    n := len(s)

    // Queue for BFS traversal
    queue := []int{0}

    // Visited array
    visited := make([]bool, n)
    visited[0] = true

    // Pointer for queue front
    front := 0

    // Farthest processed index
    far := 0

    for front < len(queue) {

        i := queue[front]
        front++

        // If last index is reached
        if i == n-1 {
            return true
        }

        // Calculate valid jump range
        start := max(i+minJump, far+1)
        end := min(i+maxJump, n-1)

        // Explore all possible next positions
        for j := start; j <= end; j++ {

            // Only move to positions with '0'
            if s[j] == '0' && !visited[j] {
                visited[j] = true
                queue = append(queue, j)
            }
        }

        // Update farthest processed position
        if end > far {
            far = end
        }
    }

    return false
}

// Helper function for maximum
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Helper function for minimum
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}