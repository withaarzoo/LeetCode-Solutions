function countIslands(grid: number[][], k: number): number {
  const m = grid.length,
    n = grid[0].length;
  let ans = 0;
  const dfs = (i: number, j: number): number => {
    if (i < 0 || i >= m || j < 0 || j >= n || grid[i][j] === 0) return 0;
    let sum = grid[i][j];
    grid[i][j] = 0;
    sum += dfs(i + 1, j);
    sum += dfs(i - 1, j);
    sum += dfs(i, j + 1);
    sum += dfs(i, j - 1);
    return sum;
  };
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (grid[i][j] > 0) {
        const s = dfs(i, j);
        if (s % k === 0) ans++;
      }
    }
  }
  return ans;
}
