/**
 * @param {number} k
 * @param {number[]} arrival
 * @param {number[]} load
 * @return {number[]}
 */
var busiestServers = function(k, arrival, load) {
    class MinHeap {
        constructor() { this.a = []; }
        size() { return this.a.length; }
        peek() { return this.a[0]; }
        push(x) {
            this.a.push(x);
            let i = this.a.length - 1;
            while (i > 0) {
                let p = (i - 1) >> 1;
                if (this.a[p] <= this.a[i]) break;
                [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
                i = p;
            }
        }
        pop() {
            const top = this.a[0];
            const last = this.a.pop();
            if (this.a.length === 0) return top;
            this.a[0] = last;
            let i = 0;
            while (true) {
                let l = i * 2 + 1, r = l + 1, smallest = i;
                if (l < this.a.length && this.a[l] < this.a[smallest]) smallest = l;
                if (r < this.a.length && this.a[r] < this.a[smallest]) smallest = r;
                if (smallest === i) break;
                [this.a[i], this.a[smallest]] = [this.a[smallest], this.a[i]];
                i = smallest;
            }
            return top;
        }
    }
    const available = new MinHeap();
    for (let i = 0; i < k; i++) available.push(i);
    const busy = new MinHeap(); // stores endTime * 1e6 + server ( suffices because endTime <= 2e9 )
    const cnt = new Array(k).fill(0);
    const n = arrival.length;
    for (let i = 0; i < n; i++) {
        const t = arrival[i];
        while (busy.size() && Math.floor(busy.peek() / 1000000) <= t) {
            const x = busy.pop() % 1000000;
            available.push(i + ((x - i) % k + k) % k);
        }
        if (available.size() === 0) continue;
        const j = available.pop() % k;
        cnt[j]++;
        busy.push((t + load[i]) * 1000000 + j);
    }
    let mx = 0;
    for (const c of cnt) if (c > mx) mx = c;
    const ans = [];
    for (let i = 0; i < k; i++) if (cnt[i] === mx) ans.push(i);
    return ans;
};