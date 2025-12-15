class Solution {
    public long getDescentPeriods(int[] prices) {
        long ans = 1; // first day
        long len = 1; // current descent length

        for (int i = 1; i < prices.length; i++) {
            if (prices[i] == prices[i - 1] - 1) {
                len++;
            } else {
                len = 1;
            }
            ans += len;
        }
        return ans;
    }
}
