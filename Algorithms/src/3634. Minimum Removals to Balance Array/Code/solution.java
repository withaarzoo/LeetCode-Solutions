class Solution {
    public int minRemoval(int[] nums, int k) {
        Arrays.sort(nums);

        int n = nums.length;
        int left = 0;
        int maxKeep = 1;

        for (int right = 0; right < n; right++) {
            while ((long) nums[right] > (long) nums[left] * k) {
                left++;
            }
            maxKeep = Math.max(maxKeep, right - left + 1);
        }

        return n - maxKeep;
    }
}
