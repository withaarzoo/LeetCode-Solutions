class Solution {
public:
    char findKthBit(int n, int k) {
        // Base case
        if (n == 1) return '0';

        int length = (1 << n) - 1;        // 2^n - 1
        int mid = (length + 1) / 2;       // middle position

        if (k == mid) {
            return '1';
        } 
        else if (k < mid) {
            // Left half
            return findKthBit(n - 1, k);
        } 
        else {
            // Right half
            char bit = findKthBit(n - 1, length - k + 1);
            // Invert the result
            return bit == '0' ? '1' : '0';
        }
    }
};