class Solution {
    public int minPartitions(String n) {
        int maxDigit = 0; // Store the maximum digit

        // Traverse each character
        for (int i = 0; i < n.length(); i++) {
            int digit = n.charAt(i) - '0'; // Convert char to int

            // Update max
            if (digit > maxDigit) {
                maxDigit = digit;
            }

            // If 9 found, we can stop early
            if (maxDigit == 9) {
                break;
            }
        }

        return maxDigit;
    }
}