class Solution {
    public int uniqueXorTriplets(int[] nums) {
        // Get the size of the permutation
        int n = nums.length;

        // Small cases
        if (n <= 2)
            return n;

        // Count the number of bits in n
        int bits = 0;
        int x = n;
        while (x > 0) {
            bits++;
            x >>= 1;
        }

        // Every value in [0, 2^bits - 1] is possible
        return 1 << bits;
    }
}