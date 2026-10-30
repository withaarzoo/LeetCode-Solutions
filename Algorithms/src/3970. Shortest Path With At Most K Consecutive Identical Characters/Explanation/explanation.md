# Shortest Path With At Most K Consecutive Identical Characters – LeetCode 3970 Solution

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

You are given a directed weighted graph with n nodes numbered from 0 to n-1. The graph is described by a list of edges, where each edge tells you the starting node, the ending node, and the weight of that edge.

Every node also has a character label. You are given a string labels of length n, so labels[i] is the character sitting on node i. You also get an integer k.

Your job is to find the path from node 0 to node n-1 that has the smallest total edge weight, but with one extra rule: when you write down the labels of the nodes you visit in order, that string must never contain more than k identical characters in a row.

If no such path exists, return -1.

This is a classic shortest-path problem on a graph with an extra constraint on consecutive characters, so normal Dijkstra is not enough by itself.

## Constraints

- 1 <= n == labels.length <= 5 * 10^4
- 0 <= edges.length <= 5 * 10^4
- edges[i] = [u_i, v_i, w_i]
- 0 <= u_i, v_i <= n - 1
- u_i != v_i
- 1 <= w_i <= 10^4
- labels consists of lowercase English letters
- 1 <= k <= 50

## Intuition

The first thing I noticed is that the constraint only cares about the last few characters of the path, not the whole string. If I am standing at a node and I already know how many times the current label has repeated at the end of the path, that is enough information to decide whether the next edge is legal.

Because the graph can have tens of thousands of nodes and edges, I cannot afford to store every possible path. Tracking only the current node and the current consecutive count keeps the state space small: n * k states. That number is manageable for Dijkstra.

So the natural idea is to run Dijkstra on states of the form (node, consecutive count) instead of just on nodes.

## Approach

I first turn the edge list into an adjacency list so I can quickly look up the outgoing edges of any node.

Then I create a two-dimensional distance array. dist[u][c] stores the cheapest cost I have found so far to reach node u while the last c characters of the path are all equal to labels[u].

I put the starting state into a priority queue: distance 0, node 0, consecutive count 1. The count starts at 1 because the path already contains the single character of node 0.

While the priority queue is not empty I pull the state with the smallest distance. If that distance is worse than the best recorded distance for the same state, I skip it. If the node is the target node n-1, I can return the distance immediately because Dijkstra always expands the cheapest path first.

For every neighbor of the current node I calculate the new consecutive count. If the labels of the two nodes are the same I add one to the old count; otherwise I reset the count to 1. When the new count is larger than k I simply ignore that edge.

If the new cost is better than the previously stored cost for that (neighbor, new-count) state, I update the distance and push the improved state into the priority queue.

If the queue empties without ever reaching node n-1, no valid path exists and I return -1.

## Data Structures Used

- Adjacency list: stores the graph so edge lookups are fast.
- Two-dimensional distance array: records the best cost for every (node, consecutive-count) pair.
- Priority queue (min-heap): always gives me the next cheapest state to process, which is required for Dijkstra.

## Operations & Behavior Summary

1. Build the adjacency list from the given edges.
2. Initialize every entry of the distance matrix to a large infinity value.
3. Set the distance of the start state (node 0, count 1) to zero and push it into the priority queue.
4. While the queue still has states:
   - Pop the state with the smallest distance.
   - Discard it if a better path to the same state was already found.
   - If the current node is the target, return the distance.
   - For each outgoing edge compute the new consecutive count.
   - Skip the edge when the new count exceeds k.
   - If a cheaper cost is discovered for the neighbor state, update the distance matrix and push the new state.
5. When the queue is empty and the target was never reached, return -1.

## Complexity

| Type              | Complexity                          | Explanation                                                                 |
|-------------------|-------------------------------------|-----------------------------------------------------------------------------|
| Time Complexity   | O((n·k + m) · log(n·k))            | n·k states, each edge is examined a constant number of times, and each heap operation costs logarithmic time. |
| Space Complexity  | O(n·k + m)                          | Distance matrix needs n·k cells; adjacency list needs space proportional to the number of edges. |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int shortestPath(int n, vector<vector<int>>& edges, string labels, int k) {
        vector<vector<pair<int,int>>> g(n);
        for (auto& e : edges) {
            g[e[0]].emplace_back(e[1], e[2]);
        }
        const int INF = 1e9;
        vector<vector<int>> dist(n, vector<int>(k + 1, INF));
        priority_queue<tuple<int,int,int>, vector<tuple<int,int,int>>, greater<>> pq;
        dist[0][1] = 0;
        pq.emplace(0, 0, 1);
        while (!pq.empty()) {
            auto [d, u, cnt] = pq.top();
            pq.pop();
            if (d > dist[u][cnt]) continue;
            if (u == n - 1) return d;
            for (auto [v, w] : g[u]) {
                int ncnt = (labels[u] == labels[v]) ? cnt + 1 : 1;
                if (ncnt > k) continue;
                if (d + w < dist[v][ncnt]) {
                    dist[v][ncnt] = d + w;
                    pq.emplace(dist[v][ncnt], v, ncnt);
                }
            }
        }
        return -1;
    }
};
```

### Java

```java
class Solution {
    public int shortestPath(int n, int[][] edges, String labels, int k) {
        List<int[]>[] g = new ArrayList[n];
        for (int i = 0; i < n; i++) g[i] = new ArrayList<>();
        for (int[] e : edges) {
            g[e[0]].add(new int[]{e[1], e[2]});
        }
        int INF = 1000000000;
        int[][] dist = new int[n][k + 1];
        for (int i = 0; i < n; i++) Arrays.fill(dist[i], INF);
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]);
        dist[0][1] = 0;
        pq.offer(new int[]{0, 0, 1});
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            int d = cur[0], u = cur[1], cnt = cur[2];
            if (d > dist[u][cnt]) continue;
            if (u == n - 1) return d;
            for (int[] edge : g[u]) {
                int v = edge[0], w = edge[1];
                int ncnt = labels.charAt(u) == labels.charAt(v) ? cnt + 1 : 1;
                if (ncnt > k) continue;
                if (d + w < dist[v][ncnt]) {
                    dist[v][ncnt] = d + w;
                    pq.offer(new int[]{dist[v][ncnt], v, ncnt});
                }
            }
        }
        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {string} labels
 * @param {number} k
 * @return {number}
 */
