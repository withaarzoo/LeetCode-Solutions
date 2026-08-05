# 3310. Remove Methods From Project

A clean and beginner-friendly solution for **LeetCode 3310 - Remove Methods From Project** using **Graph Theory** and **Depth-First Search (DFS)**. This repository explains the intuition, approach, complexity analysis, and provides multi-language solutions for C++, Java, JavaScript, Python3, and Go.

---

## Table of Contents

- [Problem Summary](#problem-summary)
- [Constraints](#constraints)
- [Intuition](#intuition)
- [Approach](#approach)
- [Data Structures Used](#data-structures-used)
- [Operations & Behavior Summary](#operations--behavior-summary)
- [Complexity](#complexity)
- [Multi-language Solutions](#multi-language-solutions)
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

You are given a project containing `n` methods numbered from `0` to `n - 1`. Every invocation tells you that one method calls another method.

A bug is found in method `k`. Because of that, method `k` and every method reachable from it through direct or indirect invocations become **suspicious**.

The goal is to remove every suspicious method, but there is one important condition.

If even one method outside the suspicious group calls a suspicious method, then removing them would break the project. In that case, no method should be removed.

Your task is to return the list of methods that remain after applying these rules.

This problem is a great example of using **Graph Theory**, **DFS Traversal**, and **Reachability** to solve dependency-related problems efficiently.

---

## Constraints

| Constraint | Value |
| ------------ | ------- |
| `1 <= n <= 10^5` | Number of methods |
| `0 <= k <= n - 1` | Buggy method |
| `0 <= invocations.length <= 2 × 10^5` | Number of edges |
| `0 <= ai, bi <= n - 1` | Valid method numbers |
| `ai != bi` | No self-loop |
| All invocation pairs are unique | Yes |

---

## Intuition

The first thing I noticed was that finding suspicious methods is actually straightforward.

If a method can be reached from the buggy method `k`, then it is suspicious. That immediately suggested using DFS because DFS naturally visits every reachable node in a graph.

The real challenge comes after finding those methods.

Even if I know which methods are suspicious, I still have to check whether removing them is actually allowed. The problem clearly says that if a safe method still depends on a suspicious one, then the entire removal becomes invalid.

So instead of making the problem more complicated, I split it into two simple tasks.

First, find every suspicious method.

Second, verify whether any non-suspicious method points into that suspicious group.

If that never happens, removing the suspicious methods is completely safe.

---

## Approach

I start by building a directed graph using the given invocations.

Next, I perform a Depth-First Search starting from method `k`.

Every method visited during DFS is marked as suspicious.

After that, I iterate through every invocation again.

For each edge, I check whether it starts from a non-suspicious method and ends at a suspicious method.

If I ever find such an edge, I immediately know that removal is impossible, so I return every method in the project.

Otherwise, I simply collect every method that was never marked suspicious and return them.

The entire solution only scans the graph a couple of times, making it efficient even for the largest constraints.

---

## Data Structures Used

### Adjacency List

I use an adjacency list to represent the directed graph because it stores outgoing edges efficiently and makes DFS very fast.

### Visited Array

A boolean array keeps track of whether a method has already been visited during DFS.

It also serves as the list of suspicious methods.

### Result Array

At the end, I store every remaining method inside a result array before returning it.

---

## Operations & Behavior Summary

1. Read every invocation.
2. Build a directed graph.
3. Start DFS from the buggy method.
4. Mark every reachable method as suspicious.
5. Scan every invocation once more.
6. If a safe method calls a suspicious method, stop immediately and return all methods.
7. Otherwise, collect only the non-suspicious methods.
8. Return the final answer.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n + m)** | Every method and every invocation is processed at most once. |
| Space Complexity | **O(n + m)** | Extra space is used for the adjacency list, visited array, and DFS recursion stack. |

Where:

- `n` = Number of methods
- `m` = Number of invocations

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> remainingMethods(int n, int k, vector<vector<int>>& invocations) {
        // Build the directed graph
        vector<vector<int>> graph(n);
        for (auto &edge : invocations) {
            graph[edge[0]].push_back(edge[1]);
        }

        // Marks whether a method is suspicious
        vector<bool> vis(n, false);

        // DFS to mark every method reachable from k
        function<void(int)> dfs = [&](int u) {
            vis[u] = true;

            // Visit every invoked method
            for (int v : graph[u]) {
                if (!vis[v]) {
                    dfs(v);
                }
            }
        };

        dfs(k);

        // Check whether any non-suspicious method calls a suspicious one
        for (auto &edge : invocations) {
            int u = edge[0];
            int v = edge[1];

            if (!vis[u] && vis[v]) {
                // Removal is impossible, return all methods
                vector<int> ans;
                for (int i = 0; i < n; i++) {
                    ans.push_back(i);
                }
                return ans;
            }
        }

        // Keep only non-suspicious methods
        vector<int> ans;
        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                ans.push_back(i);
            }
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public List<Integer> remainingMethods(int n, int k, int[][] invocations) {

        // Build the adjacency list
        List<Integer>[] graph = new ArrayList[n];
        for (int i = 0; i < n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] edge : invocations) {
            graph[edge[0]].add(edge[1]);
        }

        // Marks suspicious methods
        boolean[] vis = new boolean[n];

        // DFS from method k
        dfs(k, graph, vis);

        // If a safe method invokes a suspicious one,
        // removal is not allowed
        for (int[] edge : invocations) {
            int u = edge[0];
            int v = edge[1];

            if (!vis[u] && vis[v]) {
                List<Integer> ans = new ArrayList<>();
                for (int i = 0; i < n; i++) {
                    ans.add(i);
                }
                return ans;
            }
        }

        // Return remaining methods
        List<Integer> ans = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (!vis[i]) {
                ans.add(i);
            }
        }

        return ans;
    }

    // DFS to mark all reachable methods
    private void dfs(int u, List<Integer>[] graph, boolean[] vis) {
        vis[u] = true;

        for (int v : graph[u]) {
            if (!vis[v]) {
                dfs(v, graph, vis);
            }
        }
    }
}
```

### JavaScript

```javascript
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
```

### Python3

```python
class Solution:
    def remainingMethods(self, n: int, k: int, invocations: List[List[int]]) -> List[int]:

        # Build the graph
        graph = [[] for _ in range(n)]

        for u, v in invocations:
            graph[u].append(v)

        # Marks suspicious methods
        vis = [False] * n

        # DFS to visit every reachable method
        def dfs(u):
            vis[u] = True

            for v in graph[u]:
                if not vis[v]:
                    dfs(v)

        dfs(k)

        # If any safe method calls a suspicious method,
        # removal is impossible
        for u, v in invocations:
            if not vis[u] and vis[v]:
                return list(range(n))

        # Return only non-suspicious methods
        ans = []

        for i in range(n):
            if not vis[i]:
                ans.append(i)

        return ans
```

### Go

```go
func remainingMethods(n int, k int, invocations [][]int) []int {

 // Build adjacency list
 graph := make([][]int, n)

 for _, edge := range invocations {
  u := edge[0]
  v := edge[1]
  graph[u] = append(graph[u], v)
 }

 // Marks suspicious methods
 vis := make([]bool, n)

 // DFS from method k
 var dfs func(int)

 dfs = func(u int) {
  vis[u] = true

  for _, v := range graph[u] {
   if !vis[v] {
    dfs(v)
   }
  }
 }

 dfs(k)

 // Check whether removal is allowed
 for _, edge := range invocations {
  u := edge[0]
  v := edge[1]

  if !vis[u] && vis[v] {
   ans := make([]int, n)

   for i := 0; i < n; i++ {
    ans[i] = i
   }

   return ans
  }
 }

 // Collect remaining methods
 ans := []int{}

 for i := 0; i < n; i++ {
  if !vis[i] {
   ans = append(ans, i)
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in all five languages. Only the syntax changes.

### Step 1 — Build the Graph

The first step is converting the invocation list into an adjacency list.

Instead of repeatedly searching through the input, every method directly stores the methods it invokes.

This allows DFS to explore neighbors efficiently.

---

### Step 2 — Mark Suspicious Methods

Starting from method `k`, DFS visits every reachable method.

Whenever DFS reaches a method, it is marked as suspicious.

DFS continues until there are no more reachable methods.

At this point, the visited array contains exactly the suspicious methods.

---

### Step 3 — Validate the Removal

Finding suspicious methods alone is not enough.

The problem requires checking whether any safe method still invokes one of them.

I iterate through every invocation.

For each edge `(u, v)`:

- If `u` is safe
- And `v` is suspicious

then removing the suspicious group would leave a broken dependency.

In that situation, I immediately return every method because the problem states that nothing should be removed.

---

### Step 4 — Build the Final Answer

If no invalid dependency exists, the suspicious group can safely be removed.

I iterate through every method.

Every method that was never marked suspicious is added to the answer.

Finally, I return the result.

---

### Why DFS Works Well

DFS naturally explores every method reachable from the buggy method.

Since every reachable method must be considered suspicious, DFS perfectly matches the problem requirements.

---

### Why the Second Scan is Necessary

DFS only tells us which methods are suspicious.

It does not tell us whether outside methods still depend on them.

The second scan guarantees that the removal rule is respected.

Without this check, the answer could become invalid.

---

### Language Differences

The algorithm never changes.

Only implementation details differ.

- C++ uses `vector` and recursive functions.
- Java uses `ArrayList` and helper methods.
- JavaScript uses arrays and recursive functions.
- Python uses lists and nested DFS.
- Go uses slices and recursive functions.

The overall behavior and complexity remain exactly the same.

---

## Examples

### Example 1

**Input**

```text
n = 4
k = 1
invocations = [[1,2],[0,1],[3,2]]
```

**Output**

```text
[0,1,2,3]
```

**Explanation**

DFS marks methods `1` and `2` as suspicious.

However, method `0` calls method `1`, which is suspicious.

Since a safe method depends on a suspicious method, no removal is allowed.

---

### Example 2

**Input**

```text
n = 5
k = 0
invocations = [[1,2],[0,2],[0,1],[3,4]]
```

**Output**

```text
[3,4]
```

**Explanation**

Methods `0`, `1`, and `2` become suspicious.

No safe method calls any of them.

Therefore, they can all be removed.

---

### Example 3

**Input**

```text
n = 3
k = 2
invocations = [[1,2],[0,1],[2,0]]
```

**Output**

```text
[]
```

**Explanation**

DFS reaches every method in the graph.

Every method becomes suspicious.

Since there are no safe methods left, the entire project can be removed.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project folder.

```bash
cd <repository-folder>
```

Choose the language you want to run.

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
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

---

### Go

Run

```bash
go run solution.go
```

---

## Notes & Optimizations

- DFS and BFS both work for this problem because both find all reachable nodes.
- An adjacency list is much more memory-efficient than an adjacency matrix for sparse graphs.
- The second scan over the invocation list is essential. Without it, the solution would fail on cases where safe methods still depend on suspicious ones.
- Every edge is processed only a constant number of times, making the algorithm scalable for large inputs.
- The solution naturally handles disconnected graphs.
- If the buggy method cannot reach any other method, the algorithm still works correctly.
- If every method becomes suspicious, the algorithm correctly returns an empty list when removal is valid.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
