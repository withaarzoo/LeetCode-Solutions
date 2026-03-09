func numberOfStableArrays(zero int, one int, limit int) int {

    const MOD int = 1e9 + 7
    n := zero + one

    fact := make([]int, n+1)
    invFact := make([]int, n+1)

    fact[0] = 1

    for i := 1; i <= n; i++ {
        fact[i] = fact[i-1] * i % MOD
    }

    modPow := func(a, b int) int {
        res := 1
        base := a

        for b > 0 {
            if b&1 == 1 {
                res = res * base % MOD
            }
            base = base * base % MOD
            b >>= 1
        }
        return res
    }

    invFact[n] = modPow(fact[n], MOD-2)

    for i := n - 1; i >= 0; i-- {
        invFact[i] = invFact[i+1] * (i + 1) % MOD
    }

    C := func(n, k int) int {
        if k < 0 || k > n {
            return 0
        }
        return fact[n] * invFact[k] % MOD * invFact[n-k] % MOD
    }

    F := func(N, K int) int {

        if K <= 0 || K > N {
            return 0
        }

        ans := 0

        maxJ := (N - K) / limit

        for j := 0; j <= maxJ; j++ {

            term := C(K, j) * C(N-j*limit-1, K-1) % MOD

            if j%2 == 1 {
                ans = (ans - term + MOD) % MOD
            } else {
                ans = (ans + term) % MOD
            }
        }

        return ans
    }

    maxK := zero
    if one+1 < maxK {
        maxK = one + 1
    }

    oneWays := make([]int, maxK+3)

    for k := 1; k <= maxK+1; k++ {
        oneWays[k] = F(one, k)
    }

    ans := 0

    for k := 1; k <= maxK; k++ {

        z := F(zero, k)

        o := (oneWays[k-1] + 2*oneWays[k] + oneWays[k+1]) % MOD

        ans = (ans + z*o%MOD) % MOD
    }

    return ans
}