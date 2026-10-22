/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function(nums) {
    const n = nums.length;
    const starts = [];
    let inc = true, dec = true;
    for (let i = 0; i < n; i++) {
        if (nums[(i + 1) % n] !== (nums[i] + 1) % n) inc = false;
        if (nums[(i + 1) % n] !== (nums[i] - 1 + n) % n) dec = false;
    }
    if (inc) starts.push([0, nums[0]]);
    if (dec) starts.push([1, (n - 1 - nums[0] + n) % n]);
    if (starts.length === 0) return -1;

    const dist = Array.from({length: 2}, () => Array(n).fill(-1));
    const q = [];
    for (const [t, s] of starts) {
        dist[t][s] = 0;
        q.push([t, s]);
    }
    let head = 0;
    while (head < q.length) {
        const [t, s] = q[head++];
        if (t === 0 && s === 0) return dist[t][s];
        const d = dist[t][s];
        let nt = t, ns = (s + 1) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
        nt = 1 - t;
        ns = (n - s) % n;
        if (dist[nt][ns] === -1) {
            dist[nt][ns] = d + 1;
            q.push([nt, ns]);
        }
    }
    return -1;
};