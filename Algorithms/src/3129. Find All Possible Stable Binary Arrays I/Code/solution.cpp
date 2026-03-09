class Solution
{
public:
    const long long MOD = 1e9 + 7;

    long long modPow(long long a, long long b)
    {
        long long res = 1;
        while (b)
        {
            if (b & 1)
                res = res * a % MOD;
            a = a * a % MOD;
            b >>= 1;
        }
        return res;
    }

    int numberOfStableArrays(int zero, int one, int limit)
    {

        int n = zero + one;

        vector<long long> fact(n + 1), invFact(n + 1);

        fact[0] = 1;
        for (int i = 1; i <= n; i++)
            fact[i] = fact[i - 1] * i % MOD;

        invFact[n] = modPow(fact[n], MOD - 2);

        for (int i = n - 1; i >= 0; i--)
            invFact[i] = invFact[i + 1] * (i + 1) % MOD;

        auto C = [&](int n, int k) -> long long
        {
            if (k < 0 || k > n)
                return 0;
            return fact[n] * invFact[k] % MOD * invFact[n - k] % MOD;
        };

        auto F = [&](int N, int K) -> long long
        {
            if (K <= 0 || K > N)
                return 0;

            long long ans = 0;

            int maxJ = (N - K) / limit;

            for (int j = 0; j <= maxJ; j++)
            {

                long long ways = C(K, j) * C(N - j * limit - 1, K - 1) % MOD;

                if (j % 2)
                    ans = (ans - ways + MOD) % MOD;
                else
                    ans = (ans + ways) % MOD;
            }

            return ans;
        };

        int maxK = min(zero, one + 1);

        vector<long long> oneWays(maxK + 3);

        for (int k = 1; k <= maxK + 1; k++)
            oneWays[k] = F(one, k);

        long long ans = 0;

        for (int k = 1; k <= maxK; k++)
        {

            long long zWays = F(zero, k);

            long long oWays = (oneWays[k - 1] + 2 * oneWays[k] + oneWays[k + 1]) % MOD;

            ans = (ans + zWays * oWays) % MOD;
        }

        return ans;
    }
};