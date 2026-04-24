class Solution
{
public:
    int furthestDistanceFromOrigin(string moves)
    {
        int left = 0, right = 0, blank = 0;

        // Count occurrences
        for (char c : moves)
        {
            if (c == 'L')
                left++;
            else if (c == 'R')
                right++;
            else
                blank++; // '_'
        }

        // Current position = right - left
        int position = right - left;

        // Max distance = absolute position + blanks
        return abs(position) + blank;
    }
};