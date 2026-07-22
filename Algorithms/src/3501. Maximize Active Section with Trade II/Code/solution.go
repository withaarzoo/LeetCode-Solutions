func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
    n := len(s) // Original string length
    totalOnes := 0 // Counter for default number of '1's in the string
    
    // Establishing the baseline score
    for i := 0; i < n; i++ {
        if s[i] == '1' {
            totalOnes++
        }
    }
    
    // Arrays tracking block contents
    typeArr := []int{}
    start := []int{}
    endIdx := []int{}
    
    // Condense repeating string characters into metadata blocks
    for i := 0; i < n; {
        j := i
        // Keep marching j while characters are uniform
        for j < n && s[j] == s[i] {
            j++
        }
        val := 0
        if s[i] == '1' {
            val = 1
        }
        typeArr = append(typeArr, val)
        start = append(start, i)
        endIdx = append(endIdx, j-1)
        i = j // Snap i to the end
    }
    
    N := len(typeArr) // Length of compressed structure
    
    // Creates a lookup table translating raw character index into block ID
    posToSeg := make([]int, n)
    for i := 0; i < N; i++ {
        for j := start[i]; j <= endIdx[i]; j++ {
            posToSeg[j] = i
        }
    }
    
    // Precompute full theoretical yield for trading any given 1-segment
    ans := make([]int, N)
    for i := 1; i < N-1; i++ {
        if typeArr[i] == 1 {
            // Net addition is the sum of enclosing 0-segments
            ans[i] = (endIdx[i-1] - start[i-1] + 1) + (endIdx[i+1] - start[i+1] + 1)
        }
    }
    
    // Populate logarithmic exponents for RMQ table sizing
    logTable := make([]int, N+1)
    for i := 2; i <= N; i++ {
        logTable[i] = logTable[i/2] + 1
    }
    
    K := logTable[N] + 1
    st := make([][]int, K)
    for i := 0; i < K; i++ {
        st[i] = make([]int, N)
    }
    
    // Fill the 0th layer of the Sparse Table
    for i := 0; i < N; i++ {
        st[0][i] = ans[i]
    }
    
    // Fill remaining layers by combining overlapping powers of 2
    for j := 1; j < K; j++ {
        for i := 0; i+(1<<j) <= N; i++ {
            a := st[j-1][i]
            b := st[j-1][i+(1<<(j-1))]
            if b > a {
                st[j][i] = b
            } else {
                st[j][i] = a
            }
        }
    }
    
    // Function returning range maximum in O(1)
    queryRMQ := func(L, R int) int {
        if L > R {
            return 0
        }
        j := logTable[R-L+1]
        a := st[j][L]
        b := st[j][R-(1<<j)+1]
        if b > a {
            return b
        }
        return a
    }
    
    maxFn := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    
    // Custom evaluation logic for blocks that are cut off by L or R edges
    eval := func(i, L, R, segL, segR int) int {
        // Drop invalid or out-of-bounds queries
        if i <= segL || i >= segR {
            return 0
        }
        if typeArr[i] == 0 {
            return 0
        }
        
        leftLen := 0
        // Snip the left 0-segment exactly at L if it crosses over
        if i-1 == segL {
            leftLen = maxFn(0, endIdx[i-1]-L+1)
        } else {
            leftLen = endIdx[i-1] - start[i-1] + 1
        }
        
        rightLen := 0
        // Snip the right 0-segment exactly at R if it crosses over
        if i+1 == segR {
            rightLen = maxFn(0, R-start[i+1]+1)
        } else {
            rightLen = endIdx[i+1] - start[i+1] + 1
        }
        
        return leftLen + rightLen
    }
    
    res := make([]int, len(queries))
    // Loop over queries
    for idx, q := range queries {
        L := q[0]
        R := q[1]
        
        segL := posToSeg[L]
        segR := posToSeg[R]
        
        // If there isn't enough room to pull off a trade, record baseline score
        if segR-segL < 2 {
            res[idx] = totalOnes
            continue
        }
        
        maxGain := 0
        // Only evaluate blocks immediately touching bounds manually
        maxGain = maxFn(maxGain, eval(segL+1, L, R, segL, segR))
        maxGain = maxFn(maxGain, eval(segR-1, L, R, segL, segR))
        
        // Let the Sparse Table figure out the highest gain among fully encased segments
        if segL+2 <= segR-2 {
            maxGain = maxFn(maxGain, queryRMQ(segL+2, segR-2))
        }
        
        // Store total calculated yield
        res[idx] = totalOnes + maxGain
    }
    
    return res
}