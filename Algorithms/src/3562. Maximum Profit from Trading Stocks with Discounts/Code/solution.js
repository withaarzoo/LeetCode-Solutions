var maxProfit = function (n, present, future, hierarchy, budget) {
  const tree = Array.from({ length: n }, () => []);
  for (const [u, v] of hierarchy) tree[u - 1].push(v - 1);

  const dp = Array.from({ length: n }, () =>
    Array.from({ length: 2 }, () => Array(budget + 1).fill(0))
  );

  const merge = (A, B) => {
    const C = Array(budget + 1).fill(-1e9);
    for (let i = 0; i <= budget; i++) {
      if (A[i] < 0) continue;
      for (let j = 0; i + j <= budget; j++) {
        C[i + j] = Math.max(C[i + j], A[i] + B[j]);
      }
    }
    return C;
  };

  const dfs = (u) => {
    for (const v of tree[u]) dfs(v);

    for (let parentBought = 0; parentBought <= 1; parentBought++) {
      const price = parentBought ? Math.floor(present[u] / 2) : present[u];
      const profit = future[u] - price;

      let skip = Array(budget + 1).fill(0);
      for (const v of tree[u]) skip = merge(skip, dp[v][0]);

      let take = Array(budget + 1).fill(-1e9);
      if (price <= budget) {
        let base = Array(budget + 1).fill(0);
        for (const v of tree[u]) base = merge(base, dp[v][1]);
        for (let b = price; b <= budget; b++)
          take[b] = base[b - price] + profit;
      }

      for (let b = 0; b <= budget; b++)
        dp[u][parentBought][b] = Math.max(skip[b], take[b]);
    }
  };

  dfs(0);
  return Math.max(...dp[0][0]);
};
