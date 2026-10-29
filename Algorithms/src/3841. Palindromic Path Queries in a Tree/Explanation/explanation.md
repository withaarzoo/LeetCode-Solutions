# Palindromic Path Queries in a Tree – LeetCode 3841 Solution

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
  - [C++](#c)
  - [Java](#java)
  - [JavaScript](#javascript)
  - [TypeScript](#typescript)
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-typescript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

## Problem Summary

You are given an undirected tree with n nodes. Each node has a lowercase English letter. You also receive a list of queries.  

Some queries ask you to change the letter on a node. Other queries ask whether the unique path between two nodes can be rearranged into a palindrome.  

A string can form a palindrome after rearranging if and only if at most one character appears an odd number of times.  

Your task is to process all queries online and return a boolean answer for every path-query.  

This is a classic hard LeetCode tree problem that mixes path queries, updates, and palindrome checks.

## Constraints

- 1 ≤ n = s.length ≤ 5 × 10⁴
- edges.length == n − 1
- 0 ≤ uᵢ, vᵢ ≤ n − 1
- s consists of lowercase English letters
- The given edges form a valid tree
- 1 ≤ queries.length ≤ 5 × 10⁴
- Each query is either “update u c” or “query u v”
- 0 ≤ u, v ≤ n − 1
- c is a lowercase English letter

## Intuition

The first useful observation is that we only care about the parity of each letter on a path. With 26 letters we can pack those parities into a single 26-bit integer and combine them with XOR.  

A path can become a palindrome exactly when the final XOR mask has at most one bit set.  

Because the tree is static but we have both updates and path queries, we need a fast way to maintain the XOR of every root-to-node path and to answer “what is the XOR on the path between u and v”. Binary lifting for LCA plus a Fenwick tree that supports range XOR updates gives us exactly that in O(log n) per operation.

## Approach

Root the tree at node 0. Run a DFS that records entry and exit times for every node (Euler-tour timestamps), depths, and the binary-lifting parent table.  

Build a Fenwick tree that can XOR a value onto a whole range and answer a point query. For every node u, XOR its initial letter-mask onto the interval that covers its whole subtree. After this step a point query at the entry time of any node returns the XOR of all letters from the root down to that node.  

When a letter is updated, compute the delta between the old and new masks and XOR that delta onto the same subtree interval. Every root-path that contained the node is corrected automatically.  

For a path query u–v compute the LCA with binary lifting, then evaluate  

rootXOR(u) XOR rootXOR(v) XOR mask[lca].  

The resulting mask has ≤ 1 bit set if and only if the path can form a palindrome.

## Data Structures Used

- Adjacency list – stores the tree edges for DFS.
- Binary-lifting table – answers LCA queries in O(log n).
- Entry / exit time arrays – turn every subtree into a contiguous range.
- Fenwick tree (Binary Indexed Tree) – supports range XOR updates and point queries in O(log n).
- Current-mask array – keeps the latest 26-bit mask of every node so we can include the LCA correctly.

## Operations & Behavior Summary

1. Build the tree and run DFS to fill timestamps, depths and the lifting table.  
2. Initialize the Fenwick tree by range-XORing every node’s mask onto its subtree.  
3. For an update: calculate the delta mask, write the new mask, and range-XOR the delta onto the subtree.  
4. For a query: find LCA, fetch the two root-to-node XORs, combine them with the LCA mask, and test the pop-count.  
5. Collect every boolean answer and return the list.

## Complexity

| Type              | Complexity          | Explanation |
|-------------------|---------------------|-------------|
| Time Complexity   | O((n + q) log n)    | DFS and lifting table cost O(n log n). Each of the q operations does a constant number of O(log n) Fenwick or LCA steps. |
| Space Complexity  | O(n log n)          | Binary-lifting table dominates; Fenwick tree and auxiliary arrays are linear. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<bool> palindromePath(int n, vector<vector<int>>& edges, string s, vector<string>& queries) {
        vector<vector<int>> g(n);
        for (auto& e : edges) {
            g[e[0]].push_back(e[1]);
            g[e[1]].push_back(e[0]);
        }
        const int LOG = 16;
        vector<vector<int>> up(LOG, vector<int>(n, -1));
        vector<int> depth(n), in(n), out(n), parent(n, -1);
        vector<int> mask(n);
        for (int i = 0; i < n; i++) mask[i] = 1 << (s[i] - 'a');
        int timer = 1;
        function<void(int, int)> dfs = [&](int u, int p) {
            in[u] = timer++;
            parent[u] = p;
            up[0][u] = p;
            for (int k = 1; k < LOG; k++) {
                if (up[k - 1][u] != -1) up[k][u] = up[k - 1][up[k - 1][u]];
            }
            for (int v : g[u]) if (v != p) {
                depth[v] = depth[u] + 1;
                dfs(v, u);
            }
            out[u] = timer;
        };
        dfs(0, -1);
        vector<int> bit(n + 2, 0);
        auto upd = [&](int i, int val) {
            for (; i <= n; i += i & -i) bit[i] ^= val;
        };
        auto qry = [&](int i) {
            int res = 0;
            for (; i > 0; i -= i & -i) res ^= bit[i];
            return res;
        };
        auto rangeXor = [&](int l, int r, int val) {
            if (l > r) return;
            upd(l, val);
            upd(r + 1, val);
        };
        for (int i = 0; i < n; i++) rangeXor(in[i], out[i] - 1, mask[i]);
        auto getLca = [&](int u, int v) {
            if (depth[u] < depth[v]) swap(u, v);
            int diff = depth[u] - depth[v];
            for (int k = 0; k < LOG; k++) if (diff >> k & 1) u = up[k][u];
            if (u == v) return u;
            for (int k = LOG - 1; k >= 0; k--) {
                if (up[k][u] != up[k][v]) {
                    u = up[k][u];
                    v = up[k][v];
                }
            }
            return parent[u];
        };
        auto pathXor = [&](int u, int v) {
            int l = getLca(u, v);
            return qry(in[u]) ^ qry(in[v]) ^ mask[l];
        };
        vector<bool> ans;
        for (auto& q : queries) {
            if (q[0] == 'u') {
                int u = 0, i = 7;
                while (q[i] != ' ') u = u * 10 + (q[i++] - '0');
                char c = q.back();
                int nm = 1 << (c - 'a');
                int delta = mask[u] ^ nm;
                mask[u] = nm;
                rangeXor(in[u], out[u] - 1, delta);
            } else {
                int u = 0, v = 0, i = 6;
                while (q[i] != ' ') u = u * 10 + (q[i++] - '0');
                i++;
                while (i < (int)q.size()) v = v * 10 + (q[i++] - '0');
                int m = pathXor(u, v);
                ans.push_back(__builtin_popcount(m) <= 1);
            }
        }
        return ans;
    }
};
```

### Java

```java
class Solution {
    public List<Boolean> palindromePath(int n, int[][] edges, String s, String[] queries) {
        List<Integer>[] g = new ArrayList[n];
        for (int i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (int[] e : edges) {
            g[e[0]].add(e[1]);
            g[e[1]].add(e[0]);
        }
        final int LOG = 16;
        int[][] up = new int[LOG][n];
        for (int[] row : up) Arrays.fill(row, -1);
        int[] depth = new int[n], in = new int[n], out = new int[n], parent = new int[n];
        Arrays.fill(parent, -1);
        int[] mask = new int[n];
        for (int i = 0; i < n; i++) mask[i] = 1 << (s.charAt(i) - 'a');
        int[] timer = {1};
        dfs(0, -1, g, up, depth, in, out, parent, timer, LOG);
        int[] bit = new int[n + 2];
        for (int i = 0; i < n; i++) rangeXor(in[i], out[i] - 1, mask[i], bit, n);
        List<Boolean> ans = new ArrayList<>();
        for (String q : queries) {
            if (q.charAt(0) == 'u') {
                int u = 0, i = 7;
                while (q.charAt(i) != ' ') u = u * 10 + (q.charAt(i++) - '0');
                char c = q.charAt(q.length() - 1);
                int nm = 1 << (c - 'a');
                int delta = mask[u] ^ nm;
                mask[u] = nm;
                rangeXor(in[u], out[u] - 1, delta, bit, n);
            } else {
                int u = 0, v = 0, i = 6;
                while (q.charAt(i) != ' ') u = u * 10 + (q.charAt(i++) - '0');
                i++;
                while (i < q.length()) v = v * 10 + (q.charAt(i++) - '0');
                int m = pathXor(u, v, in, mask, up, depth, parent, bit, LOG);
                ans.add(Integer.bitCount(m) <= 1);
            }
        }
        return ans;
    }
    void dfs(int u, int p, List<Integer>[] g, int[][] up, int[] depth, int[] in, int[] out, int[] parent, int[] timer, int LOG) {
        in[u] = timer[0]++;
        parent[u] = p;
        up[0][u] = p;
        for (int k = 1; k < LOG; k++) {
            if (up[k - 1][u] != -1) up[k][u] = up[k - 1][up[k - 1][u]];
        }
        for (int v : g[u]) if (v != p) {
            depth[v] = depth[u] + 1;
            dfs(v, u, g, up, depth, in, out, parent, timer, LOG);
        }
        out[u] = timer[0];
    }
    void upd(int i, int val, int[] bit, int n) {
        for (; i <= n; i += i & -i) bit[i] ^= val;
    }
    int qry(int i, int[] bit) {
        int res = 0;
        for (; i > 0; i -= i & -i) res ^= bit[i];
        return res;
    }
    void rangeXor(int l, int r, int val, int[] bit, int n) {
        if (l > r) return;
        upd(l, val, bit, n);
        upd(r + 1, val, bit, n);
    }
    int getLca(int u, int v, int[][] up, int[] depth, int[] parent, int LOG) {
        if (depth[u] < depth[v]) {
            int t = u; u = v; v = t;
        }
        int diff = depth[u] - depth[v];
        for (int k = 0; k < LOG; k++) if (((diff >> k) & 1) == 1) u = up[k][u];
        if (u == v) return u;
        for (int k = LOG - 1; k >= 0; k--) {
            if (up[k][u] != up[k][v]) {
                u = up[k][u];
                v = up[k][v];
            }
        }
        return parent[u];
    }
    int pathXor(int u, int v, int[] in, int[] mask, int[][] up, int[] depth, int[] parent, int[] bit, int LOG) {
        int l = getLca(u, v, up, depth, parent, LOG);
        return qry(in[u], bit) ^ qry(in[v], bit) ^ mask[l];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {string} s
 * @param {string[]} queries
 * @return {boolean[]}
 */
var palindromePath = function(n, edges, s, queries) {
    const g = Array.from({length: n}, () => []);
    for (const [u, v] of edges) {
        g[u].push(v);
        g[v].push(u);
    }
    const LOG = 16;
    const up = Array.from({length: LOG}, () => new Array(n).fill(-1));
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
        for (const v of g[u]) if (v !== p) {
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
        if (q[0] === 'u') {
            let u = 0, i = 7;
            while (q[i] !== ' ') u = u * 10 + (q.charCodeAt(i++) - 48);
            const c = q[q.length - 1];
            const nm = 1 << (c.charCodeAt(0) - 97);
            const delta = mask[u] ^ nm;
            mask[u] = nm;
            rangeXor(inn[u], out[u] - 1, delta);
        } else {
            let u = 0, v = 0, i = 6;
            while (q[i] !== ' ') u = u * 10 + (q.charCodeAt(i++) - 48);
            i++;
            while (i < q.length) v = v * 10 + (q.charCodeAt(i++) - 48);
            const m = pathXor(u, v);
            ans.push((m.toString(2).split('1').length - 1) <= 1);
        }
    }
    return ans;
};
```

### TypeScript

```typescript
function palindromePath(n: number, edges: number[][], s: string, queries: string[]): boolean[] {
    const g: number[][] = Array.from({length: n}, () => []);
    for (const [u, v] of edges) {
        g[u].push(v);
        g[v].push(u);
    }
    const LOG = 16;
    const up: number[][] = Array.from({length: LOG}, () => new Array(n).fill(-1));
    const depth = new Array(n).fill(0);
    const inn = new Array(n);
    const out = new Array(n);
    const parent = new Array(n).fill(-1);
    const mask = new Array(n);
    for (let i = 0; i < n; i++) mask[i] = 1 << (s.charCodeAt(i) - 97);
    let timer = 1;
    function dfs(u: number, p: number) {
        inn[u] = timer++;
        parent[u] = p;
        up[0][u] = p;
        for (let k = 1; k < LOG; k++) {
            if (up[k - 1][u] !== -1) up[k][u] = up[k - 1][up[k - 1][u]];
        }
        for (const v of g[u]) if (v !== p) {
            depth[v] = depth[u] + 1;
            dfs(v, u);
        }
        out[u] = timer;
    }
    dfs(0, -1);
    const bit = new Array(n + 2).fill(0);
    function upd(i: number, val: number) {
        for (; i <= n; i += i & -i) bit[i] ^= val;
    }
    function qry(i: number): number {
        let res = 0;
        for (; i > 0; i -= i & -i) res ^= bit[i];
        return res;
    }
    function rangeXor(l: number, r: number, val: number) {
        if (l > r) return;
        upd(l, val);
        upd(r + 1, val);
    }
    for (let i = 0; i < n; i++) rangeXor(inn[i], out[i] - 1, mask[i]);
    function getLca(u: number, v: number): number {
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
    function pathXor(u: number, v: number): number {
        const l = getLca(u, v);
        return qry(inn[u]) ^ qry(inn[v]) ^ mask[l];
    }
    const ans: boolean[] = [];
    for (const q of queries) {
        if (q[0] === 'u') {
            let u = 0, i = 7;
            while (q[i] !== ' ') u = u * 10 + (q.charCodeAt(i++) - 48);
            const c = q[q.length - 1];
            const nm = 1 << (c.charCodeAt(0) - 97);
            const delta = mask[u] ^ nm;
            mask[u] = nm;
            rangeXor(inn[u], out[u] - 1, delta);
        } else {
            let u = 0, v = 0, i = 6;
            while (q[i] !== ' ') u = u * 10 + (q.charCodeAt(i++) - 48);
            i++;
            while (i < q.length) v = v * 10 + (q.charCodeAt(i++) - 48);
            const m = pathXor(u, v);
            ans.push((m.toString(2).split('1').length - 1) <= 1);
        }
    }
    return ans;
}
```

### Python3

```python
class Solution:
    def palindromePath(self, n: int, edges: list[list[int]], s: str, queries: list[str]) -> list[bool]:
        g = [[] for _ in range(n)]
        for u, v in edges:
            g[u].append(v)
            g[v].append(u)
        LOG = 16
        up = [[-1] * n for _ in range(LOG)]
        depth = [0] * n
        inn = [0] * n
        out = [0] * n
        parent = [-1] * n
        mask = [1 << (ord(c) - 97) for c in s]
        timer = [1]
        def dfs(u, p):
            inn[u] = timer[0]
            timer[0] += 1
            parent[u] = p
            up[0][u] = p
            for k in range(1, LOG):
                if up[k - 1][u] != -1:
                    up[k][u] = up[k - 1][up[k - 1][u]]
            for v in g[u]:
                if v != p:
                    depth[v] = depth[u] + 1
                    dfs(v, u)
            out[u] = timer[0]
        dfs(0, -1)
        bit = [0] * (n + 2)
        def upd(i, val):
            while i <= n:
                bit[i] ^= val
                i += i & -i
        def qry(i):
            res = 0
            while i > 0:
                res ^= bit[i]
                i -= i & -i
            return res
        def range_xor(l, r, val):
            if l > r:
                return
            upd(l, val)
            upd(r + 1, val)
        for i in range(n):
            range_xor(inn[i], out[i] - 1, mask[i])
        def get_lca(u, v):
            if depth[u] < depth[v]:
                u, v = v, u
            diff = depth[u] - depth[v]
            for k in range(LOG):
                if (diff >> k) & 1:
                    u = up[k][u]
            if u == v:
                return u
            for k in range(LOG - 1, -1, -1):
                if up[k][u] != up[k][v]:
                    u = up[k][u]
                    v = up[k][v]
            return parent[u]
        def path_xor(u, v):
            l = get_lca(u, v)
            return qry(inn[u]) ^ qry(inn[v]) ^ mask[l]
        ans = []
        for q in queries:
            if q[0] == 'u':
                u = 0
                i = 7
                while q[i] != ' ':
                    u = u * 10 + (ord(q[i]) - 48)
                    i += 1
                c = q[-1]
                nm = 1 << (ord(c) - 97)
                delta = mask[u] ^ nm
                mask[u] = nm
                range_xor(inn[u], out[u] - 1, delta)
            else:
                u = 0
                v = 0
                i = 6
                while q[i] != ' ':
                    u = u * 10 + (ord(q[i]) - 48)
                    i += 1
                i += 1
                while i < len(q):
                    v = v * 10 + (ord(q[i]) - 48)
                    i += 1
                m = path_xor(u, v)
                ans.append(bin(m).count('1') <= 1)
        return ans
```

### Go

```go
func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
    g := make([][]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], e[1])
        g[e[1]] = append(g[e[1]], e[0])
    }
    const LOG = 16
    up := make([][]int, LOG)
    for i := range up {
        up[i] = make([]int, n)
        for j := range up[i] {
            up[i][j] = -1
        }
    }
    depth := make([]int, n)
    inn := make([]int, n)
    out := make([]int, n)
    parent := make([]int, n)
    for i := range parent {
        parent[i] = -1
    }
    mask := make([]int, n)
    for i := 0; i < n; i++ {
        mask[i] = 1 << (s[i] - 'a')
    }
    timer := 1
    var dfs func(int, int)
    dfs = func(u, p int) {
        inn[u] = timer
        timer++
        parent[u] = p
        up[0][u] = p
        for k := 1; k < LOG; k++ {
            if up[k-1][u] != -1 {
                up[k][u] = up[k-1][up[k-1][u]]
            }
        }
        for _, v := range g[u] {
            if v != p {
                depth[v] = depth[u] + 1
                dfs(v, u)
            }
        }
        out[u] = timer
    }
    dfs(0, -1)
    bit := make([]int, n+2)
    upd := func(i, val int) {
        for ; i <= n; i += i & -i {
            bit[i] ^= val
        }
    }
    qry := func(i int) int {
        res := 0
        for ; i > 0; i -= i & -i {
            res ^= bit[i]
        }
        return res
    }
    rangeXor := func(l, r, val int) {
        if l > r {
            return
        }
        upd(l, val)
        upd(r+1, val)
    }
    for i := 0; i < n; i++ {
        rangeXor(inn[i], out[i]-1, mask[i])
    }
    getLca := func(u, v int) int {
        if depth[u] < depth[v] {
            u, v = v, u
        }
        diff := depth[u] - depth[v]
        for k := 0; k < LOG; k++ {
            if (diff>>k)&1 == 1 {
                u = up[k][u]
            }
        }
        if u == v {
            return u
        }
        for k := LOG - 1; k >= 0; k-- {
            if up[k][u] != up[k][v] {
                u = up[k][u]
                v = up[k][v]
            }
        }
        return parent[u]
    }
    pathXor := func(u, v int) int {
        l := getLca(u, v)
        return qry(inn[u]) ^ qry(inn[v]) ^ mask[l]
    }
    ans := []bool{}
    for _, q := range queries {
        if q[0] == 'u' {
            u := 0
            i := 7
            for q[i] != ' ' {
                u = u*10 + int(q[i]-'0')
                i++
            }
            c := q[len(q)-1]
            nm := 1 << (c - 'a')
            delta := mask[u] ^ nm
            mask[u] = nm
            rangeXor(inn[u], out[u]-1, delta)
        } else {
            u := 0
            v := 0
            i := 6
            for q[i] != ' ' {
                u = u*10 + int(q[i]-'0')
                i++
            }
            i++
            for i < len(q) {
                v = v*10 + int(q[i]-'0')
                i++
            }
            m := pathXor(u, v)
            cnt := 0
            for m > 0 {
                cnt += m & 1
                m >>= 1
            }
            ans = append(ans, cnt <= 1)
        }
    }
    return ans
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The same algorithmic idea is implemented in every language; only syntax and a few library details differ.

First we turn the edge list into an adjacency list so the DFS can walk the tree.  

We allocate a binary-lifting table of height 16 (enough for n ≤ 5 × 10⁴), depth, entry-time and exit-time arrays, and a parent array. While the DFS runs we also fill the first row of the lifting table and then complete the higher rows. After the DFS finishes every subtree occupies a contiguous interval of timestamps.  

A Fenwick tree of size n + 2 is created. The classic difference-array trick turns a range-XOR update into two point updates. For every node we XOR its initial mask onto the whole interval that belongs to its subtree. From that moment a point query at the entry time of a node returns the XOR of the path from the root down to that node.  

When a letter changes we compute the delta, store the new mask, and range-XOR the delta onto the same interval. All affected root-paths are updated in O(log n).  

For a path query we lift the deeper endpoint up to the same depth as the shallower one, then lift both endpoints together until their parents coincide; that parent is the LCA. The path XOR is then simply the root-XOR of the two endpoints XORed with the mask of the LCA. Because XOR is its own inverse the segments above the LCA cancel and the LCA itself is counted exactly once.  

Finally we test the mask with a pop-count (or an equivalent bit-counting loop). If the count is zero or one we record true, otherwise false.

Edge cases such as u == v, the root being one of the endpoints, or a node being updated many times are all handled correctly by the same formulas.

## Examples

**Example 1**  
Input: n = 3, edges = [[0,1],[1,2]], s = "aac", queries = ["query 0 2","update 1 b","query 0 2"]  
Output: [true, false]  

- Path 0-1-2 has letters a,a,c → two a’s and one c → can form a palindrome.  
- After changing node 1 to ‘b’ the path becomes a,b,c → three distinct letters → cannot form a palindrome.

**Example 2**  
Input: n = 4, edges = [[0,1],[0,2],[0,3]], s = "abca", queries = ["query 1 2","update 0 b","query 2 3","update 3 a","query 1 3"]  
Output: [false, false, true]  

- Path 1-0-2 = b,a,c → three odd counts.  
- After node 0 becomes ‘b’ the path 2-0-3 = c,b,a → still three odds.  
- After node 3 becomes ‘a’ the path 1-0-3 = b,b,a → one odd count → true.

## How to Use / Run Locally

1. Copy the code of the language you want into a file (Solution.cpp, Solution.java, \ldots).  
2. C++: `g++ -std=c++17 -O2 Solution.cpp -o sol && ./sol`  
3. Java: `javac Solution.java && java Solution`  
4. JavaScript: `node solution.js`  
5. TypeScript: `tsc solution.ts && node solution.js`  
6. Python3: `python3 solution.py`  
7. Go: `go run solution.go`  

Replace the hard-coded sample inputs with your own test cases or read from standard input as needed.

## Notes & Optimizations

- The Fenwick + binary-lifting approach runs in O(log n) per operation, which is asymptotically better than classic Heavy-Light Decomposition (O(log² n)).  
- Using a 26-bit integer mask is both space-efficient and extremely fast for the palindrome test.  
- If the letter alphabet were larger than 26 the same idea still works with a larger bit-set or a hash of frequencies, but for the given constraints a single integer is perfect.  
- The solution is online: every query is answered immediately after it is read, which matches the problem requirements.  
- Memory usage stays comfortably under typical contest limits even for the maximum constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
