func numberOfStableArrays(zero int, one int, limit int) int {

    const MOD int = 1e9 + 7

    dp := make([][][]int, zero+1)

    for i := range dp {
        dp[i] = make([][]int, one+1)
        for j := range dp[i] {
            dp[i][j] = make([]int,2)
        }
    }

    for i:=1;i<=min(zero,limit);i++{
        dp[i][0][0] = 1
    }

    for j:=1;j<=min(one,limit);j++{
        dp[0][j][1] = 1
    }

    for i:=1;i<=zero;i++{
        for j:=1;j<=one;j++{

            over0 := 0
            if i-limit-1 >= 0 {
                over0 = dp[i-limit-1][j][1]
            }

            over1 := 0
            if j-limit-1 >= 0 {
                over1 = dp[i][j-limit-1][0]
            }

            dp[i][j][0] =
                (dp[i-1][j][0] + dp[i-1][j][1] - over0 + MOD) % MOD

            dp[i][j][1] =
                (dp[i][j-1][0] + dp[i][j-1][1] - over1 + MOD) % MOD
        }
    }

    return (dp[zero][one][0] + dp[zero][one][1]) % MOD
}

func min(a,b int) int{
    if a<b { return a }
    return b
}