import java.util.HashSet;
import java.util.Set;

class Solution {
    public int minExtraChar(String s, String[] dictionary) {
        // Create a set to store all dictionary words for quick lookup
        Set<String> dict = new HashSet<>();

        // Populate the set with words from the dictionary
        for (String word : dictionary) {
            dict.add(word); // Add each word to the set
        }

        // Get the length of the input string `s`
        int n = s.length();

        // Create a DP (Dynamic Programming) array to store the minimum extra characters
        // required for each prefix of the string
        int[] dp = new int[n + 1]; // DP array has size n+1, where dp[i] represents the minimum
                                   // extra characters required for the first `i` characters of `s`

        // Initialize the DP array. Start by assuming the worst case, that each
        // character
        // in the string is extra, so set dp[i] = n for all i.
        for (int i = 0; i <= n; i++) {
            dp[i] = n; // Initially assume maximum extra characters
        }

        // Base case: no extra characters are needed for an empty string
        dp[0] = 0;

        // Iterate through the string from the first character to the nth character
        for (int i = 1; i <= n; i++) {
            // For each position `i`, check all possible substrings that end at `i`
            // This loop checks each prefix substring s[j:i] where `j` goes from 0 to i-1
            for (int j = 0; j < i; j++) {
                // Extract the substring `sub` that starts from index `j` and ends at index `i`
                String sub = s.substring(j, i); // s[j:i] is the substring from index j to i-1

                // Check if the substring exists in the dictionary
                if (dict.contains(sub)) {
                    // If the substring is in the dictionary, update dp[i] with the minimum value
                    // dp[i] can either be the current value or the value of dp[j]
                    // because dp[j] represents the minimum extra characters required for s[0:j].
                    dp[i] = Math.min(dp[i], dp[j]); // No extra character if found in the dictionary
                }
            }

            // Even if no valid dictionary word ends at `i`, we can consider the current
            // character
            // as an extra character. Thus, we update dp[i] to be at least dp[i-1] + 1,
            // where
            // dp[i-1] is the minimum extra characters required for the first `i-1`
            // characters.
            dp[i] = Math.min(dp[i], dp[i - 1] + 1); // Consider current character as extra
        }

        // The answer will be in dp[n], representing the minimum extra characters
        // required
        // for the entire string `s`.
        return dp[n];
    }
}
