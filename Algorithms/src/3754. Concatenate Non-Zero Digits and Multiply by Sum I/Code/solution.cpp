class Solution
{
public:
    long long sumAndMultiply(int n)
    {
        // This will store the number made from all non-zero digits.
        long long x = 0;

        // This will store the sum of all non-zero digits.
        long long sum = 0;

        // Find the highest place value so I can read digits left to right.
        int divisor = 1;
        while (n / divisor >= 10)
        {
            divisor *= 10;
        }

        // Process every digit from left to right.
        while (divisor > 0)
        {
            // Get the digit at the current place value.
            int digit = n / divisor;

            // Remove the current digit from n.
            n %= divisor;

            // Only non-zero digits are added to x and sum.
            if (digit != 0)
            {
                // Shift x left by one decimal place and append the digit.
                x = x * 10 + digit;

                // Add the same digit to the digit sum.
                sum += digit;
            }

            // Move to the next smaller place value.
            divisor /= 10;
        }

        // Multiply the concatenated number by its digit sum.
        return x * sum;
    }
};