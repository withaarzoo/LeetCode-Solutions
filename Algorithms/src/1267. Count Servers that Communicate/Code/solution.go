func countServers(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])
    rowCount := make([]int, rows)
    colCount := make([]int, cols)
    
    // First pass: Count servers in each row and column
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 {
                rowCount[i]++
                colCount[j]++
            }
        }
    }
    
    // Second pass: Count communicable servers
    count := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == 1 && (rowCount[i] > 1 || colCount[j] > 1) {
                count++
            }
        }
    }
    return count
}
