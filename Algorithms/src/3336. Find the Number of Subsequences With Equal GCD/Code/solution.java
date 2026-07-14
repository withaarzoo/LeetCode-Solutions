class Solution {
    public int subsequencePairCount(int[] nums) {
        final int MOD = 1_000_000_007;
        final int MAX = 200;

        // Current DP table.
        int[][] dp = new int[MAX + 1][MAX + 1];
        dp[0][0] = 1;

        for (int x : nums) {
            // Next DP table after processing current number.
            int[][] next = new int[MAX + 1][MAX + 1];

            for (int g1 = 0; g1 <= MAX; g1++) {
                for (int g2 = 0; g2 <= MAX; g2++) {
                    if (dp[g1][g2] == 0) continue;

                    long ways = dp[g1][g2];

                    // Choice 1: Skip current number.
                    next[g1][g2] = (int)((next[g1][g2] + ways) % MOD);

                    // Choice 2: Put into first subsequence.
                    int ng1 = (g1 == 0) ? x : gcd(g1, x);
                    next[ng1][g2] = (int)((next[ng1][g2] + ways) % MOD);

                    // Choice 3: Put into second subsequence.
                    int ng2 = (g2 == 0) ? x : gcd(g2, x);
                    next[g1][ng2] = (int)((next[g1][ng2] + ways) % MOD);
                }
            }

            dp = next;
        }

        long ans = 0;

        // Count states where both GCDs are equal and non-zero.
        for (int g = 1; g <= MAX; g++) {
            ans = (ans + dp[g][g]) % MOD;
        }

        return (int)ans;
    }

    // Computes GCD using Euclid's algorithm.
    private int gcd(int a, int b) {
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }
}