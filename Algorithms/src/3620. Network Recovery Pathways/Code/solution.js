/**
 * @param {number[][]} edges
 * @param {boolean[]} online
 * @param {number} k
 * @return {number}
 */
var findMaxPathScore = function (edges, online, k) {
  const n = online.length;

  // Build graph
  const graph = Array.from({ length: n }, () => []);
  const indegree = Array(n).fill(0);

  for (const [u, v, w] of edges) {
    graph[u].push([v, w]);
    indegree[v]++;
  }

  // Topological order
  const queue = [];
  let head = 0;

  for (let i = 0; i < n; i++) if (indegree[i] === 0) queue.push(i);

  const topo = [];

  while (head < queue.length) {
    const u = queue[head++];
    topo.push(u);

    for (const [v] of graph[u]) {
      indegree[v]--;
      if (indegree[v] === 0) queue.push(v);
    }
  }

  const check = (limit) => {
    const INF = Number.MAX_SAFE_INTEGER;
    const dp = Array(n).fill(INF);

    // Minimum cost to reach node 0 is zero
    dp[0] = 0;

    for (const u of topo) {
      // Skip unreachable nodes
      if (dp[u] === INF) continue;

      // Skip offline intermediate nodes
      if (u !== 0 && u !== n - 1 && !online[u]) continue;

      for (const [v, w] of graph[u]) {
        // Edge is too small
        if (w < limit) continue;

        // Cannot enter offline intermediate node
        if (v !== n - 1 && !online[v]) continue;

        dp[v] = Math.min(dp[v], dp[u] + w);
      }
    }

    return dp[n - 1] <= k;
  };

  let left = 0;
  let right = 1000000000;
  let ans = -1;

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);

    if (check(mid)) {
      ans = mid;
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return ans;
};
