class Solution {
    public char findKthBit(int n, int k) {
        // Base case
        if (n == 1) return '0';

        int length = (1 << n) - 1;   // 2^n - 1
        int mid = (length + 1) / 2;

        if (k == mid) {
            return '1';
        } 
        else if (k < mid) {
            return findKthBit(n - 1, k);
        } 
        else {
            char bit = findKthBit(n - 1, length - k + 1);
            return bit == '0' ? '1' : '0';
        }
    }
}