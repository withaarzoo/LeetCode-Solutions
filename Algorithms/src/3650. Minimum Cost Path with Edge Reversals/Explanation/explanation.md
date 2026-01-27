# Minimum Cost Path with Edge Reversals

**LeetCode Problem 3650**

---

## Table of Contents

* Problem Title
* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a **directed weighted graph** with `n` nodes.

Each edge has:

* a direction
* a cost

Special rule:

* Every node has a **switch**
* I can use the switch **at most once per node**
* When I reach a node, I can **reverse one incoming edge**
* That reversed edge can be used **immediately**
* Reversing an edge costs **2 × original cost**
* The reversal is valid for **only one move**

My goal is to find the **minimum cost** to go from node `0` to node `n - 1`.

If it’s not possible, I return `-1`.

---

## Constraints

* `2 ≤ n ≤ 5 × 10⁴`
* `1 ≤ edges.length ≤ 10⁵`
* `edges[i] = [ui, vi, wi]`
* `0 ≤ ui, vi < n`
* `1 ≤ wi ≤ 1000`

---

## Intuition

When I first saw the problem, I realized this is clearly a **shortest path problem**.

Normally, I would directly use **Dijkstra**.

The confusing part was the **edge reversal rule**.
But then I understood something important:

> Reversing an edge is just another valid move with a different cost.

So instead of tracking switches and states,
I decided to **pre-add reversed edges** into the graph with cost `2 × w`.

Once I do that, the problem becomes a **normal shortest path problem**.

That insight made the solution very clean.

---

## Approach

I solved the problem using these steps:

1. I created an adjacency list.
2. For every original edge `u → v` with cost `w`:

   * I added `u → v` with cost `w`
   * I added `v → u` with cost `2 × w` (reversed edge)
3. Now all valid movements exist in the graph.
4. I ran **Dijkstra’s algorithm** from node `0`.
5. If node `n - 1` is reachable, I returned its distance.
6. Otherwise, I returned `-1`.

No extra state.
No switch tracking.
Just clean graph logic.

---

## Data Structures Used

* Adjacency List
* Min Heap / Priority Queue
* Distance Array

---

## Operations & Behavior Summary

* Original edges behave normally
* Reversed edges cost double
* Each reversal is naturally limited by Dijkstra
* Shortest path always wins
* If destination is unreachable, return `-1`

---

## Complexity

**Time Complexity:**
`O((n + m) log n)`

* `n` = number of nodes
* `m` = number of edges

**Space Complexity:**
`O(n + m)`

* Graph storage
* Distance array
* Priority queue

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minCost(int n, vector<vector<int>>& edges) {
        vector<vector<pair<int,int>>> graph(n);

        for (auto &e : edges) {
            int u = e[0], v = e[1], w = e[2];
            graph[u].push_back({v, w});
            graph[v].push_back({u, 2 * w});
        }

        vector<long long> dist(n, LLONG_MAX);
        priority_queue<pair<long long,int>, vector<pair<long long,int>>, greater<>> pq;

        dist[0] = 0;
        pq.push({0, 0});

        while (!pq.empty()) {
            auto [cost, node] = pq.top();
            pq.pop();

            if (cost > dist[node]) continue;

            for (auto &[next, w] : graph[node]) {
                if (dist[next] > cost + w) {
                    dist[next] = cost + w;
                    pq.push({dist[next], next});
                }
            }
        }

        return dist[n - 1] == LLONG_MAX ? -1 : dist[n - 1];
    }
};
```

---

### Java

```java
class Solution {
    public int minCost(int n, int[][] edges) {
        List<int[]>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++) graph[i] = new ArrayList<>();

        for (int[] e : edges) {
            graph[e[0]].add(new int[]{e[1], e[2]});
            graph[e[1]].add(new int[]{e[0], 2 * e[2]});
        }

        long[] dist = new long[n];
        Arrays.fill(dist, Long.MAX_VALUE);
        dist[0] = 0;

        PriorityQueue<long[]> pq =
            new PriorityQueue<>(Comparator.comparingLong(a -> a[0]));
        pq.offer(new long[]{0, 0});

