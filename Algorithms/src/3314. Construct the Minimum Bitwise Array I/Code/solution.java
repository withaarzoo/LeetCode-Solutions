class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int[] ans = new int[nums.size()];

        for (int i = 0; i < nums.size(); i++) {
            int p = nums.get(i);
            int found = -1;

            for (int x = 0; x <= p; x++) {
                if ((x | (x + 1)) == p) {
                    found = x;
                    break;
                }
            }

            ans[i] = found;
        }

        return ans;
    }
}
