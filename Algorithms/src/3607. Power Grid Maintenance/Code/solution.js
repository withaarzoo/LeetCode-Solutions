/**
 * @param {number} c
 * @param {number[][]} connections
 * @param {number[][]} queries
 * @return {number[]}
 */
var processQueries = function (c, connections, queries) {
  // DSU (Union-Find)
  const p = Array(c + 1)
    .fill(0)
    .map((_, i) => i);
  const sz = Array(c + 1).fill(1);
  const find = (x) => (p[x] === x ? x : (p[x] = find(p[x])));
  const unite = (a, b) => {
    a = find(a);
    b = find(b);
    if (a === b) return;
    if (sz[a] < sz[b]) [a, b] = [b, a];
    p[b] = a;
    sz[a] += sz[b];
  };
  for (const [u, v] of connections) unite(u, v);

  // Min-heap implementation
  class MinHeap {
    constructor() {
      this.a = [];
    }
    size() {
      return this.a.length;
    }
    peek() {
      return this.a[0];
    }
    push(x) {
      const a = this.a;
      a.push(x);
      let i = a.length - 1;
      while (i > 0) {
        let p = (i - 1) >> 1;
        if (a[p] <= a[i]) break;
        [a[p], a[i]] = [a[i], a[p]];
        i = p;
      }
    }
    pop() {
      const a = this.a;
      if (a.length === 0) return undefined;
      const top = a[0],
        last = a.pop();
      if (a.length) {
        a[0] = last;
        let i = 0;
        while (true) {
          let l = i * 2 + 1,
            r = l + 1,
            m = i;
          if (l < a.length && a[l] < a[m]) m = l;
          if (r < a.length && a[r] < a[m]) m = r;
          if (m === i) break;
          [a[i], a[m]] = [a[m], a[i]];
          i = m;
        }
      }
      return top;
    }
  }

  // root -> heap
  const heap = new Map();
  for (let i = 1; i <= c; i++) {
    const r = find(i);
    if (!heap.has(r)) heap.set(r, new MinHeap());
    heap.get(r).push(i);
  }

  const offline = Array(c + 1).fill(false);
  const ans = [];

  for (const [t, x] of queries) {
    if (t === 2) {
      offline[x] = true;
    } else {
      if (!offline[x]) {
        ans.push(x);
      } else {
        const r = find(x);
        const pq = heap.get(r);
        if (!pq) {
          ans.push(-1);
          continue;
        }
        while (pq.size() && offline[pq.peek()]) pq.pop();
        ans.push(pq.size() ? pq.peek() : -1);
      }
    }
  }
  return ans;
};
