/**
 * @param {number} n
 * @param {number} k
 * @param {number[][]} invocations
 * @return {number[]}
 */
var remainingMethods = function(n, k, invocations) {

    // Build adjacency list
    const graph = Array.from({ length: n }, () => []);

    for (const [u, v] of invocations) {
        graph[u].push(v);
    }

    // Marks suspicious methods
    const vis = new Array(n).fill(false);

    // DFS from method k
    function dfs(u) {
        vis[u] = true;

        for (const v of graph[u]) {
            if (!vis[v]) {
                dfs(v);
            }
        }
    }

    dfs(k);

    // Check whether removal is valid
    for (const [u, v] of invocations) {
        if (!vis[u] && vis[v]) {
            const ans = [];
            for (let i = 0; i < n; i++) {
                ans.push(i);
            }
            return ans;
        }
    }

    // Return remaining methods
    const ans = [];

    for (let i = 0; i < n; i++) {
        if (!vis[i]) {
            ans.push(i);
        }
    }

    return ans;
};