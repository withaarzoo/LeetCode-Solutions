# 1345. Jump Game IV – BFS Graph Solution | LeetCode Explained

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

---

# Problem Summary

LeetCode 1345: Jump Game IV is a graph and BFS-based problem where we need to find the minimum number of jumps required to reach the last index of an array.

From any index `i`, we can move to:

* `i + 1`
* `i - 1`
* Any index `j` where `arr[i] == arr[j]`

The goal is to return the smallest number of moves needed to reach the last index starting from index `0`.

This problem looks simple at first, but handling repeated values efficiently is the real challenge. A normal brute-force solution becomes too slow because the same value can appear many times.

That is why an optimized BFS solution works best here.

---

# Constraints

| Constraint    | Value                         |
| ------------- | ----------------------------- |
| Array Length  | `1 <= arr.length <= 5 * 10^4` |
| Element Range | `-10^8 <= arr[i] <= 10^8`     |

---

# Intuition

The first thing I noticed was that this problem behaves like a graph.

Every index acts like a node.

From one node, I can jump:

* left
* right
* same-value indices

Whenever I see a shortest path problem in an unweighted graph, BFS is usually the best choice because BFS guarantees the minimum number of moves.

The tricky part is handling duplicate values efficiently.

If I repeatedly visit all indices having the same value, the solution becomes very slow for large inputs.

So the key optimization is:

After processing all indices for a value once, I remove them from the hashmap so they are never processed again.

That single optimization reduces the complexity dramatically.

---

# Approach

I solved this problem using Breadth First Search (BFS).

Step-by-step process:

1. Create a hashmap where:

   * key = array value
   * value = list of indices containing that value

2. Start BFS from index `0`.

3. For every current index:

   * move left
   * move right
   * jump to all same-value indices

4. Use a visited array to avoid revisiting indices.

5. After processing all indices for a value, clear that list from the hashmap.

   * This prevents repeated work.

6. The first time BFS reaches the last index, return the current number of steps.

This gives the minimum jumps needed.

---

# Data Structures Used

| Data Structure         | Purpose                            |
| ---------------------- | ---------------------------------- |
| HashMap / Dictionary   | Stores all indices for each value  |
| Queue                  | Used for BFS traversal             |
| Visited Array          | Prevents revisiting the same index |
| Dynamic Arrays / Lists | Stores grouped indices efficiently |

---

# Operations & Behavior Summary

Here is what the algorithm does internally:

1. Store all same-value indices together.
2. Push index `0` into the BFS queue.
3. Process nodes level by level.
4. From each index:

   * try left jump
   * try right jump
   * try same-value jumps
5. Mark every visited index immediately.
6. Clear processed same-value groups.
7. Continue until the last index is reached.

Because BFS processes level-by-level, the first valid path is always the shortest one.

---

# Complexity

| Type             | Complexity | Explanation                                           |
| ---------------- | ---------- | ----------------------------------------------------- |
| Time Complexity  | `O(n)`     | Every index and value-group is processed at most once |
| Space Complexity | `O(n)`     | HashMap, queue, and visited array use extra memory    |

Where `n` is the size of the array.

---

# Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minJumps(vector<int>& arr) {
        int n = arr.size();

        // If array has only one element, no jump is needed
        if (n == 1) return 0;

        // Map each value to all its indices
        unordered_map<int, vector<int>> mp;

        for (int i = 0; i < n; i++) {
            mp[arr[i]].push_back(i);
        }

        // Queue for BFS
        queue<int> q;

        // Visited array to avoid revisiting indices
        vector<bool> visited(n, false);

        q.push(0);
        visited[0] = true;

        int steps = 0;

        while (!q.empty()) {

            int size = q.size();

            // Process one BFS level at a time
            while (size--) {

                int idx = q.front();
                q.pop();

                // If last index reached, return answer
                if (idx == n - 1) {
                    return steps;
                }

                // Move to index - 1
                if (idx - 1 >= 0 && !visited[idx - 1]) {
                    visited[idx - 1] = true;
                    q.push(idx - 1);
                }

                // Move to index + 1
                if (idx + 1 < n && !visited[idx + 1]) {
                    visited[idx + 1] = true;
                    q.push(idx + 1);
                }

                // Move to all same-value indices
                for (int nextIdx : mp[arr[idx]]) {

                    if (!visited[nextIdx]) {
                        visited[nextIdx] = true;
                        q.push(nextIdx);
                    }
                }

                // Clear the list so we never process
                // same-value indices again
                mp[arr[idx]].clear();
            }

            // One BFS level completed
            steps++;
        }

        return -1;
    }
};
```

### Java

```java
class Solution {
    public int minJumps(int[] arr) {

        int n = arr.length;

        // No jump needed if only one element exists
        if (n == 1) return 0;

        // Store all indices for every value
        HashMap<Integer, List<Integer>> map = new HashMap<>();

        for (int i = 0; i < n; i++) {
            map.computeIfAbsent(arr[i], k -> new ArrayList<>()).add(i);
        }

        // Queue for BFS
        Queue<Integer> queue = new LinkedList<>();

        // Visited array
        boolean[] visited = new boolean[n];

        queue.offer(0);
        visited[0] = true;

        int steps = 0;

        while (!queue.isEmpty()) {

            int size = queue.size();

            // Process current BFS level
            while (size-- > 0) {

                int idx = queue.poll();

                // Last index reached
                if (idx == n - 1) {
                    return steps;
                }

                // Move left
                if (idx - 1 >= 0 && !visited[idx - 1]) {
                    visited[idx - 1] = true;
                    queue.offer(idx - 1);
                }

                // Move right
                if (idx + 1 < n && !visited[idx + 1]) {
                    visited[idx + 1] = true;
                    queue.offer(idx + 1);
                }

                // Move to same-value indices
                for (int nextIdx : map.get(arr[idx])) {

                    if (!visited[nextIdx]) {
                        visited[nextIdx] = true;
                        queue.offer(nextIdx);
                    }
                }

                // Remove repeated processing
                map.get(arr[idx]).clear();
            }

            // One level completed
            steps++;
        }

        return -1;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @return {number}
 */
var minJumps = function(arr) {

    const n = arr.length;

    // No jump needed
    if (n === 1) return 0;

    // Store indices for each value
    const map = new Map();

    for (let i = 0; i < n; i++) {

        if (!map.has(arr[i])) {
            map.set(arr[i], []);
        }

        map.get(arr[i]).push(i);
    }

    // BFS queue
    const queue = [0];

    // Visited array
    const visited = new Array(n).fill(false);

    visited[0] = true;

    let steps = 0;

    while (queue.length > 0) {

        let size = queue.length;

        // Process one BFS level
        while (size--) {

            const idx = queue.shift();

            // Last index reached
            if (idx === n - 1) {
                return steps;
            }

            // Move left
            if (idx - 1 >= 0 && !visited[idx - 1]) {
                visited[idx - 1] = true;
                queue.push(idx - 1);
            }

            // Move right
            if (idx + 1 < n && !visited[idx + 1]) {
                visited[idx + 1] = true;
                queue.push(idx + 1);
            }

            // Jump to same-value indices
            for (const nextIdx of map.get(arr[idx])) {

                if (!visited[nextIdx]) {
                    visited[nextIdx] = true;
                    queue.push(nextIdx);
                }
            }

            // Clear processed indices
            map.set(arr[idx], []);
        }

        // Next BFS level
        steps++;
    }

    return -1;
};
```

### Python3

```python
class Solution:
    def minJumps(self, arr: List[int]) -> int:

        n = len(arr)

        # No jump needed
        if n == 1:
            return 0

        from collections import defaultdict, deque

        # Store all indices for every value
        mp = defaultdict(list)

        for i, val in enumerate(arr):
            mp[val].append(i)

        # BFS queue
        q = deque([0])

        # Visited array
        visited = [False] * n

        visited[0] = True

        steps = 0

        while q:

            size = len(q)

            # Process one BFS level
            for _ in range(size):

                idx = q.popleft()

                # Last index reached
                if idx == n - 1:
                    return steps

                # Move left
                if idx - 1 >= 0 and not visited[idx - 1]:
                    visited[idx - 1] = True
                    q.append(idx - 1)

                # Move right
                if idx + 1 < n and not visited[idx + 1]:
                    visited[idx + 1] = True
                    q.append(idx + 1)

                # Move to same-value indices
                for next_idx in mp[arr[idx]]:

                    if not visited[next_idx]:
                        visited[next_idx] = True
                        q.append(next_idx)

                # Clear processed group
                mp[arr[idx]].clear()

            # Next BFS level
            steps += 1

        return -1
```

### Go

```go
func minJumps(arr []int) int {

 n := len(arr)

 // No jump needed
 if n == 1 {
  return 0
 }

 // Store all indices for every value
 mp := make(map[int][]int)

 for i, val := range arr {
  mp[val] = append(mp[val], i)
 }

 // BFS queue
 queue := []int{0}

 // Visited array
 visited := make([]bool, n)

 visited[0] = true

 steps := 0

 for len(queue) > 0 {

  size := len(queue)

  // Process one BFS level
  for i := 0; i < size; i++ {

   idx := queue[0]
   queue = queue[1:]

   // Last index reached
   if idx == n-1 {
    return steps
   }

   // Move left
   if idx-1 >= 0 && !visited[idx-1] {
    visited[idx-1] = true
    queue = append(queue, idx-1)
   }

   // Move right
   if idx+1 < n && !visited[idx+1] {
    visited[idx+1] = true
    queue = append(queue, idx+1)
   }

   // Move to same-value indices
   for _, nextIdx := range mp[arr[idx]] {

    if !visited[nextIdx] {
     visited[nextIdx] = true
     queue = append(queue, nextIdx)
    }
   }

   // Clear processed group
   mp[arr[idx]] = []int{}
  }

  // One BFS level completed
  steps++
 }

 return -1
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same across all five languages.

Only syntax changes.

### Building the HashMap

The first step is grouping indices by their values.

Example:

```text
arr = [7,6,9,6,9,6,9,7]

7 -> [0,7]
6 -> [1,3,5]
9 -> [2,4,6]
```

This allows instant access to every same-value jump.

Without this structure, we would need to scan the entire array every time, which would be too slow.

---

### Starting BFS

BFS starts from index `0`.

The queue stores indices that still need processing.

A visited array is also necessary because:

* we should never process the same index twice
* repeated visits create unnecessary loops
* BFS performance depends heavily on avoiding duplicate work

---

### BFS Level Traversal

BFS works level-by-level.

Each level represents one jump distance.

That means:

* first layer = 0 jumps
* second layer = 1 jump
* third layer = 2 jumps

So when the last index is reached, the current BFS level directly becomes the answer.

---

### Left and Right Movement

For every index:

* check `i - 1`
* check `i + 1`

Before moving:

* ensure the index stays inside array bounds
* ensure it was not visited before

This handles normal adjacent movement.

---

### Same-value Jumps

This is the most important part of the problem.

If:

```text
arr[i] == arr[j]
```

then we can jump directly from `i` to `j`.

The hashmap gives all these indices instantly.

This creates shortcut paths that make BFS efficient.

---

### Why Clearing the HashMap Matters

Suppose a value appears thousands of times.

Without optimization:

* every occurrence would repeatedly process the same group
* total operations would explode

So after processing:

```text
map[arr[i]]
```

we immediately clear it.

That guarantees every value-group is processed only once.

This is the reason the final solution runs in linear time.

---

### Language-specific Notes

#### C++

* `unordered_map` gives average O(1) lookup time.
* `queue<int>` is perfect for BFS.
* `vector<bool>` works for visited tracking.

#### Java

* `HashMap<Integer, List<Integer>>` stores grouped indices.
* `Queue<Integer>` handles BFS traversal cleanly.
* `boolean[]` is memory efficient.

#### JavaScript

* `Map` stores grouped indices.
* Arrays work as queues, though `shift()` can be slightly slower.
* Logic remains exactly the same.

#### Python3

* `defaultdict(list)` simplifies hashmap creation.
* `deque` is important because normal lists are slower for popping from the front.
* Python syntax makes BFS very compact.

#### Go

* Maps store grouped indices.
* Slices work as queues.
* Manual queue handling is required compared to higher-level languages.

---

# Examples

## Example 1

### Input

```text
arr = [100,-23,-23,404,100,23,23,23,3,404]
```

### Output

```text
3
```

### Explanation

One optimal path:

```text
0 -> 4 -> 3 -> 9
```

Steps:

* Jump from index 0 to index 4 because both contain `100`
* Jump from index 4 to index 3
* Jump from index 3 to index 9 because both contain `404`

Total jumps = 3

---

## Example 2

### Input

```text
arr = [7]
```

### Output

```text
0
```

### Explanation

The starting index is already the last index.

No movement is needed.

---

## Example 3

### Input

```text
arr = [7,6,9,6,9,6,9,7]
```

### Output

```text
1
```

### Explanation

We can directly jump:

```text
0 -> 7
```

because both indices contain `7`.

---

# How to Use / Run Locally

## C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

## Java

Compile:

```bash
javac Solution.java
```

Run:

```bash
java Solution
```

---

## JavaScript

Run using Node.js:

```bash
node solution.js
```

---

## Python3

Run:

```bash
python solution.py
```

---

## Go

Run:

```bash
go run main.go
```

---

# Notes & Optimizations

* BFS is the best approach because this is a shortest path problem in an unweighted graph.
* Clearing processed same-value groups is the most important optimization.
* Without hashmap cleanup, the solution may become too slow on large duplicate-heavy inputs.
* A DFS solution would not guarantee the minimum jumps.
* A brute-force approach would fail due to time limits.
* Using a deque in Python is important for efficient BFS performance.
* This problem is a great example of combining graph traversal with hashmap optimization.

---

# Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
