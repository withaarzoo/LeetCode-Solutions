class Solution {
    public int numberOfStableArrays(int zero, int one, int limit) {

        int MOD = 1_000_000_007;

        long[][][] dp = new long[zero + 1][one + 1][2];

        for (int i = 1; i <= Math.min(zero, limit); i++)
            dp[i][0][0] = 1;

        for (int j = 1; j <= Math.min(one, limit); j++)
            dp[0][j][1] = 1;

        for (int i = 1; i <= zero; i++) {
            for (int j = 1; j <= one; j++) {

                long over0 = (i - limit - 1 >= 0) ? dp[i - limit - 1][j][1] : 0;
                long over1 = (j - limit - 1 >= 0) ? dp[i][j - limit - 1][0] : 0;

                dp[i][j][0] = (dp[i - 1][j][0] + dp[i - 1][j][1] - over0 + MOD) % MOD;

                dp[i][j][1] = (dp[i][j - 1][0] + dp[i][j - 1][1] - over1 + MOD) % MOD;
            }
        }

        return (int) ((dp[zero][one][0] + dp[zero][one][1]) % MOD);
    }
}