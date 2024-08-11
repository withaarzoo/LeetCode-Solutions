class Solution
{
public:
    int minDays(vector<vector<int>> &grid)
    {
        // Check if the grid is already disconnected.
        if (isDisconnected(grid))
            return 0;

        int m = grid.size();    // Number of rows in the grid
        int n = grid[0].size(); // Number of columns in the grid

        // Try removing one cell at a time and check if it disconnects the grid.
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (grid[i][j] == 1)
                {                   // If the cell is land (1)
                    grid[i][j] = 0; // Temporarily remove the land (turn it into water)
                    if (isDisconnected(grid))
                        return 1;   // If disconnected, return 1 (one day)
                    grid[i][j] = 1; // Revert the change if it didn't disconnect
                }
            }
        }

        // Try removing two cells at a time and check if it disconnects the grid.
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (grid[i][j] == 1)
                {                   // If the cell is land (1)
                    grid[i][j] = 0; // Temporarily remove the land (turn it into water)
                    for (int x = 0; x < m; ++x)
                    {
                        for (int y = 0; y < n; ++y)
                        {
                            if (grid[x][y] == 1)
                            {                   // If the second cell is land (1)
                                grid[x][y] = 0; // Temporarily remove the second piece of land
                                if (isDisconnected(grid))
                                    return 2;   // If disconnected, return 2 (two days)
                                grid[x][y] = 1; // Revert the change if it didn't disconnect
                            }
                        }
                    }
                    grid[i][j] = 1; // Revert the first change
                }
            }
        }

        // If no single or double removal disconnects the grid, it requires two days.
        return 2;
    }

private:
    bool isDisconnected(vector<vector<int>> &grid)
    {
        int m = grid.size();                               // Number of rows in the grid
        int n = grid[0].size();                            // Number of columns in the grid
        vector<vector<int>> visited(m, vector<int>(n, 0)); // Visited grid to track cells that have been explored

        int landCount = 0; // Count the number of land pieces in the grid
        for (int i = 0; i < m; ++i)
        {
            for (int j = 0; j < n; ++j)
            {
                if (grid[i][j] == 1)
                {                // If the cell is land (1)
                    ++landCount; // Increment the land count
                    if (!visited[i][j])
                    { // If the land hasn't been visited yet
                        if (landCount > 1)
                            return true;          // If more than one land mass exists, grid is disconnected
                        bfs(grid, visited, i, j); // Explore all connected land cells using BFS
                    }
                }
            }
        }
        return landCount == 0; // If there's no land, the grid is considered disconnected
    }

    void bfs(vector<vector<int>> &grid, vector<vector<int>> &visited, int i, int j)
    {
        int m = grid.size();     // Number of rows in the grid
        int n = grid[0].size();  // Number of columns in the grid
        queue<pair<int, int>> q; // Queue to perform BFS
        q.push({i, j});          // Start BFS from the given cell
        visited[i][j] = 1;       // Mark the starting cell as visited

        vector<int> dirX = {-1, 1, 0, 0}; // Direction vectors for row movement
        vector<int> dirY = {0, 0, -1, 1}; // Direction vectors for column movement

        // BFS to explore all connected land cells
        while (!q.empty())
        {
            auto [x, y] = q.front(); // Get the current cell from the queue
            q.pop();

            // Explore all 4 possible directions
            for (int d = 0; d < 4; ++d)
            {
                int newX = x + dirX[d];
                int newY = y + dirY[d];
                // Check if the new cell is within bounds and is land, and hasn't been visited yet
                if (newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 && !visited[newX][newY])
                {
                    visited[newX][newY] = 1; // Mark the new cell as visited
                    q.push({newX, newY});    // Add the new cell to the queue to continue BFS
                }
            }
        }
    }
};
