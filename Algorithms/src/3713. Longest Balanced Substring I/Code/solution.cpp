class Solution
{
public:
    int longestBalanced(string s)
    {
        int n = s.size();
        int ans = 0;

        // Try every starting index
        for (int i = 0; i < n; i++)
        {
            vector<int> freq(26, 0);
            int distinct = 0;
            int maxFreq = 0;

            // Expand substring from i to j
            for (int j = i; j < n; j++)
            {
                int idx = s[j] - 'a';

                // Increase frequency
                if (freq[idx] == 0)
                    distinct++;

                freq[idx]++;
                maxFreq = max(maxFreq, freq[idx]);

                int length = j - i + 1;

                // Check balanced condition
                if (length == distinct * maxFreq)
                    ans = max(ans, length);
            }
        }

        return ans;
    }
};
