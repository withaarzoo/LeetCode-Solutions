# 1306. Jump Game III

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

## Problem Summary

LeetCode 1306: Jump Game III is a graph traversal and DFS/BFS based problem where we are given an array of non-negative integers and a starting index.

From any index `i`, we can jump to:

* `i + arr[i]`
* `i - arr[i]`

The goal is to determine whether it is possible to reach any index that contains the value `0`.

The challenge is that some indexes may create cycles. If we do not track visited positions, the algorithm can keep moving forever between the same indexes.

This problem is a great example of:

* Depth First Search (DFS)
* Breadth First Search (BFS)
* Graph traversal using arrays
* Visited state tracking
* Recursive problem solving

---

## Constraints

| Constraint                    | Value                          |
| ----------------------------- | ------------------------------ |
| `1 <= arr.length <= 5 * 10^4` | Array size can be large        |
| `0 <= arr[i] < arr.length`    | Every value is valid for jumps |
| `0 <= start < arr.length`     | Starting index is always valid |

---

## Intuition

The first thing I noticed was that every index behaves like a node in a graph.

From one index, I can move to at most two other indexes:

* forward jump
* backward jump

That immediately made this look like a graph traversal problem.

Then I realized something important.
If I keep revisiting the same indexes, the recursion can get stuck in an infinite loop.

So I needed a way to remember which indexes were already explored.

Once I added a visited array, the problem became simple:

* explore reachable indexes
* stop if I find a `0`
* avoid revisiting old positions

This naturally fits DFS or BFS.

---

## Approach

I used Depth First Search (DFS) to solve the problem.

Here is the full strategy:

1. Create a `visited` array.
2. Start DFS from the given `start` index.
3. If the current index goes outside the array, stop that path.
4. If the index was already visited, skip it.
5. If the current value is `0`, return `true`.
6. Mark the current index as visited.
7. Explore:

   * `index + arr[index]`
   * `index - arr[index]`
8. If either path reaches `0`, return `true`.
9. If all paths fail, return `false`.

This guarantees that every index is visited at most once.

---

## Data Structures Used

| Data Structure  | Purpose                                        |
| --------------- | ---------------------------------------------- |
| Array / List    | Stores input values                            |
| Visited Array   | Prevents revisiting indexes and infinite loops |
| Recursion Stack | Used by DFS traversal                          |

---

## Operations & Behavior Summary

The algorithm works like this:

1. Start from the given index.
2. Check if the current index is valid.
3. Check whether this position was already visited.
4. If the value at the current position is `0`, stop and return `true`.
5. Mark the current index as visited.
6. Try jumping forward.
7. Try jumping backward.
8. Continue recursively until:

   * a `0` is found
   * or all possible paths are exhausted

This is basically graph traversal using recursion.

---

## Complexity

| Type             | Complexity | Explanation                             |
| ---------------- | ---------- | --------------------------------------- |
| Time Complexity  | `O(n)`     | Each index is visited only once         |
| Space Complexity | `O(n)`     | Extra visited array and recursion stack |

Where:

* `n` = total number of elements in the array

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    
    // DFS function to explore reachable indexes
    bool dfs(vector<int>& arr, int index, vector<bool>& visited) {
        
        // If index goes outside array, this path is invalid
        if (index < 0 || index >= arr.size()) {
            return false;
        }

        // If already visited, no need to process again
        if (visited[index]) {
            return false;
        }

        // If current value is 0, answer is found
        if (arr[index] == 0) {
            return true;
        }

        // Mark current index as visited
        visited[index] = true;

        // Move forward
        int forward = index + arr[index];

        // Move backward
        int backward = index - arr[index];

        // Return true if any path reaches value 0
        return dfs(arr, forward, visited) || dfs(arr, backward, visited);
    }

    bool canReach(vector<int>& arr, int start) {
        
        // Track visited indexes
        vector<bool> visited(arr.size(), false);

        // Start DFS from given index
        return dfs(arr, start, visited);
    }
};
```

### Java

```java
class Solution {
    
    // DFS helper function
    private boolean dfs(int[] arr, int index, boolean[] visited) {

        // Invalid index
        if (index < 0 || index >= arr.length) {
            return false;
        }

        // Skip already visited indexes
        if (visited[index]) {
            return false;
        }

        // Found value 0
        if (arr[index] == 0) {
            return true;
        }

        // Mark current index as visited
        visited[index] = true;

        // Explore both directions
        int forward = index + arr[index];
        int backward = index - arr[index];

        // Return true if any direction reaches 0
        return dfs(arr, forward, visited) || dfs(arr, backward, visited);
    }

