class Solution
{
public:
    string makeLargestSpecial(string s)
    {
        vector<string> parts; // Store special substrings
        int count = 0;
        int start = 0;

        // Split into primitive special substrings
        for (int i = 0; i < s.size(); i++)
        {
            if (s[i] == '1')
                count++;
            else
                count--;

            // Found a balanced substring
            if (count == 0)
            {
                // Recursively process inner part
                string inner = makeLargestSpecial(s.substr(start + 1, i - start - 1));
                parts.push_back("1" + inner + "0");
                start = i + 1;
            }
        }

        // Sort in descending order for lexicographically largest
        sort(parts.begin(), parts.end(), greater<string>());

        // Combine all parts
        string result = "";
        for (auto &p : parts)
        {
            result += p;
        }

        return result;
    }
};