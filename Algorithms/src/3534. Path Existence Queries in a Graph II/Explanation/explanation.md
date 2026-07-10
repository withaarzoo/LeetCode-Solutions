# 3534. Path Existence Queries in a Graph II

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

In LeetCode 3534, Path Existence Queries in a Graph II, we are given `n` nodes numbered from `0` to `n - 1`.

We also have an integer array `nums` and an integer `maxDiff`.

An undirected edge exists between two nodes `i` and `j` when:

`|nums[i] - nums[j]| <= maxDiff`

For every query `[u, v]`, we need to find the minimum number of edges required to travel from node `u` to node `v`.

If no path exists between the two nodes, the answer is `-1`.

The main challenge is the input size. Both the number of nodes and the number of queries can reach `100,000`, so building the full graph or running BFS for every query would be too slow.

The optimized solution uses sorting, two pointers, greedy jumps, and binary lifting with a sparse table. This allows each shortest path query to be answered in `O(log n)` time.

## Constraints

| Constraint                 | Value                           |
| -------------------------- | ------------------------------- |
| Number of nodes            | `1 <= n == nums.length <= 10^5` |
| Node values                | `0 <= nums[i] <= 10^5`          |
| Maximum allowed difference | `0 <= maxDiff <= 10^5`          |
| Number of queries          | `1 <= queries.length <= 10^5`   |
| Query format               | `queries[i] == [u_i, v_i]`      |
| Query node range           | `0 <= u_i, v_i < n`             |

These constraints rule out solutions that explicitly create all possible edges or run a graph traversal for every query.

## Intuition

My first thought was to build the graph and run BFS for each query.

That works for small inputs, but it becomes impossible here. In the worst case, many pairs of nodes can be connected, and there can also be `100,000` queries.

The key observation is that an edge depends only on the values inside `nums`.

The original positions of the nodes do not decide whether an edge exists. Only the difference between their values matters.

This made sorting the natural first step.

After sorting the nodes by their `nums` values, all nodes that can be reached directly from a position form one continuous range.

For a sorted position `i`, I only need to know the farthest position `r` such that:

`sortedValue[r] - sortedValue[i] <= maxDiff`

Every position between `i` and `r` is directly reachable.

Now the graph problem becomes a jumping problem on a sorted array.

To minimize the number of edges, I always want to jump as far as possible. Choosing a shorter reachable position uses the same one edge but makes less progress toward the destination.

The only remaining issue is speed. Following these greedy jumps one by one can still take too long.

Binary lifting solves that problem by letting me skip `1`, `2`, `4`, `8`, and larger powers of two greedy jumps at once.

## Approach

I solve the problem in five main stages.

### 1. Sort the nodes by value

I keep each node's original index and sort the nodes according to `nums[i]`.

This gives me a sorted sequence of values.

I also build a `pos` array where:

`pos[node] = position of that node in sorted order`

This is needed because every query uses original node indices.

### 2. Find the farthest one-edge destination

For every sorted position `i`, I find the farthest position that can be reached using one edge.

I use two pointers.

The right pointer only moves forward, so all one-step destinations can be calculated in linear time after sorting.

I store the result as:

`jump[0][i]`

This means the farthest sorted position reachable from `i` in one greedy jump.

### 3. Build the binary lifting table

I build a sparse table where:

`jump[p][i]`

stores the farthest position reachable from `i` after at most `2^p` greedy jumps.

The transition is:

`jump[p][i] = jump[p - 1][jump[p - 1][i]]`

In simple words, I make a large jump by combining two smaller jumps.

### 4. Convert each query into sorted positions

For a query `[u, v]`, I convert both original nodes using the `pos` array.

Because the graph is undirected, I always process the query from the smaller sorted position to the larger sorted position.

If both positions are the same, the answer is immediately `0`.

### 5. Use binary lifting to count the minimum jumps

I check the binary lifting levels from largest to smallest.

