class Solution {
    public boolean hasAlternatingBits(int n) {

        int prev = n & 1;
        n >>= 1;

        while (n > 0) {
            int curr = n & 1;

            if (curr == prev) {
                return false;
            }

            prev = curr;
            n >>= 1;
        }

        return true;
    }
}
