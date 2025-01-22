var highestPeak = function (isWater) {
  const m = isWater.length,
    n = isWater[0].length;
  const height = Array.from({ length: m }, () => Array(n).fill(-1));
  const queue = [];

  // Initialize water cells
  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (isWater[i][j] === 1) {
        height[i][j] = 0;
        queue.push([i, j]);
      }
    }
  }

  // Directions for BFS
  const directions = [
    [0, 1],
    [0, -1],
    [1, 0],
    [-1, 0],
  ];

  // BFS
  while (queue.length > 0) {
    const [x, y] = queue.shift();
    for (const [dx, dy] of directions) {
      const nx = x + dx,
        ny = y + dy;
      if (nx >= 0 && ny >= 0 && nx < m && ny < n && height[nx][ny] === -1) {
        height[nx][ny] = height[x][y] + 1;
        queue.push([nx, ny]);
      }
    }
  }

  return height;
};
