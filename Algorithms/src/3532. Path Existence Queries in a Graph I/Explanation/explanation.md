# 3532. Path Existence Queries in a Graph I

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
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

## Problem Summary

In LeetCode 3532, Path Existence Queries in a Graph I, we are given `n` nodes labeled from `0` to `n - 1`.

We also get a sorted integer array called `nums`. An undirected edge exists between two nodes `i` and `j` when the absolute difference between their values is at most `maxDiff`.

In other words, nodes `i` and `j` are directly connected when:

`|nums[i] - nums[j]| <= maxDiff`

We are also given multiple queries. Each query contains two nodes `[u, v]`.

For every query, we need to determine whether there is any path from node `u` to node `v`.

The important part is that the path does not need to be a direct edge. The two nodes can also be connected through other nodes.

The output is a boolean array where:

* `true` means a path exists between the two queried nodes.
* `false` means no path exists.

Since `nums` is already sorted, we can solve this graph connectivity problem without building the full graph.

## Constraints

| Constraint                 | Value                                    |
| -------------------------- | ---------------------------------------- |
| Number of nodes            | `1 <= n == nums.length <= 10^5`          |
| Value of each element      | `0 <= nums[i] <= 10^5`                   |
| Array order                | `nums` is sorted in non-decreasing order |
| Maximum allowed difference | `0 <= maxDiff <= 10^5`                   |
| Number of queries          | `1 <= queries.length <= 10^5`            |
| Query format               | `queries[i] = [u_i, v_i]`                |
| Node range                 | `0 <= u_i, v_i < n`                      |

The large input size means checking every pair of nodes or running a graph traversal for every query would be too slow.

## Intuition

My first thought was to build the graph and connect every pair of nodes whose values differ by at most `maxDiff`.

That approach is unnecessary and can become very expensive.

The key observation comes from the fact that `nums` is sorted.

Suppose I find two consecutive positions where:

`nums[i] - nums[i - 1] > maxDiff`

This gap separates the graph into two parts.

Because the array is sorted, every node before index `i` has a value less than or equal to `nums[i - 1]`. Every node from index `i` onward has a value greater than or equal to `nums[i]`.

If even the closest values across this boundary differ by more than `maxDiff`, then no edge can cross this gap.

That means every large consecutive gap starts a new connected component.

On the other hand, if consecutive values differ by at most `maxDiff`, those nodes are directly connected. A chain of such valid consecutive edges keeps all those nodes inside the same connected component.

So the whole graph can be reduced to continuous segments of connected nodes.

Once I know the component of every node, each path existence query becomes a simple comparison.

## Approach

I solve the problem in two main stages.

First, I preprocess the connected components.

I create an array called `component`, where `component[i]` stores the connected component ID of node `i`.

Node `0` starts in component `0`.

Then I scan `nums` from left to right.

For every index `i` from `1` to `n - 1`, I compare:

`nums[i] - nums[i - 1]`

If the difference is greater than `maxDiff`, I know no edge can cross this gap. I start a new connected component by increasing the component ID.

Otherwise, the current node stays in the same component as the previous node.

After this preprocessing step, I process every query `[u, v]`.

If:

`component[u] == component[v]`

then both nodes belong to the same connected component, so a path exists.

Otherwise, they are separated by at least one gap larger than `maxDiff`, so no path exists.

This approach answers every query in constant time after one linear preprocessing pass.

## Data Structures Used

### Component Array

I use an integer array of size `n`.

`component[i]` stores the connected component ID of node `i`.

This lets me answer a path existence query with one direct comparison instead of running DFS, BFS, or another graph search.

### Answer Array

I use a boolean array to store the result of every query.

Each position contains `true` if the queried nodes are connected and `false` otherwise.

### Query Array

The input already provides the queries as pairs of node indices.

I process each query once and do not need to modify or sort them.

## Operations & Behavior Summary

The algorithm works like this:

1. Create a component array for all `n` nodes.
2. Put the first node in component `0`.
3. Start scanning from the second node.
4. Compare each value with the value immediately before it.
5. If the difference is greater than `maxDiff`, start a new component.
6. Store the current component ID for the current node.
7. After preprocessing all nodes, process the queries.
8. For each query `[u, v]`, compare the component IDs of `u` and `v`.
9. Return `true` if the IDs match.
10. Return `false` if the IDs are different.
11. Return the final boolean answer array.

