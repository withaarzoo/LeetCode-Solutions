class Solution
{
public:
    int countPrefixSuffixPairs(vector<string> &words)
    {
        int n = words.size();
        int count = 0;

        for (int i = 0; i < n; ++i)
        {
            for (int j = i + 1; j < n; ++j)
            {
                // Check if words[i] is prefix and suffix of words[j]
                int len = words[i].size();
                if (words[j].substr(0, len) == words[i] && words[j].substr(words[j].size() - len) == words[i])
                {
                    count++;
                }
            }
        }

        return count;
    }
};
