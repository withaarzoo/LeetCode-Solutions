# 2685. Count the Number of Complete Components

A clean and optimized solution for LeetCode 2685 - Count the Number of Complete Components using Graph Theory and Depth-First Search (DFS). This repository explains the intuition, approach, complexity analysis, and provides solutions in C++, Java, JavaScript, Python, and Go.

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

You are given an undirected graph with `n` vertices numbered from `0` to `n - 1` and a list of edges connecting some of those vertices.

The graph may contain multiple connected components.

Your task is to count how many of those connected components are **complete graphs**.

A connected component is considered complete if every pair of vertices inside that component has a direct edge between them.

### Input

- An integer `n` representing the number of vertices.
- A 2D array `edges` where each pair `[u, v]` represents an undirected edge.

### Output

- Return the number of complete connected components in the graph.

This problem is a good example of Graph Theory, Connected Components, and Depth-First Search (DFS).

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= n <= 50` | Number of vertices |
| `0 <= edges.length <= n * (n - 1) / 2` | Number of edges |
| `edges[i].length == 2` | Every edge contains two vertices |
| `0 <= ai, bi < n` | Valid vertex index |
| `ai != bi` | No self-loop |
| No repeated edges | Every edge is unique |

---

## Intuition

The first thing I noticed was that I don't need to compare every pair of vertices directly.

Instead, I can process one connected component at a time.

For each connected component, I only need to know two things:

- How many vertices it contains.
- How many edges belong to that component.

If a component has `m` vertices, then a complete graph must contain exactly:

`m × (m − 1) / 2`

edges.

So after finding a connected component using DFS, I simply compare the actual number of edges with the expected number. If they match, that component is complete.

---

## Approach

I solved the problem using Depth-First Search.

The steps are straightforward.

1. Build an adjacency list from the given edges.
2. Create a visited array.
3. Start DFS from every unvisited vertex.
4. During DFS:
   - Count the number of vertices.
   - Count the total degree of all vertices.
5. Divide the total degree by 2 because every undirected edge is counted twice.
6. Compare the edge count with the expected number of edges.
7. If both values are equal, increase the answer.
8. Continue until every connected component has been processed.

This solution visits every vertex only once and is efficient for the given constraints.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Adjacency List | Stores the graph efficiently |
| Visited Array | Prevents revisiting vertices during DFS |
| DFS Recursion Stack | Explores one connected component completely |
| Integer Counters | Count vertices and edges inside each component |

---

## Operations & Behavior Summary

The algorithm works in several simple stages.

1. Read all edges.
2. Build an adjacency list.
3. Visit every unvisited vertex.
4. Run DFS to discover one complete connected component.
5. Count:
   - Number of vertices
   - Sum of degrees
6. Convert degree sum into edge count.
7. Check whether:

   `edges = vertices × (vertices − 1) / 2`

8. If true, count this component.
9. Return the final answer.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n + m)** | Every vertex and every edge is visited once. |
| Space Complexity | **O(n + m)** | Extra memory is used for the adjacency list, visited array, and DFS recursion stack. |

Where:

- `n` = number of vertices
- `m` = number of edges

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    // DFS to visit one connected component
    void dfs(int node, vector<vector<int>>& graph, vector<bool>& vis,
             int& vertices, int& degreeSum) {

        // Mark current node as visited
        vis[node] = true;

        // Count this vertex
        vertices++;

        // Add its degree
        degreeSum += graph[node].size();

        // Visit all unvisited neighbors
        for (int next : graph[node]) {
            if (!vis[next]) {
                dfs(next, graph, vis, vertices, degreeSum);
            }
        }
    }

    int countCompleteComponents(int n, vector<vector<int>>& edges) {

        // Build adjacency list
        vector<vector<int>> graph(n);

        for (auto &e : edges) {
            graph[e[0]].push_back(e[1]);
            graph[e[1]].push_back(e[0]);
        }

        // Keep track of visited nodes
        vector<bool> vis(n, false);

        int answer = 0;

        // Process every connected component
        for (int i = 0; i < n; i++) {

            if (vis[i]) continue;

            int vertices = 0;
            int degreeSum = 0;

            // Find all nodes in this component
            dfs(i, graph, vis, vertices, degreeSum);

            // Every edge is counted twice
            int edgeCount = degreeSum / 2;

            // Check whether this component is complete
            if (edgeCount == vertices * (vertices - 1) / 2) {
                answer++;
            }
        }

        return answer;
    }
};
```

