/**
 * @param {number[][]} heights
 * @return {number[][]}
 */
var pacificAtlantic = function (heights) {
  if (!heights || heights.length === 0) return [];
  const m = heights.length,
    n = heights[0].length;

  const pac = Array.from({ length: m }, () => Array(n).fill(false));
  const atl = Array.from({ length: m }, () => Array(n).fill(false));

  const bfs = (queue, visited) => {
    let head = 0;
    const dirs = [
      [1, 0],
      [-1, 0],
      [0, 1],
      [0, -1],
    ];
    while (head < queue.length) {
      const [r, c] = queue[head++];
      for (const [dr, dc] of dirs) {
        const nr = r + dr,
          nc = c + dc;
        if (nr < 0 || nr >= m || nc < 0 || nc >= n) continue;
        if (visited[nr][nc]) continue;
        if (heights[nr][nc] < heights[r][c]) continue;
        visited[nr][nc] = true;
        queue.push([nr, nc]);
      }
    }
  };

  // Pacific: top row and left column
  const q1 = [];
  for (let j = 0; j < n; ++j) {
    pac[0][j] = true;
    q1.push([0, j]);
  }
  for (let i = 1; i < m; ++i) {
    pac[i][0] = true;
    q1.push([i, 0]);
  }
  bfs(q1, pac);

  // Atlantic: bottom row and right column
  const q2 = [];
  for (let j = 0; j < n; ++j) {
    atl[m - 1][j] = true;
    q2.push([m - 1, j]);
  }
  for (let i = 0; i < m - 1; ++i) {
    atl[i][n - 1] = true;
    q2.push([i, n - 1]);
  }
  bfs(q2, atl);

  const res = [];
  for (let i = 0; i < m; ++i) {
    for (let j = 0; j < n; ++j) {
      if (pac[i][j] && atl[i][j]) res.push([i, j]);
    }
  }
  return res;
};
