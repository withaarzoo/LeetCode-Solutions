class Solution {
public:
    int numOfWays(vector<int>& nums) {
        const int MOD = 1e9 + 7;
        int n = nums.size();
        vector<long long> fact(n + 1, 1), inv(n + 1, 1);
        for (int i = 1; i <= n; i++) fact[i] = fact[i - 1] * i % MOD;
        inv[n] = modPow(fact[n], MOD - 2, MOD);
        for (int i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * (i + 1) % MOD;
        auto comb = [&](int a, int b) -> long long {
            if (b < 0 || b > a) return 0;
            return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
        };
        function<long long(vector<int>&)> ways = [&](vector<int>& arr) -> long long {
            int m = arr.size();
            if (m <= 2) return 1;
            vector<int> left, right;
            for (int x : arr) {
                if (x < arr[0]) left.push_back(x);
                else if (x > arr[0]) right.push_back(x);
            }
            return comb(m - 1, left.size()) * ways(left) % MOD * ways(right) % MOD;
        };
        return (ways(nums) - 1 + MOD) % MOD;
    }
private:
    long long modPow(long long base, long long exp, long long mod) {
        long long res = 1;
        while (exp) {
            if (exp & 1) res = res * base % mod;
            base = base * base % mod;
            exp >>= 1;
        }
        return res;
    }
};