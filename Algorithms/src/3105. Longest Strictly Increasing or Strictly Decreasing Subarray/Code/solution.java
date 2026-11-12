class Solution {
    public int longestMonotonicSubarray(int[] nums) {
        int n = nums.length;
        int maxLen = 1, inc = 1, dec = 1;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[i - 1]) {
                inc++;
                dec = 1;
            } else if (nums[i] < nums[i - 1]) {
                dec++;
                inc = 1;
            } else {
                inc = 1;
                dec = 1;
            }
            maxLen = Math.max(maxLen, Math.max(inc, dec));
        }
        return maxLen;
    }
}