        while (!pq.isEmpty()) {
            long[] cur = pq.poll();
            long cost = cur[0];
            int node = (int) cur[1];

            if (cost > dist[node]) continue;

            for (int[] nxt : graph[node]) {
                if (dist[nxt[0]] > cost + nxt[1]) {
                    dist[nxt[0]] = cost + nxt[1];
                    pq.offer(new long[]{dist[nxt[0]], nxt[0]});
                }
            }
        }

        return dist[n - 1] == Long.MAX_VALUE ? -1 : (int) dist[n - 1];
    }
}
```

---

### JavaScript

```javascript
var minCost = function(n, edges) {
    const graph = Array.from({ length: n }, () => []);

    for (const [u, v, w] of edges) {
        graph[u].push([v, w]);
        graph[v].push([u, 2 * w]);
    }

    const dist = Array(n).fill(Infinity);
    dist[0] = 0;
    const pq = [[0, 0]];

    while (pq.length) {
        pq.sort((a, b) => a[0] - b[0]);
        const [cost, node] = pq.shift();

        if (cost > dist[node]) continue;

        for (const [next, w] of graph[node]) {
            if (dist[next] > cost + w) {
                dist[next] = cost + w;
                pq.push([dist[next], next]);
            }
        }
    }

    return dist[n - 1] === Infinity ? -1 : dist[n - 1];
};
```

---

### Python3

```python
import heapq

class Solution:
    def minCost(self, n, edges):
        graph = [[] for _ in range(n)]

        for u, v, w in edges:
            graph[u].append((v, w))
            graph[v].append((u, 2 * w))

        dist = [float('inf')] * n
        dist[0] = 0
        pq = [(0, 0)]

        while pq:
            cost, node = heapq.heappop(pq)
            if cost > dist[node]:
                continue

            for nxt, w in graph[node]:
                if dist[nxt] > cost + w:
                    dist[nxt] = cost + w
                    heapq.heappush(pq, (dist[nxt], nxt))

        return -1 if dist[n - 1] == float('inf') else dist[n - 1]
```

---

### Go

```go
func minCost(n int, edges [][]int) int {
    graph := make([][][2]int, n)

    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        graph[u] = append(graph[u], [2]int{v, w})
        graph[v] = append(graph[v], [2]int{u, 2 * w})
    }

    const INF = int(1e18)
    dist := make([]int, n)
    for i := range dist {
        dist[i] = INF
    }
    dist[0] = 0

    pq := [][]int{{0, 0}}

    for len(pq) > 0 {
        minIdx := 0
        for i := range pq {
            if pq[i][0] < pq[minIdx][0] {
                minIdx = i
            }
        }

        cur := pq[minIdx]
        pq = append(pq[:minIdx], pq[minIdx+1:]...)
        cost, node := cur[0], cur[1]

        if cost > dist[node] {
            continue
        }

        for _, e := range graph[node] {
            next, w := e[0], e[1]
            if dist[next] > cost+w {
                dist[next] = cost + w
                pq = append(pq, []int{dist[next], next})
            }
        }
    }

    if dist[n-1] == INF {
        return -1
    }
    return dist[n-1]
}
```

---

## Step-by-step Detailed Explanation

1. I converted every reversal into a real edge.
2. Reversed edges always cost double.
3. The graph now contains all valid paths.
4. Dijkstra always picks the cheapest next move.
5. This naturally respects the “use once” rule.
6. When I reach the destination, the cost is guaranteed minimum.

---

## Examples

**Example 1**

```bash
Input:
n = 4
edges = [[0,1,3],[3,1,1],[2,3,4],[0,2,2]]

Output:
5
```

**Example 2**

```bash
Input:
n = 4
edges = [[0,2,1],[2,1,1],[1,3,1],[2,3,3]]

Output:
3
```

---

## How to use / Run locally

1. Copy the solution code
2. Paste into your local environment
3. Compile or run based on the language
4. Provide input through function calls (LeetCode style)

---

## Notes & Optimizations

* No state tracking needed
* Clean Dijkstra solution
* Works within all constraints
* Interview-friendly and scalable

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
