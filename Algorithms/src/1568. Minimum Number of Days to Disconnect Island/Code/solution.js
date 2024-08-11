/**
 * This function determines the minimum number of days needed to disconnect a grid of islands.
 * @param {number[][]} grid - A 2D grid representing islands (1) and water (0).
 * @return {number} - The minimum number of days required to disconnect the grid.
 */
var minDays = function (grid) {
  // Step 1: Check if the grid is already disconnected
  if (isDisconnected(grid)) return 0;

  const m = grid.length; // Number of rows in the grid
  const n = grid[0].length; // Number of columns in the grid

  // Step 2: Try removing one cell to see if it disconnects the grid
  for (let i = 0; i < m; ++i) {
    for (let j = 0; j < n; ++j) {
      if (grid[i][j] === 1) {
        // If the cell is land
        grid[i][j] = 0; // Temporarily change it to water
        if (isDisconnected(grid)) return 1; // If it disconnects, return 1 day
        grid[i][j] = 1; // Otherwise, revert the cell back to land
      }
    }
  }

  // Step 3: Try removing two cells to see if it disconnects the grid
  for (let i = 0; i < m; ++i) {
    for (let j = 0; j < n; ++j) {
      if (grid[i][j] === 1) {
        // If the cell is land
        grid[i][j] = 0; // Temporarily change it to water
        for (let x = 0; x < m; ++x) {
          for (let y = 0; y < n; ++y) {
            if (grid[x][y] === 1) {
              // If another cell is land
              grid[x][y] = 0; // Temporarily change it to water
              if (isDisconnected(grid)) return 2; // If it disconnects, return 2 days
              grid[x][y] = 1; // Otherwise, revert the second cell back to land
            }
          }
        }
        grid[i][j] = 1; // Revert the first cell back to land
      }
    }
  }

  // If no disconnection occurs after trying to remove up to two cells, return 2 days
  return 2;
};

/**
 * This function checks if the grid is disconnected (i.e., more than one isolated landmass).
 * @param {number[][]} grid - The grid to check.
 * @return {boolean} - True if the grid is disconnected, otherwise false.
 */
function isDisconnected(grid) {
  const m = grid.length;
  const n = grid[0].length;
  const visited = Array.from({ length: m }, () => Array(n).fill(false));

  let landCount = 0; // Count of land cells encountered
  for (let i = 0; i < m; ++i) {
    for (let j = 0; j < n; ++j) {
      if (grid[i][j] === 1) {
        // If the cell is land
        ++landCount; // Increment the land count
        if (!visited[i][j]) {
          // If the cell hasn't been visited
          if (landCount > 1) return true; // If more than one landmass is found, return true (disconnected)
          bfs(grid, visited, i, j); // Perform BFS to visit all connected land cells
        }
      }
    }
  }

  // If no land is found or only one landmass, the grid is not disconnected
  return landCount === 0;
}

/**
 * This function performs a BFS to explore all connected land cells.
 * @param {number[][]} grid - The grid containing land and water.
 * @param {boolean[][]} visited - A grid to track visited cells.
 * @param {number} i - The starting row index.
 * @param {number} j - The starting column index.
 */
function bfs(grid, visited, i, j) {
  const m = grid.length;
  const n = grid[0].length;
  const queue = [[i, j]]; // Queue for BFS, starting with the initial land cell
  visited[i][j] = true; // Mark the starting cell as visited

  // Direction vectors for moving up, down, left, and right
  const dirX = [-1, 1, 0, 0];
  const dirY = [0, 0, -1, 1];

  // Perform BFS to visit all connected land cells
  while (queue.length > 0) {
    const [x, y] = queue.shift(); // Dequeue the current cell

    for (let d = 0; d < 4; ++d) {
      // Check all four possible directions
      const newX = x + dirX[d];
      const newY = y + dirY[d];
      // Check if the new cell is within bounds, is land, and hasn't been visited
      if (
        newX >= 0 &&
        newX < m &&
        newY >= 0 &&
        newY < n &&
        grid[newX][newY] === 1 &&
        !visited[newX][newY]
      ) {
        visited[newX][newY] = true; // Mark the new cell as visited
        queue.push([newX, newY]); // Enqueue the new cell
      }
    }
  }
}
