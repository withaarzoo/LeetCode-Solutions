class Solution {
    public int maxProduct(int n) {

        // Store the largest digit
        int first = 0;

        // Store the second largest digit
        int second = 0;

        // Process every digit
        while (n > 0) {

            // Extract the last digit
            int digit = n % 10;

            // Update largest and second largest if needed
            if (digit >= first) {
                second = first;
                first = digit;
            } else if (digit > second) {
                second = digit;
            }

            // Remove the last digit
            n /= 10;
        }

        // Return the answer
        return first * second;
    }
}