The full graph is never created.

That is the main reason this solution stays fast even when `n` and the number of queries are both as large as `10^5`.

## Complexity

| Complexity       | Value      | Explanation                                                                                                               |
| ---------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n + k)` | I scan the `n` nodes once and process all `k` queries once. Each query takes `O(1)` time.                                 |
| Space Complexity | `O(n + k)` | I store `n` component IDs and return `k` boolean answers. Excluding the required output array, the extra space is `O(n)`. |

Here, `n` is the number of nodes and `k` is the number of queries.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<bool> pathExistenceQueries(int n, vector<int>& nums, int maxDiff, vector<vector<int>>& queries) {
        // component[i] stores which connected component node i belongs to.
        vector<int> component(n, 0);

        // Start with component 0 for the first node.
        int componentId = 0;

        // Check every gap between two consecutive sorted values.
        for (int i = 1; i < n; i++) {
            // A gap larger than maxDiff separates the graph into two parts.
            if (nums[i] - nums[i - 1] > maxDiff) {
                componentId++;
            }

            // Store the component of the current node.
            component[i] = componentId;
        }

        // Store the result of every query.
        vector<bool> answer;
        answer.reserve(queries.size());

        // Two nodes have a path exactly when their component IDs are equal.
        for (const auto& query : queries) {
            int u = query[0];
            int v = query[1];

            answer.push_back(component[u] == component[v]);
        }

        return answer;
    }
}; 
```

### Java

```java
class Solution {
    public boolean[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        // component[i] stores which connected component node i belongs to.
        int[] component = new int[n];

        // Start with component 0 for the first node.
        int componentId = 0;

        // Check every gap between two consecutive sorted values.
        for (int i = 1; i < n; i++) {
            // A gap larger than maxDiff separates the graph into two parts.
            if (nums[i] - nums[i - 1] > maxDiff) {
                componentId++;
            }

            // Store the component of the current node.
            component[i] = componentId;
        }

        // Create one answer for every query.
        boolean[] answer = new boolean[queries.length];

        // Two nodes have a path exactly when their component IDs are equal.
        for (int i = 0; i < queries.length; i++) {
            int u = queries[i][0];
            int v = queries[i][1];

            answer[i] = component[u] == component[v];
        }

        return answer;
    }
} 
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number[]} nums
 * @param {number} maxDiff
 * @param {number[][]} queries
 * @return {boolean[]}
 */
var pathExistenceQueries = function(n, nums, maxDiff, queries) {
    // component[i] stores which connected component node i belongs to.
    const component = new Array(n).fill(0);

    // Start with component 0 for the first node.
    let componentId = 0;

    // Check every gap between two consecutive sorted values.
    for (let i = 1; i < n; i++) {
        // A gap larger than maxDiff separates the graph into two parts.
        if (nums[i] - nums[i - 1] > maxDiff) {
            componentId++;
        }

        // Store the component of the current node.
        component[i] = componentId;
    }

    // Create one answer for every query.
    const answer = new Array(queries.length);

    // Two nodes have a path exactly when their component IDs are equal.
    for (let i = 0; i < queries.length; i++) {
        const u = queries[i][0];
        const v = queries[i][1];

        answer[i] = component[u] === component[v];
    }

    return answer;
}; 
```

### Python3

```python
class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[bool]:
        # component[i] stores which connected component node i belongs to.
        component = [0] * n

        # Start with component 0 for the first node.
        component_id = 0

        # Check every gap between two consecutive sorted values.
        for i in range(1, n):
            # A gap larger than maxDiff separates the graph into two parts.
            if nums[i] - nums[i - 1] > maxDiff:
                component_id += 1

            # Store the component of the current node.
            component[i] = component_id

        # Two nodes have a path exactly when their component IDs are equal.
        answer = []
        for u, v in queries:
            answer.append(component[u] == component[v])

        return answer
```

### Go

