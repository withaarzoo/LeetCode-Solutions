class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int[] ans = new int[nums.size()];

        for (int i = 0; i < nums.size(); i++) {
            int p = nums.get(i);

            int removable = ((p + 1) & ~p) >> 1;

            if (removable == 0) {
                ans[i] = -1;
            } else {
                ans[i] = p ^ removable;
            }
        }

        return ans;
    }
}
