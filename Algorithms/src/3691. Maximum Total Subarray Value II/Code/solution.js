/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxTotalValue = function (nums, k) {
  const n = nums.length;

  const lg = new Array(n + 1).fill(0);

  for (let i = 2; i <= n; i++) {
    lg[i] = lg[i >> 1] + 1;
  }

  const K = lg[n] + 1;

  const mx = Array.from({ length: K }, () => Array(n).fill(0));
  const mn = Array.from({ length: K }, () => Array(n).fill(0));

  for (let i = 0; i < n; i++) {
    mx[0][i] = nums[i];
    mn[0][i] = nums[i];
  }

  for (let j = 1; j < K; j++) {
    for (let i = 0; i + (1 << j) <= n; i++) {
      mx[j][i] = Math.max(mx[j - 1][i], mx[j - 1][i + (1 << (j - 1))]);

      mn[j][i] = Math.min(mn[j - 1][i], mn[j - 1][i + (1 << (j - 1))]);
    }
  }

  const getValue = (l, r) => {
    const len = r - l + 1;
    const p = lg[len];

    const mxVal = Math.max(mx[p][l], mx[p][r - (1 << p) + 1]);

    const mnVal = Math.min(mn[p][l], mn[p][r - (1 << p) + 1]);

    return mxVal - mnVal;
  };

  class MaxHeap {
    constructor() {
      this.heap = [];
    }

    push(x) {
      this.heap.push(x);

      let i = this.heap.length - 1;

      while (i > 0) {
        let p = (i - 1) >> 1;

        if (this.heap[p][0] >= this.heap[i][0]) break;

        [this.heap[p], this.heap[i]] = [this.heap[i], this.heap[p]];

        i = p;
      }
    }

    pop() {
      const top = this.heap[0];
      const last = this.heap.pop();

      if (this.heap.length) {
        this.heap[0] = last;

        let i = 0;

        while (true) {
          let largest = i;
          let l = i * 2 + 1;
          let r = i * 2 + 2;

          if (l < this.heap.length && this.heap[l][0] > this.heap[largest][0]) {
            largest = l;
          }

          if (r < this.heap.length && this.heap[r][0] > this.heap[largest][0]) {
            largest = r;
          }

          if (largest === i) break;

          [this.heap[i], this.heap[largest]] = [
            this.heap[largest],
            this.heap[i],
          ];

          i = largest;
        }
      }

      return top;
    }
  }

  const pq = new MaxHeap();

  for (let l = 0; l < n; l++) {
    pq.push([getValue(l, n - 1), l, n - 1]);
  }

  let ans = 0;

  while (k--) {
    const [val, l, r] = pq.pop();

    ans += val;

    if (r > l) {
      pq.push([getValue(l, r - 1), l, r - 1]);
    }
  }

  return ans;
};
