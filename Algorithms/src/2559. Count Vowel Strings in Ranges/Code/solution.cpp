#include <vector>
#include <string>
#include <unordered_set>
using namespace std;

class Solution
{
public:
    vector<int> vowelStrings(vector<string> &words, vector<vector<int>> &queries)
    {
        unordered_set<char> vowels = {'a', 'e', 'i', 'o', 'u'};
        int n = words.size();
        vector<int> prefix(n, 0);

        // Precompute the prefix sum
        for (int i = 0; i < n; ++i)
        {
            if (vowels.count(words[i].front()) && vowels.count(words[i].back()))
            {
                prefix[i] = 1;
            }
            if (i > 0)
            {
                prefix[i] += prefix[i - 1];
            }
        }

        // Answer the queries
        vector<int> result;
        for (const auto &query : queries)
        {
            int l = query[0], r = query[1];
            result.push_back(prefix[r] - (l > 0 ? prefix[l - 1] : 0));
        }
        return result;
    }
};