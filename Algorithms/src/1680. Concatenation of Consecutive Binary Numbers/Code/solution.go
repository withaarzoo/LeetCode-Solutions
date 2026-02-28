func concatenatedBinary(n int) int {
    const MOD = 1000000007
    var ans int64 = 0
    bitLength := 0
    
    for i := 1; i <= n; i++ {
        
        // If i is power of 2
        if (i & (i - 1)) == 0 {
            bitLength++
        }
        
        // Shift and add
        ans = ((ans << bitLength) % MOD + int64(i)) % MOD
    }
    
    return int(ans)
}