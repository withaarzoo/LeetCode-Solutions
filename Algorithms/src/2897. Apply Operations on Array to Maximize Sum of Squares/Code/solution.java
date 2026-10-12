class Solution {
    public int maxSum(List<Integer> nums, int k) {
        final int MOD = 1000000007;
        int n = nums.size();
        int[] cnt = new int[30];
        for (int x : nums) {
            for (int b = 0; b < 30; b++) {
                if ((x & (1 << b)) != 0) cnt[b]++;
            }
        }
        long[] a = new long[n];
        for (int b = 29; b >= 0; b--) {
            int c = cnt[b];
            for (int i = 0; i < c; i++) {
                a[i] |= (1L << b);
            }
        }
        long ans = 0;
        for (int i = 0; i < k; i++) {
            ans = (ans + a[i] * a[i]) % MOD;
        }
        return (int) ans;
    }
}