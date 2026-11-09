# Find Servers That Handled Most Number of Requests - LeetCode 1606 Solution

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

You are given `k` servers numbered from `0` to `k-1`. Each server can handle only one request at a time, but it has unlimited capacity otherwise.

Requests arrive at strictly increasing times given in the array `arrival`. Each request also has a processing time given in the array `load`.

When the `i`-th request arrives, the system tries to assign it to server `i % k`.  
If that server is busy, the system looks for the next free server in order (wrapping around to `0` if needed).  
If every server is busy, the request is dropped.

Your task is to return the list of server IDs that handled the largest number of requests. If several servers share the same maximum count, return all of them in any order.

This is a classic load-balancing simulation problem that appears frequently in interviews and competitive programming contests.

## Constraints

- `1 <= k <= 10^5`
- `1 <= arrival.length, load.length <= 10^5`
- `arrival.length == load.length`
- `1 <= arrival[i], load[i] <= 10^9`
- `arrival` is strictly increasing

Because both `k` and the number of requests can be as large as `10^5`, any solution that scans all servers for every request will time out. An efficient O(n log k) approach is required.

## Intuition

The first thing I noticed is that a naive loop checking every server for every request is too slow under the given constraints.  

The preferred server is always `i % k`, and I only need the next free server after that point (or the smallest free server if I have to wrap around).  

If I can keep the free servers in a sorted structure, I can jump straight to the correct server in logarithmic time. At the same time I need a way to know when busy servers become free again. A min-heap ordered by finish time solves that part cleanly.

Putting these two ideas together gives a clean and fast simulation.

## Approach

1. Put every server into a free collection at the start.
2. Maintain a min-heap of busy servers. Each entry stores the time the server will finish its current job and the server id.
3. For every incoming request:
   - Free all servers whose finish time is less than or equal to the current arrival time and put them back into the free collection.
   - If no free server exists, drop the request.
   - Otherwise find the smallest free server whose id is greater than or equal to the preferred id (`i % k`). If none exists, take the overall smallest free server (wrap-around).
   - Assign the request to that server, increase its request counter, remove it from the free collection, and push it onto the busy heap with the new finish time.
4. After processing all requests, find the maximum counter value and collect every server that reached that value.

In languages without a built-in ordered set I replace the ordered set by a carefully maintained min-heap of “virtual” indices. When a server is freed at request `i` I push the value `i + (server - i) % k`. This keeps the invariant that every value still in the heap lies between `i` and `i + k - 1`, so the smallest value modulo `k` is exactly the next free server I need.

## Data Structures Used

- **Ordered set / Sorted list / TreeSet** – stores currently free server ids so I can find the next suitable server in O(log k).
- **Min-heap (Priority Queue)** – stores busy servers ordered by their finish time so I can free them efficiently.
- **Count array** – simple integer array of size `k` that records how many requests each server handled.
- **Virtual-index min-heap** (used in JavaScript, TypeScript and Go) – replaces the ordered set while preserving the same asymptotic complexity.

## Operations & Behavior Summary

- Initialize free servers and empty busy heap.
- For each request in arrival order:
  - Release finished servers from the busy heap into the free collection.
  - If free collection is empty → drop request.
  - Locate the preferred or next free server.
  - Assign the request, update count, move server to busy heap.
- Scan the count array and collect all servers that share the maximum load.

## Complexity

