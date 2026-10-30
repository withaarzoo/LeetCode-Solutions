function shortestPath(
  n: number,
  edges: number[][],
  labels: string,
  k: number,
): number {
  const g: [number, number][][] = Array.from({ length: n }, () => []);
  for (const [u, v, w] of edges) {
    g[u].push([v, w]);
  }
  const INF = 1e9;
  const dist: number[][] = Array.from({ length: n }, () =>
    Array(k + 1).fill(INF),
  );
  class MinHeap {
    a: [number, number, number][] = [];
    push(x: [number, number, number]) {
      this.a.push(x);
      this._up(this.a.length - 1);
    }
    pop(): [number, number, number] {
      const top = this.a[0];
      const last = this.a.pop()!;
      if (this.a.length) {
        this.a[0] = last;
        this._down(0);
      }
      return top;
    }
    get size() {
      return this.a.length;
    }
    _up(i: number) {
      while (i > 0) {
        const p = (i - 1) >> 1;
        if (this.a[p][0] <= this.a[i][0]) break;
        [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
        i = p;
      }
    }
    _down(i: number) {
      const n = this.a.length;
      while (true) {
        let s = i;
        const l = i * 2 + 1,
          r = l + 1;
        if (l < n && this.a[l][0] < this.a[s][0]) s = l;
        if (r < n && this.a[r][0] < this.a[s][0]) s = r;
        if (s === i) break;
        [this.a[s], this.a[i]] = [this.a[i], this.a[s]];
        i = s;
      }
    }
  }
  const pq = new MinHeap();
  dist[0][1] = 0;
  pq.push([0, 0, 1]);
  while (pq.size) {
    const [d, u, cnt] = pq.pop();
    if (d > dist[u][cnt]) continue;
    if (u === n - 1) return d;
    for (const [v, w] of g[u]) {
      const ncnt = labels[u] === labels[v] ? cnt + 1 : 1;
      if (ncnt > k) continue;
      if (d + w < dist[v][ncnt]) {
        dist[v][ncnt] = d + w;
        pq.push([dist[v][ncnt], v, ncnt]);
      }
    }
  }
  return -1;
}
