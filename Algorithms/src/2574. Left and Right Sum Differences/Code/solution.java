class Solution {
    public int[] leftRightDifference(int[] nums) {

        int n = nums.length;

        // Store total array sum
        int rightSum = 0;
        for (int num : nums) {
            rightSum += num;
        }

        // Sum of elements on the left
        int leftSum = 0;

        // Result array
        int[] ans = new int[n];

        for (int i = 0; i < n; i++) {

            // Remove current element from right side sum
            rightSum -= nums[i];

            // Store absolute difference
            ans[i] = Math.abs(leftSum - rightSum);

            // Include current element in left side sum
            leftSum += nums[i];
        }

        return ans;
    }
}