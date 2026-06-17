class Solution
{
public:
    char processStr(string s, long long k)
    {
        int n = s.size();

        // len[i] = length of the result string after processing s[i]
        vector<long long> len(n);
        long long curLen = 0;

        for (int i = 0; i < n; i++)
        {
            char c = s[i];

            if (c >= 'a' && c <= 'z')
            {
                // Appending a letter increases length by 1
                curLen++;
            }
            else if (c == '*')
            {
                // Remove last character if it exists
                if (curLen > 0)
                    curLen--;
            }
            else if (c == '#')
            {
                // Duplicate the whole string
                curLen *= 2;
            }
            else
            { // '%'
                // Reversing does not change length
            }

            len[i] = curLen;
        }

        // k is outside the final string
        if (k >= curLen)
            return '.';

        // Walk backwards and undo operations
        for (int i = n - 1; i >= 0; i--)
        {
            char c = s[i];

            long long after = len[i];
            long long before = (i == 0 ? 0 : len[i - 1]);

            if (c >= 'a' && c <= 'z')
            {
                // This letter was appended at index "before"
                if (k == before)
                    return c;
            }
            else if (c == '#')
            {
                // T + T -> map back into the first copy
                if (before > 0)
                    k %= before;
            }
            else if (c == '%')
            {
                // Reverse operation
                k = before - 1 - k;
            }
            else
            {
                // '*' removed the last character.
                // All remaining indices stay unchanged.
            }
        }

        return '.';
    }
};