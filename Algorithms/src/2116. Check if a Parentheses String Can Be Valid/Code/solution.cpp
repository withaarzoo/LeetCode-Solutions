class Solution
{
public:
    bool canBeValid(string s, string locked)
    {
        if (s.size() % 2 != 0)
            return false; // Odd length can't be balanced

        int open = 0, flexible = 0;
        // Left-to-right pass
        for (int i = 0; i < s.size(); i++)
        {
            if (locked[i] == '1')
            {
                open += (s[i] == '(' ? 1 : -1);
            }
            else
            {
                flexible++;
            }
            if (open + flexible < 0)
                return false;
        }

        open = 0, flexible = 0;
        // Right-to-left pass
        for (int i = s.size() - 1; i >= 0; i--)
        {
            if (locked[i] == '1')
            {
                open += (s[i] == ')' ? 1 : -1);
            }
            else
            {
                flexible++;
            }
            if (open + flexible < 0)
                return false;
        }

        return true;
    }
};
