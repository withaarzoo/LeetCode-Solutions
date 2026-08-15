class Solution {
    public int longestSubsequence(int[] nums) {
        int xorValue = 0; // Stores the XOR of all elements.
        boolean hasNonZero = false; // Tracks whether at least one element is non-zero.

        for (int x : nums) {
            xorValue ^= x; // Add the current element to the total XOR.

            if (x != 0) {
                hasNonZero = true; // A non-zero element can be removed if needed.
            }
        }

        if (xorValue != 0) {
            return nums.length; // The entire array already has a non-zero XOR.
        }

        if (hasNonZero) {
            return nums.length - 1; // Remove one non-zero element to make XOR non-zero.
        }

        return 0; // All elements are zero, so every subsequence has XOR zero.
    }
}