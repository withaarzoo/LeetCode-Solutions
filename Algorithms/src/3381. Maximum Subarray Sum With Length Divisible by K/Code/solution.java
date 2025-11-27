class Solution {
    public long maxSubarraySum(int[] nums, int k) {
        int n = nums.length;

        long INF = (long) 4e18; // large number
        long[] minPref = new long[k];
        for (int i = 0; i < k; i++) {
            minPref[i] = INF;
        }

        long prefix = 0;
        long ans = -INF;

        // prefix index 0 has sum = 0 and remainder 0
        minPref[0] = 0;

        for (int i = 0; i < n; i++) {
            prefix += nums[i];
            int rem = (i + 1) % k; // prefix index is i+1

            if (minPref[rem] != INF) {
                ans = Math.max(ans, prefix - minPref[rem]);
            }

            if (prefix < minPref[rem]) {
                minPref[rem] = prefix;
            }
        }

        return ans;
    }
}
