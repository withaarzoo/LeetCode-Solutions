/**
 * @param {number[]} source
 * @param {number[]} target
 * @param {number[][]} allowedSwaps
 * @return {number}
 */
var minimumHammingDistance = function (source, target, allowedSwaps) {
  const n = source.length;

  const parent = Array.from({ length: n }, (_, i) => i);
  const rank = Array(n).fill(0);

  function find(x) {
    if (parent[x] !== x) {
      parent[x] = find(parent[x]);
    }
    return parent[x];
  }

  function union(a, b) {
    let pa = find(a);
    let pb = find(b);

    if (pa === pb) return;

    if (rank[pa] < rank[pb]) {
      parent[pa] = pb;
    } else if (rank[pb] < rank[pa]) {
      parent[pb] = pa;
    } else {
      parent[pb] = pa;
      rank[pa]++;
    }
  }

  // Build connected components
  for (const [u, v] of allowedSwaps) {
    union(u, v);
  }

  // Group indices by root
  const groups = new Map();

  for (let i = 0; i < n; i++) {
    const root = find(i);

    if (!groups.has(root)) {
      groups.set(root, []);
    }

    groups.get(root).push(i);
  }

  let answer = 0;

  // Process each component
  for (const indices of groups.values()) {
    const freq = new Map();

    // Count source values
    for (const idx of indices) {
      freq.set(source[idx], (freq.get(source[idx]) || 0) + 1);
    }

    // Match target values
    for (const idx of indices) {
      const val = target[idx];

      if ((freq.get(val) || 0) > 0) {
        freq.set(val, freq.get(val) - 1);
      } else {
        answer++;
      }
    }
  }

  return answer;
};
