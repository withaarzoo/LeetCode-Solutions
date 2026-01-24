class Solution {
    public int minPairSum(int[] nums) {
        // Step 1: Sort the array
        Arrays.sort(nums);

        int left = 0;
        int right = nums.length - 1;
        int maxPairSum = 0;

        // Step 2: Pair smallest with largest
        while (left < right) {
            int pairSum = nums[left] + nums[right];
            maxPairSum = Math.max(maxPairSum, pairSum);
            left++;
            right--;
        }

        return maxPairSum;
    }
}
