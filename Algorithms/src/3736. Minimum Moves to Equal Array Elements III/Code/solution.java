class Solution {
    public int minMoves(int[] nums) {
        int mx = nums[0];
        for (int x : nums)
            if (x > mx)
                mx = x;
        int ans = 0;
        for (int x : nums)
            ans += mx - x;
        return ans;
    }
}