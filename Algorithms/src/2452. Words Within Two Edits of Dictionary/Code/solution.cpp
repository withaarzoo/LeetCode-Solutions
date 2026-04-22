class Solution
{
public:
    vector<string> twoEditWords(vector<string> &queries, vector<string> &dictionary)
    {
        vector<string> result;

        // Check every query word
        for (string query : queries)
        {

            // Compare current query with every dictionary word
            for (string word : dictionary)
            {
                int diff = 0;

                // Count different characters
                for (int i = 0; i < query.size(); i++)
                {
                    if (query[i] != word[i])
                    {
                        diff++;
                    }

                    // No need to continue if already more than 2 edits
                    if (diff > 2)
                    {
                        break;
                    }
                }

                // If current dictionary word can match within 2 edits
                if (diff <= 2)
                {
                    result.push_back(query);
                    break;
                }
            }
        }

        return result;
    }
};