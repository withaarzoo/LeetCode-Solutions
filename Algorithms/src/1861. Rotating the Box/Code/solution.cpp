class Solution
{
public:
    vector<vector<char>> rotateTheBox(vector<vector<char>> &boxGrid)
    {
        int m = boxGrid.size();
        int n = boxGrid[0].size();

        // Process every row independently
        for (int row = 0; row < m; row++)
        {

            // This points to the rightmost empty position
            // where the next stone can fall
            int emptyCol = n - 1;

            // Traverse from right to left
            for (int col = n - 1; col >= 0; col--)
            {

                // Obstacle blocks movement
                if (boxGrid[row][col] == '*')
                {

                    // Stones can only fall before obstacle
                    emptyCol = col - 1;
                }

                // Found a stone
                else if (boxGrid[row][col] == '#')
                {

                    // Remove stone from current position
                    boxGrid[row][col] = '.';

                    // Put stone at the valid empty position
                    boxGrid[row][emptyCol] = '#';

                    // Next stone should go one step left
                    emptyCol--;
                }
            }
        }

        // Create rotated matrix
        vector<vector<char>> rotated(n, vector<char>(m));

        // Rotate clockwise
        for (int i = 0; i < m; i++)
        {
            for (int j = 0; j < n; j++)
            {

                // Standard clockwise rotation formula
                rotated[j][m - 1 - i] = boxGrid[i][j];
            }
        }

        return rotated;
    }
};