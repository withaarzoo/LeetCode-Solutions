var minimumCost = function (source, target, original, changed, cost) {
  const INF = BigInt("18446744073709551615");
  const id = new Map();
  const lens = new Set();
  let sz = 0;

  const dist = Array.from({ length: 201 }, () => Array(201).fill(INF));

  for (let i = 0; i < original.length; i++) {
    if (!id.has(original[i])) {
      id.set(original[i], sz++);
      lens.add(original[i].length);
    }
    if (!id.has(changed[i])) {
      id.set(changed[i], sz++);
    }
    const u = id.get(original[i]);
    const v = id.get(changed[i]);
    dist[u][v] = BigInt(Math.min(Number(dist[u][v]), cost[i]));
  }

  for (let i = 0; i < sz; i++) dist[i][i] = 0n;

  for (let k = 0; k < sz; k++)
    for (let i = 0; i < sz; i++)
      if (dist[i][k] !== INF)
        for (let j = 0; j < sz; j++)
          if (dist[k][j] !== INF)
            dist[i][j] = BigInt(
              Math.min(Number(dist[i][j]), Number(dist[i][k] + dist[k][j])),
            );

  const n = source.length;
  const dp = Array(n + 1).fill(INF);
  dp[0] = 0n;

  for (let i = 0; i < n; i++) {
    if (dp[i] === INF) continue;

    if (source[i] === target[i])
      dp[i + 1] = dp[i + 1] < dp[i] ? dp[i + 1] : dp[i];

    for (const L of lens) {
      if (i + L > n) continue;
      const s = source.substr(i, L);
      const t = target.substr(i, L);

      if (id.has(s) && id.has(t)) {
        const d = dist[id.get(s)][id.get(t)];
        if (d !== INF)
          dp[i + L] = dp[i + L] < dp[i] + d ? dp[i + L] : dp[i] + d;
      }
    }
  }

  return dp[n] === INF ? -1 : Number(dp[n]);
};
