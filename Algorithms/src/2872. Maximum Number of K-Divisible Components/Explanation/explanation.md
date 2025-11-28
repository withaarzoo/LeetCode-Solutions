# 2872. Maximum Number of K-Divisible Components

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

We are given:

* An **undirected tree** with `n` nodes labeled from `0` to `n-1`.
* An array `edges` of length `n-1` where `edges[i] = [a, b]` means there is an edge between nodes `a` and `b`.
* An integer array `values`, where `values[i]` is the value of the `i`-th node.
* An integer `k`.

A **valid split** of the tree is made by removing some edges (possibly none) so that:

* Every resulting **connected component** has its **sum of node values divisible by `k`**.

We need to return the **maximum possible number of connected components** in any valid split.

The problem guarantees that **total sum of all node values is divisible by `k`**.

---

## Constraints

* `1 <= n <= 3 * 10^4`
* `edges.length == n - 1`
* `edges[i].length == 2`
* `0 <= a, b < n`
* `values.length == n`
* `0 <= values[i] <= 10^9`
* `1 <= k <= 10^9`
* Sum of all `values[i]` is divisible by `k`.
* `edges` always form a valid tree (connected, acyclic).

---

## Intuition

I started by thinking:

* I can cut edges anywhere.
* After cutting, each connected part must have **sum divisible by `k`**.
* I want to **maximize** the number of such parts.

For a **tree**, a very natural idea is:

> “What if I root the tree and think in terms of **subtrees**?”

If I look at any node:

* The **subtree** under it (the node + all its descendants) has a fixed sum.
* If that subtree sum is divisible by `k`, I can cut the edge above this node and treat the whole subtree as **one valid component**.

So the core question becomes:

> For each node, what is `subtreeSum % k`?

If I can compute that for every node using DFS, I can decide where to cut.

Also, since the **total sum is divisible by `k`**, the root’s subtree will definitely form a valid component at the end, so the count will be correct.

---

## Approach

Step-by-step (from my point of view):

1. **Build the adjacency list**

   * Convert `edges` into a graph structure: `adj[u]` holds all neighbors of `u`.

2. **Root the tree at node 0**

   * I will run DFS starting from node `0`.
   * I pass the `parent` in DFS so that I don’t go back up and create a cycle.

3. **DFS logic**
   For each node `u`:

   * Start with `sum = values[u] % k`.
   * For every child `v` (neighbors except the parent):

     * Recursively compute `childRem = dfs(v, u)` which gives `subtreeSum(v) % k`.
     * Add it: `sum = (sum + childRem) % k`.

4. **Cutting rule**
   After processing children of `u`:

   * If `sum % k == 0`, then the entire subtree at `u` has sum divisible by `k`.

     * I **count one component**.
     * I **return 0** to the parent, meaning:
       “My whole subtree is complete and separated; I don’t add anything to your sum.”
   * Else:

     * I **cannot cut** here.
     * I return `sum` upwards to be combined with the parent’s subtree.

5. **Answer**

   * I keep a global counter `ans`.
   * Every time a subtree remainder is 0, I increment `ans`.
   * After the DFS from root `0`, `ans` is exactly the **maximum number of `k`-divisible components**.

I only store sums **modulo `k`** to avoid overflow and to keep everything efficient, because I only care about divisibility.

---

## Data Structures Used

* **Adjacency list** (`vector<vector<int>>`, `List<List<Integer>>`, arrays of arrays, etc.)
  To represent the tree efficiently.

* **Recursion / explicit stack**
  DFS over the tree to compute subtree sums.

* **Global / outer variables**
  To store:

  * `ans` → number of components.
  * Input arrays and `k`, in some languages.

---

## Operations & Behavior Summary

For each node:

1. Visit node.
2. Start with its own value `values[u] % k`.
3. Visit each child once using DFS.
4. Add children’s remainders modulo `k`.
5. If the final remainder at this node is `0`:

   * Increment global `ans`.
   * Return `0` to parent to signal “this subtree is closed”.
6. Otherwise:

   * Return the remainder to the parent.

Because we traverse each node and edge exactly once, everything is linear in `n`.

---

## Complexity

Let `n` = number of nodes.

* **Time Complexity:** `O(n)`

  * Each node is processed once.
  * Each edge is considered from both ends once.

