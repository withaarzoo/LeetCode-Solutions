class Solution
{
public:
    string generateString(string str1, string str2)
    {
        int n = str1.size();
        int m = str2.size();
        int len = n + m - 1;

        string ans(len, '?');
        vector<bool> fixed(len, false);

        // Apply all 'T' constraints
        for (int i = 0; i < n; i++)
        {
            if (str1[i] == 'T')
            {
                for (int j = 0; j < m; j++)
                {
                    int pos = i + j;

                    if (ans[pos] != '?' && ans[pos] != str2[j])
                    {
                        return "";
                    }

                    ans[pos] = str2[j];
                    fixed[pos] = true;
                }
            }
        }

        // Fill remaining positions with 'a'
        for (int i = 0; i < len; i++)
        {
            if (ans[i] == '?')
                ans[i] = 'a';
        }

        // Process all 'F' constraints
        for (int i = 0; i < n; i++)
        {
            if (str1[i] == 'F')
            {
                bool same = true;

                for (int j = 0; j < m; j++)
                {
                    if (ans[i + j] != str2[j])
                    {
                        same = false;
                        break;
                    }
                }

                // Already different
                if (!same)
                    continue;

                bool changed = false;

                // Try changing from right to left
                for (int j = m - 1; j >= 0; j--)
                {
                    int pos = i + j;

                    if (fixed[pos])
                        continue;

                    // Change to smallest possible different character
                    for (char c = 'a'; c <= 'z'; c++)
                    {
                        if (c != ans[pos] && c != str2[j])
                        {
                            ans[pos] = c;
                            changed = true;
                            break;
                        }
                    }

                    if (changed)
                        break;
                }

                if (!changed)
                    return "";
            }
        }

        return ans;
    }
};