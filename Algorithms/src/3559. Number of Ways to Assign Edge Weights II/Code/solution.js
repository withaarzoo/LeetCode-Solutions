/**
 * @param {number[][]} edges
 * @param {number[][]} queries
 * @return {number[]}
 */
var assignEdgeWeights = function (edges, queries) {
  const MOD = 1000000007;
  const n = edges.length + 1;

  let LOG = 1;
  while (1 << LOG <= n) LOG++;

  // Adjacency list
  const graph = Array.from({ length: n + 1 }, () => []);

  for (const [u, v] of edges) {
    graph[u].push(v);
    graph[v].push(u);
  }

  const depth = Array(n + 1).fill(0);
  const up = Array.from({ length: n + 1 }, () => Array(LOG).fill(1));

  // DFS
  const dfs = (node, parent) => {
    up[node][0] = parent;

    for (let j = 1; j < LOG; j++) {
      up[node][j] = up[up[node][j - 1]][j - 1];
    }

    for (const next of graph[node]) {
      if (next === parent) continue;

      depth[next] = depth[node] + 1;
      dfs(next, node);
    }
  };

  dfs(1, 1);

  // LCA
  const lca = (a, b) => {
    if (depth[a] < depth[b]) {
      [a, b] = [b, a];
    }

    let diff = depth[a] - depth[b];

    for (let j = LOG - 1; j >= 0; j--) {
      if ((diff >> j) & 1) {
        a = up[a][j];
      }
    }

    if (a === b) return a;

    for (let j = LOG - 1; j >= 0; j--) {
      if (up[a][j] !== up[b][j]) {
        a = up[a][j];
        b = up[b][j];
      }
    }

    return up[a][0];
  };

  // Powers of two modulo MOD
  const pow2 = Array(n + 1).fill(1);

  for (let i = 1; i <= n; i++) {
    pow2[i] = (pow2[i - 1] * 2) % MOD;
  }

  const ans = [];

  for (const [u, v] of queries) {
    const ancestor = lca(u, v);

    const dist = depth[u] + depth[v] - 2 * depth[ancestor];

    ans.push(dist === 0 ? 0 : pow2[dist - 1]);
  }

  return ans;
};
