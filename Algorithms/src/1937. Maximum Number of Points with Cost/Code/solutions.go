func maxPoints(points [][]int) int64 {
    // m: Number of rows in the points matrix
    // n: Number of columns in the points matrix
    m, n := len(points), len(points[0])

    // dp: Array to store the maximum points obtainable for each column in the current row
    dp := make([]int64, n)

    // Initialize dp with the values from the first row of the points matrix
    for j := 0; j < n; j++ {
        dp[j] = int64(points[0][j])
    }

    // Iterate over each row starting from the second row
    for i := 1; i < m; i++ {
        // leftMax: Array to store the maximum points obtainable moving leftward in the row
        leftMax := make([]int64, n)
        // rightMax: Array to store the maximum points obtainable moving rightward in the row
        rightMax := make([]int64, n)
        // newDp: Array to store the updated maximum points for the current row
        newDp := make([]int64, n)

        // Calculate the leftMax array:
        // Start with the first element, as there are no elements to the left of it
        leftMax[0] = dp[0]
        // Fill the leftMax array by moving left to right
        for j := 1; j < n; j++ {
            // The value at leftMax[j] is the maximum of the previous leftMax or the current dp value plus the column index
            leftMax[j] = max(leftMax[j-1], dp[j]+int64(j))
        }

        // Calculate the rightMax array:
        // Start with the last element, as there are no elements to the right of it
        rightMax[n-1] = dp[n-1] - int64(n-1)
        // Fill the rightMax array by moving right to left
        for j := n-2; j >= 0; j-- {
            // The value at rightMax[j] is the maximum of the next rightMax or the current dp value minus the column index
            rightMax[j] = max(rightMax[j+1], dp[j]-int64(j))
        }

        // Calculate the newDp array for the current row:
        for j := 0; j < n; j++ {
            // The value at newDp[j] is the maximum of leftMax or rightMax adjusted by the column index, plus the current points value
            newDp[j] = max(leftMax[j]-int64(j), rightMax[j]+int64(j)) + int64(points[i][j])
        }

        // Update dp to be the newDp for the next iteration
        dp = newDp
    }

    // Find the maximum value in the dp array which represents the maximum points obtainable
    result := dp[0]
    for j := 1; j < n; j++ {
        result = max(result, dp[j])
    }

    // Return the maximum points obtainable
    return result
}

// max: Helper function to return the maximum of two int64 values
func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
