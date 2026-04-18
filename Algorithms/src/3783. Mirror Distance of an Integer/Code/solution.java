class Solution {
    public int mirrorDistance(int n) {
        int rev = 0;
        int temp = n;

        // Reverse the digits of n
        while (temp > 0) {
            int digit = temp % 10; // Get last digit
            rev = rev * 10 + digit; // Add digit to reversed number
            temp /= 10; // Remove last digit
        }

        // Return absolute difference
        return Math.abs(n - rev);
    }
}