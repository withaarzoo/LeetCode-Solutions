# Number of Ways to Assign Edge Weights II - LeetCode 3559 Solution

A complete explanation of LeetCode 3559, including intuition, LCA (Lowest Common Ancestor), Binary Lifting, tree distance calculation, complexity analysis, and multi-language implementations in C++, Java, JavaScript, Python, and Go.

## Table of Contents

* [Problem Summary](#problem-summary)
* [Constraints](#constraints)
* [Intuition](#intuition)
* [Approach](#approach)
* [Data Structures Used](#data-structures-used)
* [Operations & Behavior Summary](#operations--behavior-summary)
* [Complexity](#complexity)
* [Multi-language Solutions](#multi-language-solutions)

  * [C++](#c)
  * [Java](#java)
  * [JavaScript](#javascript)
  * [Python3](#python3)
  * [Go](#go)
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given an undirected tree containing `n` nodes.

Every edge must be assigned either:

* Weight `1`
* Weight `2`

For each query `[u, v]`, you only consider the path between node `u` and node `v`.

The goal is to find how many different edge-weight assignments make the total path cost odd.

Since the answer can become very large, return it modulo:

`10^9 + 7`

This problem combines several important data structure and algorithm concepts:

* Trees
* Lowest Common Ancestor (LCA)
* Binary Lifting
* Path Distance Queries
* Modular Arithmetic
* Parity Observations

---

## Constraints

| Constraint                    | Value             |
| ----------------------------- | ----------------- |
| `2 <= n <= 10^5`              | Number of nodes   |
| `edges.length == n - 1`       | Valid tree        |
| `1 <= queries.length <= 10^5` | Number of queries |
| `1 <= ui, vi <= n`            | Query nodes       |
| Tree is connected             | Guaranteed        |

---

## Intuition

The first thing I noticed was that the actual values `1` and `2` are not as important as their parity.

* `1` is odd
* `2` is even

The total path cost becomes odd only when an odd number of edges receive weight `1`.

For a path containing `L` edges:

* Every edge has 2 choices
* Total assignments = `2^L`

Among all possible assignments, exactly half produce an odd sum and half produce an even sum.

That means:

* Odd assignments = `2^(L-1)`
* Even assignments = `2^(L-1)`

for every path where `L > 0`.

So the entire problem becomes:

"How quickly can I find the number of edges between two nodes?"

That is exactly what LCA and Binary Lifting help us do efficiently.

---

## Approach

### Step 1: Build the Tree

Convert the edge list into an adjacency list representation.

This allows efficient DFS traversal.

### Step 2: Preprocess Depth and Ancestors

Run DFS from node `1`.

Store:

* Depth of every node
* Immediate parent
* Binary lifting table

### Step 3: Build Binary Lifting Table

Store:

`up[node][j]`

which represents the `2^j`th ancestor of a node.

This allows jumping upward in logarithmic time.

### Step 4: Precompute Powers of Two

Since every answer is:

`2^(distance - 1)`

precompute all powers of two modulo `10^9 + 7`.

### Step 5: Process Queries

For each query:

1. Find LCA of `u` and `v`
2. Calculate path length
3. If path length is zero, answer is `0`
4. Otherwise return:

`2^(distance - 1)`

---

## Data Structures Used

### Adjacency List

Used to represent the tree efficiently.

### Depth Array

Stores depth of every node from the root.

### Binary Lifting Table

Stores ancestors for fast LCA queries.

### Power Array

Stores precomputed powers of two modulo `10^9 + 7`.

### Answer Array

Stores the result for each query.

---

## Operations & Behavior Summary

1. Read tree edges.
2. Build adjacency list.
3. Start DFS from node `1`.
4. Compute depth values.
5. Build ancestor table.
6. Precompute powers of two.
7. For every query:

   * Find LCA.
   * Compute path distance.
   * Convert distance into number of valid assignments.
8. Return all answers.

---

## Complexity

| Operation         | Complexity       |
| ----------------- | ---------------- |
| Tree Construction | O(n)             |
| DFS Preprocessing | O(n log n)       |
| LCA Query         | O(log n)         |
| Query Processing  | O(q log n)       |
| Total Complexity  | O((n + q) log n) |

### Time Complexity

**O((n + q) log n)**

Where:

* `n` = number of nodes
* `q` = number of queries

### Space Complexity

**O(n log n)**

Extra memory is used for:

* Adjacency list
* Depth array
* Binary lifting table
* Power array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static const int MOD = 1000000007;

    vector<int> depth;
    vector<vector<int>> up;
    vector<vector<int>> graph;
    int LOG;

    // DFS to fill depth and binary lifting table
    void dfs(int node, int parent) {
        up[node][0] = parent;

        for (int j = 1; j < LOG; j++) {
            up[node][j] = up[up[node][j - 1]][j - 1];
        }

        for (int next : graph[node]) {
            if (next == parent) continue;

            depth[next] = depth[node] + 1;
            dfs(next, node);
        }
    }

    // Find LCA using binary lifting
    int lca(int a, int b) {
        if (depth[a] < depth[b]) swap(a, b);

        int diff = depth[a] - depth[b];

        for (int j = LOG - 1; j >= 0; j--) {
            if ((diff >> j) & 1) {
                a = up[a][j];
            }
        }

        if (a == b) return a;

        for (int j = LOG - 1; j >= 0; j--) {
            if (up[a][j] != up[b][j]) {
                a = up[a][j];
                b = up[b][j];
            }
        }

        return up[a][0];
    }

    vector<int> assignEdgeWeights(vector<vector<int>>& edges, vector<vector<int>>& queries) {
        int n = edges.size() + 1;

        LOG = 17;
        while ((1 << LOG) <= n) LOG++;

        graph.assign(n + 1, {});
        for (auto &e : edges) {
            int u = e[0];
            int v = e[1];

            graph[u].push_back(v);
            graph[v].push_back(u);
        }

        depth.assign(n + 1, 0);
        up.assign(n + 1, vector<int>(LOG, 1));

        dfs(1, 1);

        // Precompute powers of 2 modulo MOD
        vector<int> pow2(n + 1, 1);
        for (int i = 1; i <= n; i++) {
            pow2[i] = (long long)pow2[i - 1] * 2 % MOD;
        }

        vector<int> ans;

        for (auto &q : queries) {
            int u = q[0];
            int v = q[1];

            int ancestor = lca(u, v);

            int dist = depth[u] + depth[v] - 2 * depth[ancestor];

            if (dist == 0) {
                ans.push_back(0);
            } else {
                ans.push_back(pow2[dist - 1]);
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    static final int MOD = 1000000007;

    int LOG;
    int[] depth;
    int[][] up;
    java.util.List<Integer>[] graph;

    // DFS to compute depth and ancestors
    void dfs(int node, int parent) {
        up[node][0] = parent;

        for (int j = 1; j < LOG; j++) {
            up[node][j] = up[up[node][j - 1]][j - 1];
        }

        for (int next : graph[node]) {
            if (next == parent) continue;

            depth[next] = depth[node] + 1;
            dfs(next, node);
        }
    }

    // Binary lifting LCA
    int lca(int a, int b) {
        if (depth[a] < depth[b]) {
            int temp = a;
            a = b;
            b = temp;
        }

        int diff = depth[a] - depth[b];

        for (int j = LOG - 1; j >= 0; j--) {
            if (((diff >> j) & 1) == 1) {
                a = up[a][j];
            }
        }

        if (a == b) return a;

        for (int j = LOG - 1; j >= 0; j--) {
            if (up[a][j] != up[b][j]) {
                a = up[a][j];
                b = up[b][j];
            }
        }

        return up[a][0];
    }

    public int[] assignEdgeWeights(int[][] edges, int[][] queries) {
        int n = edges.length + 1;

        LOG = 17;
        while ((1 << LOG) <= n) LOG++;

        graph = new ArrayList[n + 1];
        for (int i = 0; i <= n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] e : edges) {
            int u = e[0];
            int v = e[1];

            graph[u].add(v);
            graph[v].add(u);
        }

        depth = new int[n + 1];
        up = new int[n + 1][LOG];

        dfs(1, 1);

        int[] pow2 = new int[n + 1];
        pow2[0] = 1;

        for (int i = 1; i <= n; i++) {
            pow2[i] = (int)((long)pow2[i - 1] * 2 % MOD);
        }

        int[] ans = new int[queries.length];

        for (int i = 0; i < queries.length; i++) {
            int u = queries[i][0];
            int v = queries[i][1];

            int ancestor = lca(u, v);

            int dist = depth[u] + depth[v] - 2 * depth[ancestor];

            ans[i] = (dist == 0) ? 0 : pow2[dist - 1];
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} edges
 * @param {number[][]} queries
 * @return {number[]}
 */
var assignEdgeWeights = function(edges, queries) {
    const MOD = 1000000007;
    const n = edges.length + 1;

    let LOG = 1;
    while ((1 << LOG) <= n) LOG++;

    // Adjacency list
    const graph = Array.from({ length: n + 1 }, () => []);

    for (const [u, v] of edges) {
        graph[u].push(v);
        graph[v].push(u);
    }

    const depth = Array(n + 1).fill(0);
    const up = Array.from({ length: n + 1 }, () => Array(LOG).fill(1));

    // DFS
    const dfs = (node, parent) => {
        up[node][0] = parent;

        for (let j = 1; j < LOG; j++) {
            up[node][j] = up[up[node][j - 1]][j - 1];
        }

        for (const next of graph[node]) {
            if (next === parent) continue;

            depth[next] = depth[node] + 1;
            dfs(next, node);
        }
    };

    dfs(1, 1);

    // LCA
    const lca = (a, b) => {
        if (depth[a] < depth[b]) {
            [a, b] = [b, a];
        }

        let diff = depth[a] - depth[b];

        for (let j = LOG - 1; j >= 0; j--) {
            if ((diff >> j) & 1) {
                a = up[a][j];
            }
        }

        if (a === b) return a;

        for (let j = LOG - 1; j >= 0; j--) {
            if (up[a][j] !== up[b][j]) {
                a = up[a][j];
                b = up[b][j];
            }
        }

        return up[a][0];
    };

    // Powers of two modulo MOD
    const pow2 = Array(n + 1).fill(1);

    for (let i = 1; i <= n; i++) {
        pow2[i] = (pow2[i - 1] * 2) % MOD;
    }

    const ans = [];

    for (const [u, v] of queries) {
        const ancestor = lca(u, v);

        const dist =
            depth[u] + depth[v] - 2 * depth[ancestor];

        ans.push(dist === 0 ? 0 : pow2[dist - 1]);
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def assignEdgeWeights(self, edges: List[List[int]], queries: List[List[int]]) -> List[int]:
        MOD = 1000000007

        n = len(edges) + 1

        LOG = 1
        while (1 << LOG) <= n:
            LOG += 1

        # Build adjacency list
        graph = [[] for _ in range(n + 1)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        depth = [0] * (n + 1)
        up = [[1] * LOG for _ in range(n + 1)]

        # DFS preprocessing
        def dfs(node: int, parent: int) -> None:
            up[node][0] = parent

            for j in range(1, LOG):
                up[node][j] = up[up[node][j - 1]][j - 1]

            for nxt in graph[node]:
                if nxt == parent:
                    continue

                depth[nxt] = depth[node] + 1
                dfs(nxt, node)

        dfs(1, 1)

        # Binary lifting LCA
        def lca(a: int, b: int) -> int:
            if depth[a] < depth[b]:
                a, b = b, a

            diff = depth[a] - depth[b]

            for j in range(LOG - 1, -1, -1):
                if (diff >> j) & 1:
                    a = up[a][j]

            if a == b:
                return a

            for j in range(LOG - 1, -1, -1):
                if up[a][j] != up[b][j]:
                    a = up[a][j]
                    b = up[b][j]

            return up[a][0]

        # Precompute powers of two
        pow2 = [1] * (n + 1)

        for i in range(1, n + 1):
            pow2[i] = (pow2[i - 1] * 2) % MOD

        ans = []

        for u, v in queries:
            ancestor = lca(u, v)

            dist = depth[u] + depth[v] - 2 * depth[ancestor]

            if dist == 0:
                ans.append(0)
            else:
                ans.append(pow2[dist - 1])

        return ans
```

### Go

```go
func assignEdgeWeights(edges [][]int, queries [][]int) []int {
 const MOD = 1000000007

 n := len(edges) + 1

 // Compute required binary lifting height
 LOG := 1
 for (1 << LOG) <= n {
  LOG++
 }

 // Build adjacency list
 graph := make([][]int, n+1)

 for _, e := range edges {
  u, v := e[0], e[1]

  graph[u] = append(graph[u], v)
  graph[v] = append(graph[v], u)
 }

 depth := make([]int, n+1)

 // up[node][j] = 2^j-th ancestor
 up := make([][]int, n+1)
 for i := 0; i <= n; i++ {
  up[i] = make([]int, LOG)
  for j := 0; j < LOG; j++ {
   up[i][j] = 1
  }
 }

 // DFS preprocessing
 var dfs func(int, int)

 dfs = func(node, parent int) {
  up[node][0] = parent

  for j := 1; j < LOG; j++ {
   up[node][j] = up[up[node][j-1]][j-1]
  }

  for _, next := range graph[node] {
   if next == parent {
    continue
   }

   depth[next] = depth[node] + 1
   dfs(next, node)
  }
 }

 dfs(1, 1)

 // LCA using binary lifting
 lca := func(a, b int) int {
  if depth[a] < depth[b] {
   a, b = b, a
  }

  diff := depth[a] - depth[b]

  for j := LOG - 1; j >= 0; j-- {
   if ((diff >> j) & 1) == 1 {
    a = up[a][j]
   }
  }

  if a == b {
   return a
  }

  for j := LOG - 1; j >= 0; j-- {
   if up[a][j] != up[b][j] {
    a = up[a][j]
    b = up[b][j]
   }
  }

  return up[a][0]
 }

 // Precompute powers of 2
 pow2 := make([]int, n+1)
 pow2[0] = 1

 for i := 1; i <= n; i++ {
  pow2[i] = int((int64(pow2[i-1]) * 2) % MOD)
 }

 ans := make([]int, len(queries))

 for i, q := range queries {
  u, v := q[0], q[1]

  ancestor := lca(u, v)

  dist := depth[u] + depth[v] - 2*depth[ancestor]

  if dist == 0 {
   ans[i] = 0
  } else {
   ans[i] = pow2[dist-1]
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical across all five languages.

### Building the Tree

The edge list is converted into an adjacency list.

Since the tree is undirected, every edge is added in both directions.

This makes traversal easy and efficient.

### DFS Traversal

Starting from node `1`:

* Compute depth of each node.
* Store parent information.
* Fill the first level of the binary lifting table.

### Binary Lifting Preparation

For every node:

`up[node][j]`

stores the ancestor located `2^j` levels above.

This preprocessing allows fast upward jumps.

### Finding LCA

When two nodes are given:

1. Move the deeper node upward.
2. Make both nodes reach the same depth.
3. Lift both nodes together.
4. The first common ancestor becomes the LCA.

This process takes only `O(log n)` time.

### Calculating Distance

Once LCA is known:

distance = depth[u] + depth[v] − 2 × depth[lca]

This gives the number of edges on the path.

### Counting Valid Assignments

For a path with `L` edges:

* Total assignments = `2^L`
* Odd-sum assignments = `2^(L−1)`

Special case:

If `L = 0`, there are no edges.

The path sum is always `0`.

Therefore the answer is `0`.

---

## Examples

### Example 1

#### Input

```text
edges = [[1,2]]
queries = [[1,1],[1,2]]
```

#### Output

```text
[0,1]
```

#### Explanation

Query `[1,1]`

* Distance = 0
* No edges exist
* Answer = 0

Query `[1,2]`

* Distance = 1
* Answer = 2^(1-1)
* Answer = 1

---

### Example 2

#### Input

```text
edges = [[1,2],[1,3],[3,4],[3,5]]
queries = [[1,4],[3,4],[2,5]]
```

#### Output

```text
[2,1,4]
```

#### Explanation

Path lengths:

* 1 → 4 = 2 edges
* 3 → 4 = 1 edge
* 2 → 5 = 3 edges

Answers:

* 2^(2−1) = 2
* 2^(1−1) = 1
* 2^(3−1) = 4

---

### Example 3

#### Input

```text
edges = [[1,2],[2,3]]
queries = [[1,3]]
```

#### Output

```text
[2]
```

#### Explanation

Distance between nodes:

```text
1 -> 2 -> 3
```

Path length = 2

Valid odd-sum assignments:

```text
(1,2)
(2,1)
```

Total = 2

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ solution.cpp -O2 -std=c++17
```

Run:

```bash
./a.out
```

---

### Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run:

```bash
node solution.js
```

---

### Python3

Run:

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

Build:

```bash
go build solution.go
```

---

## Notes & Optimizations

* The key observation is parity.
* Actual edge values do not matter beyond odd/even behavior.
* LCA reduces path-distance queries from linear time to logarithmic time.
* Binary Lifting is the most practical approach for large constraints.
* Precomputing powers of two removes repeated modular exponentiation.
* The solution easily handles up to `10^5` nodes and `10^5` queries.

### Edge Cases

* Query where `u == v`
* Tree with only one edge
* Deep tree structures
* Large number of queries
* Queries involving the root node

### Alternative Approaches

Other LCA techniques include:

* Euler Tour + RMQ
* Heavy-Light Decomposition

However, Binary Lifting is simpler and perfectly fits the constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
