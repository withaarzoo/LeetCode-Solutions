class Solution {

    // Function to find GCD using the Euclidean Algorithm
    private int gcd(int a, int b) {
        // Continue until b becomes 0
        while (b != 0) {
            int temp = b; // Store current b
            b = a % b; // Update b with remainder
            a = temp; // Move previous b into a
        }

        // a now contains the GCD
        return a;
    }

    public int findGCD(int[] nums) {

        // Initialize minimum and maximum with the first element
        int min = nums[0];
        int max = nums[0];

        // Find the minimum and maximum values
        for (int num : nums) {
            min = Math.min(min, num);
            max = Math.max(max, num);
        }

        // Return their GCD
        return gcd(min, max);
    }
}