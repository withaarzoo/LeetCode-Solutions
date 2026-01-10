func minimumDeleteSum(s1 string, s2 string) int {
    n, m := len(s1), len(s2)
    dp := make([]int, m+1)

    for j := m - 1; j >= 0; j-- {
        dp[j] = dp[j+1] + int(s2[j])
    }

    for i := n - 1; i >= 0; i-- {
        prev := dp[m]
        dp[m] += int(s1[i])

        for j := m - 1; j >= 0; j-- {
            temp := dp[j]
            if s1[i] == s2[j] {
                dp[j] = prev
            } else {
                a := int(s1[i]) + dp[j]
                b := int(s2[j]) + dp[j+1]
                if a < b {
                    dp[j] = a
                } else {
                    dp[j] = b
                }
            }
            prev = temp
        }
    }
    return dp[0]
}
