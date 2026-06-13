class Solution
{
public:
    string mapWordWeights(vector<string> &words, vector<int> &weights)
    {
        string result;

        // Process each word independently
        for (const string &word : words)
        {
            int sumWeight = 0;

            // Add the weight of every character
            for (char ch : word)
            {
                sumWeight += weights[ch - 'a'];
            }

            // Reduce the weight into range [0, 25]
            int value = sumWeight % 26;

            // Reverse mapping:
            // 0 -> z, 1 -> y, ..., 25 -> a
            result.push_back(char('z' - value));
        }

        return result;
    }
};