### Java

```java
class Solution {

    // DFS to visit one connected component
    private void dfs(int node, List<Integer>[] graph, boolean[] visited,
                     int[] info) {

        // Mark node as visited
        visited[node] = true;

        // Count one more vertex
        info[0]++;

        // Add current node's degree
        info[1] += graph[node].size();

        // Visit all neighbors
        for (int next : graph[node]) {
            if (!visited[next]) {
                dfs(next, graph, visited, info);
            }
        }
    }

    public int countCompleteComponents(int n, int[][] edges) {

        // Build adjacency list
        List<Integer>[] graph = new ArrayList[n];

        for (int i = 0; i < n; i++) {
            graph[i] = new ArrayList<>();
        }

        for (int[] edge : edges) {
            graph[edge[0]].add(edge[1]);
            graph[edge[1]].add(edge[0]);
        }

        boolean[] visited = new boolean[n];

        int answer = 0;

        // Process every component
        for (int i = 0; i < n; i++) {

            if (visited[i]) continue;

            // info[0] = vertices
            // info[1] = total degree
            int[] info = new int[2];

            dfs(i, graph, visited, info);

            int vertices = info[0];
            int edgeCount = info[1] / 2;

            // Check whether this component is complete
            if (edgeCount == vertices * (vertices - 1) / 2) {
                answer++;
            }
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[][]} edges
 * @return {number}
 */
var countCompleteComponents = function(n, edges) {

    // Build adjacency list
    const graph = Array.from({ length: n }, () => []);

    for (const [u, v] of edges) {
        graph[u].push(v);
        graph[v].push(u);
    }

    // Track visited nodes
    const visited = new Array(n).fill(false);

    let answer = 0;

    // DFS for one connected component
    function dfs(node, data) {

        // Mark current node
        visited[node] = true;

        // Count vertex
        data.vertices++;

        // Add degree
        data.degreeSum += graph[node].length;

        // Visit neighbors
        for (const next of graph[node]) {
            if (!visited[next]) {
                dfs(next, data);
            }
        }
    }

    // Traverse every component
    for (let i = 0; i < n; i++) {

        if (visited[i]) continue;

        const data = {
            vertices: 0,
            degreeSum: 0
        };

        dfs(i, data);

        // Every edge was counted twice
        const edgeCount = data.degreeSum / 2;

        // Check completeness
        if (edgeCount === data.vertices * (data.vertices - 1) / 2) {
            answer++;
        }
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def countCompleteComponents(self, n: int, edges: List[List[int]]) -> int:

        # Build adjacency list
        graph = [[] for _ in range(n)]

        for u, v in edges:
            graph[u].append(v)
            graph[v].append(u)

        # Keep track of visited nodes
        visited = [False] * n

        # DFS to explore one connected component
        def dfs(node):

            # Mark current node
            visited[node] = True

            # Count current vertex
            vertices = 1

            # Add current node's degree
            degree_sum = len(graph[node])

            # Visit all neighbors
            for nxt in graph[node]:
                if not visited[nxt]:
                    v, d = dfs(nxt)
                    vertices += v
                    degree_sum += d

            return vertices, degree_sum

        answer = 0

        # Process every component
        for i in range(n):

            if visited[i]:
                continue

            vertices, degree_sum = dfs(i)

            # Every edge appears twice
            edge_count = degree_sum // 2

            # Check if this component is complete
            if edge_count == vertices * (vertices - 1) // 2:
                answer += 1

        return answer
```

