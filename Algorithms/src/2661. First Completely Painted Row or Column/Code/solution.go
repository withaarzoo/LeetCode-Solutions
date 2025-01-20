func firstCompleteIndex(arr []int, mat [][]int) int {
    m, n := len(mat), len(mat[0])
    position := make(map[int][2]int)
    rowCount := make([]int, m)
    colCount := make([]int, n)

    // Map matrix values to their positions
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            position[mat[i][j]] = [2]int{i, j}
        }
    }

    // Iterate through the array and simulate painting
    for i, val := range arr {
        pos := position[val]
        row, col := pos[0], pos[1]
        rowCount[row]++
        colCount[col]++

        if rowCount[row] == n || colCount[col] == m {
            return i
        }
    }
    return -1
}
