func minOperations(boxes string) []int {
    n := len(boxes)
    answer := make([]int, n)

    // Left-to-right pass
    balls, operations := 0, 0
    for i := 0; i < n; i++ {
        answer[i] += operations
        if boxes[i] == '1' {
            balls++ // Count balls
        }
        operations += balls // Add the current number of balls to operations
    }

    // Right-to-left pass
    balls, operations = 0, 0
    for i := n - 1; i >= 0; i-- {
        answer[i] += operations
        if boxes[i] == '1' {
            balls++ // Count balls
        }
        operations += balls // Add the current number of balls to operations
    }

    return answer
}
