# Process Tasks Using Servers - LeetCode 1882 Solution

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

You are given two arrays: `servers` and `tasks`.  
`servers[i]` tells you the weight of the i-th server.  
`tasks[j]` tells you how many seconds the j-th task needs once it starts running.

Tasks arrive one by one. The j-th task becomes available exactly at second j.  
Whenever a free server exists, you must assign the next waiting task to the free server that has the smallest weight. If two free servers have the same weight, pick the one with the smaller index.

If no server is free, you simply wait until the earliest server finishes its current work and then assign the next task immediately.

Your job is to return an array `ans` where `ans[j]` is the index of the server that finally processes task j.

This is the classic Process Tasks Using Servers problem (LeetCode 1882). The goal is to simulate the assignment correctly while staying fast enough for large inputs.

## Constraints

- `servers.length == n`
- `tasks.length == m`
- `1 <= n, m <= 2 * 10^5`
- `1 <= servers[i], tasks[j] <= 2 * 10^5`

Because both n and m can be as large as 200 000, any solution slower than roughly O((n + m) log n) will time out.

## Intuition

The first thing that stands out is the selection rule: always take the free server with the smallest weight, and break ties by the smallest index. That sounds exactly like a min-heap ordered by (weight, index).

At the same time we also need to know when the next server becomes free. That is another min-heap ordered by free time.

So the natural idea is to maintain two heaps:

- one heap that always gives us the best free server in O(log n)
- one heap that always tells us the earliest time any busy server will become free

With these two structures we can process every task in logarithmic time and still respect the exact rules of the problem.

## Approach

1. Put every server into a free-server min-heap. Each entry stores (weight, index).
2. Create an empty busy-server min-heap that will store (freeTime, weight, index).
3. Walk through the tasks from left to right. For task j (which arrives at time j):
   - Move every busy server whose freeTime is ≤ j into the free heap.
   - If the free heap is not empty, take the top server, record its index as the answer for this task, and push it into the busy heap with freeTime = j + tasks[j].
   - If the free heap is empty, take the server that becomes free earliest, assign the current task to it, and push it back with the updated freeTime.
4. After processing all tasks, return the answer array.

This simulation follows the problem statement exactly while staying efficient.

## Data Structures Used

- Free-server min-heap  
  Ordered by (weight, index). Always gives the server the problem wants us to use next.

- Busy-server min-heap  
  Ordered by (freeTime, weight, index). Lets us know the next moment a server becomes available and which server it is.

These two priority queues (heaps) are the only extra structures needed. Together they hold at most n servers at any moment.

## Operations & Behavior Summary

- Initialize the free heap with all n servers.
- For each task j from 0 to m-1:
  - Release every server that has already finished by time j.
  - If any free server exists, assign the best one and mark it busy until j + tasks[j].
  - Otherwise jump to the earliest free time, assign that server, and update its free time.
- Collect the server index chosen for every task.

The whole process is a clean event-driven simulation that never wastes time scanning all servers.

## Complexity

| Complexity       | Value              | Explanation |
|------------------|--------------------|-------------|
| Time Complexity  | O((n + m) log n)   | Each of the m tasks causes a constant number of heap operations, and each operation costs O(log n). Initializing the free heap costs O(n log n). |
| Space Complexity | O(n)               | Both heaps together store at most n servers. The answer array of size m is required by the problem. |

## Multi-language Solutions

### C++
```cpp
class Solution {
public:
    vector<int> assignTasks(vector<int>& servers, vector<int>& tasks) {
        using pii = pair<int, int>;
        using arr3 = array<long long, 3>;
        priority_queue<pii, vector<pii>, greater<pii>> free;
        priority_queue<arr3, vector<arr3>, greater<arr3>> busy;
        for (int i = 0; i < (int)servers.size(); ++i) {
            free.push({servers[i], i});
        }
        int m = tasks.size();
        vector<int> ans(m);
        for (int j = 0; j < m; ++j) {
            while (!busy.empty() && busy.top()[0] <= j) {
                auto [_, w, i] = busy.top();
                busy.pop();
                free.push({(int)w, (int)i});
            }
            if (!free.empty()) {
                auto [w, i] = free.top();
                free.pop();
                ans[j] = i;
                busy.push({(long long)j + tasks[j], w, i});
            } else {
                auto [ft, w, i] = busy.top();
                busy.pop();
                ans[j] = (int)i;
                busy.push({ft + tasks[j], w, i});
            }
        }
        return ans;
    }
};
```

