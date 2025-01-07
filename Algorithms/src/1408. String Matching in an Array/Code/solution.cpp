class Solution
{
public:
    vector<string> stringMatching(vector<string> &words)
    {
        // Sort words by length
        sort(words.begin(), words.end(), [](const string &a, const string &b)
             { return a.size() < b.size(); });

        vector<string> result;

        // Check for substrings
        for (int i = 0; i < words.size(); ++i)
        {
            for (int j = i + 1; j < words.size(); ++j)
            {
                if (words[j].find(words[i]) != string::npos)
                {
                    result.push_back(words[i]);
                    break;
                }
            }
        }

        return result;
    }
};
