class Solution
{
public:
    string shiftingLetters(string s, vector<vector<int>> &shifts)
    {
        int n = s.length();
        vector<int> diff(n + 1, 0);

        // Build the difference array
        for (auto &shift : shifts)
        {
            int start = shift[0], end = shift[1], direction = shift[2];
            int delta = (direction == 1) ? 1 : -1;
            diff[start] += delta;
            if (end + 1 < n)
                diff[end + 1] -= delta;
        }

        // Calculate cumulative shifts
        int shift = 0;
        for (int i = 0; i < n; ++i)
        {
            shift += diff[i];
            shift = (shift % 26 + 26) % 26; // Normalize shift to [0, 25]
            s[i] = 'a' + (s[i] - 'a' + shift) % 26;
        }

        return s;
    }
};
