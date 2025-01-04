class Solution
{
public:
    int countPalindromicSubsequence(string s)
    {
        vector<int> first(26, -1), last(26, -1);
        int n = s.size();

        // Record first and last occurrences
        for (int i = 0; i < n; i++)
        {
            int index = s[i] - 'a';
            if (first[index] == -1)
                first[index] = i;
            last[index] = i;
        }

        int result = 0;

        // Count unique middle characters for each letter
        for (int i = 0; i < 26; i++)
        {
            if (first[i] != -1 && last[i] > first[i])
            {
                unordered_set<char> middleChars;
                for (int j = first[i] + 1; j < last[i]; j++)
                {
                    middleChars.insert(s[j]);
                }
                result += middleChars.size();
            }
        }

        return result;
    }
};
