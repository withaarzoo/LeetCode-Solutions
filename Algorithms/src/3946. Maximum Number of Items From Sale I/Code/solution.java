class Solution {
    public int maximumSaleItems(int[][] items, int budget) {
        int[] mult = new int[1501];
        for (int[] it : items) {
            int fac = it[0];
            for (int f = 1; f <= 1500; ++f) {
                if (fac % f == 0) ++mult[f];
            }
        }
        int[] dp = new int[budget + 1];
        for (int[] it : items) {
            int f = it[0], p = it[1];
            int[] ndp = dp.clone();
            for (int b = budget; b >= p; --b) {
                ndp[b] = Math.max(ndp[b], dp[b - p] + mult[f]);
            }
            for (int b = p; b <= budget; ++b) {
                ndp[b] = Math.max(ndp[b], ndp[b - p] + 1);
            }
            dp = ndp;
        }
        int ans = 0;
        for (int v : dp) ans = Math.max(ans, v);
        return ans;
    }
}