class Solution
{
public:
    int minSwaps(vector<vector<int>> &grid)
    {
        int n = grid.size();

        // Step 1: Count trailing zeros in each row
        vector<int> trailing(n, 0);

        for (int i = 0; i < n; i++)
        {
            int count = 0;
            for (int j = n - 1; j >= 0; j--)
            {
                if (grid[i][j] == 0)
                    count++;
                else
                    break; // stop when we see first 1
            }
            trailing[i] = count;
        }

        int swaps = 0;

        // Step 2: Try placing rows one by one
        for (int i = 0; i < n; i++)
        {
            int required = n - 1 - i;
            int j = i;

            // Find a row that satisfies requirement
            while (j < n && trailing[j] < required)
                j++;

            if (j == n)
                return -1; // no valid row found

            // Bring row j up to position i
            while (j > i)
            {
                swap(trailing[j], trailing[j - 1]);
                swaps++;
                j--;
            }
        }

        return swaps;
    }
};