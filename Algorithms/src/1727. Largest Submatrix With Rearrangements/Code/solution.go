import "sort"

func largestSubmatrix(matrix [][]int) int {
    m, n := len(matrix), len(matrix[0])

    // Step 1: Build heights
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 1 {
                matrix[i][j] += matrix[i-1][j]
            }
        }
    }

    maxArea := 0

    // Step 2 & 3
    for i := 0; i < m; i++ {
        row := make([]int, n)
        copy(row, matrix[i])

        sort.Sort(sort.Reverse(sort.IntSlice(row)))

        for j := 0; j < n; j++ {
            area := row[j] * (j + 1)
            if area > maxArea {
                maxArea = area
            }
        }
    }

    return maxArea
}