/**
 * @param {number[][]} grid
 * @return {boolean}
 */
var hasValidPath = function (grid) {
  const m = grid.length;
  const n = grid[0].length;

  // Directions: left, right, up, down
  const dirs = [
    [0, -1], // left
    [0, 1], // right
    [-1, 0], // up
    [1, 0], // down
  ];

  // For each street type, which directions it supports.
  // 0 = left, 1 = right, 2 = up, 3 = down
  const streetDirs = [
    [], // dummy
    [0, 1], // 1: left-right
    [2, 3], // 2: up-down
    [0, 3], // 3: left-down
    [1, 3], // 4: right-down
    [0, 2], // 5: left-up
    [1, 2], // 6: right-up
  ];

  const opposite = [1, 0, 3, 2];

  const visited = Array.from({ length: m }, () => Array(n).fill(false));
  const queue = [[0, 0]];
  let head = 0;
  visited[0][0] = true;

  while (head < queue.length) {
    const [r, c] = queue[head++];

    if (r === m - 1 && c === n - 1) {
      return true;
    }

    const type = grid[r][c];

    for (const d of streetDirs[type]) {
      const nr = r + dirs[d][0];
      const nc = c + dirs[d][1];

      if (nr < 0 || nr >= m || nc < 0 || nc >= n || visited[nr][nc]) {
        continue;
      }

      const nextType = grid[nr][nc];

      let ok = false;
      for (const nd of streetDirs[nextType]) {
        if (nd === opposite[d]) {
          ok = true;
          break;
        }
      }

      if (ok) {
        visited[nr][nc] = true;
        queue.push([nr, nc]);
      }
    }
  }

  return false;
};
