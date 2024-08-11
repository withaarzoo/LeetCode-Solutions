class Solution {
    public int minDays(int[][] grid) {
        // Check if the grid is already disconnected
        if (isDisconnected(grid))
            return 0; // If already disconnected, no need to remove any land

        int m = grid.length; // Number of rows in the grid
        int n = grid[0].length; // Number of columns in the grid

        // First pass: Try removing one piece of land at a time
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) { // Check if the current cell is land
                    grid[i][j] = 0; // Temporarily remove the land
                    if (isDisconnected(grid))
                        return 1; // If the grid becomes disconnected, return 1
                    grid[i][j] = 1; // Restore the land
                }
            }
        }

        // Second pass: Try removing two pieces of land at a time
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) { // Check if the current cell is land
                    grid[i][j] = 0; // Temporarily remove the land
                    for (int x = 0; x < m; ++x) {
                        for (int y = 0; y < n; ++y) {
                            if (grid[x][y] == 1) { // Check if another cell is land
                                grid[x][y] = 0; // Temporarily remove the second piece of land
                                if (isDisconnected(grid))
                                    return 2; // If the grid becomes disconnected, return 2
                                grid[x][y] = 1; // Restore the second piece of land
                            }
                        }
                    }
                    grid[i][j] = 1; // Restore the first piece of land
                }
            }
        }

        return 2; // If the grid cannot be disconnected by removing one or two pieces of land,
                  // return 2
    }

    private boolean isDisconnected(int[][] grid) {
        int m = grid.length; // Number of rows in the grid
        int n = grid[0].length; // Number of columns in the grid
        boolean[][] visited = new boolean[m][n]; // Visited array to track the cells that have been explored

        int landCount = 0; // Counter to track the number of land pieces
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) { // If the current cell is land
                    ++landCount; // Increment the land counter
                    if (!visited[i][j]) { // If the current cell has not been visited
                        if (landCount > 1)
                            return true; // If more than one land piece is found, the grid is disconnected
                        bfs(grid, visited, i, j); // Perform BFS to mark all connected land pieces
                    }
                }
            }
        }

        return landCount == 0; // Return true if no land is left, meaning the grid is disconnected
    }

    private void bfs(int[][] grid, boolean[][] visited, int i, int j) {
        int m = grid.length; // Number of rows in the grid
        int n = grid[0].length; // Number of columns in the grid
        Queue<int[]> q = new LinkedList<>(); // Queue for BFS
        q.offer(new int[] { i, j }); // Add the starting cell to the queue
        visited[i][j] = true; // Mark the starting cell as visited

        // Directions for moving up, down, left, and right
        int[] dirX = { -1, 1, 0, 0 };
        int[] dirY = { 0, 0, -1, 1 };

        // BFS loop
        while (!q.isEmpty()) {
            int[] current = q.poll(); // Get the current cell from the queue
            int x = current[0];
            int y = current[1];

            // Explore the neighboring cells
            for (int d = 0; d < 4; ++d) {
                int newX = x + dirX[d];
                int newY = y + dirY[d];
                // Check if the neighboring cell is within bounds, is land, and has not been
                // visited
                if (newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 && !visited[newX][newY]) {
                    visited[newX][newY] = true; // Mark the neighboring cell as visited
                    q.offer(new int[] { newX, newY }); // Add the neighboring cell to the queue
                }
            }
        }
    }
}
