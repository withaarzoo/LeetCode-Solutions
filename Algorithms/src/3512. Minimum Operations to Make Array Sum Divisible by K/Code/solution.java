class Solution {
    public int minOperations(int[] nums, int k) {
        long longSum = 0;

        // Calculate total sum of the array
        for (int x : nums) {
            longSum += x;
        }

        // Minimum operations is the remainder of the sum when divided by k
        int remainder = (int) (longSum % k);

        return remainder;
    }
}
