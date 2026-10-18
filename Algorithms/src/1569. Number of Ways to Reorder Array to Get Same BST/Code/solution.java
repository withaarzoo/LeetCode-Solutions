class Solution {
    public int numOfWays(int[] nums) {
        final int MOD = 1000000007;
        int n = nums.length;
        long[] fact = new long[n + 1];
        long[] inv = new long[n + 1];
        fact[0] = 1;
        for (int i = 1; i <= n; i++) fact[i] = fact[i - 1] * i % MOD;
        inv[n] = modPow(fact[n], MOD - 2, MOD);
        for (int i = n - 1; i >= 0; i--) inv[i] = inv[i + 1] * (i + 1) % MOD;
        return (int)((ways(nums, fact, inv, MOD) - 1 + MOD) % MOD);
    }
    private long ways(int[] arr, long[] fact, long[] inv, int MOD) {
        int m = arr.length;
        if (m <= 2) return 1;
        List<Integer> left = new ArrayList<>();
        List<Integer> right = new ArrayList<>();
        for (int x : arr) {
            if (x < arr[0]) left.add(x);
            else if (x > arr[0]) right.add(x);
        }
        int[] l = left.stream().mapToInt(i -> i).toArray();
        int[] r = right.stream().mapToInt(i -> i).toArray();
        return comb(m - 1, l.length, fact, inv, MOD) * ways(l, fact, inv, MOD) % MOD * ways(r, fact, inv, MOD) % MOD;
    }
    private long comb(int a, int b, long[] fact, long[] inv, int MOD) {
        if (b < 0 || b > a) return 0;
        return fact[a] * inv[b] % MOD * inv[a - b] % MOD;
    }
    private long modPow(long base, long exp, long mod) {
        long res = 1;
        while (exp > 0) {
            if ((exp & 1) == 1) res = res * base % mod;
            base = base * base % mod;
            exp >>= 1;
        }
        return res;
    }
}