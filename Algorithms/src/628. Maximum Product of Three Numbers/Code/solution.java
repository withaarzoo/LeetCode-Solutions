class Solution {
    public int maximumProduct(int[] nums) {

        // Store the three largest numbers
        int max1 = Integer.MIN_VALUE;
        int max2 = Integer.MIN_VALUE;
        int max3 = Integer.MIN_VALUE;

        // Store the two smallest numbers
        int min1 = Integer.MAX_VALUE;
        int min2 = Integer.MAX_VALUE;

        // Traverse the array once
        for (int num : nums) {

            // Update the three largest numbers
            if (num >= max1) {
                max3 = max2;
                max2 = max1;
                max1 = num;
            } else if (num >= max2) {
                max3 = max2;
                max2 = num;
            } else if (num >= max3) {
                max3 = num;
            }

            // Update the two smallest numbers
            if (num <= min1) {
                min2 = min1;
                min1 = num;
            } else if (num <= min2) {
                min2 = num;
            }
        }

        // Product of the three largest numbers
        int product1 = max1 * max2 * max3;

        // Product of the two smallest numbers and the largest number
        int product2 = min1 * min2 * max1;

        // Return the maximum product
        return Math.max(product1, product2);
    }
}