func champagneTower(poured int, query_row int, query_glass int) float64 {
    
    // Create DP table
    dp := make([][]float64, 101)
    for i := range dp {
        dp[i] = make([]float64, 101)
    }
    
    dp[0][0] = float64(poured)
    
    // Simulate flow
    for r := 0; r <= query_row; r++ {
        for c := 0; c <= r; c++ {
            
            if dp[r][c] > 1.0 {
                overflow := (dp[r][c] - 1.0) / 2.0
                
                dp[r+1][c] += overflow
                dp[r+1][c+1] += overflow
                
                dp[r][c] = 1.0
            }
        }
    }
    
    if dp[query_row][query_glass] > 1.0 {
        return 1.0
    }
    
    return dp[query_row][query_glass]
}
