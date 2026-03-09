class Solution {

    static final long MOD = 1000000007;

    long modPow(long a, long b) {
        long res = 1;
        while (b > 0) {
            if ((b & 1) == 1)
                res = res * a % MOD;
            a = a * a % MOD;
            b >>= 1;
        }
        return res;
    }

    public int numberOfStableArrays(int zero, int one, int limit) {

        int n = zero + one;

        long[] fact = new long[n + 1];
        long[] invFact = new long[n + 1];

        fact[0] = 1;

        for (int i = 1; i <= n; i++)
            fact[i] = fact[i - 1] * i % MOD;

        invFact[n] = modPow(fact[n], MOD - 2);

        for (int i = n - 1; i >= 0; i--)
            invFact[i] = invFact[i + 1] * (i + 1) % MOD;

        java.util.function.BiFunction<Integer, Integer, Long> C = (nn, kk) -> {
            if (kk < 0 || kk > nn)
                return 0L;
            return fact[nn] * invFact[kk] % MOD * invFact[nn - kk] % MOD;
        };

        java.util.function.BiFunction<Integer, Integer, Long> F = (N, K) -> {

            if (K <= 0 || K > N)
                return 0L;

            long ans = 0;

            int maxJ = (N - K) / limit;

            for (int j = 0; j <= maxJ; j++) {

                long term = C.apply(K, j) * C.apply(N - j * limit - 1, K - 1) % MOD;

                if (j % 2 == 1)
                    ans = (ans - term + MOD) % MOD;
                else
                    ans = (ans + term) % MOD;
            }

            return ans;
        };

        int maxK = Math.min(zero, one + 1);

        long[] oneWays = new long[maxK + 3];

        for (int k = 1; k <= maxK + 1; k++)
            oneWays[k] = F.apply(one, k);

        long ans = 0;

        for (int k = 1; k <= maxK; k++) {

            long z = F.apply(zero, k);

            long o = (oneWays[k - 1] + 2 * oneWays[k] + oneWays[k + 1]) % MOD;

            ans = (ans + z * o) % MOD;
        }

        return (int) ans;
    }
}