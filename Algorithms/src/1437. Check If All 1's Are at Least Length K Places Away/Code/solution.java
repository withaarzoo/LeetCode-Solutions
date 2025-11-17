class Solution {
    public boolean kLengthApart(int[] nums, int k) {
        int prev = -1; // index of last seen 1; -1 means none
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                if (prev != -1) {
                    // zeros between two 1s = i - prev - 1
                    if (i - prev - 1 < k)
                        return false;
                }
                prev = i;
            }
        }
        return true;
    }
}
