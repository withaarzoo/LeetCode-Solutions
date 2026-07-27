class Solution {
    public int maxProduct(int[] nums) {

        // Store the largest value found so far
        int first = 0;

        // Store the second largest value found so far
        int second = 0;

        // Traverse every element once
        for (int num : nums) {

            // If current number becomes the largest
            if (num >= first) {
                // Old largest becomes second largest
                second = first;

                // Update largest
                first = num;
            }
            // Otherwise update second largest if needed
            else if (num > second) {
                second = num;
            }
        }

        // Return the required product
        return (first - 1) * (second - 1);
    }
}