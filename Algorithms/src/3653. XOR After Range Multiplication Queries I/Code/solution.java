class Solution {
    public int xorAfterQueries(int[] nums, int[][] queries) {
        long MOD = 1000000007L;

        // Process each query
        for (int[] q : queries) {
            int l = q[0];
            int r = q[1];
            int k = q[2];
            int v = q[3];

            // Visit indices: l, l+k, l+2k, ... <= r
            for (int i = l; i <= r; i += k) {
                nums[i] = (int) ((1L * nums[i] * v) % MOD);
            }
        }

        // Compute XOR of all final values
        int ans = 0;
        for (int num : nums) {
            ans ^= num;
        }

        return ans;
    }
}