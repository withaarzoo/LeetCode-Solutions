/**
 * @param {number} m
 * @param {number} n
 * @param {number[][]} guards
 * @param {number[][]} walls
 * @return {number}
 */
var countUnguarded = function (m, n, guards, walls) {
  // 0 = empty, 1 = guard, 2 = wall, 3 = guarded
  const grid = Array.from({ length: m }, () => Array(n).fill(0));
  for (const w of walls) grid[w[0]][w[1]] = 2;
  for (const g of guards) grid[g[0]][g[1]] = 1;

  const dirs = [
    [-1, 0],
    [1, 0],
    [0, -1],
    [0, 1],
  ];
  for (const g of guards) {
    const r = g[0],
      c = g[1];
    for (const d of dirs) {
      let nr = r + d[0],
        nc = c + d[1];
      while (nr >= 0 && nr < m && nc >= 0 && nc < n) {
        if (grid[nr][nc] === 2 || grid[nr][nc] === 1) break;
        if (grid[nr][nc] === 0) grid[nr][nc] = 3;
        nr += d[0];
        nc += d[1];
      }
    }
  }

  let ans = 0;
  for (let i = 0; i < m; ++i)
    for (let j = 0; j < n; ++j) if (grid[i][j] === 0) ans++;
  return ans;
};
