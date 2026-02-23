class Solution
{
public:
    bool hasAllCodes(string s, int k)
    {
        int n = s.size();

        // If string is too small, impossible
        if (n < k)
            return false;

        int total = 1 << k; // total possible binary codes
        if (n - k + 1 < total)
            return false;

        vector<bool> seen(total, false);

        int mask = total - 1; // keeps only last k bits
        int curr = 0;
        int count = 0;

        // Build first k-length number
        for (int i = 0; i < k; i++)
        {
            curr = (curr << 1) | (s[i] - '0');
        }

        if (!seen[curr])
        {
            seen[curr] = true;
            count++;
        }

        // Sliding window
        for (int i = k; i < n; i++)
        {
            // Shift left and add new bit
            curr = ((curr << 1) & mask) | (s[i] - '0');

            if (!seen[curr])
            {
                seen[curr] = true;
                count++;
                if (count == total)
                    return true;
            }
        }

        return count == total;
    }
};