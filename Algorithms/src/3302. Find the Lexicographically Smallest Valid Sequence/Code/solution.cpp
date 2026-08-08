class Solution {
public:
    vector<int> validSequence(string word1, string word2) {
        int n = word1.size(); // I store the length of word1 for boundary checks.
        int m = word2.size(); // I store the length of word2 because the answer needs m indices.

        vector<int> last(m, -1); // I store the position that can match each suffix character of word2.

        int i = n - 1; // I start from the last character of word1.
        int j = m - 1; // I start by matching the last character of word2.

        // I build a valid matching arrangement from right to left.
        // This tells me how far left I must stay if I use a mismatch earlier.
        while (i >= 0 && j >= 0) {
            // If the current characters match, I can use this position for word2[j].
            if (word1[i] == word2[j]) {
                last[j] = i; // I remember this position for the suffix starting at j.
                --j;         // I now need to match the previous character of word2.
            }

            --i; // I continue searching toward the beginning of word1.
        }

        vector<int> ans; // I store the lexicographically smallest sequence here.
        ans.reserve(m); // I reserve exactly m positions because the answer has length m.

        bool canSkip = true; // I have not used my one allowed mismatch yet.
        j = 0;               // I start matching word2 from its first character.

        // I scan from left to right so that I always prefer the smallest possible index.
        for (i = 0; i < n && j < m; ++i) {
            // If the characters match, taking this earliest index is always safe.
            if (word1[i] == word2[j]) {
                ans.push_back(i); // I choose this index for word2[j].
                ++j;              // I move to the next character of word2.
            }
            // Otherwise, I try to spend my one allowed mismatch.
            else if (canSkip &&
                     (j == m - 1 || i < last[j + 1])) {
                canSkip = false;  // I use the only allowed character change.
                ans.push_back(i); // I choose this earliest possible index.
                ++j;              // I move to the next character of word2.
            }
        }

        // I return the answer only if every character of word2 was matched.
        if (j == m) {
            return ans;
        }

        // If I could not match all of word2, no valid sequence exists.
        return {};
    }
};