class Solution
{
public:
    int minDeletionSize(vector<string> &strs)
    {
        int rows = strs.size();
        int cols = strs[0].size();
        int deletions = 0;

        // Check each column
        for (int c = 0; c < cols; c++)
        {
            // Compare characters row by row
            for (int r = 0; r < rows - 1; r++)
            {
                if (strs[r][c] > strs[r + 1][c])
                {
                    deletions++; // Column is not sorted
                    break;       // No need to check further in this column
                }
            }
        }
        return deletions;
    }
};
