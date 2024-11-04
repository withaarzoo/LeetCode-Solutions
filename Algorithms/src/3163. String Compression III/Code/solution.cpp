class Solution
{
public:
    string compressedString(string word)
    {
        string comp;
        int count = 1;

        for (int i = 1; i <= word.size(); ++i)
        {
            if (i == word.size() || word[i] != word[i - 1] || count == 9)
            {
                comp += to_string(count) + word[i - 1];
                count = 1;
            }
            else
            {
                count++;
            }
        }

        return comp;
    }
};
