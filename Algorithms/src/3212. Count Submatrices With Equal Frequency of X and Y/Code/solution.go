func numberOfSubmatrices(grid [][]byte) int {
    n, m := len(grid), len(grid[0])
    
    sum := make([][]int, 2)
    countX := make([][]int, 2)
    
    for i := 0; i < 2; i++ {
        sum[i] = make([]int, m+1)
        countX[i] = make([]int, m+1)
    }
    
    ans := 0
    
    for i := 0; i < n; i++ {
        cur := i % 2
        prev := 1 - cur
        
        for j := 0; j < m; j++ {
            val := 0
            if grid[i][j] == 'X' {
                val = 1
            } else if grid[i][j] == 'Y' {
                val = -1
            }
            
            isX := 0
            if grid[i][j] == 'X' {
                isX = 1
            }
            
            sum[cur][j+1] = val +
                sum[cur][j] +
                sum[prev][j+1] -
                sum[prev][j]
            
            countX[cur][j+1] = isX +
                countX[cur][j] +
                countX[prev][j+1] -
                countX[prev][j]
            
            if sum[cur][j+1] == 0 && countX[cur][j+1] > 0 {
                ans++
            }
        }
    }
    
    return ans
}