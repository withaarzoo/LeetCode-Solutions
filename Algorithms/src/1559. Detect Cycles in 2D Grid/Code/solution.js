/**
 * @param {character[][]} grid
 * @return {boolean}
 */
var containsCycle = function (grid) {
  const m = grid.length;
  const n = grid[0].length;
  const visited = Array.from({ length: m }, () => Array(n).fill(false));

  const dr = [1, -1, 0, 0];
  const dc = [0, 0, 1, -1];

  for (let r = 0; r < m; r++) {
    for (let c = 0; c < n; c++) {
      if (visited[r][c]) continue;

      const stack = [[r, c, -1, -1]];
      visited[r][c] = true;

      while (stack.length > 0) {
        const [cr, cc, pr, pc] = stack.pop();

        for (let k = 0; k < 4; k++) {
          const nr = cr + dr[k];
          const nc = cc + dc[k];

          if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
          if (grid[nr][nc] !== grid[cr][cc]) continue;
          if (nr === pr && nc === pc) continue;

          if (visited[nr][nc]) return true;

          visited[nr][nc] = true;
          stack.push([nr, nc, cr, cc]);
        }
      }
    }
  }

  return false;
};
