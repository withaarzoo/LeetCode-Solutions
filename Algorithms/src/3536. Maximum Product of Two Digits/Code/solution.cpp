class Solution
{
public:
    int maxProduct(int n)
    {
        // Store the largest digit found so far
        int first = 0;

        // Store the second largest digit found so far
        int second = 0;

        // Process every digit of the number
        while (n > 0)
        {
            // Extract the last digit
            int digit = n % 10;

            // If this digit becomes the new largest
            if (digit >= first)
            {
                // Old largest becomes second largest
                second = first;
                first = digit;
            }
            // Otherwise check if it should become second largest
            else if (digit > second)
            {
                second = digit;
            }

            // Remove the last digit
            n /= 10;
        }

        // Return the maximum product
        return first * second;
    }
};