### Java
```java
class Solution {
    public int[] assignTasks(int[] servers, int[] tasks) {
        PriorityQueue<int[]> free = new PriorityQueue<>((a, b) -> a[0] != b[0] ? a[0] - b[0] : a[1] - b[1]);
        for (int i = 0; i < servers.length; i++) {
            free.offer(new int[]{servers[i], i});
        }
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> {
            if (a[0] != b[0]) return Long.compare(a[0], b[0]);
            if (a[1] != b[1]) return Long.compare(a[1], b[1]);
            return Long.compare(a[2], b[2]);
        });
        int m = tasks.length;
        int[] ans = new int[m];
        for (int j = 0; j < m; j++) {
            while (!busy.isEmpty() && busy.peek()[0] <= j) {
                long[] cur = busy.poll();
                free.offer(new int[]{(int)cur[1], (int)cur[2]});
            }
            if (!free.isEmpty()) {
                int[] cur = free.poll();
                ans[j] = cur[1];
                busy.offer(new long[]{j + (long)tasks[j], cur[0], cur[1]});
            } else {
                long[] cur = busy.poll();
                ans[j] = (int)cur[2];
                busy.offer(new long[]{cur[0] + tasks[j], cur[1], cur[2]});
            }
        }
        return ans;
    }
}
```

### JavaScript
```javascript
/**
 * @param {number[]} servers
 * @param {number[]} tasks
 * @return {number[]}
 */
var assignTasks = function(servers, tasks) {
    const free = new MinPriorityQueue({ compare: (a, b) => a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1] });
    for (let i = 0; i < servers.length; i++) free.enqueue([servers[i], i]);
    const busy = new MinPriorityQueue({ compare: (a, b) => {
        if (a[0] !== b[0]) return a[0] - b[0];
        if (a[1] !== b[1]) return a[1] - b[1];
        return a[2] - b[2];
    }});
    const m = tasks.length;
    const ans = new Array(m);
    for (let j = 0; j < m; j++) {
        while (!busy.isEmpty() && busy.front()[0] <= j) {
            const cur = busy.dequeue();
            free.enqueue([cur[1], cur[2]]);
        }
        if (!free.isEmpty()) {
            const cur = free.dequeue();
            ans[j] = cur[1];
            busy.enqueue([j + tasks[j], cur[0], cur[1]]);
        } else {
            const cur = busy.dequeue();
            ans[j] = cur[2];
            busy.enqueue([cur[0] + tasks[j], cur[1], cur[2]]);
        }
    }
    return ans;
};
```

### TypeScript
```typescript
function assignTasks(servers: number[], tasks: number[]): number[] {
    const free = new MinPriorityQueue({ compare: (a: number[], b: number[]) => a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1] });
    for (let i = 0; i < servers.length; i++) free.enqueue([servers[i], i]);
    const busy = new MinPriorityQueue({ compare: (a: number[], b: number[]) => {
        if (a[0] !== b[0]) return a[0] - b[0];
        if (a[1] !== b[1]) return a[1] - b[1];
        return a[2] - b[2];
    }});
    const m = tasks.length;
    const ans = new Array(m);
    for (let j = 0; j < m; j++) {
        while (!busy.isEmpty() && busy.front()[0] <= j) {
            const cur = busy.dequeue();
            free.enqueue([cur[1], cur[2]]);
        }
        if (!free.isEmpty()) {
            const cur = free.dequeue();
            ans[j] = cur[1];
            busy.enqueue([j + tasks[j], cur[0], cur[1]]);
        } else {
            const cur = busy.dequeue();
            ans[j] = cur[2];
            busy.enqueue([cur[0] + tasks[j], cur[1], cur[2]]);
        }
    }
    return ans;
};
```

### Python3
```python
class Solution:
    def assignTasks(self, servers: list[int], tasks: list[int]) -> list[int]:
        import heapq
        free = [(w, i) for i, w in enumerate(servers)]
        heapq.heapify(free)
        busy = []
        ans = [0] * len(tasks)
        for j, t in enumerate(tasks):
            while busy and busy[0][0] <= j:
                _, w, i = heapq.heappop(busy)
                heapq.heappush(free, (w, i))
            if free:
                w, i = heapq.heappop(free)
                ans[j] = i
                heapq.heappush(busy, (j + t, w, i))
            else:
                ft, w, i = heapq.heappop(busy)
                ans[j] = i
                heapq.heappush(busy, (ft + t, w, i))
        return ans
```

