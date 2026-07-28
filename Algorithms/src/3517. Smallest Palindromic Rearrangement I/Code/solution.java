class Solution {
    public String smallestPalindrome(String s) {

        // Store frequency of every lowercase letter
        int[] freq = new int[26];

        // Count character frequencies
        for (char c : s.toCharArray())
            freq[c - 'a']++;

        StringBuilder left = new StringBuilder();
        String middle = "";

        // Build the left half and find the middle character
        for (int i = 0; i < 26; i++) {

            // Add half of the occurrences to the left side
            for (int j = 0; j < freq[i] / 2; j++)
                left.append((char) ('a' + i));

            // Odd frequency character becomes the center
            if (freq[i] % 2 == 1)
                middle = String.valueOf((char) ('a' + i));
        }

        // Right half is the reverse of the left half
        String right = new StringBuilder(left).reverse().toString();

        // Return the complete palindrome
        return left.toString() + middle + right;
    }
}