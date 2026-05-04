func rotate(matrix [][]int) {
    n := len(matrix)

    // Step 1: Transpose
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            // Swap elements across diagonal
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // Step 2: Reverse each row
    for i := 0; i < n; i++ {
        left, right := 0, n-1
        for left < right {
            matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
            left++
            right--
        }
    }
}