* **Space Complexity:** `O(n)`

  * Adjacency list stores `2 * (n-1)` entries → `O(n)`.
  * Recursion stack (or iterative DFS stack) may go up to `O(n)` in worst case (skewed tree).
  * Extra arrays like `parent`, `remainder` (for JS iterative version) are also `O(n)`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int maxKDivisibleComponents(int n, vector<vector<int>>& edges, vector<int>& values, int k) {
        // Build adjacency list
        vector<vector<int>> adj(n);
        for (auto &e : edges) {
            int u = e[0], v = e[1];
            adj[u].push_back(v);
            adj[v].push_back(u);
        }

        long long ans = 0;

        // DFS: returns subtree sum % k for node u
        function<long long(int, int)> dfs = [&](int u, int parent) -> long long {
            long long sum = values[u] % k;

            for (int v : adj[u]) {
                if (v == parent) continue;
                long long childRem = dfs(v, u);
                sum = (sum + childRem) % k;
            }

            if (sum % k == 0) {
                ans++;      // this subtree is a valid component
                return 0;   // nothing contributed upwards
            }
            return sum;     // pass remainder to parent
        };

        dfs(0, -1);
        return (int)ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {

    private List<List<Integer>> adj;
    private int[] values;
    private int k;
    private int ans;

    public int maxKDivisibleComponents(int n, int[][] edges, int[] values, int k) {
        this.values = values;
        this.k = k;
        this.ans = 0;

        // Build adjacency list
        adj = new ArrayList<>();
        for (int i = 0; i < n; i++) adj.add(new ArrayList<>());

        for (int[] e : edges) {
            int u = e[0], v = e[1];
            adj.get(u).add(v);
            adj.get(v).add(u);
        }

        dfs(0, -1);  // root at 0
        return ans;
    }

    // DFS returns subtree sum % k for node u
    private long dfs(int u, int parent) {
        long sum = values[u] % (long)k;

        for (int v : adj.get(u)) {
            if (v == parent) continue;
            long childRem = dfs(v, u);
            sum = (sum + childRem) % k;
        }

        if (sum % k == 0) {
            ans++;       // one more valid component
            return 0;    // subtree closed
        }
        return sum;       // pass remainder to parent
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number[]} values
 * @param {number} k
 * @return {number}
 */
var maxKDivisibleComponents = function(n, edges, values, k) {
    // Build adjacency list
    const adj = Array.from({ length: n }, () => []);
    for (const [u, v] of edges) {
        adj[u].push(v);
        adj[v].push(u);
    }

    let ans = 0;

    // We'll do iterative DFS to avoid recursion depth issues
    const parent = Array(n).fill(-1);
    const order = [];
    const stack = [0];
    parent[0] = -2; // mark root specially

    // First DFS to get nodes in traversal order
    while (stack.length > 0) {
        const u = stack.pop();
        order.push(u);
        for (const v of adj[u]) {
            if (v === parent[u]) continue;
            parent[v] = u;
            stack.push(v);
        }
    }

    const remainder = Array(n).fill(0);

    // Process nodes in reverse order (post-order)
    for (let i = order.length - 1; i >= 0; i--) {
        const u = order[i];
        let sum = values[u] % k;

        for (const v of adj[u]) {
            if (v === parent[u]) continue;
            sum = (sum + remainder[v]) % k;
        }

        if (sum % k === 0) {
            ans++;
            remainder[u] = 0;     // subtree is its own component
        } else {
            remainder[u] = sum;   // contribute to parent
        }
    }

    return ans;
};
```

---

### Python3

```python
from typing import List
import sys

class Solution:
    def maxKDivisibleComponents(self, n: int, edges: List[List[int]], values: List[int], k: int) -> int:
        # For very deep trees, increase recursion limit
        sys.setrecursionlimit(10**6)

        # Build adjacency list
        adj = [[] for _ in range(n)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)

        self.ans = 0

        def dfs(u: int, parent: int) -> int:
            """
            Returns: subtree sum % k for node u
            """
            total = values[u] % k

            for v in adj[u]:
                if v == parent:
                    continue
                child_rem = dfs(v, u)
                total = (total + child_rem) % k

            if total % k == 0:
                self.ans += 1
                return 0        # this subtree forms its own component
            return total        # pass remainder to parent

        dfs(0, -1)
        return self.ans
```

---

### Go

```go
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
    // Build adjacency list
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    ans := 0

    var dfs func(u, parent int) int64
    dfs = func(u, parent int) int64 {
        // start with current node value modulo k
        sum := int64(values[u] % k)

        for _, v := range adj[u] {
            if v == parent {
                continue
            }
            childRem := dfs(v, u)
            sum = (sum + childRem) % int64(k)
        }

        if sum%int64(k) == 0 {
            ans++
            return 0
        }
        return sum
    }

    dfs(0, -1)
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I’ll explain with the Python version (logic is same in all languages):

### 1. Building the graph

```python
adj = [[] for _ in range(n)]
for u, v in edges:
    adj[u].append(v)
    adj[v].append(u)
```

* I create `adj` – a list of neighbors for each node.
* Since the tree is undirected, I add both `u -> v` and `v -> u`.

Same idea in C++/Java/JS/Go using vectors/lists/arrays.

---

### 2. Preparing the answer

```python
self.ans = 0
```

* This is my global counter for how many components I can create.

In other languages, I use a global or outer variable `ans`.

---

### 3. DFS definition

```python
def dfs(u: int, parent: int) -> int:
    total = values[u] % k
    ...
```

* I define a DFS that takes current node `u` and its `parent`.
* I start with the current node’s own value modulo `k`.

C++:

```cpp
function<long long(int, int)> dfs = [&](int u, int parent) -> long long {
    long long sum = values[u] % k;
    ...
};
```

Java uses a separate `dfs` method; Go uses an inner function `dfs`.

---

### 4. Processing children

```python
for v in adj[u]:
    if v == parent:
        continue
    child_rem = dfs(v, u)
    total = (total + child_rem) % k
```

* I loop through neighbors `v` of `u`.
* If `v` is the parent, I skip it to avoid going back up.
* Recursively call DFS on the child.
* Add the child’s remainder into `total` and keep it modulo `k`.

All other languages follow the same concept.

---

### 5. Checking divisibility and cutting

```python
if total % k == 0:
    self.ans += 1
    return 0
return total
```

* After all children are processed:

  * If `total` is divisible by `k`, the full subtree under `u` is valid as one component.
  * I increase `ans` and return `0` to parent, meaning I’ve cut this subtree from above.
* Otherwise:

  * I return the remainder to be added by the parent node.

Again, exactly the same in C++, Java, JS, Go.

---

### 6. Starting DFS and returning result

```python
dfs(0, -1)
return self.ans
```

* I root the tree at `0`.
* Use `-1` as a dummy parent.
* After DFS is finished, `self.ans` holds the maximum number of K-divisible components.

Same for C++/Java/Go.
In JavaScript, I simulate DFS iteratively but the post-order logic is equivalent.

---

## Examples

### Example 1

```text
n = 5
edges = [[0,2],[1,2],[1,3],[2,4]]
values = [1,8,1,4,4]
k = 6
```

Tree structure (rooted at 0 for understanding):

* Node 2 connected to 0,1,4
* Node 1 connected to 2 and 3

Subtree sums:

* Subtree(3) = 4 → 4 % 6 = 4
* Subtree(4) = 4 → 4 % 6 = 4
* Subtree(1) = 8 + 4 = 12 → 12 % 6 = 0 → **component**
* Subtree(2) = 1 (node 2) + 4 (from node 4) + 1 (from 0) = 6 → **component**

Answer: `2`.

---

### Example 2

```text
n = 7
edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
values = [3,0,6,1,5,2,1]
k = 3
```

* We can cut edges so that components are:

  * {0} with sum 3
  * {2,5,6} with sum 9
  * {1,3,4} with sum 6

Answer: `3`.

---

## How to use / Run locally

### C++

```bash
g++ -std=c++17 -O2 main.cpp -o main
./main
```

Your `main.cpp` should include:

```cpp
#include <bits/stdc++.h>
using namespace std;

// Paste the Solution class here

int main() {
    // create test cases and call Solution().maxKDivisibleComponents(...)
    return 0;
}
```

---

### Java

```bash
javac Main.java
java Main
```

`Main.java`:

```java
public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int n = 5;
        int[][] edges = {{0,2},{1,2},{1,3},{2,4}};
        int[] values = {1,8,1,4,4};
        int k = 6;
        System.out.println(sol.maxKDivisibleComponents(n, edges, values, k));
    }
}

// Paste the Solution class below
```

---

### JavaScript (Node.js)

```bash
node main.js
```

`main.js`:

```javascript
// Paste the JS function here

const n = 5;
const edges = [[0,2],[1,2],[1,3],[2,4]];
const values = [1,8,1,4,4];
const k = 6;

console.log(maxKDivisibleComponents(n, edges, values, k));
```

---

### Python3

```bash
python3 main.py
```

`main.py`:

```python
from typing import List

# Paste the Solution class here

if __name__ == "__main__":
    n = 5
    edges = [[0,2],[1,2],[1,3],[2,4]]
    values = [1,8,1,4,4]
    k = 6

    sol = Solution()
    print(sol.maxKDivisibleComponents(n, edges, values, k))
```

---

### Go

```bash
go run main.go
```

`main.go`:

```go
package main

import "fmt"

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
    // Paste the Go implementation here
    adj := make([][]int, n)
    for _, e := range edges {
        u, v := e[0], e[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    ans := 0

    var dfs func(int, int) int64
    dfs = func(u, parent int) int64 {
        sum := int64(values[u] % k)
        for _, v := range adj[u] {
            if v == parent {
                continue
            }
            childRem := dfs(v, u)
            sum = (sum + childRem) % int64(k)
        }
        if sum%int64(k) == 0 {
            ans++
            return 0
        }
        return sum
    }

    dfs(0, -1)
    return ans
}

func main() {
    n := 5
    edges := [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}
    values := []int{1, 8, 1, 4, 4}
    k := 6
    fmt.Println(maxKDivisibleComponents(n, edges, values, k))
}
```

---

## Notes & Optimizations

* I only store **subtree sums modulo `k`**, not the full sums.
  This avoids overflow (especially with `values[i]` up to `1e9`) and keeps math simple.
* The algorithm is **pure DFS** and runs in **linear time**, which is optimal for trees.
* Using recursion is simplest; in JS I used an **iterative DFS** to be safe from stack limits.
* The guarantee that **total sum is divisible by `k`** ensures the root also forms a valid component, so we don’t need extra checks for “leftover” remainder at the end.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)

Feel free to fork, star ⭐, or open issues if you see any improvements!
