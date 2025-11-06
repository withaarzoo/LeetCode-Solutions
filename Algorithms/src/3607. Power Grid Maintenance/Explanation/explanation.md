Below is a polished, copy-paste ready **README.md** you can put on GitHub for **LeetCode 3607 – Power Grid Maintenance**.
It includes my clear approach, complexity, step-by-step notes, and working code in **C++, Java, JavaScript, Python3, and Go**.

---

# Power Grid Maintenance (LeetCode 3607)

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
* [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

You are given `c` stations (1…c) connected by undirected cables. All stations start **online**.
You receive queries of two types:

1. `[1, x]` – Maintenance check for station `x`.

   * If `x` is online → answer `x`.
   * If `x` is offline → answer the **smallest-id online station** in the same connected component as `x`. If none exists → `-1`.

2. `[2, x]` – Station `x` goes **offline**.

Taking a station offline does **not** change the connectivity; offline nodes remain part of their components.

Return answers for queries of type `[1, x]` in order.

---

## Constraints

* `1 ≤ c ≤ 2 * 10^5` (typical for DSU tasks; exact limits vary but are large)
* `0 ≤ connections.length ≤ 2 * 10^5`
* `1 ≤ queries.length ≤ 2 * 10^5`
* Stations are labeled `1…c`
* Graph is static (edges never change)

---

## Intuition

I noticed that turning a station offline doesn’t break the cables. So the **connected components never change**.
That means I can:

* Build components once using **Union–Find (DSU)**.
* For each component, keep a **min-heap** of its station IDs so I can always get the **smallest online** station quickly.
* When a station goes offline, I won’t delete it immediately from the heap; instead, I’ll remove it lazily when it appears on top (called **lazy deletion**).

---

## Approach

1. **Build components with DSU** by uniting all edges.
2. **Create a min-heap per component** and push all member station IDs into it.
3. Keep a boolean `offline[i]` array (initially all `false`).
4. Process each query:

   * Type 2: set `offline[x] = true`.
   * Type 1:

     * If `x` is online → return `x`.
     * Else:

       * Find the root of `x`, clean the heap top while it points to an offline station (lazy deletion).
       * If heap becomes empty → `-1`, else heap’s top is the answer.

This works because each station can be popped from its heap at most once across all queries.

---

## Data Structures Used

* **Disjoint Set Union (Union–Find)** with path compression + union by size.
* **Min-heaps (priority queues)** per connected component.
* **Boolean array** to track offline stations.

---

## Operations & Behavior Summary

* **DSU find/union**: Build static components in near O(1) amortized.
* **Heap push**: Once per station during preprocessing.
* **Lazy deletion**: During a type-1 query for an offline `x`, pop offline IDs from the component heap until the top is online.

---

## Complexity

* **Time Complexity:**

  * DSU build: `O((c + m) α(c))`, where `m` = number of connections.
  * Heap initialization: `O(c log c)` total.
  * Queries: Each station is popped at most once → `O(q log c)` amortized.
  * **Overall:** `O((c + m) α(c) + c log c + q log c)`.
* **Space Complexity:**

  * DSU arrays + heaps store each node once + status arrays → **O(c)**.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    struct DSU {
        vector<int> p, sz;
        DSU(int n=0): p(n+1), sz(n+1,1){ iota(p.begin(), p.end(), 0); }
        int find(int x){ return p[x]==x? x : p[x]=find(p[x]); }
        void unite(int a, int b){
            a=find(a); b=find(b);
            if(a==b) return;
            if(sz[a]<sz[b]) swap(a,b);
            p[b]=a; sz[a]+=sz[b];
        }
    };

    vector<int> processQueries(int c, vector<vector<int>>& connections, vector<vector<int>>& queries) {
        DSU dsu(c);
        for (auto &e: connections) dsu.unite(e[0], e[1]);

        unordered_map<int, priority_queue<int, vector<int>, greater<int>>> heap;
        heap.reserve(c*2);
        for (int i=1;i<=c;i++) heap[dsu.find(i)].push(i);

        vector<char> offline(c+1,false);
        vector<int> ans; ans.reserve(queries.size());

        for (auto &q: queries){
            int t=q[0], x=q[1];
            if (t==2){ offline[x]=true; continue; }
            if (!offline[x]) { ans.push_back(x); continue; }

            int r=dsu.find(x);
            auto &pq = heap[r];
            while(!pq.empty() && offline[pq.top()]) pq.pop();
            ans.push_back(pq.empty()? -1 : pq.top());
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    static class DSU {
        int[] p, sz;
        DSU(int n){ p=new int[n+1]; sz=new int[n+1];
            for(int i=0;i<=n;i++){ p[i]=i; sz[i]=1; } }
        int find(int x){ return p[x]==x? x : (p[x]=find(p[x])); }
        void unite(int a,int b){
            a=find(a); b=find(b);
            if(a==b) return;
            if(sz[a]<sz[b]){ int t=a; a=b; b=t; }
            p[b]=a; sz[a]+=sz[b];
        }
    }

    public int[] processQueries(int c, int[][] connections, int[][] queries) {
        DSU dsu = new DSU(c);
        for (int[] e: connections) dsu.unite(e[0], e[1]);

        Map<Integer, PriorityQueue<Integer>> heap = new HashMap<>();
        for (int i=1;i<=c;i++) heap.computeIfAbsent(dsu.find(i), k->new PriorityQueue<>()).offer(i);

        boolean[] offline = new boolean[c+1];
        int[] ans = new int[(int)Arrays.stream(queries).filter(q->q[0]==1).count()];
        int idx=0;

        for (int[] q: queries){
            int t=q[0], x=q[1];
            if (t==2){ offline[x]=true; }
            else{
                if (!offline[x]) ans[idx++]=x;
                else{
                    int r=dsu.find(x);
                    PriorityQueue<Integer> pq = heap.get(r);
                    if (pq==null){ ans[idx++]=-1; continue; }
                    while(!pq.isEmpty() && offline[pq.peek()]) pq.poll();
                    ans[idx++]= pq.isEmpty()? -1 : pq.peek();
                }
            }
        }
        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} c
 * @param {number[][]} connections
 * @param {number[][]} queries
 * @return {number[]}
 */
var processQueries = function(c, connections, queries) {
  // DSU
  const p = Array(c+1).fill(0).map((_,i)=>i);
  const sz = Array(c+1).fill(1);
  const find = x => p[x]===x ? x : (p[x]=find(p[x]));
  const unite = (a,b) => {
    a=find(a); b=find(b);
    if (a===b) return;
    if (sz[a]<sz[b]) [a,b]=[b,a];
    p[b]=a; sz[a]+=sz[b];
  };
  for (const [u,v] of connections) unite(u,v);

  // Simple min-heap
  class MinHeap{
    constructor(){ this.a=[]; }
    size(){ return this.a.length; }
    peek(){ return this.a[0]; }
    push(x){ const a=this.a; a.push(x);
      for(let i=a.length-1;i>0;){ let j=(i-1)>>1; if (a[j]<=a[i]) break; [a[i],a[j]]=[a[j],a[i]]; i=j; } }
    pop(){ const a=this.a; if(!a.length) return;
      const top=a[0], last=a.pop(); if(a.length){ a[0]=last;
        for(let i=0;;){ let l=i*2+1, r=l+1, m=i;
          if(l<a.length && a[l]<a[m]) m=l;
          if(r<a.length && a[r]<a[m]) m=r;
          if(m===i) break; [a[i],a[m]]=[a[m],a[i]]; i=m; } }
      return top; }
  }

  const heap = new Map();
  for (let i=1;i<=c;i++){
    const r=find(i);
    if (!heap.has(r)) heap.set(r, new MinHeap());
    heap.get(r).push(i);
  }

  const offline = Array(c+1).fill(false);
  const ans = [];

  for (const [t,x] of queries){
    if (t===2){ offline[x]=true; continue; }
    if (!offline[x]) { ans.push(x); continue; }
    const r=find(x);
    const pq = heap.get(r);
    if (!pq){ ans.push(-1); continue; }
    while (pq.size() && offline[pq.peek()]) pq.pop();
    ans.push(pq.size()? pq.peek() : -1);
  }
  return ans;
};
```

### Python3

```python
from typing import List
import heapq

class Solution:
    def processQueries(self, c: int, connections: List[List[int]], queries: List[List[int]]) -> List[int]:
        # DSU
        parent = list(range(c + 1))
        size = [1] * (c + 1)

        def find(x: int) -> int:
            while parent[x] != x:
                parent[x] = parent[parent[x]]
                x = parent[x]
            return x

        def unite(a: int, b: int) -> None:
            ra, rb = find(a), find(b)
            if ra == rb: return
            if size[ra] < size[rb]: ra, rb = rb, ra
            parent[rb] = ra
            size[ra] += size[rb]

        for u, v in connections:
            unite(u, v)

        heaps = {}
        for i in range(1, c + 1):
            r = find(i)
            heaps.setdefault(r, []).append(i)
        for r in heaps:
            heapq.heapify(heaps[r])

        offline = [False] * (c + 1)
        ans = []

        for t, x in queries:
            if t == 2:
                offline[x] = True
            else:
                if not offline[x]:
                    ans.append(x)
                else:
                    r = find(x)
                    h = heaps.get(r, [])
                    while h and offline[h[0]]:
                        heapq.heappop(h)
                    ans.append(h[0] if h else -1)

        return ans
```

### Go

```go
package main

import "container/heap"

type DSU struct {
 p  []int
 sz []int
}

func NewDSU(n int) *DSU {
 p := make([]int, n+1)
 sz := make([]int, n+1)
 for i := 0; i <= n; i++ { p[i] = i; sz[i] = 1 }
 return &DSU{p, sz}
}
func (d *DSU) Find(x int) int {
 if d.p[x] != x { d.p[x] = d.Find(d.p[x]) }
 return d.p[x]
}
func (d *DSU) Unite(a, b int) {
 ra, rb := d.Find(a), d.Find(b)
 if ra == rb { return }
 if d.sz[ra] < d.sz[rb] { ra, rb = rb, ra }
 d.p[rb] = ra
 d.sz[ra] += d.sz[rb]
}

// Min-heap of ints
type IntHeap []int
func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{}   { old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x }
func (h IntHeap) Peek() int           { return h[0] }

func processQueries(c int, connections [][]int, queries [][]int) []int {
 dsu := NewDSU(c)
 for _, e := range connections { dsu.Unite(e[0], e[1]) }

 heaps := make(map[int]*IntHeap, c*2)
 for i := 1; i <= c; i++ {
  r := dsu.Find(i)
  if heaps[r] == nil { heaps[r] = &IntHeap{} }
  heap.Push(heaps[r], i)
 }
 for _, h := range heaps { heap.Init(h) }

 offline := make([]bool, c+1)
 ans := make([]int, 0, len(queries))

 for _, q := range queries {
  t, x := q[0], q[1]
  if t == 2 { offline[x] = true; continue }
  if !offline[x] { ans = append(ans, x); continue }

  r := dsu.Find(x)
  h := heaps[r]
  if h == nil { ans = append(ans, -1); continue }
  for h.Len() > 0 && offline[h.Peek()] { heap.Pop(h) }
  if h.Len() == 0 { ans = append(ans, -1) } else { ans = append(ans, h.Peek()) }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. **Union–Find setup**

   * Initialize `parent[i]=i`, `size[i]=1`.
   * For each edge `(u,v)`, call `unite(u,v)` to merge components.

2. **Build heaps**

   * For each station `i` from `1..c`, find its root `r = find(i)` and push `i` into `heap[r]`.
   * Heaps are min-heaps, so top is the smallest station id in that component.

3. **Track status**

   * `offline[i] = false` for all `i`. We flip it to `true` on query type `[2, i]`.

4. **Process queries**

   * For `[2, x]`: mark `x` offline.
   * For `[1, x]`:

     * If `x` is online → return `x`.
     * If `x` is offline:

       * Let `r = find(x)` and get the heap for `r`.
       * While heap top is offline, pop it (lazy deletion).
       * If heap is empty → `-1`, else return the top.

5. **Why lazy deletion works**

   * We never push again after initialization.
   * Each station is removed from a heap at most once when it first comes to the top while offline.
   * Hence total pops ≤ `c`.

---

## Examples

### Example 1

```
c = 5
connections = [[1,2],[2,3],[3,4],[4,5]]
queries = [[1,3],[2,1],[1,1],[1,1],[2,2],[1,2],[1,2]]

Output: [3,2,3]
Explanation:
- [1,3] -> 3 is online => 3
- [2,1] -> 1 goes offline
- [1,1] -> 1 offline, smallest online in its comp is 2 => 2
- [1,1] -> still 2
- [2,2] -> 2 goes offline
- [1,2] -> 2 offline, smallest online now 3 => 3
- [1,2] -> 3
```

### Example 2

```
c = 3
connections = []
queries = [[1,1],[2,1],[1,1]]

Output: [1,-1]
```

---

## How to use / Run locally

> These are generic run commands. Adapt to your local setup.

* **C++**

  ```bash
  g++ -std=gnu++17 solution.cpp -O2 && ./a.out
  ```

* **Java**

  ```bash
  javac Solution.java && java Solution
  ```

* **JavaScript (Node.js)**

  ```bash
  node solution.js
  ```

* **Python3**

  ```bash
  python3 solution.py
  ```

* **Go**

  ```bash
  go run solution.go
  ```

Each file should define the provided function signature used by your driver or unit tests. For LeetCode, paste the class/function body into the editor as required.

---

## Notes & Optimizations

* **Why not rebuild components when nodes go offline?** Because the problem states the grid **preserves** its structure. An offline node doesn’t break connectivity.
* **Could we store only online nodes in heaps?** We could, but that requires deletion from arbitrary positions (costly). **Lazy deletion** keeps code simple and fast.
* DSU with **path compression** and **union by size/rank** is crucial for near O(1) find/union.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)

If you use this, a ⭐ on the repo would be awesome!
