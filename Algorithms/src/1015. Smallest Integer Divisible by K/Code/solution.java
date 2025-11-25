class Solution {
    public int smallestRepunitDivByK(int k) {
        // If k has factor 2 or 5, impossible.
        if (k % 2 == 0 || k % 5 == 0)
            return -1;

        int rem = 0; // current remainder
        for (int length = 1; length <= k; length++) {
            // Update remainder when appending digit '1'
            rem = (rem * 10 + 1) % k;
            if (rem == 0) {
                return length;
            }
        }
        return -1;
    }
}
