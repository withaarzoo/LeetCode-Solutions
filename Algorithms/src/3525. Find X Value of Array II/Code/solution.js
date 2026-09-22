/**
 * @param {number[]} nums
 * @param {number} k
 * @param {number[][]} queries
 * @return {number[]}
 */
var resultArray = function(nums, k, queries) {
    class Node {
        constructor() {
            this.prod = 1;
            this.freq = new Array(k).fill(0);
        }
    }
    
    const n = nums.length;
    const tree = new Array(4 * n);
    for (let i = 0; i < tree.length; i++) tree[i] = new Node();
    
    function merge(L, R) {
        const res = new Node();
        res.prod = (L.prod * R.prod) % k;
        for (let i = 0; i < k; i++) res.freq[i] = L.freq[i];
        for (let r = 0; r < k; r++) {
            if (R.freq[r]) {
                const nr = (L.prod * r) % k;
                res.freq[nr] += R.freq[r];
            }
        }
        return res;
    }
    
    function build(v, tl, tr) {
        if (tl === tr) {
            tree[v].prod = nums[tl] % k;
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        build(v * 2, tl, tm);
        build(v * 2 + 1, tm + 1, tr);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    function update(v, tl, tr, pos, val) {
        if (tl === tr) {
            tree[v].prod = val % k;
            tree[v].freq.fill(0);
            tree[v].freq[tree[v].prod] = 1;
            return;
        }
        const tm = (tl + tr) >> 1;
        if (pos <= tm) update(v * 2, tl, tm, pos, val);
        else update(v * 2 + 1, tm + 1, tr, pos, val);
        tree[v] = merge(tree[v * 2], tree[v * 2 + 1]);
    }
    
    function query(v, tl, tr, l, r) {
        if (l > r) return new Node();
        if (l === tl && r === tr) return tree[v];
        const tm = (tl + tr) >> 1;
        return merge(query(v * 2, tl, tm, l, Math.min(r, tm)),
                     query(v * 2 + 1, tm + 1, tr, Math.max(l, tm + 1), r));
    }
    
    build(1, 0, n - 1);
    const ans = [];
    for (const q of queries) {
        const idx = q[0], val = q[1], start = q[2], x = q[3];
        update(1, 0, n - 1, idx, val);
        const res = query(1, 0, n - 1, start, n - 1);
        ans.push(res.freq[x]);
    }
    return ans;
};