/**
 * @param {number[][]} grid
 * @return {number}
 */
var maximumSafenessFactor = function (grid) {
  const n = grid.length;

  const dist = Array.from({ length: n }, () => Array(n).fill(-1));
  const queue = [];

  // Push every thief into the queue
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (grid[i][j] === 1) {
        dist[i][j] = 0;
        queue.push([i, j]);
      }
    }
  }

  const dir = [-1, 0, 1, 0, -1];
  let head = 0;

  // Multi-source BFS
  while (head < queue.length) {
    const [x, y] = queue[head++];

    for (let k = 0; k < 4; k++) {
      const nx = x + dir[k];
      const ny = y + dir[k + 1];

      if (nx < 0 || ny < 0 || nx >= n || ny >= n) continue;

      if (dist[nx][ny] !== -1) continue;

      dist[nx][ny] = dist[x][y] + 1;
      queue.push([nx, ny]);
    }
  }

  function canReach(limit) {
    if (dist[0][0] < limit || dist[n - 1][n - 1] < limit) return false;

    const vis = Array.from({ length: n }, () => Array(n).fill(false));
    const bfs = [[0, 0]];
    let idx = 0;

    vis[0][0] = true;

    while (idx < bfs.length) {
      const [x, y] = bfs[idx++];

      if (x === n - 1 && y === n - 1) return true;

      for (let k = 0; k < 4; k++) {
        const nx = x + dir[k];
        const ny = y + dir[k + 1];

        if (nx < 0 || ny < 0 || nx >= n || ny >= n) continue;

        if (vis[nx][ny] || dist[nx][ny] < limit) continue;

        vis[nx][ny] = true;
        bfs.push([nx, ny]);
      }
    }

    return false;
  }

  let left = 0;
  let right = 2 * n;
  let ans = 0;

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);

    if (canReach(mid)) {
      ans = mid;
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return ans;
};
