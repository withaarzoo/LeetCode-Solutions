class Solution {
    public int minimumCost(int[] nums) {
        int first = nums[0];

        // Sort from index 1 onward
        Arrays.sort(nums, 1, nums.length);

        return first + nums[1] + nums[2];
    }
}
