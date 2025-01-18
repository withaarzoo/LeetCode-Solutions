var minCost = function (grid) {
  const m = grid.length,
    n = grid[0].length;
  const directions = [
    [0, 1],
    [0, -1],
    [1, 0],
    [-1, 0],
  ];
  const cost = Array.from({ length: m }, () => Array(n).fill(Infinity));
  const deque = [];
  deque.push([0, 0]);
  cost[0][0] = 0;

  while (deque.length) {
    const [x, y] = deque.shift();
    for (let i = 0; i < 4; i++) {
      const nx = x + directions[i][0];
      const ny = y + directions[i][1];
      const newCost = cost[x][y] + (grid[x][y] !== i + 1 ? 1 : 0);

      if (nx >= 0 && ny >= 0 && nx < m && ny < n && newCost < cost[nx][ny]) {
        cost[nx][ny] = newCost;
        if (grid[x][y] === i + 1) deque.unshift([nx, ny]);
        else deque.push([nx, ny]);
      }
    }
  }
  return cost[m - 1][n - 1];
};
