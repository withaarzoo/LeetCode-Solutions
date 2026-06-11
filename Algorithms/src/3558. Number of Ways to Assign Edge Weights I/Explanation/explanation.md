# LeetCode 3558. Number of Ways to Assign Edge Weights I

A clean and optimized solution for **LeetCode 3558 - Number of Ways to Assign Edge Weights I** using **Tree DFS**, **Graph Traversal**, **Depth Calculation**, **Parity Observation**, and **Modular Exponentiation**.

This repository contains explanations, complexity analysis, and multi-language implementations in C++, Java, JavaScript, Python, and Go.

---

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

You are given an undirected tree with `n` nodes rooted at node `1`.

Initially, every edge has weight `0`. You must assign each edge a weight of either:

* `1`
* `2`

Only the path from the root node (`1`) to a node located at the maximum depth matters.

The goal is to determine how many different valid weight assignments make the total cost of that path odd.

Since the answer can become very large, it must be returned modulo:

```text
1,000,000,007
```

This problem combines several common competitive programming concepts:

* Trees
* Graph Traversal
* DFS
* Maximum Depth Calculation
* Parity Mathematics
* Modular Exponentiation

---

## Constraints

| Constraint    | Value          |
| ------------- | -------------- |
| n             | 2 ≤ n ≤ 10⁵    |
| edges.length  | n - 1          |
| Edge Format   | [uᵢ, vᵢ]       |
| Node Values   | 1 ≤ uᵢ, vᵢ ≤ n |
| Tree Property | Valid Tree     |

---

## Intuition

The first thing I noticed was that the actual edge values are not very important.

Each edge can only be assigned:

* 1 (odd)
* 2 (even)

Since the question only asks whether the total path cost is odd, I only need to think about parity.

Adding a weight of `2` never changes parity because it is even.

Only weight `1` affects whether the final sum becomes odd or even.

That means the path cost is odd if the number of edges assigned weight `1` is odd.

Once I realized that, the problem became much simpler.

All I needed was:

1. Find the maximum depth of the tree.
2. Count how many assignments produce an odd number of `1`s on that path.

---

## Approach

1. Build an adjacency list from the given edges.
2. Start DFS from node `1`.
3. Compute the depth of every node.
4. Track the maximum depth found.
5. Let that depth be `d`.
6. The path contains exactly `d` edges.
7. Each edge has two possible choices:

   * weight `1`
   * weight `2`
8. Total assignments become:

```text
2^d
```

1. Exactly half of those assignments have odd parity.
2. Therefore the answer becomes:

```text
2^(d - 1)
```

1. Compute the value using fast modular exponentiation.

---

## Data Structures Used

### Adjacency List

Used to represent the tree efficiently.

Reason:

* Fast traversal
* O(n) memory
* Standard representation for sparse graphs

### Stack / DFS Traversal

Used to explore the tree and calculate node depths.

Reason:

* Visits every node exactly once
* Efficient for depth calculations

### Visited Array

Used to avoid revisiting nodes.

Reason:

* Prevents moving back to parent nodes
* Guarantees O(n) traversal

---

## Operations & Behavior Summary

The algorithm performs the following major steps:

### Step 1: Build the Tree

Convert edge list into an adjacency list.

### Step 2: Traverse the Tree

Run DFS from node `1`.

### Step 3: Compute Depths

Keep track of how far each node is from the root.

### Step 4: Find Maximum Depth

Store the deepest level encountered.

### Step 5: Apply Mathematical Observation

If maximum depth is `d`:

```text
Answer = 2^(d - 1)
```

### Step 6: Return Answer

Use modular exponentiation to avoid overflow.

---

## Complexity

| Type             | Complexity | Explanation                                  |
| ---------------- | ---------- | -------------------------------------------- |
| Time Complexity  | O(n)       | Every node and edge is visited once          |
| Space Complexity | O(n)       | Adjacency list, DFS stack, and visited array |

Where:

