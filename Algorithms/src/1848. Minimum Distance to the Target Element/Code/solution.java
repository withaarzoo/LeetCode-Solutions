class Solution {
    public int getMinDistance(int[] nums, int target, int start) {
        // Store the minimum distance found so far
        int answer = Integer.MAX_VALUE;

        // Traverse through the array
        for (int i = 0; i < nums.length; i++) {
            // Check if current element is the target
            if (nums[i] == target) {
                // Update the minimum distance
                answer = Math.min(answer, Math.abs(i - start));
            }
        }

        return answer;
    }
}