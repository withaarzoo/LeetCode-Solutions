func minDays(grid [][]int) int {
    // Check if the grid is already disconnected (no continuous land)
    // If disconnected, return 0 as no need to remove any land
    if isDisconnected(grid) {
        return 0
    }

    m, n := len(grid), len(grid[0])

    // Try removing one land cell at a time to see if it disconnects the grid
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                // Temporarily remove the land cell
                grid[i][j] = 0

                // Check if the removal disconnects the grid
                if isDisconnected(grid) {
                    return 1 // Return 1 if the grid becomes disconnected
                }

                // Restore the land cell
                grid[i][j] = 1
            }
        }
    }

    // If removing one land cell doesn't disconnect the grid,
    // try removing two land cells one by one
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                // Temporarily remove the first land cell
                grid[i][j] = 0

                // Try removing a second land cell
                for x := 0; x < m; x++ {
                    for y := 0; y < n; y++ {
                        if grid[x][y] == 1 {
                            // Temporarily remove the second land cell
                            grid[x][y] = 0

                            // Check if the grid becomes disconnected
                            if isDisconnected(grid) {
                                return 2 // Return 2 if the grid becomes disconnected
                            }

                            // Restore the second land cell
                            grid[x][y] = 1
                        }
                    }
                }

                // Restore the first land cell
                grid[i][j] = 1
            }
        }
    }

    // If removing two land cells doesn't disconnect the grid, return 2
    return 2
}

// Helper function to check if the grid is disconnected
func isDisconnected(grid [][]int) bool {
    m, n := len(grid), len(grid[0])

    // Initialize a visited array to track which cells have been visited
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    landCount := 0 // Counter to track the number of land cells

    // Traverse the grid to find land cells
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                landCount++ // Increment land cell count
                if !visited[i][j] {
                    // If more than one land cell has been found, check for disconnection
                    if landCount > 1 {
                        return true // Return true if grid is disconnected
                    }
                    // Perform BFS to mark all connected land cells
                    bfs(grid, visited, i, j)
                }
            }
        }
    }

    // If there are no land cells left, return true (considered disconnected)
    return landCount == 0
}

// Helper function to perform BFS and mark all connected land cells
func bfs(grid [][]int, visited [][]bool, i, j int) {
    m, n := len(grid), len(grid[0])

    // Directions for moving up, down, left, and right
    dirX := []int{-1, 1, 0, 0}
    dirY := []int{0, 0, -1, 1}

    // Initialize queue for BFS and mark the starting cell as visited
    queue := [][]int{{i, j}}
    visited[i][j] = true

    // Process each cell in the queue
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]

        // Explore all 4 possible directions (up, down, left, right)
        for d := 0; d < 4; d++ {
            newX, newY := x+dirX[d], y+dirY[d]
            // Check if the new cell is within bounds and is land and not visited
            if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == 1 && !visited[newX][newY] {
                visited[newX][newY] = true // Mark the new cell as visited
                queue = append(queue, []int{newX, newY}) // Add the new cell to the queue
            }
        }
    }
}
