class Solution {
    public long maxProfit(int[] prices, int[] strategy, int k) {
        int n = prices.length;

        long base = 0;
        for (int i = 0; i < n; i++) {
            base += (long) strategy[i] * prices[i];
        }

        long[] prefPrice = new long[n + 1];
        long[] prefProfit = new long[n + 1];

        for (int i = 0; i < n; i++) {
            prefPrice[i + 1] = prefPrice[i] + prices[i];
            prefProfit[i + 1] = prefProfit[i] + (long) strategy[i] * prices[i];
        }

        long bestDelta = 0;
        int half = k / 2;

        for (int l = 0; l + k <= n; l++) {
            int m = l + half;
            int r = l + k;

            long oldProfit = prefProfit[r] - prefProfit[l];
            long newProfit = prefPrice[r] - prefPrice[m];

            bestDelta = Math.max(bestDelta, newProfit - oldProfit);
        }

        return base + bestDelta;
    }
}
