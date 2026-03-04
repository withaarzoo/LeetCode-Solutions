func numSpecial(mat [][]int) int {
    m := len(mat)
    n := len(mat[0])

    rowCount := make([]int, m)
    colCount := make([]int, n)

    // Count number of 1s in each row and column
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                rowCount[i]++
                colCount[j]++
            }
        }
    }

    special := 0

    // Check for special positions
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 && rowCount[i] == 1 && colCount[j] == 1 {
                special++
            }
        }
    }

    return special
}