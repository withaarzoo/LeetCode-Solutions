var findMaxFish = function (grid) {
  const m = grid.length,
    n = grid[0].length;
  const visited = Array.from({ length: m }, () => Array(n).fill(false));
  let maxFish = 0;

  const dfs = (r, c) => {
    if (r < 0 || c < 0 || r >= m || c >= n || visited[r][c] || grid[r][c] === 0)
      return 0;
    visited[r][c] = true;
    let fish = grid[r][c];
    fish += dfs(r + 1, c);
    fish += dfs(r - 1, c);
    fish += dfs(r, c + 1);
    fish += dfs(r, c - 1);
    return fish;
  };

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (!visited[i][j] && grid[i][j] > 0) {
        maxFish = Math.max(maxFish, dfs(i, j));
      }
    }
  }
  return maxFish;
};
