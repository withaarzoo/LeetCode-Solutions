class Solution {
    public int bitwiseComplement(int n) {

        // Edge case
        if (n == 0)
            return 1;

        int mask = 0;

        // Create mask of all 1s
        while (mask < n) {
            mask = (mask << 1) | 1;
        }

        // XOR to flip bits
        return mask ^ n;
    }
}