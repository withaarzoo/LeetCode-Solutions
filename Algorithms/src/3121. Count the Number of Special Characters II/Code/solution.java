class Solution {
    public int numberOfSpecialChars(String word) {

        // Store last occurrence of lowercase letters
        int[] lower = new int[26];

        // Store first occurrence of uppercase letters
        int[] upper = new int[26];

        // Initialize arrays with -1
        Arrays.fill(lower, -1);
        Arrays.fill(upper, -1);

        // Traverse the string
        for (int i = 0; i < word.length(); i++) {

            char ch = word.charAt(i);

            // If lowercase letter
            if (ch >= 'a' && ch <= 'z') {

                // Update last occurrence
                lower[ch - 'a'] = i;
            } else {

                int idx = ch - 'A';

                // Store only first occurrence
                if (upper[idx] == -1) {
                    upper[idx] = i;
                }
            }
        }

        int ans = 0;

        // Check all letters
        for (int i = 0; i < 26; i++) {

            // Both lowercase and uppercase must exist
            if (lower[i] != -1 && upper[i] != -1) {

                // Lowercase must appear before uppercase
                if (lower[i] < upper[i]) {
                    ans++;
                }
            }
        }

        return ans;
    }
}