/**
 * @param {number} n
 * @param {number[][]} logs
 * @param {number} x
 * @param {number[]} queries
 * @return {number[]}
 */
var countServers = function(n, logs, x, queries) {
    logs.sort((a, b) => a[1] - b[1]);
    const m = queries.length;
    const q = queries.map((t, i) => [t, i]);
    q.sort((a, b) => a[0] - b[0]);
    const ans = new Array(m);
    const freq = new Map();
    let active = 0;
    let left = 0, right = 0;
    const L = logs.length;
    for (const [time, idx] of q) {
        while (right < L && logs[right][1] <= time) {
            const s = logs[right][0];
            freq.set(s, (freq.get(s) || 0) + 1);
            if (freq.get(s) === 1) active++;
            right++;
        }
        while (left < right && logs[left][1] < time - x) {
            const s = logs[left][0];
            freq.set(s, freq.get(s) - 1);
            if (freq.get(s) === 0) active--;
            left++;
        }
        ans[idx] = n - active;
    }
    return ans;
};