Whenever a group of jumps still leaves me strictly before the target, I take it.

This counts as many jumps as possible without reaching the destination too early.

After that, I check whether one final edge can reach the target.

If it can, I add one to the distance.

Otherwise, no path exists, so the answer is `-1`.

## Data Structures Used

### Sorted node list

I store each node together with its original index and sort by value.

This changes the graph into an ordered structure where reachable nodes form continuous ranges.

### Position mapping array

The `pos` array maps an original node index to its position after sorting.

Without this array, every query would need to search for both nodes again.

### Sorted values array

I keep the node values in sorted order.

This makes it easy to check whether two positions differ by at most `maxDiff`.

### Two pointers

I use a left position and a right pointer to find the farthest direct destination for every node.

The right pointer never moves backward, which keeps this stage at `O(n)`.

### Binary lifting table

The sparse table stores destinations after powers of two greedy jumps.

This is the main structure that reduces each query from potentially `O(n)` time to `O(log n)` time.

### Answer array

I store one result for every query in the same order as the input queries.

## Operations & Behavior Summary

The complete algorithm works like this:

1. Create a list containing every node index.
2. Sort the node indices by their corresponding values in `nums`.
3. Build a mapping from original node indices to sorted positions.
4. Store all node values in sorted order.
5. Use two pointers to find the farthest position reachable from every position in one edge.
6. Store these destinations in the first level of the binary lifting table.
7. Build the remaining sparse table levels by combining smaller jumps.
8. For each query, convert both nodes to sorted positions.
9. Swap them if necessary so the search always moves from left to right.
10. Return `0` immediately if both positions are equal.
11. Check binary lifting levels from largest to smallest.
12. Take a group of jumps whenever it still stops before the target.
13. Check whether one final edge reaches the target.
14. Return the minimum distance if reachable.
15. Return `-1` if the nodes belong to different connected components.

## Complexity

| Type             | Complexity             | Explanation                                                                                                                    |
| ---------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n log n + q log n)` | Sorting takes `O(n log n)`, building the binary lifting table takes `O(n log n)`, and each of the `q` queries takes `O(log n)` |
| Space Complexity | `O(n log n)`           | The binary lifting table stores `O(log n)` levels for all `n` sorted positions                                                 |

Here, `n` is the number of nodes and `q` is the number of queries.

The two-pointer stage takes only `O(n)` because the right pointer never moves backward.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> pathExistenceQueries(int n, vector<int>& nums, int maxDiff, vector<vector<int>>& queries) {
        // Store each value together with its original node index.
        vector<pair<int, int>> nodes;
        nodes.reserve(n);

        for (int i = 0; i < n; ++i) {
            nodes.push_back({nums[i], i});
        }

        // Sorting turns the graph problem into a left-to-right jumping problem.
        sort(nodes.begin(), nodes.end());

        // pos[node] gives the position of an original node in sorted order.
        vector<int> pos(n);

        // values stores only the sorted nums values for easier comparisons.
        vector<int> values(n);

        for (int i = 0; i < n; ++i) {
            values[i] = nodes[i].first;
            pos[nodes[i].second] = i;
        }

        // LOG is the number of levels needed for binary lifting.
        int LOG = 1;
        while ((1 << LOG) <= n) {
            ++LOG;
        }

        // jump[p][i] is the farthest position reachable after at most 2^p greedy jumps.
        vector<vector<int>> jump(LOG, vector<int>(n));

        // Use two pointers to find the farthest node reachable in one edge.
        int r = 0;

        for (int i = 0; i < n; ++i) {
            // r never moves backward, so this whole loop is O(n).
            if (r < i) {
                r = i;
            }

            // Extend r while the value difference still allows a direct edge.
            while (r + 1 < n && values[r + 1] - values[i] <= maxDiff) {
                ++r;
            }

            // This is the farthest position reachable from i in one jump.
            jump[0][i] = r;
        }

        // Build larger jumps from smaller jumps.
        for (int p = 1; p < LOG; ++p) {
            for (int i = 0; i < n; ++i) {
                // Two jumps of size 2^(p-1) make one jump of size 2^p.
                jump[p][i] = jump[p - 1][jump[p - 1][i]];
            }
        }

        // Store one answer for every query.
        vector<int> answer;
        answer.reserve(queries.size());

        for (const auto& query : queries) {
            // Convert original node indices into sorted positions.
            int left = pos[query[0]];
            int right = pos[query[1]];

            // The graph is undirected, so I always move from left to right.
            if (left > right) {
                swap(left, right);
            }

            // A node has distance 0 from itself.
            if (left == right) {
                answer.push_back(0);
                continue;
            }

            int current = left;
            int distance = 0;

            // Take the largest groups of jumps that still stop before the target.
            for (int p = LOG - 1; p >= 0; --p) {
                if (jump[p][current] < right) {
                    // Skip 2^p greedy jumps at once.
                    current = jump[p][current];
                    distance += (1 << p);
                }
            }

            // One final edge must now reach the target.
            if (jump[0][current] >= right) {
                answer.push_back(distance + 1);
            } else {
                // If even the farthest one-step jump cannot move forward enough,
                // the target lies in a different connected component.
                answer.push_back(-1);
            }
        }

        return answer;
    }
};
```

