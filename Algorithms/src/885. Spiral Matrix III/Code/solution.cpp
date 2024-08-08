class Solution
{
public:
    vector<vector<int>> spiralMatrixIII(int rows, int cols, int rStart, int cStart)
    {
        // Initialize the result vector to store the coordinates
        vector<vector<int>> result;

        // Define the four possible directions: right, down, left, up
        vector<vector<int>> directions = {{0, 1}, {1, 0}, {0, -1}, {-1, 0}};

        // Start with 1 step to move
        int steps = 1;

        // Direction index (0=right, 1=down, 2=left, 3=up)
        int d = 0;

        // Starting coordinates
        int r = rStart, c = cStart;

        // Add the starting position to the result
        result.push_back({r, c});

        // Loop until we have visited all cells in the matrix
        while (result.size() < rows * cols)
        {
            // There are two legs per each step increment
            for (int i = 0; i < 2; ++i)
            {
                // Move in the current direction 'steps' times
                for (int j = 0; j < steps; ++j)
                {
                    // Move to the next cell
                    r += directions[d][0];
                    c += directions[d][1];

                    // Check if the new position is within bounds
                    if (r >= 0 && r < rows && c >= 0 && c < cols)
                    {
                        // If it is, add it to the result
                        result.push_back({r, c});
                    }
                }
                // Change direction: right -> down -> left -> up -> right -> ...
                d = (d + 1) % 4;
            }
            // After completing two legs, increase the number of steps
            ++steps;
        }

        // Return the final result containing all the coordinates in spiral order
        return result;
    }
};
