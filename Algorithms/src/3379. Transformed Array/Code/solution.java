class Solution {
    public int[] constructTransformedArray(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];

        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) {
                result[i] = nums[i];
            } else {
                int target = (i + nums[i]) % n;
                if (target < 0)
                    target += n;
                result[i] = nums[target];
            }
        }
        return result;
    }
}
