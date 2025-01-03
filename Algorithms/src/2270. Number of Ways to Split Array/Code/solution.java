class Solution {
    public int waysToSplitArray(int[] nums) {
        long totalSum = 0;
        for (int num : nums) {
            totalSum += num;
        }

        long prefixSum = 0;
        int count = 0;

        for (int i = 0; i < nums.length - 1; i++) {
            prefixSum += nums[i];
            long rightSum = totalSum - prefixSum;
            if (prefixSum >= rightSum) {
                count++;
            }
        }

        return count;
    }
}
