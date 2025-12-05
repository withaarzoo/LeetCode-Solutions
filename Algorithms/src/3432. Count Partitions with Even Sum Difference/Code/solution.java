class Solution {
    public int countPartitions(int[] nums) {
        long total = 0;
        // Compute the total sum of the array
        for (int x : nums) {
            total += x;
        }

        // If total sum is odd, no valid partition
        if ((total & 1L) == 1L)
            return 0;

        // If total is even, every position between elements is a valid partition
        int n = nums.length;
        return n - 1;
    }
}
