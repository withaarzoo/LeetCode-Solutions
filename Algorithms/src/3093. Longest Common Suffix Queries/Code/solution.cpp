class Solution
{
    // Structure to define a node in the Trie
    struct TrieNode
    {
        int children[26];
        int bestLen;
        int bestIdx;

        // Initialize children to -1 and best trackers to a large number
        TrieNode()
        {
            fill(begin(children), end(children), -1);
            bestLen = 1e9;
            bestIdx = 1e9;
        }
    };

public:
    vector<int> stringIndices(vector<string> &wordsContainer, vector<string> &wordsQuery)
    {
        // Using a vector of nodes instead of pointers to avoid memory overhead and MLE
        vector<TrieNode> trie;
        trie.emplace_back(); // Push root node

        // Insert each string from the container into the Trie
        for (int i = 0; i < wordsContainer.size(); i++)
        {
            int len = wordsContainer[i].length();
            int curr = 0; // Start at root

            // Update root with the absolute best string (in case of zero matches later)
            if (len < trie[curr].bestLen || (len == trie[curr].bestLen && i < trie[curr].bestIdx))
            {
                trie[curr].bestLen = len;
                trie[curr].bestIdx = i;
            }

            // Traverse the string backwards to simulate suffix matching as prefix matching
            for (int j = len - 1; j >= 0; j--)
            {
                int charIdx = wordsContainer[i][j] - 'a';

                // If child path doesn't exist, create a new node
                if (trie[curr].children[charIdx] == -1)
                {
                    trie[curr].children[charIdx] = trie.size();
                    trie.emplace_back();
                }

                // Move down the Trie
                curr = trie[curr].children[charIdx];

                // Update the best string properties for this specific prefix path
                if (len < trie[curr].bestLen || (len == trie[curr].bestLen && i < trie[curr].bestIdx))
                {
                    trie[curr].bestLen = len;
                    trie[curr].bestIdx = i;
                }
            }
        }

        vector<int> ans;
        ans.reserve(wordsQuery.size());

        // Process each query
        for (const string &query : wordsQuery)
        {
            int curr = 0; // Start at root
            int len = query.length();

            // Traverse backwards down the Trie
            for (int j = len - 1; j >= 0; j--)
            {
                int charIdx = query[j] - 'a';
                // Stop if the matching path ends
                if (trie[curr].children[charIdx] == -1)
                {
                    break;
                }
                curr = trie[curr].children[charIdx];
            }
            // The current node holds the index of the best matching string
            ans.push_back(trie[curr].bestIdx);
        }

        return ans;
    }
};