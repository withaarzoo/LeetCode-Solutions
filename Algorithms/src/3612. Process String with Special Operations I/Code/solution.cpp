class Solution
{
public:
    string processStr(string s)
    {
        // Stores the current result being built
        string result;

        for (char c : s)
        {
            // Lowercase letter -> append to result
            if (c >= 'a' && c <= 'z')
            {
                result.push_back(c);
            }
            // Remove last character if it exists
            else if (c == '*')
            {
                if (!result.empty())
                {
                    result.pop_back();
                }
            }
            // Duplicate current result
            else if (c == '#')
            {
                result += result;
            }
            // Reverse current result
            else if (c == '%')
            {
                reverse(result.begin(), result.end());
            }
        }

        return result;
    }
};