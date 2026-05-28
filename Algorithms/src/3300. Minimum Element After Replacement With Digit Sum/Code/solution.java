class Solution {

    // Helper function to calculate digit sum
    private int digitSum(int num) {
        int sum = 0;

        // Process every digit
        while (num > 0) {
            sum += num % 10; // Add last digit
            num /= 10; // Remove last digit
        }

        return sum;
    }

    public int minElement(int[] nums) {
        int ans = Integer.MAX_VALUE;

        // Check digit sum of every number
        for (int num : nums) {
            ans = Math.min(ans, digitSum(num));
        }

        return ans;
    }
}