| Complexity       | Value          | Explanation |
|------------------|----------------|-------------|
| Time Complexity  | O(n log k)     | Each of the n requests performs a constant number of ordered-set and heap operations, each costing O(log k). |
| Space Complexity | O(k)           | Free collection, busy heap and count array all hold at most k elements. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<int> busiestServers(int k, vector<int>& arrival, vector<int>& load) {
        set<int> free;
        for (int i = 0; i < k; ++i) free.insert(i);
        priority_queue<pair<long long, int>, vector<pair<long long, int>>, greater<>> busy;
        vector<int> cnt(k);
        int n = arrival.size();
        for (int i = 0; i < n; ++i) {
            long long t = arrival[i];
            while (!busy.empty() && busy.top().first <= t) {
                free.insert(busy.top().second);
                busy.pop();
            }
            if (free.empty()) continue;
            int pref = i % k;
            auto it = free.lower_bound(pref);
            if (it == free.end()) it = free.begin();
            int s = *it;
            free.erase(it);
            ++cnt[s];
            busy.emplace(t + load[i], s);
        }
        int mx = *max_element(cnt.begin(), cnt.end());
        vector<int> ans;
        for (int i = 0; i < k; ++i)
            if (cnt[i] == mx) ans.push_back(i);
        return ans;
    }
};
```

### Java
```java
class Solution {
    public List<Integer> busiestServers(int k, int[] arrival, int[] load) {
        TreeSet<Integer> free = new TreeSet<>();
        for (int i = 0; i < k; i++) free.add(i);
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> Long.compare(a[0], b[0]));
        int[] cnt = new int[k];
        int n = arrival.length;
        for (int i = 0; i < n; i++) {
            long t = arrival[i];
            while (!busy.isEmpty() && busy.peek()[0] <= t) {
                free.add((int) busy.poll()[1]);
            }
            if (free.isEmpty()) continue;
            int pref = i % k;
            Integer s = free.ceiling(pref);
            if (s == null) s = free.first();
            free.remove(s);
            cnt[s]++;
            busy.offer(new long[]{t + load[i], s});
        }
        int mx = 0;
        for (int c : cnt) mx = Math.max(mx, c);
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < k; i++)
            if (cnt[i] == mx) ans.add(i);
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number} k
 * @param {number[]} arrival
 * @param {number[]} load
 * @return {number[]}
 */
