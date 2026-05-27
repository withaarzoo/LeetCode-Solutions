class Solution
{
public:
    int numberOfSpecialChars(string word)
    {

        // Store last occurrence of lowercase letters
        vector<int> lower(26, -1);

        // Store first occurrence of uppercase letters
        vector<int> upper(26, -1);

        // Traverse the string
        for (int i = 0; i < word.size(); i++)
        {

            char ch = word[i];

            // If lowercase letter
            if (ch >= 'a' && ch <= 'z')
            {

                // Update last occurrence
                lower[ch - 'a'] = i;
            }
            else
            {

                // Convert uppercase letter to index
                int idx = ch - 'A';

                // Store only first occurrence
                if (upper[idx] == -1)
                {
                    upper[idx] = i;
                }
            }
        }

        int ans = 0;

        // Check all 26 letters
        for (int i = 0; i < 26; i++)
        {

            // Both lowercase and uppercase must exist
            if (lower[i] != -1 && upper[i] != -1)
            {

                // All lowercase must come before uppercase
                if (lower[i] < upper[i])
                {
                    ans++;
                }
            }
        }

        return ans;
    }
};