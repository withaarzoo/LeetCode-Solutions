class Solution {
    public int minOperations(int[] nums, int x) {
        long total = 0;
        for (int num : nums) total += num;
        if (total < x) return -1;
        long target = total - x;
        if (target == 0) return nums.length;
        int n = nums.length;
        int left = 0;
        long sum = 0;
        int maxLen = -1;
        for (int right = 0; right < n; ++right) {
            sum += nums[right];
            while (sum > target && left <= right) {
                sum -= nums[left];
                ++left;
            }
            if (sum == target) {
                maxLen = Math.max(maxLen, right - left + 1);
            }
        }
        return maxLen == -1 ? -1 : n - maxLen;
    }
}