var shortestPath = function(n, edges, labels, k) {
    const g = Array.from({length: n}, () => []);
    for (const [u, v, w] of edges) {
        g[u].push([v, w]);
    }
    const INF = 1e9;
    const dist = Array.from({length: n}, () => Array(k + 1).fill(INF));
    class MinHeap {
        constructor() {
            this.a = [];
        }
        push(x) {
            this.a.push(x);
            this._up(this.a.length - 1);
        }
        pop() {
            const top = this.a[0];
            const last = this.a.pop();
            if (this.a.length) {
                this.a[0] = last;
                this._down(0);
            }
            return top;
        }
        get size() {
            return this.a.length;
        }
        _up(i) {
            while (i > 0) {
                const p = (i - 1) >> 1;
                if (this.a[p][0] <= this.a[i][0]) break;
                [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
                i = p;
            }
        }
        _down(i) {
            const n = this.a.length;
            while (true) {
                let s = i;
                const l = i * 2 + 1, r = l + 1;
                if (l < n && this.a[l][0] < this.a[s][0]) s = l;
                if (r < n && this.a[r][0] < this.a[s][0]) s = r;
                if (s === i) break;
                [this.a[s], this.a[i]] = [this.a[i], this.a[s]];
                i = s;
            }
        }
    }
    const pq = new MinHeap();
    dist[0][1] = 0;
    pq.push([0, 0, 1]);
    while (pq.size) {
        const [d, u, cnt] = pq.pop();
        if (d > dist[u][cnt]) continue;
        if (u === n - 1) return d;
        for (const [v, w] of g[u]) {
            const ncnt = labels[u] === labels[v] ? cnt + 1 : 1;
            if (ncnt > k) continue;
            if (d + w < dist[v][ncnt]) {
                dist[v][ncnt] = d + w;
                pq.push([dist[v][ncnt], v, ncnt]);
            }
        }
    }
    return -1;
};
```

### TypeScript

```typescript
function shortestPath(n: number, edges: number[][], labels: string, k: number): number {
    const g: [number, number][][] = Array.from({length: n}, () => []);
    for (const [u, v, w] of edges) {
        g[u].push([v, w]);
    }
    const INF = 1e9;
    const dist: number[][] = Array.from({length: n}, () => Array(k + 1).fill(INF));
    class MinHeap {
        a: [number, number, number][] = [];
        push(x: [number, number, number]) {
            this.a.push(x);
            this._up(this.a.length - 1);
        }
        pop(): [number, number, number] {
            const top = this.a[0];
            const last = this.a.pop()!;
            if (this.a.length) {
                this.a[0] = last;
                this._down(0);
            }
            return top;
        }
        get size() {
            return this.a.length;
        }
        _up(i: number) {
            while (i > 0) {
                const p = (i - 1) >> 1;
                if (this.a[p][0] <= this.a[i][0]) break;
                [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
                i = p;
            }
        }
        _down(i: number) {
            const n = this.a.length;
            while (true) {
                let s = i;
                const l = i * 2 + 1, r = l + 1;
                if (l < n && this.a[l][0] < this.a[s][0]) s = l;
                if (r < n && this.a[r][0] < this.a[s][0]) s = r;
                if (s === i) break;
                [this.a[s], this.a[i]] = [this.a[i], this.a[s]];
                i = s;
            }
        }
    }
    const pq = new MinHeap();
    dist[0][1] = 0;
    pq.push([0, 0, 1]);
    while (pq.size) {
        const [d, u, cnt] = pq.pop();
        if (d > dist[u][cnt]) continue;
        if (u === n - 1) return d;
        for (const [v, w] of g[u]) {
            const ncnt = labels[u] === labels[v] ? cnt + 1 : 1;
            if (ncnt > k) continue;
            if (d + w < dist[v][ncnt]) {
                dist[v][ncnt] = d + w;
                pq.push([dist[v][ncnt], v, ncnt]);
            }
        }
    }
    return -1;
}
```

### Python3

```python
class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        g = [[] for _ in range(n)]
        for u, v, w in edges:
            g[u].append((v, w))
        INF = 10**9
        dist = [[INF] * (k + 1) for _ in range(n)]
        pq = [(0, 0, 1)]
        dist[0][1] = 0
        while pq:
            d, u, cnt = heapq.heappop(pq)
            if d > dist[u][cnt]:
                continue
            if u == n - 1:
                return d
            for v, w in g[u]:
                ncnt = cnt + 1 if labels[u] == labels[v] else 1
                if ncnt > k:
                    continue
                if d + w < dist[v][ncnt]:
                    dist[v][ncnt] = d + w
                    heapq.heappush(pq, (dist[v][ncnt], v, ncnt))
        return -1
```

### Go

```go
type item struct {
    d, u, cnt int
}

type hp []item

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].d < h[j].d }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *hp) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
    g := make([][][2]int, n)
    for _, e := range edges {
        g[e[0]] = append(g[e[0]], [2]int{e[1], e[2]})
    }
    const INF = 1000000000
    dist := make([][]int, n)
    for i := range dist {
        dist[i] = make([]int, k+1)
        for j := range dist[i] {
            dist[i][j] = INF
        }
    }
    pq := &hp{}
    heap.Init(pq)
    dist[0][1] = 0
    heap.Push(pq, item{0, 0, 1})
    for pq.Len() > 0 {
        cur := heap.Pop(pq).(item)
        d, u, cnt := cur.d, cur.u, cur.cnt
        if d > dist[u][cnt] {
            continue
        }
        if u == n-1 {
            return d
        }
        for _, e := range g[u] {
            v, w := e[0], e[1]
            ncnt := 1
            if labels[u] == labels[v] {
                ncnt = cnt + 1
            }
            if ncnt > k {
                continue
            }
            if d+w < dist[v][ncnt] {
                dist[v][ncnt] = d + w
                heap.Push(pq, item{dist[v][ncnt], v, ncnt})
            }
        }
    }
    return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical in every language: Dijkstra on an expanded state space.

