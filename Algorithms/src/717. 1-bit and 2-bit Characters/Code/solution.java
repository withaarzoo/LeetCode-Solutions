class Solution {
    public boolean isOneBitCharacter(int[] bits) {
        int n = bits.length;
        int i = 0;
        // iterate until index reaches the last element
        while (i < n - 1) {
            if (bits[i] == 1) {
                // 1 must be the start of a two-bit character
                i += 2;
            } else {
                // 0 is a one-bit character
                i += 1;
            }
        }
        // if i == last index, last character is 1-bit
        return i == n - 1;
    }
}
