# 3620. Network Recovery Pathways | Binary Search + DAG Dynamic Programming Solution

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
  - [Python3](#python3)
  - [Go](#go)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

In this problem, we are given a **Directed Acyclic Graph (DAG)** where every edge has a recovery cost. Some nodes may be offline, and every valid path must avoid using offline intermediate nodes.

The goal is to find a path from node `0` to node `n - 1` such that:

- The total recovery cost does not exceed `k`.
- Every intermediate node is online.
- Among all valid paths, the minimum edge cost on the path is as large as possible.

In other words, instead of minimizing the total path cost, we want to maximize the weakest edge on the chosen path while staying within the allowed total cost.

This is a great problem that combines **Binary Search**, **Dynamic Programming**, **Topological Sort**, and **Graph Algorithms**.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Number of nodes (`n`) | `2 <= n <= 5 × 10^4` |
| Number of edges (`m`) | `0 <= m <= min(10^5, n × (n - 1) / 2)` |
| Edge cost | `0 <= cost <= 10^9` |
| Maximum allowed cost (`k`) | `0 <= k <= 5 × 10^13` |
| Graph Type | Directed Acyclic Graph (DAG) |
| Node Status | Online or Offline |
| Node `0` and `n-1` | Always Online |

---

## Intuition

The first thing I noticed was that I wasn't trying to minimize the total cost. Instead, I wanted to maximize the smallest edge on the selected path.

That immediately suggested using **Binary Search on the answer**.

If I assume that the answer is some value `X`, then every edge on the chosen path must have weight at least `X`.

That means I can simply ignore every edge whose weight is smaller than `X`.

Now the problem becomes much simpler:

Can I still reach the destination while keeping the total cost within `k`?

Since the graph is already a DAG, I can process nodes in topological order and use dynamic programming to compute the minimum cost to every node.

---

## Approach

I solved the problem using the following steps.

1. Build the graph using an adjacency list.
2. Compute one topological ordering of the DAG.
3. Binary search the answer over possible edge weights.
4. For every binary search value:
   - Ignore edges whose weight is smaller than the current value.
   - Ignore offline intermediate nodes.
   - Run dynamic programming in topological order.
   - Store the minimum cost required to reach every node.
5. If the destination can be reached with cost at most `k`, try a larger answer.
6. Otherwise, search for a smaller answer.

The binary search eventually finds the maximum possible minimum edge weight.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Adjacency List | Stores all outgoing edges efficiently |
| Queue | Generates topological order using Kahn's Algorithm |
| Topological Order Array | Processes nodes in dependency order |
| DP Array | Stores the minimum recovery cost to each node |
| Binary Search | Finds the largest valid minimum edge weight |

---

## Operations & Behavior Summary

The algorithm works in several stages.

- Read all edges and build the graph.
- Compute a topological ordering once.
- Start binary searching the answer.
- For every candidate answer:
  - Skip all edges below the current threshold.
  - Skip offline intermediate nodes.
  - Traverse the graph in topological order.
  - Update the minimum cost required to reach each node.
- Check whether the destination can be reached within the given budget.
- Continue binary search until the best answer is found.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O((n + m) × log W)** | `n` is the number of nodes, `m` is the number of edges, and `W` is the maximum edge weight. Each binary search iteration scans the DAG once. |
| Space Complexity | **O(n + m)** | Used for the adjacency list, topological order, queue, and DP array. |

---

# Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int findMaxPathScore(vector<vector<int>>& edges, vector<bool>& online, long long k) {
        int n = online.size();

        // Build graph and indegree for topological sorting
        vector<vector<pair<int,int>>> graph(n);
        vector<int> indegree(n, 0);

        for (auto &e : edges) {
            graph[e[0]].push_back({e[1], e[2]});
            indegree[e[1]]++;
        }

        // Compute topological order once because the graph never changes
        queue<int> q;
        for (int i = 0; i < n; i++)
            if (indegree[i] == 0)
                q.push(i);

        vector<int> topo;
        while (!q.empty()) {
            int u = q.front();
            q.pop();
            topo.push_back(u);

            for (auto &[v, w] : graph[u]) {
                if (--indegree[v] == 0)
                    q.push(v);
            }
        }

        // Check whether a minimum edge value "limit" is possible
        auto check = [&](int limit) {
            const long long INF = (1LL << 60);

            // dp[i] = minimum total cost to reach node i
            vector<long long> dp(n, INF);
            dp[0] = 0;

            for (int u : topo) {

                // Skip unreachable nodes
                if (dp[u] == INF)
                    continue;

                // Offline intermediate nodes cannot be used
                if (u != 0 && u != n - 1 && !online[u])
                    continue;

                for (auto &[v, w] : graph[u]) {

                    // Edge is too small
                    if (w < limit)
                        continue;

                    // Cannot move into an offline intermediate node
                    if (v != n - 1 && !online[v])
                        continue;

                    if (dp[u] + w < dp[v])
                        dp[v] = dp[u] + w;
                }
            }

            return dp[n - 1] <= k;
        };

        int left = 0, right = 1000000000;
        int ans = -1;

        // Binary search on the answer
        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (check(mid)) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int findMaxPathScore(int[][] edges, boolean[] online, long k) {
        int n = online.length;

        // Build graph
        ArrayList<int[]>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++) graph[i] = new ArrayList<>();

        int[] indegree = new int[n];

        for (int[] e : edges) {
            graph[e[0]].add(new int[]{e[1], e[2]});
            indegree[e[1]]++;
        }

        // Topological order
        Queue<Integer> q = new ArrayDeque<>();
        for (int i = 0; i < n; i++)
            if (indegree[i] == 0)
                q.offer(i);

        ArrayList<Integer> topo = new ArrayList<>();

        while (!q.isEmpty()) {
            int u = q.poll();
            topo.add(u);

            for (int[] edge : graph[u]) {
                if (--indegree[edge[0]] == 0)
                    q.offer(edge[0]);
            }
        }

        int left = 0, right = 1_000_000_000;
        int ans = -1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            long INF = Long.MAX_VALUE / 4;
            long[] dp = new long[n];
            Arrays.fill(dp, INF);
            dp[0] = 0;

            for (int u : topo) {

                // Skip unreachable nodes
                if (dp[u] == INF)
                    continue;

                // Skip offline intermediate nodes
                if (u != 0 && u != n - 1 && !online[u])
                    continue;

                for (int[] edge : graph[u]) {
                    int v = edge[0];
                    int w = edge[1];

                    // Ignore small edges
                    if (w < mid)
                        continue;

                    // Cannot enter offline intermediate nodes
                    if (v != n - 1 && !online[v])
                        continue;

                    dp[v] = Math.min(dp[v], dp[u] + w);
                }
            }

            if (dp[n - 1] <= k) {
                ans = mid;
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} edges
 * @param {boolean[]} online
 * @param {number} k
 * @return {number}
 */
var findMaxPathScore = function(edges, online, k) {
    const n = online.length;

    // Build graph
    const graph = Array.from({ length: n }, () => []);
    const indegree = Array(n).fill(0);

    for (const [u, v, w] of edges) {
        graph[u].push([v, w]);
        indegree[v]++;
    }

    // Topological order
    const queue = [];
    let head = 0;

    for (let i = 0; i < n; i++)
        if (indegree[i] === 0)
            queue.push(i);

    const topo = [];

    while (head < queue.length) {
        const u = queue[head++];
        topo.push(u);

        for (const [v] of graph[u]) {
            indegree[v]--;
            if (indegree[v] === 0)
                queue.push(v);
        }
    }

    const check = (limit) => {
        const INF = Number.MAX_SAFE_INTEGER;
        const dp = Array(n).fill(INF);

        // Minimum cost to reach node 0 is zero
        dp[0] = 0;

        for (const u of topo) {

            // Skip unreachable nodes
            if (dp[u] === INF)
                continue;

            // Skip offline intermediate nodes
            if (u !== 0 && u !== n - 1 && !online[u])
                continue;

            for (const [v, w] of graph[u]) {

                // Edge is too small
                if (w < limit)
                    continue;

                // Cannot enter offline intermediate node
                if (v !== n - 1 && !online[v])
                    continue;

                dp[v] = Math.min(dp[v], dp[u] + w);
            }
        }

        return dp[n - 1] <= k;
    };

    let left = 0;
    let right = 1000000000;
    let ans = -1;

    while (left <= right) {
        const mid = Math.floor((left + right) / 2);

        if (check(mid)) {
            ans = mid;
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def findMaxPathScore(self, edges: List[List[int]], online: List[bool], k: int) -> int:
        n = len(online)

        # Build graph
        graph = [[] for _ in range(n)]
        indegree = [0] * n

        for u, v, w in edges:
            graph[u].append((v, w))
            indegree[v] += 1

        # Topological order
        from collections import deque

        q = deque()

        for i in range(n):
            if indegree[i] == 0:
                q.append(i)

        topo = []

        while q:
            u = q.popleft()
            topo.append(u)

            for v, _ in graph[u]:
                indegree[v] -= 1
                if indegree[v] == 0:
                    q.append(v)

        def check(limit):
            INF = 10 ** 30

            # Minimum cost to each node
            dp = [INF] * n
            dp[0] = 0

            for u in topo:

                # Skip unreachable nodes
                if dp[u] == INF:
                    continue

                # Skip offline intermediate nodes
                if u != 0 and u != n - 1 and not online[u]:
                    continue

                for v, w in graph[u]:

                    # Ignore small edges
                    if w < limit:
                        continue

                    # Cannot enter offline intermediate node
                    if v != n - 1 and not online[v]:
                        continue

                    if dp[u] + w < dp[v]:
                        dp[v] = dp[u] + w

            return dp[-1] <= k

        left = 0
        right = 10 ** 9
        ans = -1

        while left <= right:
            mid = (left + right) // 2

            if check(mid):
                ans = mid
                left = mid + 1
            else:
                right = mid - 1

        return ans
```

### Go

```go
func findMaxPathScore(edges [][]int, online []bool, k int64) int {
 n := len(online)

 type Edge struct {
  to int
  w  int
 }

 // Build graph
 graph := make([][]Edge, n)
 indegree := make([]int, n)

 for _, e := range edges {
  u, v, w := e[0], e[1], e[2]
  graph[u] = append(graph[u], Edge{v, w})
  indegree[v]++
 }

 // Topological order
 queue := make([]int, 0)
 for i := 0; i < n; i++ {
  if indegree[i] == 0 {
   queue = append(queue, i)
  }
 }

 topo := make([]int, 0)

 for head := 0; head < len(queue); head++ {
  u := queue[head]
  topo = append(topo, u)

  for _, e := range graph[u] {
   indegree[e.to]--
   if indegree[e.to] == 0 {
    queue = append(queue, e.to)
   }
  }
 }

 check := func(limit int) bool {
  const INF int64 = 1 << 60

  // Minimum cost to every node
  dp := make([]int64, n)
  for i := range dp {
   dp[i] = INF
  }
  dp[0] = 0

  for _, u := range topo {

   // Skip unreachable nodes
   if dp[u] == INF {
    continue
   }

   // Skip offline intermediate nodes
   if u != 0 && u != n-1 && !online[u] {
    continue
   }

   for _, e := range graph[u] {

    // Ignore small edges
    if e.w < limit {
     continue
    }

    // Cannot enter offline intermediate node
    if e.to != n-1 && !online[e.to] {
     continue
    }

    if dp[u]+int64(e.w) < dp[e.to] {
     dp[e.to] = dp[u] + int64(e.w)
    }
   }
  }

  return dp[n-1] <= k
 }

 left, right := 0, 1000000000
 ans := -1

 for left <= right {
  mid := left + (right-left)/2

  if check(mid) {
   ans = mid
   left = mid + 1
  } else {
   right = mid - 1
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The solution follows exactly the same logic in every language. Only the syntax changes.

### Step 1 — Build the Graph

The first step is creating an adjacency list.

Every node stores all of its outgoing edges along with their recovery costs.

This makes graph traversal efficient.

---

### Step 2 — Compute Topological Order

Since the graph is guaranteed to be a DAG, I compute one topological ordering using Kahn's Algorithm.

Processing nodes in topological order guarantees that every node is processed only after all of its predecessors.

This removes the need for Dijkstra's algorithm.

---

### Step 3 — Binary Search the Answer

Instead of checking every possible minimum edge value one by one, I binary search the answer.

For every middle value:

- Only keep edges whose weight is at least the current value.
- Ignore every smaller edge.

This reduces the search space from billions of possibilities to about 31 iterations.

---

### Step 4 — Dynamic Programming

I create a DP array.

`dp[i]` stores the minimum recovery cost needed to reach node `i`.

Initially:

- `dp[0] = 0`
- Every other node is unreachable.

---

### Step 5 — Process Nodes

For every node in topological order:

- Skip unreachable nodes.
- Skip offline intermediate nodes.
- Check every outgoing edge.
- Ignore edges below the binary search threshold.
- Ignore edges leading to offline intermediate nodes.
- Update the minimum recovery cost whenever a cheaper path is found.

Since every predecessor has already been processed, these updates are always correct.

---

### Step 6 — Validate the Path

After processing all nodes:

- If the destination cost is at most `k`, then the current answer is possible.
- Otherwise, it is impossible.

Binary search continues until the largest valid answer is found.

---

### Language Notes

#### C++

Uses STL containers such as `vector`, `queue`, and `pair`. The implementation is compact and efficient.

#### Java

Uses `ArrayList`, `Queue`, and primitive arrays. The logic remains identical while following Java's object-oriented syntax.

#### JavaScript

Uses arrays for both the graph and the dynamic programming table. Queue operations are implemented with an index pointer for efficiency.

#### Python3

Uses lists, tuples, and `collections.deque`. Python's clean syntax makes the dynamic programming implementation very readable.

#### Go

Uses slices and structs to represent graph edges. The implementation follows the same algorithm while taking advantage of Go's lightweight data structures.

---

## Examples

### Example 1

**Input**

```text
edges = [[0,1,5],[1,3,10],[0,2,3],[2,3,4]]
online = [true,true,true,true]
k = 10
```

**Output**

```text
3
```

**Explanation**

Two paths exist.

- `0 → 1 → 3` costs 15, so it exceeds the budget.
- `0 → 2 → 3` costs 7, which is valid.

The minimum edge on the valid path is `3`, so the answer is `3`.

---

### Example 2

**Input**

```text
edges = [[0,1,7],[1,4,5],[0,2,6],[2,3,6],[3,4,2],[2,4,6]]
online = [true,true,true,false,true]
k = 12
```

**Output**

```text
6
```

**Explanation**

Node `3` is offline, so every path through it becomes invalid.

The remaining valid path has minimum edge value `6`, which becomes the answer.

---

### Example 3

**Input**

```text
edges = []
online = [true,true]
k = 5
```

**Output**

```text
-1
```

**Explanation**

There is no path from the source to the destination.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project.

```bash
cd <repository-name>
```

### C++

Compile

```bash
g++ solution.cpp -std=c++17 -O2
```

Run

```bash
./a.out
```

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

---

## Notes & Optimizations

- Binary Search works because the answer is monotonic. If a minimum edge value is possible, then every smaller value is also possible.
- Using topological order avoids running Dijkstra's algorithm.
- Dynamic programming on a DAG is both simpler and faster for this problem.
- Offline intermediate nodes are skipped during traversal, so invalid paths are never considered.
- The solution handles disconnected graphs naturally.
- If no valid path exists, the algorithm correctly returns `-1`.
- The overall solution is efficient enough for the maximum input size provided in the constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
