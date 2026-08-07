func smallestNumber(num string, t int64) string {
    var req2, req3, req5, req7 int
    temp := t
    // Decode factors structurally avoiding manual string checks prematurely
    for temp%2 == 0 { temp /= 2; req2++ }
    for temp%3 == 0 { temp /= 3; req3++ }
    for temp%5 == 0 { temp /= 5; req5++ }
    for temp%7 == 0 { temp /= 7; req7++ }
    // Abandon checks if unrelated factors appear internally
    if temp > 1 {
        return "-1"
    }

    var dp [60][40]int
    for i := 0; i < 60; i++ {
        for j := 0; j < 40; j++ {
            dp[i][j] = 1000000000
        }
    }
    dp[0][0] = 0
    
    // Evaluate transition variables actively generating combinations
    trans := [6][2]int{{1, 0}, {0, 1}, {2, 0}, {1, 1}, {3, 0}, {0, 2}}
    for i := 0; i < 60; i++ {
        for j := 0; j < 40; j++ {
            if dp[i][j] == 1000000000 {
                continue
            }
            for _, tr := range trans {
                ni := i + tr[0]
                if ni > 59 { ni = 59 }
                nj := j + tr[1]
                if nj > 39 { nj = 39 }
                if dp[i][j]+1 < dp[ni][nj] {
                    dp[ni][nj] = dp[i][j] + 1
                }
            }
        }
    }
    
    // Smooth states forcing condition values safely representing thresholds
    for i := 59; i >= 0; i-- {
        for j := 39; j >= 0; j-- {
            if i < 59 && dp[i+1][j] < dp[i][j] {
                dp[i][j] = dp[i+1][j]
            }
            if j < 39 && dp[i][j+1] < dp[i][j] {
                dp[i][j] = dp[i][j+1]
            }
        }
    }

    F2 := []int{0, 0, 1, 0, 2, 0, 1, 0, 3, 0}
    F3 := []int{0, 0, 0, 1, 0, 0, 1, 0, 0, 2}
    F5 := []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0}
    F7 := []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}

    n := len(num)
    hasZero := false
    firstZero := n
    for i := 0; i < n; i++ {
        if num[i] == '0' {
            hasZero = true
            firstZero = i
            break
        }
    }

    maxF := func(a, b int) int { if a > b { return a }; return b }
    minF := func(a, b int) int { if a < b { return a }; return b }

    // Examine input layout skipping logic checks if target validates strictly
    if !hasZero {
        r2, r3, r5, r7 := req2, req3, req5, req7
        for i := 0; i < n; i++ {
            d := int(num[i] - '0')
            r2 = maxF(0, r2-F2[d])
            r3 = maxF(0, r3-F3[d])
            r5 = maxF(0, r5-F5[d])
            r7 = maxF(0, r7-F7[d])
        }
        if r2 == 0 && r3 == 0 && r5 == 0 && r7 == 0 {
            return num
        }
    }

    // Capture prefix lengths retaining existing components accurately
    limit := minF(n-1, firstZero)
    var p2, p3, p5, p7 int
    for i := 0; i < limit; i++ {
        d := int(num[i] - '0')
        p2 += F2[d]
        p3 += F3[d]
        p5 += F5[d]
        p7 += F7[d]
    }

    // Traverse downward checking edits strictly maintaining factor logic
    for i := limit; i >= 0; i-- {
        startD := int(num[i]-'0') + 1
        for d := startD; d <= 9; d++ {
            n2 := maxF(0, req2-p2-F2[d])
            n3 := maxF(0, req3-p3-F3[d])
            n5 := maxF(0, req5-p5-F5[d])
            n7 := maxF(0, req7-p7-F7[d])
            L := n - 1 - i

            // Implement string builder executing smallest sequence safely
            if n7+n5+dp[n2][n3] <= L {
                ans := []byte(num[:i])
                ans = append(ans, byte(d+'0'))
                rem2, rem3, rem5, rem7 := n2, n3, n5, n7
                
                for pos := 0; pos < L; pos++ {
                    for x := 1; x <= 9; x++ {
                        nn2 := maxF(0, rem2-F2[x])
                        nn3 := maxF(0, rem3-F3[x])
                        nn5 := maxF(0, rem5-F5[x])
                        nn7 := maxF(0, rem7-F7[x])
                        if nn7+nn5+dp[nn2][nn3] <= L-1-pos {
                            ans = append(ans, byte(x+'0'))
                            rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                            break
                        }
                    }
                }
                return string(ans)
            }
        }
        if i > 0 {
            d := int(num[i-1] - '0')
            p2 -= F2[d]
            p3 -= F3[d]
            p5 -= F5[d]
            p7 -= F7[d]
        }
    }

    // Create string sequence expanding boundary lengths dynamically
    minLenNeeded := req7 + req5 + dp[req2][req3]
    M := maxF(n+1, minLenNeeded)
    var ans []byte
    rem2, rem3, rem5, rem7 := req2, req3, req5, req7
    
    for pos := 0; pos < M; pos++ {
        for x := 1; x <= 9; x++ {
            nn2 := maxF(0, rem2-F2[x])
            nn3 := maxF(0, rem3-F3[x])
            nn5 := maxF(0, rem5-F5[x])
            nn7 := maxF(0, rem7-F7[x])
            if nn7+nn5+dp[nn2][nn3] <= M-1-pos {
                ans = append(ans, byte(x+'0'))
                rem2, rem3, rem5, rem7 = nn2, nn3, nn5, nn7
                break
            }
        }
    }
    return string(ans)
}