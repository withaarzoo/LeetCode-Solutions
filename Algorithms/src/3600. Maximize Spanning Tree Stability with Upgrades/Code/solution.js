var maxStability = function (n, edges, k) {
  const parent = Array(n)
    .fill(0)
    .map((_, i) => i);
  const rank = Array(n).fill(0);

  function find(x) {
    if (parent[x] !== x) parent[x] = find(parent[x]);
    return parent[x];
  }

  function union(a, b) {
    a = find(a);
    b = find(b);

    if (a === b) return false;

    if (rank[a] < rank[b]) [a, b] = [b, a];

    parent[b] = a;

    if (rank[a] === rank[b]) rank[a]++;

    return true;
  }

  let comp = n;
  let mandatoryMin = Infinity;

  let optional = [];

  for (const e of edges) {
    if (e[3] === 1) {
      if (!union(e[0], e[1])) return -1;
      comp--;
      mandatoryMin = Math.min(mandatoryMin, e[2]);
    } else optional.push(e);
  }

  optional.sort((a, b) => b[2] - a[2]);

  let used = [];

  for (const e of optional) {
    if (union(e[0], e[1])) {
      used.push(e[2]);
      comp--;
      if (comp === 1) break;
    }
  }

  if (comp > 1) return -1;

  used.sort((a, b) => a - b);

  let ans = mandatoryMin;

  for (let w of used) {
    let val = w;

    if (k > 0) {
      val *= 2;
      k--;
    }

    if (ans === Infinity) ans = val;
    else ans = Math.min(ans, val);
  }

  return ans;
};
