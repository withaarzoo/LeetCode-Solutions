class Solution {
    public int[] maxValue(int[] nums) {
        int n = nums.length;

        // suffixMin[i] = smallest value in nums[i...n-1]
        // I keep one extra slot so the boundary case is easy to handle.
        int[] suffixMin = new int[n + 1];
        suffixMin[n] = Integer.MAX_VALUE;
        for (int i = n - 1; i >= 0; i--) {
            suffixMin[i] = Math.min(nums[i], suffixMin[i + 1]);
        }

        int[] ans = new int[n];
        int l = 0;

        // I process one connected component at a time.
        while (l < n) {
            int r = l;
            int componentMax = nums[l];

            // I extend the current segment while an inversion crosses the next cut.
            while (r + 1 < n && componentMax > suffixMin[r + 1]) {
                r++;
                componentMax = Math.max(componentMax, nums[r]);
            }

            // All positions in this component share the same maximum reachable value.
            for (int i = l; i <= r; i++) {
                ans[i] = componentMax;
            }

            // Jump to the next untouched segment.
            l = r + 1;
        }

        return ans;
    }
}