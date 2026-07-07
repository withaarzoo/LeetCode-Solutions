class Solution {
    public long sumAndMultiply(int n) {
        // This stores the number formed by all non-zero digits.
        long x = 0;

        // This stores the sum of all non-zero digits.
        long sum = 0;

        // Find the highest place value to read digits from left to right.
        int divisor = 1;
        while (n / divisor >= 10) {
            divisor *= 10;
        }

        // Process every digit from left to right.
        while (divisor > 0) {
            // Extract the digit at the current place value.
            int digit = n / divisor;

            // Remove the extracted digit from n.
            n %= divisor;

            // Ignore zero because it should not be part of x.
            if (digit != 0) {
                // Append the digit to the end of x.
                x = x * 10 + digit;

                // Add the digit to the sum.
                sum += digit;
            }

            // Move to the next digit.
            divisor /= 10;
        }

        // Return the required product.
        return x * sum;
    }
}