### Java

```java
class Solution {
    public int[] pathExistenceQueries(int n, int[] nums, int maxDiff, int[][] queries) {
        // order stores original node indices before sorting them by nums value.
        Integer[] order = new Integer[n];

        for (int i = 0; i < n; i++) {
            order[i] = i;
        }

        // Sort node indices according to their nums values.
        Arrays.sort(order, (a, b) -> Integer.compare(nums[a], nums[b]));

        // pos[node] gives the sorted position of an original node.
        int[] pos = new int[n];

        // values stores nums in sorted order.
        int[] values = new int[n];

        for (int i = 0; i < n; i++) {
            values[i] = nums[order[i]];
            pos[order[i]] = i;
        }

        // Find how many binary lifting levels are needed.
        int log = 1;

        while ((1 << log) <= n) {
            log++;
        }

        // jump[p][i] stores the farthest position after at most 2^p greedy jumps.
        int[][] jump = new int[log][n];

        // Find the farthest one-edge reach for every position using two pointers.
        int r = 0;

        for (int i = 0; i < n; i++) {
            // Keep r at or after the current position.
            if (r < i) {
                r = i;
            }

            // Move r while a direct edge from i is still possible.
            while (r + 1 < n && values[r + 1] - values[i] <= maxDiff) {
                r++;
            }

            // Save the farthest one-jump destination.
            jump[0][i] = r;
        }

        // Build all larger binary jumps.
        for (int p = 1; p < log; p++) {
            for (int i = 0; i < n; i++) {
                // Apply the previous jump twice.
                jump[p][i] = jump[p - 1][jump[p - 1][i]];
            }
        }

        // Store the result of every query.
        int[] answer = new int[queries.length];

        for (int q = 0; q < queries.length; q++) {
            // Convert both original nodes to sorted positions.
            int left = pos[queries[q][0]];
            int right = pos[queries[q][1]];

            // Always process the query from the smaller position to the larger one.
            if (left > right) {
                int temp = left;
                left = right;
                right = temp;
            }

            // The distance from a node to itself is zero.
            if (left == right) {
                answer[q] = 0;
                continue;
            }

            int current = left;
            int distance = 0;

            // Use the largest possible groups of jumps first.
            for (int p = log - 1; p >= 0; p--) {
                if (jump[p][current] < right) {
                    // Take 2^p greedy jumps without passing the target.
                    current = jump[p][current];
                    distance += 1 << p;
                }
            }

            // Check whether one last edge reaches the target.
            if (jump[0][current] >= right) {
                answer[q] = distance + 1;
            } else {
                // No forward progress means the nodes are disconnected.
                answer[q] = -1;
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
 * @param {number[]} nums
 * @param {number} maxDiff
 * @param {number[][]} queries
 * @return {number[]}
 */
var pathExistenceQueries = function(n, nums, maxDiff, queries) {
    // Store original node indices so I can sort them by nums value.
    const order = Array.from({ length: n }, (_, i) => i);

    // Sort nodes according to their values.
    order.sort((a, b) => nums[a] - nums[b]);

    // pos[node] gives the sorted position of an original node.
    const pos = new Array(n);

    // values stores nums in sorted order.
    const values = new Array(n);

    for (let i = 0; i < n; i++) {
        values[i] = nums[order[i]];
        pos[order[i]] = i;
    }

    // Find the number of levels needed for binary lifting.
    let LOG = 1;

    while (2 ** LOG <= n) {
        LOG++;
    }

    // jump[p][i] is the farthest position after at most 2^p greedy jumps.
    const jump = Array.from({ length: LOG }, () => new Array(n));

    // Use two pointers to find every one-jump destination.
    let r = 0;

    for (let i = 0; i < n; i++) {
        // Keep r at or after i.
        if (r < i) {
            r = i;
        }

        // Extend r while there is still a direct edge from i.
        while (r + 1 < n && values[r + 1] - values[i] <= maxDiff) {
            r++;
        }

        // Save the farthest position reachable in one edge.
        jump[0][i] = r;
    }

    // Build the binary lifting table.
    for (let p = 1; p < LOG; p++) {
        for (let i = 0; i < n; i++) {
            // Apply the previous level twice to double the jump count.
            jump[p][i] = jump[p - 1][jump[p - 1][i]];
        }
    }

    // Store one result for each query.
    const answer = new Array(queries.length);

    for (let q = 0; q < queries.length; q++) {
        // Convert original nodes into sorted positions.
        let left = pos[queries[q][0]];
        let right = pos[queries[q][1]];

        // The graph is undirected, so always move from left to right.
        if (left > right) {
            [left, right] = [right, left];
        }

        // No edge is needed when both nodes are the same.
        if (left === right) {
            answer[q] = 0;
            continue;
        }

        let current = left;
        let distance = 0;

        // Take the largest groups of jumps that still stop before the target.
        for (let p = LOG - 1; p >= 0; p--) {
            if (jump[p][current] < right) {
                // Skip 2^p greedy jumps at once.
                current = jump[p][current];
                distance += 2 ** p;
            }
        }

        // One final jump must reach the target.
        if (jump[0][current] >= right) {
            answer[q] = distance + 1;
        } else {
            // Otherwise the target belongs to another connected component.
            answer[q] = -1;
        }
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def pathExistenceQueries(self, n: int, nums: List[int], maxDiff: int, queries: List[List[int]]) -> List[int]:
        # Sort original node indices according to their nums values.
        order = sorted(range(n), key=lambda i: nums[i])

        # pos[node] gives the sorted position of an original node.
        pos = [0] * n

        # values stores nums in sorted order.
        values = [0] * n

        for i in range(n):
            values[i] = nums[order[i]]
            pos[order[i]] = i

        # bit_length gives enough levels to represent every possible jump count.
        LOG = max(1, n.bit_length())

        # jump[p][i] is the farthest position after at most 2^p greedy jumps.
        jump = [[0] * n for _ in range(LOG)]

        # Use two pointers to find the farthest one-edge reach from every position.
        r = 0

        for i in range(n):
            # Keep r at or after the current position.
            if r < i:
                r = i

            # Extend r while a direct edge from i is still allowed.
            while r + 1 < n and values[r + 1] - values[i] <= maxDiff:
                r += 1

            # Save the farthest position reachable in one jump.
            jump[0][i] = r

        # Build larger jumps by combining two smaller jumps.
        for p in range(1, LOG):
            for i in range(n):
                # Applying level p - 1 twice gives 2^p greedy jumps.
                jump[p][i] = jump[p - 1][jump[p - 1][i]]

        # Store the answer for every query.
        answer = []

        for u, v in queries:
            # Convert original nodes into sorted positions.
            left = pos[u]
            right = pos[v]

            # The graph is undirected, so always process left to right.
            if left > right:
                left, right = right, left

            # The distance from a node to itself is zero.
            if left == right:
                answer.append(0)
                continue

            current = left
            distance = 0

            # Take the largest groups of jumps that still stop before the target.
            for p in range(LOG - 1, -1, -1):
                if jump[p][current] < right:
                    # Skip 2^p greedy jumps at once.
                    current = jump[p][current]
                    distance += 1 << p

            # Check whether one final edge reaches the target.
            if jump[0][current] >= right:
                answer.append(distance + 1)
            else:
                # No path exists if the greedy jump cannot move toward the target.
                answer.append(-1)

        return answer
```

