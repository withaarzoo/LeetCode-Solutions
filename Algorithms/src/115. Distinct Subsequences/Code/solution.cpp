class Solution
{
public:
    int numDistinct(string s, string t)
    {
        int m = t.size();

        // dp[j] stores the number of ways to form the first j characters of t.
        // dp[0] is 1 because there is exactly one way to form an empty string.
        vector<unsigned long long> dp(m + 1, 0);
        dp[0] = 1;

        // Process every character of s one by one.
        for (char c : s)
        {
            // Go from right to left so dp[j - 1] still represents
            // the previous state and the current character is not reused.
            for (int j = m; j >= 1; --j)
            {
                // If the current character can provide t[j - 1],
                // add all ways of forming the previous prefix.
                if (c == t[j - 1])
                {
                    dp[j] += dp[j - 1];
                }
            }
        }

        // dp[m] contains the number of ways to form all of t.
        return (int)dp[m];
    }
};