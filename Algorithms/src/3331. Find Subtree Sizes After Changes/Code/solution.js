/**
 * @param {number[]} parent
 * @param {string} s
 * @return {number[]}
 */
var findSubtreeSizes = function(parent, s) {
    const n = parent.length;
    const g = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) g[parent[i]].push(i);
    const np = parent.slice();
    const last = new Array(26).fill(-1);
    const dfs = (u) => {
        const c = s.charCodeAt(u) - 97;
        const prev = last[c];
        if (prev !== -1) np[u] = prev;
        last[c] = u;
        for (const v of g[u]) dfs(v);
        last[c] = prev;
    };
    dfs(0);
    const ng = Array.from({length: n}, () => []);
    for (let i = 1; i < n; i++) ng[np[i]].push(i);
    const ans = new Array(n);
    const size = (u) => {
        let sz = 1;
        for (const v of ng[u]) sz += size(v);
        return ans[u] = sz;
    };
    size(0);
    return ans;
};