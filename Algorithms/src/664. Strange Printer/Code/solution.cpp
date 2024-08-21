class Solution
{
public:
    int strangePrinter(string s)
    {
        int n = s.size(); // Get the length of the input string

        // Create a 2D vector (table) to store the minimum number of turns needed to print the substring s[i...j]
        vector<vector<int>> dp(n, vector<int>(n, 0));

        // Start filling the DP table from the end of the string towards the beginning
        for (int i = n - 1; i >= 0; --i)
        {
            dp[i][i] = 1; // A single character always needs exactly 1 turn to print

            // Consider all possible substrings starting from i to j
            for (int j = i + 1; j < n; ++j)
            {
                // Initialize the minimum turns for the substring s[i...j] as if we're printing the last character separately
                dp[i][j] = dp[i][j - 1] + 1;

                // Try to optimize by finding a previous occurrence of the current character s[j]
                for (int k = i; k < j; ++k)
                {
                    // If the character at position k matches the character at j,
                    // it means we might not need an additional turn to print s[j] since it can be printed along with s[k]
                    if (s[k] == s[j])
                    {
                        // Update the minimum turns by considering the best split at position k
                        dp[i][j] = min(dp[i][j], dp[i][k] + dp[k + 1][j - 1]);
                    }
                }
            }
        }

        // The result is stored in dp[0][n-1], representing the minimum turns to print the entire string s
        return dp[0][n - 1];
    }
};
