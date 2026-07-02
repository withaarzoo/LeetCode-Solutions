/**
 * @param {number[][]} grid
 * @param {number} health
 * @return {boolean}
 */
var findSafeWalk = function (grid, health) {
  const m = grid.length;
  const n = grid[0].length;

  // Minimum health lost for each cell
  const dist = Array.from({ length: m }, () => Array(n).fill(Infinity));

  // Deque implementation using an array and head pointer
  const deque = [];
  let head = 0;

  // Starting cost includes the starting cell
  dist[0][0] = grid[0][0];
  deque.push([0, 0]);

  const dx = [-1, 0, 1, 0];
  const dy = [0, 1, 0, -1];

  while (head < deque.length) {
    const [x, y] = deque[head++];

    // Explore all four directions
    for (let k = 0; k < 4; k++) {
      const nx = x + dx[k];
      const ny = y + dy[k];

      // Skip invalid cells
      if (nx < 0 || ny < 0 || nx >= m || ny >= n) continue;

      // Cost after entering the next cell
      const newCost = dist[x][y] + grid[nx][ny];

      // Update only if this path is better
      if (newCost < dist[nx][ny]) {
        dist[nx][ny] = newCost;

        // Weight 0 should be processed earlier
        if (grid[nx][ny] === 0) {
          deque.splice(head, 0, [nx, ny]);
        } else {
          deque.push([nx, ny]);
        }
      }
    }
  }

  // Health must stay at least 1
  return dist[m - 1][n - 1] < health;
};
