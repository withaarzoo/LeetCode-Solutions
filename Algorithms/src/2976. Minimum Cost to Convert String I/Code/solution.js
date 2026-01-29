var minimumCost = function (source, target, original, changed, cost) {
  const INF = 1e18;
  const dist = Array.from({ length: 26 }, () => Array(26).fill(INF));

  for (let i = 0; i < 26; i++) dist[i][i] = 0;

  for (let i = 0; i < original.length; i++) {
    const u = original[i].charCodeAt(0) - 97;
    const v = changed[i].charCodeAt(0) - 97;
    dist[u][v] = Math.min(dist[u][v], cost[i]);
  }

  // Floyd-Warshall
  for (let k = 0; k < 26; k++) {
    for (let i = 0; i < 26; i++) {
      for (let j = 0; j < 26; j++) {
        dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);
      }
    }
  }

  let ans = 0;
  for (let i = 0; i < source.length; i++) {
    const s = source.charCodeAt(i) - 97;
    const t = target.charCodeAt(i) - 97;
    if (dist[s][t] === INF) return -1;
    ans += dist[s][t];
  }

  return ans;
};
