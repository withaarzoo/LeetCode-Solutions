func maxDotProduct(nums1 []int, nums2 []int) int {
    n := len(nums1)
    m := len(nums2)
    
    // Initialize DP with very small values
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j <= m; j++ {
            dp[i][j] = -1 << 60
        }
    }
    
    for i := n - 1; i >= 0; i-- {
        for j := m - 1; j >= 0; j-- {
            product := nums1[i] * nums2[j]
            
            takeBoth := product
            if dp[i+1][j+1] != -1<<60 {
                if product+dp[i+1][j+1] > takeBoth {
                    takeBoth = product + dp[i+1][j+1]
                }
            }
            
            dp[i][j] = max(
                takeBoth,
                max(dp[i+1][j], dp[i][j+1]),
            )
        }
    }
    
    return dp[0][0]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