### Go

```go
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {
    // order stores original node indices before sorting by nums value.
    order := make([]int, n)

    for i := 0; i < n; i++ {
        order[i] = i
    }

    // Sort node indices according to their nums values.
    sort.Slice(order, func(i, j int) bool {
        return nums[order[i]] < nums[order[j]]
    })

    // pos[node] gives the sorted position of an original node.
    pos := make([]int, n)

    // values stores nums in sorted order.
    values := make([]int, n)

    for i := 0; i < n; i++ {
        values[i] = nums[order[i]]
        pos[order[i]] = i
    }

    // Find the number of levels needed for binary lifting.
    LOG := 1

    for (1 << LOG) <= n {
        LOG++
    }

    // jump[p][i] is the farthest position after at most 2^p greedy jumps.
    jump := make([][]int, LOG)

    for p := 0; p < LOG; p++ {
        jump[p] = make([]int, n)
    }

    // Use two pointers to find every one-jump destination.
    r := 0

    for i := 0; i < n; i++ {
        // Keep r at or after the current position.
        if r < i {
            r = i
        }

        // Extend r while a direct edge from i is still possible.
        for r+1 < n && values[r+1]-values[i] <= maxDiff {
            r++
        }

        // Save the farthest position reachable in one edge.
        jump[0][i] = r
    }

    // Build all larger binary jumps.
    for p := 1; p < LOG; p++ {
        for i := 0; i < n; i++ {
            // Apply the previous jump twice to double the number of jumps.
            jump[p][i] = jump[p-1][jump[p-1][i]]
        }
    }

    // Store one result for every query.
    answer := make([]int, len(queries))

    for q, query := range queries {
        // Convert original nodes into sorted positions.
        left := pos[query[0]]
        right := pos[query[1]]

        // The graph is undirected, so always move from left to right.
        if left > right {
            left, right = right, left
        }

        // No edge is needed when both nodes are the same.
        if left == right {
            answer[q] = 0
            continue
        }

        current := left
        distance := 0

        // Take the largest groups of jumps that still stop before the target.
        for p := LOG - 1; p >= 0; p-- {
            if jump[p][current] < right {
                // Skip 2^p greedy jumps at once.
                current = jump[p][current]
                distance += 1 << p
            }
        }

        // One final edge must reach the target.
        if jump[0][current] >= right {
            answer[q] = distance + 1
        } else {
            // Otherwise the two nodes are in different connected components.
            answer[q] = -1
        }
    }

    return answer
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The main algorithm is the same in all five languages. Only the syntax and standard library functions are different.

### Step 1: Sort nodes without losing original indices

The queries refer to original node numbers, so I cannot simply sort `nums` and forget where each value came from.

Instead, I sort node indices or pairs containing:

`(value, original index)`

After sorting, I know both the value order and the original identity of every node.

In C++, pairs work naturally for this.

In Java and JavaScript, I can sort an array of original indices using the values in `nums`.

In Python3, sorting `range(n)` with a key function gives the same result.

In Go, I store the indices in a slice and sort them with a custom comparison function.

### Step 2: Build the original-to-sorted position mapping

Suppose a node with original index `7` appears at sorted position `3`.

Then I store:

`pos[7] = 3`

This makes query conversion constant time.

For a query `[u, v]`, I can immediately find:

`left = pos[u]`

`right = pos[v]`

Without this mapping, I would need to search the sorted array for every query, which would add unnecessary work.

### Step 3: Find the farthest direct jump

For every sorted position `i`, I want the farthest position `r` where:

`values[r] - values[i] <= maxDiff`

Because the values are sorted, every position between `i` and `r` is also directly reachable.

I use a right pointer that only moves forward.

When `i` increases, I do not reset the right pointer to the beginning. This is important because resetting it would make the process quadratic in the worst case.

The result becomes the first binary lifting level:

`jump[0][i] = r`

### Step 4: Understand why the greedy jump is safe

Suppose I can directly reach positions `i + 1`, `i + 2`, and `i + 3`.

Choosing `i + 1` uses one edge.

Choosing `i + 3` also uses one edge.

Since the target is on the right, reaching `i + 3` gives at least as much progress as reaching `i + 1`.

That is why I only need the farthest direct destination.

This removes the need to store every graph edge.

### Step 5: Build larger jumps

The first level stores one greedy jump.

The next level stores two greedy jumps.

The next one stores four.

Then eight, sixteen, and so on.

If:

`jump[p - 1][i]`

is the destination after `2^(p - 1)` jumps, I apply the same level again from that destination.

So:

`jump[p][i] = jump[p - 1][jump[p - 1][i]]`

This binary lifting table is what makes fast shortest path queries possible.

### Step 6: Normalize every query

The graph is undirected.

If the first query node appears after the second one in sorted order, I swap them.

This means I only need to handle movement from left to right.

The algorithm becomes simpler because every jump table entry also moves right or stays in the same place.

### Step 7: Handle the same-node case

If both query nodes are the same, the minimum distance is `0`.

No edge is required.

This case should be handled before binary lifting.

### Step 8: Count large groups of jumps

I start from the largest binary lifting level.

Suppose a level represents `16` greedy jumps.

If taking those `16` jumps still leaves me before the target, I take them immediately.

Then I try smaller powers of two.

This is similar to building the answer using binary representation.

The important condition is that the destination must remain strictly before the target.

If I allowed a jump that reaches or passes the target during this stage, I could count more edges than necessary.

### Step 9: Check the final edge

After processing all binary lifting levels, I have taken the maximum number of jumps that still leave me before the target.

Now I check whether one direct jump reaches the target.

If:

`jump[0][current] >= target`

then the target is reachable in one more edge.

Otherwise, the current connected component cannot move far enough to reach the destination.

The answer is then `-1`.

### Language-specific behavior

The algorithm does not change between C++, Java, JavaScript, Python3, and Go.

The main differences are implementation details.

C++ uses `vector`, `pair`, and `sort`.

Java uses arrays and `Arrays.sort`.

JavaScript uses standard arrays and a custom sorting function.

Python3 uses lists, `sorted`, and `bit_length` for calculating the number of binary lifting levels.

Go uses slices and `sort.Slice`.

All five versions use the same sorted graph observation, two-pointer preprocessing, sparse table construction, and binary lifting query logic.

## Examples

### Example 1

Input:

```text
n = 5
nums = [1, 8, 3, 4, 2]
maxDiff = 3
queries = [[0, 3], [2, 4]]
```

Expected output:

```text
[1, 1]
```

After sorting the values:

```text
[1, 2, 3, 4, 8]
```

For query `[0, 3]`, the values are `1` and `4`.

Their difference is `3`, so a direct edge exists.

The minimum distance is `1`.

For query `[2, 4]`, the values are `3` and `2`.

Their difference is `1`, so they are also directly connected.

The minimum distance is `1`.

### Example 2

Input:

```text
n = 5
nums = [5, 3, 1, 9, 10]
maxDiff = 2
queries = [[0, 1], [0, 2], [2, 3], [4, 3]]
```

Expected output:

```text
[1, 2, -1, 1]
```

The sorted values are:

```text
[1, 3, 5, 9, 10]
```

Values `1`, `3`, and `5` form one connected component.

Values `9` and `10` form another connected component.

The path from value `5` to value `3` needs one edge.

The path from value `5` to value `1` needs two edges:

```text
5 -> 3 -> 1
```

There is no path from value `1` to value `9`, so the answer is `-1`.

Values `10` and `9` are directly connected, so their distance is `1`.

### Example 3

Input:

```text
n = 3
nums = [3, 6, 1]
maxDiff = 1
queries = [[0, 0], [0, 1], [1, 2]]
```

Expected output:

```text
[0, -1, -1]
```

The sorted values are:

```text
[1, 3, 6]
```

Every gap is larger than `maxDiff`.

This means no two different nodes are connected.

A node can still reach itself with distance `0`, so the first query returns `0`.

The other two queries return `-1`.

## How to Use / Run Locally

Each solution follows the LeetCode function format. If you want to run it locally, you may need to add a small test program that creates the input arrays, calls the function, and prints the result.

### C++

Save the solution as:

```text
main.cpp
```

Compile it with:

```bash
g++ -std=c++17 -O2 main.cpp -o main
```

Run it with:

```bash
./main
```

On Windows, run:

```bash
main.exe
```

### Java

Save the file as:

```text
Solution.java
```

Compile it with:

```bash
javac Solution.java
```

Run it with:

```bash
java Solution
```

For local testing, add a `main` method because LeetCode calls the solution method automatically, but a local Java program does not.

### JavaScript

Save the solution as:

```text
solution.js
```

Make sure Node.js is installed.

Run it with:

```bash
node solution.js
```

Add your own test input and `console.log` calls when testing outside LeetCode.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```bash
python3 solution.py
```

On some systems, the command may be:

```bash
python solution.py
```

Add a small test section at the bottom of the file to create a `Solution` object and call the method.

### Go

Save the solution as:

```text
main.go
```

Run it directly with:

```bash
go run main.go
```

Or build it first:

```bash
go build main.go
```

Then run the generated executable.

For local testing, add a `main` function because LeetCode only requires the solution function.

## Notes & Optimizations

The most important optimization is avoiding the actual graph.

A direct graph construction can require too many edges. Since connectivity depends only on value differences, sorting gives all the structure needed for the solution.

Running BFS for every query is also too slow. Even if the graph were stored efficiently, `100,000` separate graph searches would not fit the constraints.

A simpler approach could repeatedly take the farthest greedy jump for each query. The greedy idea is correct, but one-by-one jumping can still take `O(n)` time for a single query.

Binary lifting improves this to `O(log n)` per query.

Another possible method is to use binary search to find the farthest one-step destination for every sorted position. That would take `O(n log n)` for this preprocessing stage.

The two-pointer method is better because it finds all one-step destinations in `O(n)` after sorting.

Duplicate values are handled naturally. Their difference is `0`, so they are directly connected whenever `maxDiff >= 0`, which is always true under the given constraints.

When `maxDiff = 0`, only nodes with equal values can connect to each other.

A large gap between consecutive sorted values creates a boundary between connected components. The binary lifting process automatically stops at such a boundary, so a separate connected-component data structure is not required.

The final solution is optimized for the full LeetCode constraints and combines several important competitive programming techniques: sorting, two pointers, greedy reachability, sparse table preprocessing, and binary lifting.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
