class Solution {
    public int smallestNumber(int n, int t) {
        // Keep checking numbers starting from n
        while (true) {
            int product = 1;
            int x = n;

            // Calculate the product of all digits
            while (x > 0) {
                product *= (x % 10);
                x /= 10;
            }

            // Return the first valid number
            if (product % t == 0)
                return n;

            // Try the next number
            n++;
        }
    }
}