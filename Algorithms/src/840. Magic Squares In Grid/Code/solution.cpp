#include <vector>
using namespace std;

class Solution
{
public:
    int numMagicSquaresInside(vector<vector<int>> &grid)
    {
        // Get the number of rows and columns in the grid
        int rows = grid.size();
        int cols = grid[0].size();

        // Initialize a counter to keep track of the number of magic squares found
        int count = 0;

        // Lambda function to check if the 3x3 grid with top-left corner at (r, c) is a magic square
        auto isMagicSquare = [&](int r, int c) -> bool
        {
            // Create a vector to track the frequency of numbers 1 to 9 in the 3x3 grid
            vector<int> vals(10, 0);

            // Traverse the 3x3 subgrid
            for (int i = 0; i < 3; ++i)
            {
                for (int j = 0; j < 3; ++j)
                {
                    // Get the current number in the 3x3 grid
                    int num = grid[r + i][c + j];

                    // Check if the number is between 1 and 9, and if it has already been seen
                    if (num < 1 || num > 9 || vals[num])
                        return false;

                    // Mark this number as seen
                    vals[num] = 1;
                }
            }

            // Check if all rows, columns, and diagonals sum to 15
            return (grid[r][c] + grid[r][c + 1] + grid[r][c + 2] == 15 &&             // Row 1
                    grid[r + 1][c] + grid[r + 1][c + 1] + grid[r + 1][c + 2] == 15 && // Row 2
                    grid[r + 2][c] + grid[r + 2][c + 1] + grid[r + 2][c + 2] == 15 && // Row 3
                    grid[r][c] + grid[r + 1][c] + grid[r + 2][c] == 15 &&             // Column 1
                    grid[r][c + 1] + grid[r + 1][c + 1] + grid[r + 2][c + 1] == 15 && // Column 2
                    grid[r][c + 2] + grid[r + 1][c + 2] + grid[r + 2][c + 2] == 15 && // Column 3
                    grid[r][c] + grid[r + 1][c + 1] + grid[r + 2][c + 2] == 15 &&     // Diagonal from top-left to bottom-right
                    grid[r][c + 2] + grid[r + 1][c + 1] + grid[r + 2][c] == 15);      // Diagonal from top-right to bottom-left
        };

        // Traverse the grid, checking every possible 3x3 subgrid
        for (int i = 0; i < rows - 2; ++i)
        {
            for (int j = 0; j < cols - 2; ++j)
            {
                // If the subgrid starting at (i, j) is a magic square, increment the counter
                if (isMagicSquare(i, j))
                {
                    ++count;
                }
            }
        }

        // Return the total count of magic squares found in the grid
        return count;
    }
};
