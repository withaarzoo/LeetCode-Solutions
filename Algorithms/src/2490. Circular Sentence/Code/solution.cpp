class Solution
{
public:
    bool isCircularSentence(string sentence)
    {
        // Step 1: Split the sentence into words
        vector<string> words;
        string word;
        stringstream ss(sentence);

        while (ss >> word)
        {
            words.push_back(word);
        }

        // Step 2: Check adjacent pairs and the circular condition
        for (int i = 0; i < words.size(); ++i)
        {
            char lastChar = words[i].back();
            char firstChar = words[(i + 1) % words.size()].front();
            if (lastChar != firstChar)
            {
                return false;
            }
        }

        return true;
    }
};