### Go

```go
func countCompleteComponents(n int, edges [][]int) int {

 // Build adjacency list
 graph := make([][]int, n)

 for _, edge := range edges {
  u, v := edge[0], edge[1]
  graph[u] = append(graph[u], v)
  graph[v] = append(graph[v], u)
 }

 // Track visited nodes
 visited := make([]bool, n)

 // DFS for one connected component
 var dfs func(int, *int, *int)

 dfs = func(node int, vertices *int, degreeSum *int) {

  // Mark node as visited
  visited[node] = true

  // Count current vertex
  *vertices++

  // Add current node's degree
  *degreeSum += len(graph[node])

  // Visit neighbors
  for _, next := range graph[node] {
   if !visited[next] {
    dfs(next, vertices, degreeSum)
   }
  }
 }

 answer := 0

 // Process every component
 for i := 0; i < n; i++ {

  if visited[i] {
   continue
  }

  vertices := 0
  degreeSum := 0

  dfs(i, &vertices, &degreeSum)

  // Every edge is counted twice
  edgeCount := degreeSum / 2

  // Check whether this component is complete
  if edgeCount == vertices*(vertices-1)/2 {
   answer++
  }
 }

 return answer
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in every language. Only the syntax changes.

### Step 1

Build an adjacency list.

This allows every vertex to quickly access all of its neighboring vertices.

---

### Step 2

Create a visited array.

Without it, the DFS would keep revisiting vertices and could enter an infinite loop.

---

### Step 3

Loop through every vertex.

Whenever an unvisited vertex is found, start a new DFS.

That DFS will discover exactly one connected component.

---

### Step 4

During DFS, keep updating two counters.

The first counter stores how many vertices belong to the current component.

The second counter stores the total degree of all vertices.

---

### Step 5

Convert the degree sum into the actual number of edges.

Because the graph is undirected, every edge appears twice inside the adjacency list.

```
Degree Sum = 2 × Number of Edges
```

So,

```
Edges = Degree Sum / 2
```

---

### Step 6

Suppose the component contains `m` vertices.

A complete graph with `m` vertices must contain

```
m × (m − 1) / 2
```

edges.

If the calculated edge count matches this value, then every pair of vertices is connected.

---

### Step 7

Increase the answer.

Repeat the same process until every connected component has been processed.

Finally, return the total number of complete connected components.

---

## Examples

### Example 1

**Input**

```text
n = 6
edges = [[0,1],[0,2],[1,2],[3,4]]
```

**Output**

```text
3
```

### Trace

Component 1

```
0 - 1
 \ /
  2
```

Vertices = 3

Edges = 3

Complete

---

Component 2

```
3 - 4
```

Vertices = 2

Edges = 1

Complete

---

Component 3

```
5
```

Single isolated vertex

Complete

Answer = 3

---

### Example 2

**Input**

```text
n = 6
edges = [[0,1],[0,2],[1,2],[3,4],[3,5]]
```

**Output**

```text
1
```

### Trace

First component

```
0
|\
| \
1--2
```

Complete

Second component

```
5
 \
  3
 /
4
```

Missing edge between 4 and 5

Not Complete

Answer = 1

---

### Example 3

**Input**

```text
n = 4
edges = []
```

**Output**

```text
4
```

### Trace

Every vertex is isolated.

Each isolated vertex is considered a complete graph.

Answer = 4

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project folder.

```bash
cd your-repository
```

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

- Every connected component is visited exactly once.
- The adjacency list is more memory-efficient than an adjacency matrix for sparse graphs.
- Counting the degree sum is much simpler than checking every pair of vertices individually.
- The formula `m × (m − 1) / 2` provides an immediate way to verify whether a component is complete.
- This DFS solution is optimal for the given constraints.
- The same idea can also be implemented using Breadth-First Search (BFS) or Union-Find, although DFS is usually the simplest approach here.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
