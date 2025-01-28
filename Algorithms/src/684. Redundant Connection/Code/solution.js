var findRedundantConnection = function (edges) {
  let parent = Array(edges.length + 1)
    .fill(0)
    .map((_, i) => i);
  let rank = Array(edges.length + 1).fill(0);

  function find(node) {
    if (parent[node] !== node) parent[node] = find(parent[node]); // Path compression
    return parent[node];
  }

  function union(u, v) {
    let rootU = find(u),
      rootV = find(v);
    if (rootU === rootV) return false;
    if (rank[rootU] > rank[rootV]) parent[rootV] = rootU;
    else if (rank[rootU] < rank[rootV]) parent[rootU] = rootV;
    else {
      parent[rootV] = rootU;
      rank[rootU]++;
    }
    return true;
  }

  for (let [u, v] of edges) {
    if (!union(u, v)) return [u, v];
  }
  return [];
};
