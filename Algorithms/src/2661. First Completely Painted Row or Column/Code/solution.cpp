class Solution
{
public:
    int firstCompleteIndex(vector<int> &arr, vector<vector<int>> &mat)
    {
        int m = mat.size(), n = mat[0].size();
        unordered_map<int, pair<int, int>> position;
        vector<int> rowCount(m, 0), colCount(n, 0);

        // Map matrix values to their positions
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {
                position[mat[i][j]] = {i, j};
            }
        }

        // Iterate through the array and simulate painting
        for (int i = 0; i < arr.size(); i++)
        {
            auto [row, col] = position[arr[i]];
            rowCount[row]++;
            colCount[col]++;

            if (rowCount[row] == n || colCount[col] == m)
            {
                return i;
            }
        }
        return -1;
    }
};
