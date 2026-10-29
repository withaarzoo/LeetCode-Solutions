/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {string} s
 * @param {string[]} queries
 * @return {boolean[]}
 */
var palindromePath = function (n, edges, s, queries) {
  const g = Array.from({ length: n }, () => []);
  for (const [u, v] of edges) {
    g[u].push(v);
    g[v].push(u);
  }
  const LOG = 16;
  const up = Array.from({ length: LOG }, () => new Array(n).fill(-1));
  const depth = new Array(n).fill(0);
  const inn = new Array(n);
  const out = new Array(n);
  const parent = new Array(n).fill(-1);
  const mask = new Array(n);
  for (let i = 0; i < n; i++) mask[i] = 1 << (s.charCodeAt(i) - 97);
  let timer = 1;
  function dfs(u, p) {
    inn[u] = timer++;
    parent[u] = p;
    up[0][u] = p;
    for (let k = 1; k < LOG; k++) {
      if (up[k - 1][u] !== -1) up[k][u] = up[k - 1][up[k - 1][u]];
    }
    for (const v of g[u])
      if (v !== p) {
        depth[v] = depth[u] + 1;
        dfs(v, u);
      }
    out[u] = timer;
  }
  dfs(0, -1);
  const bit = new Array(n + 2).fill(0);
  function upd(i, val) {
    for (; i <= n; i += i & -i) bit[i] ^= val;
  }
  function qry(i) {
    let res = 0;
    for (; i > 0; i -= i & -i) res ^= bit[i];
    return res;
  }
  function rangeXor(l, r, val) {
    if (l > r) return;
    upd(l, val);
    upd(r + 1, val);
  }
  for (let i = 0; i < n; i++) rangeXor(inn[i], out[i] - 1, mask[i]);
  function getLca(u, v) {
    if (depth[u] < depth[v]) [u, v] = [v, u];
    let diff = depth[u] - depth[v];
    for (let k = 0; k < LOG; k++) if ((diff >> k) & 1) u = up[k][u];
    if (u === v) return u;
    for (let k = LOG - 1; k >= 0; k--) {
      if (up[k][u] !== up[k][v]) {
        u = up[k][u];
        v = up[k][v];
      }
    }
    return parent[u];
  }
  function pathXor(u, v) {
    const l = getLca(u, v);
    return qry(inn[u]) ^ qry(inn[v]) ^ mask[l];
  }
  const ans = [];
  for (const q of queries) {
    if (q[0] === "u") {
      let u = 0,
        i = 7;
      while (q[i] !== " ") u = u * 10 + (q.charCodeAt(i++) - 48);
      const c = q[q.length - 1];
      const nm = 1 << (c.charCodeAt(0) - 97);
      const delta = mask[u] ^ nm;
      mask[u] = nm;
      rangeXor(inn[u], out[u] - 1, delta);
    } else {
      let u = 0,
        v = 0,
        i = 6;
      while (q[i] !== " ") u = u * 10 + (q.charCodeAt(i++) - 48);
      i++;
      while (i < q.length) v = v * 10 + (q.charCodeAt(i++) - 48);
      const m = pathXor(u, v);
      ans.push(m.toString(2).split("1").length - 1 <= 1);
    }
  }
  return ans;
};
