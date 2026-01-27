var minCost = function (n, edges) {
  const graph = Array.from({ length: n }, () => []);

  for (const [u, v, w] of edges) {
    graph[u].push([v, w]);
    graph[v].push([u, 2 * w]);
  }

  const dist = Array(n).fill(Infinity);
  dist[0] = 0;

  const pq = [[0, 0]]; // [cost, node]

  while (pq.length) {
    pq.sort((a, b) => a[0] - b[0]);
    const [cost, node] = pq.shift();

    if (cost > dist[node]) continue;

    for (const [next, w] of graph[node]) {
      if (dist[next] > cost + w) {
        dist[next] = cost + w;
        pq.push([dist[next], next]);
      }
    }
  }

  return dist[n - 1] === Infinity ? -1 : dist[n - 1];
};
