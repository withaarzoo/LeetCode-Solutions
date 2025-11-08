class Solution {
    public int minimumOneBitOperations(int n) {
        // Inverse Gray code using iterative XOR and shift
        int ans = 0;
        while (n != 0) {
            ans ^= n;
            n >>= 1;
        }
        return ans;
    }
}
