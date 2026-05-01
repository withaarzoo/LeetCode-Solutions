class Solution {
    public int maxRotateFunction(int[] nums) {
        int n = nums.length;

        long sum = 0; // total sum
        long F = 0; // F(0)

        // Step 1: compute sum and F(0)
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            F += (long) i * nums[i];
        }

        long result = F;

        // Step 2: compute next rotations
        for (int k = 1; k < n; k++) {
            F = F + sum - (long) n * nums[n - k];
            result = Math.max(result, F);
        }

        return (int) result;
    }
}