First I build the graph. In C++ and Java I use vectors or ArrayLists of pairs. In JavaScript, TypeScript and Python I use ordinary arrays of arrays. In Go I use a slice of slices of fixed-size arrays. The structure is the same; only the syntax changes.

Next I allocate the distance matrix. Every cell starts with a large constant that represents infinity. Setting the start cell dist[0][1] = 0 tells the algorithm that the empty path ending at node 0 with one occurrence of its label costs nothing.

The priority queue always stores triples (current distance, node, consecutive count). Languages that do not ship a ready-made min-heap (JavaScript, TypeScript, Go) need a small custom heap or the container/heap package. The comparison is always performed on the distance field so the cheapest state comes out first.

When a state is popped I check whether a better path to the same (node, count) pair already exists. Stale entries are common in Dijkstra when the same state is pushed several times; ignoring them keeps the algorithm correct and fast.

For every neighbor I compute the new consecutive count with a simple ternary or if-statement. The check against k is the only place the original constraint appears. Everything else is ordinary shortest-path bookkeeping.

As soon as a state whose node equals n-1 is extracted, its distance is optimal, so I return it. If the queue drains completely, the answer is -1.

Edge cases such as n = 1 are handled automatically: the start state already sits at the target and the distance is zero.

## Examples

Example 1  
Input: n = 3, edges = [[0,1,1],[1,2,1],[0,2,3]], labels = "aab", k = 1  
Output: 3  

The only legal path is the direct edge from 0 to 2. The path 0 → 1 → 2 produces the string "aab" which contains two consecutive a’s, violating k = 1. The direct edge produces "ab" and costs 3.

Example 2  
Input: same graph and labels, but k = 2  
Output: 2  

Now the path 0 → 1 → 2 is legal because two consecutive a’s are allowed. Its total weight is 2, which is better than the direct edge.

Example 3  
Input: n = 3, edges = [[0,1,1],[1,2,1]], labels = "aaa", k = 2  
Output: -1  

Every possible path produces three a’s in a row, which exceeds k = 2. No valid route exists.

## How to Use / Run Locally

C++  
g++ -std=c++17 solution.cpp -o solution && ./solution

Java  
javac Solution.java && java Solution

JavaScript  
node solution.js

TypeScript  
tsc solution.ts && node solution.js

Python3  
python3 solution.py

Go  
go run solution.go

Replace the file names with whatever you named your source files. Make sure the function signature matches the LeetCode template so you can copy-paste the code directly into the online judge.

## Notes & Optimizations

- Because k is at most 50, the extra factor of k never becomes a problem.
- The algorithm works for both sparse and moderately dense graphs; the limiting factor is usually the size of the priority queue.
- An alternative would be 0-1 BFS if all edge weights were 1, but the weights are arbitrary, so Dijkstra is the correct choice.
- Early termination when the target is first extracted is already optimal; no extra pruning is needed.
- Watch out for the case n = 1: the answer is always 0 because the single-node path has consecutive count 1 and k is at least 1.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