```go
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
    // component[i] stores which connected component node i belongs to.
    component := make([]int, n)

    // Start with component 0 for the first node.
    componentId := 0

    // Check every gap between two consecutive sorted values.
    for i := 1; i < n; i++ {
        // A gap larger than maxDiff separates the graph into two parts.
        if nums[i]-nums[i-1] > maxDiff {
            componentId++
        }

        // Store the component of the current node.
        component[i] = componentId
    }

    // Create one answer for every query.
    answer := make([]bool, len(queries))

    // Two nodes have a path exactly when their component IDs are equal.
    for i, query := range queries {
        u := query[0]
        v := query[1]

        answer[i] = component[u] == component[v]
    }

    return answer
} 
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in C++, Java, JavaScript, Python3, and Go. Only the syntax and the way arrays are created differ.

### Step 1: Create the component storage

I first create an array with `n` positions.

Each position represents one node.

The value stored at index `i` tells me which connected component contains node `i`.

At the beginning, all values can be `0` because the first component ID is also `0`.

If `n = 5`, the initial component array looks like this:

`[0, 0, 0, 0, 0]`

At this point, I have not processed the gaps yet.

### Step 2: Start with the first component

I keep a variable for the current component ID.

It starts at `0`.

Node `0` automatically belongs to component `0`.

I do not need to compare the first node with anything before it because there is no previous node.

### Step 3: Scan consecutive values

I start the loop from index `1`.

For every current index `i`, I compare:

`nums[i] - nums[i - 1]`

I only need this subtraction because `nums` is sorted in non-decreasing order.

That means:

`nums[i] >= nums[i - 1]`

So the difference cannot be negative.

### Step 4: Detect a disconnected boundary

If the consecutive difference is greater than `maxDiff`, I increase the component ID.

For example:

`nums = [2, 5, 6, 8]`

`maxDiff = 2`

The first gap is:

`5 - 2 = 3`

Since `3 > 2`, node `0` cannot connect to node `1`.

It also cannot connect to nodes `2` or `3`, because those values are even larger.

So a new connected component starts at node `1`.

### Step 5: Keep connected nodes together

Now consider the next gap:

`6 - 5 = 1`

Since `1 <= 2`, nodes `1` and `2` have a direct edge.

They stay in the same component.

The next gap is:

`8 - 6 = 2`

This is also valid because it is equal to `maxDiff`.

So nodes `2` and `3` are directly connected.

The component array becomes:

`[0, 1, 1, 1]`

This tells me that node `0` is isolated from the other three nodes, while nodes `1`, `2`, and `3` are connected.

### Step 6: Understand indirect paths

A query does not require a direct edge.

For example, node `1` may not have a direct edge to node `3`, but the following path can still exist:

`1 -> 2 -> 3`

That is why storing connected components is enough.

If two nodes have the same component ID, some valid path connects them.

### Step 7: Process every query

For a query `[u, v]`, I compare:

`component[u]` and `component[v]`

If both values are equal, I store `true`.

If they are different, I store `false`.

No DFS or BFS is needed for individual queries.

### Step 8: Language-specific behavior

In C++, I can store the result in a `vector<bool>` and append each query result.

In Java, I use an `int[]` for component IDs and a `boolean[]` for the final answers.

In JavaScript, I use normal arrays. Strict equality is used when comparing component IDs.

In Python3, I use lists. Query pairs can be unpacked directly while processing them.

In Go, I use integer slices for component IDs and a boolean slice for the answers.

The graph logic and time complexity stay exactly the same in all five languages.

### Step 9: Handle a query from a node to itself

If a query is `[u, u]`, the answer is always `true`.

The component comparison handles this automatically because:

`component[u] == component[u]`

No special condition is needed.

### Step 10: Handle duplicate values

The array is sorted in non-decreasing order, so duplicate values are allowed.

If:

`nums[i] == nums[i - 1]`

then the difference is `0`.

Since `maxDiff` is always at least `0`, duplicate consecutive values stay in the same connected component.

### Step 11: Handle maxDiff equal to zero

When `maxDiff = 0`, only nodes with equal values can have edges.

The same algorithm still works.

Whenever two consecutive values are different, a new component starts.

Whenever they are equal, they remain in the same component.

No extra case is required.

## Examples

### Example 1

Input:

```text
n = 2
nums = [1, 3]
maxDiff = 1
queries = [[0, 0], [0, 1]]
```

Expected output:

```text
[true, false]
```

Trace:

The gap between the two values is:

`3 - 1 = 2`

Since `2 > 1`, a new component starts.

The component array becomes:

`[0, 1]`

For query `[0, 0]`, both nodes are the same, so the answer is `true`.

For query `[0, 1]`, the component IDs are different, so the answer is `false`.

### Example 2

Input:

```text
n = 4
nums = [2, 5, 6, 8]
maxDiff = 2
queries = [[0, 1], [0, 2], [1, 3], [2, 3]]
```

Expected output:

```text
[false, false, true, true]
```

Trace:

The consecutive gaps are:

`5 - 2 = 3`

`6 - 5 = 1`

`8 - 6 = 2`

The first gap is greater than `maxDiff`, so node `0` is separated from the remaining nodes.

The other two gaps are valid.

The component array becomes:

`[0, 1, 1, 1]`

Nodes `0` and `1` are in different components, so the answer is `false`.

Nodes `0` and `2` are also in different components, so the answer is `false`.

Nodes `1` and `3` are in the same component, so the answer is `true`.

Nodes `2` and `3` are also in the same component, so the answer is `true`.

### Example 3

Input:

```text
n = 5
nums = [1, 2, 4, 10, 11]
maxDiff = 2
queries = [[0, 2], [1, 4], [3, 4]]
```

Expected output:

```text
[true, false, true]
```

Trace:

The consecutive gaps are:

`2 - 1 = 1`

`4 - 2 = 2`

`10 - 4 = 6`

`11 - 10 = 1`

The gap between `4` and `10` is too large, so it splits the graph.

The component array becomes:

`[0, 0, 0, 1, 1]`

Nodes `0` and `2` are connected.

Nodes `1` and `4` are in different components.

Nodes `3` and `4` are connected.

## How to Use / Run Locally

The solution code for each language should be saved in its own file.

Since the original problem uses LeetCode method signatures, the submitted solution normally runs inside the LeetCode environment. For local testing, add a small main function or test runner that creates the input values, calls the solution method, and prints the result.

### C++

Save the C++ solution in a file such as:

`solution.cpp`

Compile it with:

```bash
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```bash
./solution
```

