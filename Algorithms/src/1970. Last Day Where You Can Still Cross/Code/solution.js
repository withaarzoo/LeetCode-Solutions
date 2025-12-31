var latestDayToCross = function (row, col, cells) {
  const n = row * col;
  const top = n,
    bottom = n + 1;

  const parent = Array(n + 2)
    .fill(0)
    .map((_, i) => i);
  const rank = Array(n + 2).fill(0);
  const grid = Array.from({ length: row }, () => Array(col).fill(false));

  const find = (x) => {
    if (parent[x] !== x) parent[x] = find(parent[x]);
    return parent[x];
  };

  const union = (a, b) => {
    a = find(a);
    b = find(b);
    if (a === b) return;
    if (rank[a] < rank[b]) parent[a] = b;
    else {
      parent[b] = a;
      if (rank[a] === rank[b]) rank[a]++;
    }
  };

  const dr = [1, -1, 0, 0];
  const dc = [0, 0, 1, -1];

  for (let d = n - 1; d >= 0; d--) {
    const r = cells[d][0] - 1;
    const c = cells[d][1] - 1;
    grid[r][c] = true;
    const id = r * col + c;

    if (r === 0) union(id, top);
    if (r === row - 1) union(id, bottom);

    for (let k = 0; k < 4; k++) {
      const nr = r + dr[k];
      const nc = c + dc[k];
      if (nr >= 0 && nr < row && nc >= 0 && nc < col && grid[nr][nc]) {
        union(id, nr * col + nc);
      }
    }

    if (find(top) === find(bottom)) return d;
  }
  return 0;
};
