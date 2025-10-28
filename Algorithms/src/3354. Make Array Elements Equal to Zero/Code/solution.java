class Solution {
    // helper: simulate starting at start with dir (-1 left, +1 right)
    private boolean simulate(int[] nums, int start, int dir) {
        int n = nums.length;
        int[] a = nums.clone(); // copy to mutate
        int curr = start;
        while (curr >= 0 && curr < n) {
            if (a[curr] == 0) {
                curr += dir; // move same direction
            } else {
                a[curr]--; // decrement
                dir = -dir; // reverse direction
                curr += dir; // step in new direction
            }
        }
        // check all zero
        for (int v : a)
            if (v != 0)
                return false;
        return true;
    }

    public int countValidSelections(int[] nums) {
        int n = nums.length;
        int ans = 0;
        for (int i = 0; i < n; ++i) {
            if (nums[i] != 0)
                continue;
            if (simulate(nums, i, -1))
                ans++; // left
            if (simulate(nums, i, +1))
                ans++; // right
        }
        return ans;
    }
}