* `n` = number of nodes in the tree

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static const int MOD = 1000000007;

    // Fast modular exponentiation
    long long modPow(long long base, long long exp) {
        long long result = 1;

        while (exp > 0) {
            // If current bit is set, multiply answer
            if (exp & 1) {
                result = (result * base) % MOD;
            }

            // Square the base for next bit
            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    int assignEdgeWeights(vector<vector<int>>& edges) {
        int n = edges.size() + 1;

        // Build adjacency list
        vector<vector<int>> graph(n + 1);

        for (auto &e : edges) {
            int u = e[0];
            int v = e[1];

            graph[u].push_back(v);
            graph[v].push_back(u);
        }

        int maxDepth = 0;

        // Iterative DFS: {node, depth}
        stack<pair<int, int>> st;
        st.push({1, 0});

        vector<int> visited(n + 1, 0);
        visited[1] = 1;

        while (!st.empty()) {
            auto [node, depth] = st.top();
            st.pop();

            maxDepth = max(maxDepth, depth);

            for (int nei : graph[node]) {
                if (!visited[nei]) {
                    visited[nei] = 1;
                    st.push({nei, depth + 1});
                }
            }
        }

        // Number of odd-parity assignments = 2^(maxDepth - 1)
        return (int)modPow(2, maxDepth - 1);
    }
};
```

### Java

```java
class Solution {
    private static final long MOD = 1_000_000_007L;

    // Fast modular exponentiation
    private long modPow(long base, long exp) {
        long result = 1;

        while (exp > 0) {
            // Multiply when current bit is set
            if ((exp & 1) == 1) {
                result = (result * base) % MOD;
            }

            // Square the base
            base = (base * base) % MOD;
            exp >>= 1;
        }

        return result;
    }

    public int assignEdgeWeights(int[][] edges) {
        int n = edges.length + 1;

        // Build adjacency list
        List<Integer>[] graph = new ArrayList[n + 1];

        for (int i = 0; i <= n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] e : edges) {
            int u = e[0];
            int v = e[1];

            graph[u].add(v);
            graph[v].add(u);
        }

        int maxDepth = 0;

        Deque<int[]> stack = new ArrayDeque<>();
        stack.push(new int[]{1, 0});

        boolean[] visited = new boolean[n + 1];
        visited[1] = true;

        while (!stack.isEmpty()) {
            int[] cur = stack.pop();

            int node = cur[0];
            int depth = cur[1];

            maxDepth = Math.max(maxDepth, depth);

            for (int next : graph[node]) {
                if (!visited[next]) {
                    visited[next] = true;
                    stack.push(new int[]{next, depth + 1});
                }
            }
        }

        return (int) modPow(2, maxDepth - 1);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[][]} edges
 * @return {number}
 */
var assignEdgeWeights = function(edges) {
    const MOD = 1000000007n;
    const n = edges.length + 1;

    // Build adjacency list
    const graph = Array.from({ length: n + 1 }, () => []);

    for (const [u, v] of edges) {
        graph[u].push(v);
        graph[v].push(u);
    }

    let maxDepth = 0;

    // Iterative DFS
    const stack = [[1, 0]];
    const visited = Array(n + 1).fill(false);
    visited[1] = true;

    while (stack.length) {
        const [node, depth] = stack.pop();

        maxDepth = Math.max(maxDepth, depth);

        for (const next of graph[node]) {
            if (!visited[next]) {
                visited[next] = true;
                stack.push([next, depth + 1]);
            }
        }
    }

    // Fast modular exponentiation
    let base = 2n;
    let exp = BigInt(maxDepth - 1);
    let result = 1n;

    while (exp > 0n) {
        if (exp & 1n) {
            result = (result * base) % MOD;
        }

        base = (base * base) % MOD;
        exp >>= 1n;
    }

    return Number(result);
};
```

### Python3

```python
class Solution:
    def assignEdgeWeights(self, edges: List[List[int]]) -> int:
        MOD = 10**9 + 7
        n = len(edges) + 1

        # Build adjacency list
        graph = [[] for _ in range(n + 1)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        max_depth = 0

        # Iterative DFS -> (node, depth)
        stack = [(1, 0)]
        visited = [False] * (n + 1)
        visited[1] = True

        while stack:
            node, depth = stack.pop()

            max_depth = max(max_depth, depth)

            for nxt in graph[node]:
                if not visited[nxt]:
                    visited[nxt] = True
                    stack.append((nxt, depth + 1))

        # Number of valid assignments = 2^(max_depth - 1)
        return pow(2, max_depth - 1, MOD)
```

### Go

```go
func assignEdgeWeights(edges [][]int) int {
 const MOD int64 = 1000000007

 n := len(edges) + 1

 // Build adjacency list
 graph := make([][]int, n+1)

 for _, e := range edges {
  u, v := e[0], e[1]

  graph[u] = append(graph[u], v)
  graph[v] = append(graph[v], u)
 }

 type Pair struct {
  node  int
  depth int
 }

 // Iterative DFS
 stack := []Pair{{1, 0}}

 visited := make([]bool, n+1)
 visited[1] = true

 maxDepth := 0

 for len(stack) > 0 {
  cur := stack[len(stack)-1]
  stack = stack[:len(stack)-1]

  if cur.depth > maxDepth {
   maxDepth = cur.depth
  }

  for _, next := range graph[cur.node] {
   if !visited[next] {
    visited[next] = true
    stack = append(stack, Pair{next, cur.depth + 1})
   }
  }
 }

 // Fast modular exponentiation
 base := int64(2)
 exp := maxDepth - 1
 result := int64(1)

 for exp > 0 {
  // Multiply when current bit is set
  if exp&1 == 1 {
   result = (result * base) % MOD
  }

  // Square the base
  base = (base * base) % MOD
  exp >>= 1
 }

 return int(result)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.

### Building the Graph

The input arrives as an edge list.

For efficient traversal, I convert it into an adjacency list.

This allows me to quickly access all neighboring nodes.

---

### Starting DFS

I start traversal from node `1` because the tree is rooted there.

The initial depth is:

```text
0
```

Every time I move to a child node, depth increases by one.

---

### Tracking Maximum Depth

While traversing the tree, I continuously compare:

```text
current depth
vs
maximum depth seen so far
```

Whenever a deeper node appears, I update the answer.

After traversal finishes, I know the deepest level in the tree.

---

### Understanding the Math

Suppose the maximum depth is:

```text
d
```

Then the selected path contains exactly:

```text
d edges
```

Each edge has two choices.

Therefore:

```text
Total assignments = 2^d
```

---

### Odd Sum Requirement

The total path cost becomes odd only when the number of weight `1` edges is odd.

A well-known parity property says:

```text
Half of all binary assignments have odd parity.
Half have even parity.
```

Therefore:

```text
Valid assignments = 2^d / 2
```

which simplifies to:

```text
2^(d - 1)
```

---

### Modular Exponentiation

The exponent can become large.

Instead of multiplying repeatedly, I use binary exponentiation.

Benefits:

* Much faster
* O(log d)
* Prevents overflow when combined with modulo arithmetic

---

## Examples

### Example 1

Input:

```text
edges = [[1,2]]
```

Tree:

```text
1 -- 2
```

Maximum depth:

```text
1
```

Answer:

```text
2^(1-1)
= 1
```

Output:

```text
1
```

---

### Example 2

Input:

```text
edges = [[1,2],[1,3],[3,4],[3,5]]
```

Tree:

```text
      1
     / \
    2   3
       / \
      4   5
```

Maximum depth:

```text
2
```

Answer:

```text
2^(2-1)
= 2
```

Output:

```text
2
```

---

### Example 3

Input:

```text
edges = [[1,2],[2,3],[3,4]]
```

Tree:

```text
1 - 2 - 3 - 4
```

Maximum depth:

```text
3
```

Answer:

```text
2^(3-1)
= 4
```

Output:

```text
4
```

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -O2 -std=c++17
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

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

### Important Observation

The actual edge values are irrelevant.

Only parity matters.

This reduces the problem from a combinatorics problem into a simple depth calculation plus mathematical counting.

---

### Why DFS Works

DFS visits every node exactly once.

That makes it ideal for finding the maximum depth.

A BFS solution would also work.

---

### Possible Alternative

Instead of DFS, you can use BFS and calculate levels layer by layer.

Both approaches have:

```text
O(n)
```

time complexity.

---

### Edge Case

Smallest valid tree:

```text
1 -- 2
```

Maximum depth is:

```text
1
```

Answer becomes:

```text
2^(0) = 1
```

which matches the expected result.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