    public boolean canReach(int[] arr, int start) {
        
        // Visited array to avoid cycles
        boolean[] visited = new boolean[arr.length];

        // Start DFS
        return dfs(arr, start, visited);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @param {number} start
 * @return {boolean}
 */
var canReach = function(arr, start) {
    
    // Array to track visited indexes
    const visited = new Array(arr.length).fill(false);

    // DFS function
    const dfs = (index) => {

        // Invalid index
        if (index < 0 || index >= arr.length) {
            return false;
        }

        // Skip visited indexes
        if (visited[index]) {
            return false;
        }

        // Found value 0
        if (arr[index] === 0) {
            return true;
        }

        // Mark current index as visited
        visited[index] = true;

        // Explore forward and backward
        return dfs(index + arr[index]) || dfs(index - arr[index]);
    };

    // Start DFS from given index
    return dfs(start);
};
```

### Python3

```python
class Solution:
    def canReach(self, arr: List[int], start: int) -> bool:
        
        # Visited array to avoid infinite loops
        visited = [False] * len(arr)

        # DFS function
        def dfs(index):

            # Invalid index
            if index < 0 or index >= len(arr):
                return False

            # Skip already visited indexes
            if visited[index]:
                return False

            # Found value 0
            if arr[index] == 0:
                return True

            # Mark current index as visited
            visited[index] = True

            # Explore both directions
            return dfs(index + arr[index]) or dfs(index - arr[index])

        # Start DFS from given index
        return dfs(start)
```

### Go

```go
func canReach(arr []int, start int) bool {
    
    // Visited array to prevent revisiting indexes
    visited := make([]bool, len(arr))

    // DFS function
    var dfs func(int) bool

    dfs = func(index int) bool {

        // Invalid index
        if index < 0 || index >= len(arr) {
            return false
        }

        // Skip already visited indexes
        if visited[index] {
            return false
        }

        // Found value 0
        if arr[index] == 0 {
            return true
        }

        // Mark current index as visited
        visited[index] = true

        // Explore forward and backward jumps
        return dfs(index+arr[index]) || dfs(index-arr[index])
    }

    // Start DFS from given index
    return dfs(start)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same in all five languages.
Only syntax changes.

### Step 1 — Create a Visited Array

The first thing the solution does is create a visited array.

This array keeps track of indexes that were already processed.

Without this, the algorithm could repeatedly visit the same positions forever.

Example:

```text
1 -> 4 -> 1 -> 4
```

That would create an infinite loop.

The visited array prevents that problem.

---

### Step 2 — Start DFS from the Given Index

The traversal begins from the `start` index.

At this point, the algorithm tries to explore every reachable path.

The idea is simple:

* keep moving
* stop immediately if a `0` is found

---

### Step 3 — Check Array Bounds

Before doing anything else, the algorithm checks whether the current index is valid.

Invalid indexes happen when jumps go outside the array.

Example:

```text
index = 3
arr[3] = 5
3 + 5 = 8
```

If index `8` does not exist, that path becomes invalid.

So the algorithm immediately stops exploring that direction.

---

### Step 4 — Skip Already Visited Indexes

If an index was already processed earlier, the algorithm ignores it.

This is important because revisiting old positions wastes time and may create cycles.

The visited check makes the algorithm efficient.

---

### Step 5 — Check for Value 0

This is the success condition.

If:

```text
arr[index] == 0
```

then the problem is solved.

The algorithm immediately returns `true`.

No further traversal is needed.

---

### Step 6 — Mark Current Index as Visited

The current position is marked visited before exploring neighbors.

This ordering matters.

If we delay marking, another recursive call could revisit the same index before it gets marked.

That would create duplicate work.

---

### Step 7 — Explore Both Directions

From every index, there are only two possible moves:

```text
index + arr[index]
index - arr[index]
```

The algorithm explores both.

If either path succeeds, the answer becomes `true`.

---

### Step 8 — Return Final Result

If every reachable path gets explored and no `0` is found, the algorithm returns `false`.

That means reaching a zero-value index is impossible.

---

## Examples

### Example 1

Input:

```text
arr = [4,2,3,0,3,1,2]
start = 5
```

Output:

```text
true
```

Explanation:

```text
5 -> 4 -> 1 -> 3
```

At index `3`, the value is `0`.

So the answer is `true`.

---

### Example 2

Input:

```text
arr = [4,2,3,0,3,1,2]
start = 0
```

Output:

```text
true
```

Explanation:

```text
0 -> 4 -> 1 -> 3
```

Index `3` contains `0`.

---

### Example 3

Input:

```text
arr = [3,0,2,1,2]
start = 2
```

Output:

```text
false
```

Explanation:

No valid sequence of jumps can reach index `1`, which contains `0`.

---

## How to Use / Run Locally

### C++

Compile:

```bash
g++ main.cpp -o main
```

Run:

```bash
./main
```

---

### Java

Compile:

```bash
javac Main.java
```

Run:

```bash
java Main
```

---

### JavaScript

Run using Node.js:

```bash
node main.js
```

---

### Python3

Run:

```bash
python main.py
```

---

### Go

Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* DFS and BFS both work for this problem.
* DFS is shorter and easier to write recursively.
* BFS may avoid deep recursion in some languages.
* The visited array is mandatory.
* Without visited tracking, the solution can enter infinite loops.
* Every index is processed only once, making the algorithm efficient even for large inputs.
* Since constraints are up to `5 * 10^4`, an `O(n)` solution is the best practical approach.

Possible alternative approaches:

* Breadth First Search (BFS)
* Iterative DFS using stack

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