var busiestServers = function(k, arrival, load) {
    class MinHeap {
        constructor() { this.a = []; }
        size() { return this.a.length; }
        peek() { return this.a[0]; }
        push(x) {
            this.a.push(x);
            let i = this.a.length - 1;
            while (i > 0) {
                let p = (i - 1) >> 1;
                if (this.a[p] <= this.a[i]) break;
                [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
                i = p;
            }
        }
        pop() {
            const top = this.a[0];
            const last = this.a.pop();
            if (this.a.length === 0) return top;
            this.a[0] = last;
            let i = 0;
            while (true) {
                let l = i * 2 + 1, r = l + 1, smallest = i;
                if (l < this.a.length && this.a[l] < this.a[smallest]) smallest = l;
                if (r < this.a.length && this.a[r] < this.a[smallest]) smallest = r;
                if (smallest === i) break;
                [this.a[i], this.a[smallest]] = [this.a[smallest], this.a[i]];
                i = smallest;
            }
            return top;
        }
    }
    const available = new MinHeap();
    for (let i = 0; i < k; i++) available.push(i);
    const busy = new MinHeap(); // stores endTime * 1e6 + server ( suffices because endTime <= 2e9 )
    const cnt = new Array(k).fill(0);
    const n = arrival.length;
    for (let i = 0; i < n; i++) {
        const t = arrival[i];
        while (busy.size() && Math.floor(busy.peek() / 1000000) <= t) {
            const x = busy.pop() % 1000000;
            available.push(i + ((x - i) % k + k) % k);
        }
        if (available.size() === 0) continue;
        const j = available.pop() % k;
        cnt[j]++;
        busy.push((t + load[i]) * 1000000 + j);
    }
    let mx = 0;
    for (const c of cnt) if (c > mx) mx = c;
    const ans = [];
    for (let i = 0; i < k; i++) if (cnt[i] === mx) ans.push(i);
    return ans;
};
```

### TypeScript
```typescript
function busiestServers(k: number, arrival: number[], load: number[]): number[] {
    class MinHeap {
        a: number[] = [];
        size() { return this.a.length; }
        peek() { return this.a[0]; }
        push(x: number) {
            this.a.push(x);
            let i = this.a.length - 1;
            while (i > 0) {
                let p = (i - 1) >> 1;
                if (this.a[p] <= this.a[i]) break;
                [this.a[p], this.a[i]] = [this.a[i], this.a[p]];
                i = p;
            }
        }
        pop(): number {
            const top = this.a[0];
            const last = this.a.pop()!;
            if (this.a.length === 0) return top;
            this.a[0] = last;
            let i = 0;
            while (true) {
                let l = i * 2 + 1, r = l + 1, smallest = i;
                if (l < this.a.length && this.a[l] < this.a[smallest]) smallest = l;
                if (r < this.a.length && this.a[r] < this.a[smallest]) smallest = r;
                if (smallest === i) break;
                [this.a[i], this.a[smallest]] = [this.a[smallest], this.a[i]];
                i = smallest;
            }
            return top;
        }
    }
    const available = new MinHeap();
    for (let i = 0; i < k; i++) available.push(i);
    const busy = new MinHeap();
    const cnt = new Array(k).fill(0);
    const n = arrival.length;
    for (let i = 0; i < n; i++) {
        const t = arrival[i];
        while (busy.size() && Math.floor(busy.peek() / 1000000) <= t) {
            const x = busy.pop() % 1000000;
            available.push(i + ((x - i) % k + k) % k);
        }
        if (available.size() === 0) continue;
        const j = available.pop() % k;
        cnt[j]++;
        busy.push((t + load[i]) * 1000000 + j);
    }
    let mx = 0;
    for (const c of cnt) if (c > mx) mx = c;
    const ans: number[] = [];
    for (let i = 0; i < k; i++) if (cnt[i] === mx) ans.push(i);
    return ans;
}
```

### Python3
```python
from sortedcontainers import SortedList
import heapq

class Solution:
    def busiestServers(self, k: int, arrival: list[int], load: list[int]) -> list[int]:
        free = SortedList(range(k))
        busy = []
        cnt = [0] * k
        n = len(arrival)
        for i in range(n):
            t = arrival[i]
            while busy and busy[0][0] <= t:
                free.add(heapq.heappop(busy)[1])
            if not free:
                continue
            pref = i % k
            idx = free.bisect_left(pref)
            if idx == len(free):
                idx = 0
            s = free[idx]
            free.pop(idx)
            cnt[s] += 1
            heapq.heappush(busy, (t + load[i], s))
        mx = max(cnt)
        return [i for i in range(k) if cnt[i] == mx]
```

### Go
```go
func busiestServers(k int, arrival []int, load []int) []int {
    available := &IntHeap{}
    heap.Init(available)
    for i := 0; i < k; i++ {
        heap.Push(available, i)
    }
    busy := &PairHeap{}
    heap.Init(busy)
    cnt := make([]int, k)
    n := len(arrival)
    for i := 0; i < n; i++ {
        t := arrival[i]
        for busy.Len() > 0 && (*busy)[0].end <= t {
            x := heap.Pop(busy).(Pair).server
            heap.Push(available, i+((x-i)%k+k)%k)
        }
        if available.Len() == 0 {
            continue
        }
        j := heap.Pop(available).(int) % k
        cnt[j]++
        heap.Push(busy, Pair{t + load[i], j})
    }
    mx := 0
    for _, c := range cnt {
        if c > mx {
            mx = c
        }
    }
    ans := []int{}
    for i := 0; i < k; i++ {
        if cnt[i] == mx {
            ans = append(ans, i)
        }
    }
    return ans
}

type IntHeap []int
func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type Pair struct{ end, server int }
type PairHeap []Pair
func (h PairHeap) Len() int            { return len(h) }
func (h PairHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h PairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PairHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *PairHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

**Common logic across all languages**

I begin by inserting every server id from `0` to `k-1` into the free structure. At time zero nothing is busy.

I also prepare an empty min-heap that will store pairs `(finish time, server id)` and a count array of size `k` filled with zeros.

For every request `i` I look at its arrival time `t`. While the busy heap is not empty and the earliest finish time is still ≤ `t`, I pop that server and put it back into the free structure. Those servers have just become available again.

If the free structure is empty after that step, every server is still occupied, so I drop the request and move on.

Otherwise I need the next free server that is at least `i % k` (or the smallest free server if none exists). In languages that provide an ordered set I simply ask for the lower-bound / ceiling of the preferred id and fall back to the first element when necessary. In the pure-heap version the numbers that sit in the available heap already satisfy the same ordering invariant, so I just pop the minimum and take it modulo `k`.

I remove the chosen server from the free structure, increase its counter by one, and push the pair `(t + load[i], server)` onto the busy heap so I know when it will be free again.

After the last request I walk through the count array once, remember the highest value, and collect every index that holds that value. Those indices are the answer.

**Language-specific notes**

- C++ and Java use `std::set` / `TreeSet` for free servers and a standard priority queue for busy servers.  
- Python uses `SortedList` from the sortedcontainers library together with `heapq`.  
- JavaScript, TypeScript and Go implement the same algorithm with two pure heaps (the virtual-index trick) because they lack a built-in ordered set. A small custom min-heap class is written for JavaScript and TypeScript; Go uses the standard `container/heap` package.

## Examples

**Example 1**  
Input: `k = 3`, `arrival = [1,2,3,4,5]`, `load = [5,2,3,3,3]`  
Output: `[1]`  

Trace:  
- Request 0 → server 0 (free)  
- Request 1 → server 1 (free)  
- Request 2 → server 2 (free)  
- Request 3 prefers server 0 (busy) → next free is server 1  
- Request 4 finds all servers busy → dropped  
Server 1 handled two requests, the others handled one → answer `[1]`.

**Example 2**  
Input: `k = 3`, `arrival = [1,2,3,4]`, `load = [1,2,1,2]`  
Output: `[0]`  

Trace:  
- First three requests go to servers 0, 1, 2.  
- Fourth request prefers server 0, which is already free again → assigned to 0.  
Server 0 handled two requests → answer `[0]`.

**Example 3**  
Input: `k = 3`, `arrival = [1,2,3]`, `load = [10,12,11]`  
Output: `[0,1,2]`  

Each server receives exactly one request and all three finish at different times, so they are tied for the maximum.

## How to Use / Run Locally

1. Copy the code for the language you want into a file named according to the language convention (`main.cpp`, `Solution.java`, `solution.js`, etc.).
2. Compile / run with the standard tools:

- **C++**  
  ```bash
  g++ -std=c++17 -O2 main.cpp -o main && ./main
  ```

- **Java**  
  ```bash
  javac Solution.java && java Solution
  ```

- **JavaScript**  
  ```bash
  node solution.js
  ```

- **TypeScript**  
  ```bash
  tsc solution.ts && node solution.js
  ```

- **Python3**  
  ```bash
  python3 solution.py
  ```

- **Go**  
  ```bash
  go run main.go
  ```

Add a simple `main` function or test harness that calls `busiestServers` with the sample inputs if you want to verify the output yourself.

## Notes & Optimizations

- The pure-heap formulation (used in JavaScript, TypeScript and Go) is necessary because those languages do not provide a balanced ordered set out of the box. It keeps the same O(n log k) complexity.
- Watch out for integer overflow when computing finish times (`arrival[i] + load[i]`). Use 64-bit integers where needed.
- When several servers share the maximum load, any order of the returned ids is accepted by the judge.
- An alternative approach using a segment tree that tracks free status is also possible, but the ordered-set + heap solution is shorter and easier to implement correctly under contest time pressure.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
