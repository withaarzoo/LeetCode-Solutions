class Solution
{
public:
    string lexGreaterPermutation(string s, string target)
    {
        // Store the frequency of every character available in s.
        vector<int> count(26, 0);
        for (char ch : s)
        {
            count[ch - 'a']++;
        }

        int n = s.size();
        int matched = 0;

        // Try to keep the answer exactly equal to target for as long as possible.
        while (matched < n && count[target[matched] - 'a'] > 0)
        {
            // Use target[matched] because matching it keeps the prefix smallest.
            count[target[matched] - 'a']--;
            matched++;
        }

        // If matching failed before reaching the end, first try to increase
        // exactly at the position where matching became impossible.
        int start = (matched < n ? matched : n - 1);

        // Move from right to left because changing a later position gives
        // a lexicographically smaller answer than changing an earlier one.
        for (int i = start; i >= 0; i--)
        {
            // If this position was previously matched, undo that choice
            // so its character becomes available again.
            if (i < matched)
            {
                count[target[i] - 'a']++;
            }

            // Find the smallest available character strictly greater than target[i].
            int bigger = -1;
            for (int ch = target[i] - 'a' + 1; ch < 26; ch++)
            {
                if (count[ch] > 0)
                {
                    bigger = ch;
                    break;
                }
            }

            // If a larger character exists, build the smallest possible answer.
            if (bigger != -1)
            {
                // Use the larger character to make the whole string > target.
                count[bigger]--;

                // Keep everything before i equal to target.
                string answer = target.substr(0, i);

                // Place the smallest possible larger character at position i.
                answer += char('a' + bigger);

                // Append all remaining characters in sorted order.
                for (int ch = 0; ch < 26; ch++)
                {
                    answer.append(count[ch], char('a' + ch));
                }

                return answer;
            }
        }

        // No position can be increased, so no valid permutation exists.
        return "";
    }
};