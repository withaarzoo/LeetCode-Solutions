var trapRainWater = function (heightMap) {
  const m = heightMap.length,
    n = heightMap[0].length;
  if (m < 3 || n < 3) return 0;

  const pq = new MinPriorityQueue({ priority: (cell) => cell.height });
  const visited = Array.from({ length: m }, () => Array(n).fill(false));

  // Add all boundary cells
  for (let i = 0; i < m; i++) {
    pq.enqueue({ height: heightMap[i][0], x: i, y: 0 });
    pq.enqueue({ height: heightMap[i][n - 1], x: i, y: n - 1 });
    visited[i][0] = visited[i][n - 1] = true;
  }
  for (let j = 0; j < n; j++) {
    pq.enqueue({ height: heightMap[0][j], x: 0, y: j });
    pq.enqueue({ height: heightMap[m - 1][j], x: m - 1, y: j });
    visited[0][j] = visited[m - 1][j] = true;
  }

  let result = 0;
  const directions = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ];

  while (!pq.isEmpty()) {
    const { height, x, y } = pq.dequeue();

    for (const [dx, dy] of directions) {
      const nx = x + dx,
        ny = y + dy;
      if (nx >= 0 && ny >= 0 && nx < m && ny < n && !visited[nx][ny]) {
        result += Math.max(0, height - heightMap[nx][ny]);
        pq.enqueue({
          height: Math.max(height, heightMap[nx][ny]),
          x: nx,
          y: ny,
        });
        visited[nx][ny] = true;
      }
    }
  }

  return result;
};