On Windows, run:

```bash
solution.exe
```

### Java

Save the Java solution in:

`Solution.java`

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For local testing, add a `main` method because LeetCode normally calls the solution method automatically.

### JavaScript

Save the JavaScript solution in:

`solution.js`

Run it with Node.js:

```bash
node solution.js
```

Add test inputs and a function call at the bottom of the file when testing locally.

### Python3

Save the Python solution in:

`solution.py`

Run it with:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

Add a small test case below the solution class to print the returned boolean array.

### Go

Save the Go solution in:

`main.go`

Run it directly with:

```bash
go run main.go
```

Or build it first:

```bash
go build main.go
```

Then run the generated executable.

For local testing, add a `main` function that creates the inputs and calls `pathExistenceQueries`.

## Notes & Optimizations

The most important optimization is avoiding explicit graph construction.

A direct graph-building approach could compare many pairs of nodes. In the worst case, the graph may contain a very large number of edges.

Running DFS or BFS separately for every query would also be too slow because there can be up to `10^5` queries.

Union-Find is another possible approach. I could join consecutive nodes whenever their difference is at most `maxDiff`, then check whether two nodes have the same root.

That works, but it is more complicated than needed.

Because `nums` is sorted, connected components always form continuous index segments. A simple component ID array gives the same result with less code and no Union-Find operations.

Binary search could also be used if I stored only the positions of large gaps. However, that would make each query take `O(log n)` time.

The component array is better here because preprocessing takes `O(n)` and every query takes only `O(1)`.

Important edge cases include:

* A graph with only one node.
* Queries where both endpoints are the same node.
* Duplicate values in `nums`.
* `maxDiff = 0`.
* Every node belonging to one connected component.
* Every consecutive gap being too large, making every node its own component.

The sorted order is the key reason this linear-time graph connectivity solution works.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