### Go
```go
func assignTasks(servers []int, tasks []int) []int {
    idle := hp{}
    for i, x := range servers {
        heap.Push(&idle, pair{x, i})
    }
    busy := hp2{}
    ans := make([]int, len(tasks))
    for j, t := range tasks {
        for len(busy) > 0 && busy[0].ft <= int64(j) {
            p := heap.Pop(&busy).(bpair)
            heap.Push(&idle, pair{p.w, p.i})
        }
        if idle.Len() > 0 {
            p := heap.Pop(&idle).(pair)
            ans[j] = p.i
            heap.Push(&busy, bpair{int64(j) + int64(t), p.w, p.i})
        } else {
            p := heap.Pop(&busy).(bpair)
            ans[j] = p.i
            heap.Push(&busy, bpair{p.ft + int64(t), p.w, p.i})
        }
    }
    return ans
}

type pair struct{ w, i int }
type hp []pair
func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].w < h[j].w || (h[i].w == h[j].w && h[i].i < h[j].i) }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type bpair struct {
    ft int64
    w, i int
}
type hp2 []bpair
func (h hp2) Len() int { return len(h) }
func (h hp2) Less(i, j int) bool {
    if h[i].ft != h[j].ft {
        return h[i].ft < h[j].ft
    }
    if h[i].w != h[j].w {
        return h[i].w < h[j].w
    }
    return h[i].i < h[j].i
}
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(bpair)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, TypeScript, Python3, Go)

The core idea is identical in every language. Only the syntax for priority queues changes.

We start by filling the free heap. For every server we push the pair (weight, index). Because the heap is ordered first by weight and then by index, the top element is always the server the problem wants when several are free.

The busy heap begins empty. Its entries are (freeTime, weight, index). We store freeTime as a 64-bit integer so that j + tasks[j] never overflows.

We allocate an answer array of length m.

The main loop runs once for each task j. The arrival time of the task is simply j.

While the busy heap is not empty and the earliest freeTime is ≤ j, we pop that server and push it into the free heap. This makes every server that has already finished its previous work available again at the current moment.

After the cleanup we look at the free heap.

If it still contains at least one server we pop the best one, write its index into ans[j], and push the same server into the busy heap with the new freeTime j + tasks[j].

If the free heap is empty, every server is still working at time j. We therefore pop the server that becomes free earliest. Because the busy heap is ordered by freeTime, then weight, then index, this is exactly the server the problem requires. We record its index, compute the new freeTime as old freeTime + tasks[j], and push it back.

When the loop finishes we return the answer array. Every task has been assigned exactly once and the heaps always kept the servers ordered the way the problem demands.

The only real differences across languages are:

- C++ uses `priority_queue` with a custom greater comparator and `array` or `tuple`.
- Java uses `PriorityQueue` with lambda comparators.
- JavaScript and TypeScript rely on a MinPriorityQueue utility (or a custom heap implementation).
- Python uses the built-in `heapq` module.
- Go requires explicit heap interfaces (`container/heap`) with custom types.

All of them implement the same two-heap logic and achieve the same asymptotic performance.

## Examples

**Example 1**

Input:  
servers = [3, 3, 2]  
tasks = [1, 2, 3, 2, 1, 2]

Output:  
[2, 2, 0, 2, 1, 2]

Trace:  
- Time 0: task 0 goes to server 2 (lightest).  
- Time 1: server 2 becomes free, takes task 1.  
- Time 2: server 2 is busy, so task 2 goes to server 0.  
- Time 3: server 2 frees again and takes task 3.  
- Time 4: task 4 goes to server 1.  
- Time 5: all servers free, task 5 goes back to server 2.

**Example 2**

Input:  
servers = [5, 1, 4, 3, 2]  
tasks = [2, 1, 2, 4, 5, 2, 1]

Output:  
[1, 4, 1, 4, 1, 3, 2]

The same two-heap process correctly selects the lightest free server at every step, jumping forward in time only when necessary.

## How to Use / Run Locally

1. Copy the solution for the language you want into a file (for example `main.cpp`, `Main.java`, `solution.js`, etc.).
2. Make sure you have a recent compiler or runtime installed:
   - C++: g++ or clang++
   - Java: JDK 11+
   - JavaScript / TypeScript: Node.js
   - Python: Python 3.8+
   - Go: Go 1.18+
3. Compile and run with a small test:
   - C++: `g++ -std=c++17 main.cpp -o main && ./main`
   - Java: `javac Main.java && java Main`
   - JavaScript: `node solution.js`
   - TypeScript: `tsc solution.ts && node solution.js`
   - Python: `python3 solution.py`
   - Go: `go run solution.go`
4. Replace the hard-coded test arrays with any input you want to verify.

## Notes & Optimizations

- Using 64-bit integers for freeTime is important. With m and tasks[j] both up to 2e5, the sum can exceed 32-bit range.
- The solution never advances time second-by-second. It always jumps to the next relevant event, which keeps it fast.
- An alternative approach that stores only freeTime and index in the busy heap and looks the weight up in an array also works, but the three-field version is clearer and still O((n + m) log n).
- Edge cases worth checking: all servers have the same weight, a single server, tasks that finish exactly when a new task arrives, and the maximum size of 2e5.

This two-heaps method is currently one of the cleanest and fastest ways to solve the Process Tasks Using Servers problem on LeetCode.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
