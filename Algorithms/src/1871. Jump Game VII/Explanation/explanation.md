# 1871. Jump Game VII

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

LeetCode 1871: Jump Game VII is a binary string traversal problem where we need to determine whether it is possible to reach the last index of the string.

We are given:

* A binary string `s`
* Two integers:

  * `minJump`
  * `maxJump`

We start from index `0`, and we are only allowed to jump to positions containing `'0'`.

From any current index `i`, we can jump to any index `j` such that:

* `i + minJump <= j <= i + maxJump`
* `s[j] == '0'`

The goal is to check whether the last index of the string can be reached.

This problem is a popular BFS + greedy optimization problem in Data Structures and Algorithms because a direct brute-force solution becomes too slow for large inputs.

---

## Constraints

| Constraint                           | Value                             |
| ------------------------------------ | --------------------------------- |
| `2 <= s.length <= 10^5`              | String size can be large          |
| `s[i]`                               | Either `'0'` or `'1'`             |
| `s[0] == '0'`                        | Starting position is always valid |
| `1 <= minJump <= maxJump < s.length` | Valid jump range                  |

---

## Intuition

My first thought was to use BFS because every reachable index can lead to many other indices.

But then I noticed a problem.

If I try every possible jump for every position, the same ranges get scanned again and again. That creates an almost quadratic time complexity, which will fail for large test cases.

So I needed a way to avoid rechecking indices repeatedly.

The important observation was:

Once I already scanned a range of positions, I never need to scan that range again.

That is where the `far` pointer optimization comes in. It keeps track of the farthest index already processed, making the whole traversal linear.

This turns the problem into an optimized BFS traversal problem.

---

## Approach

I used a BFS-based solution with a greedy optimization.

Step-by-step strategy:

1. Start BFS from index `0`.
2. Maintain a queue of reachable indices.
3. For every current index:

   * Calculate the valid jump range.
4. Only process positions:

   * inside the string
   * containing `'0'`
   * not visited before
5. Use a variable called `far`:

   * It stores the farthest index range already scanned.
   * This prevents duplicate range traversal.
6. If the last index is reached:

   * return `true`
7. If BFS finishes without reaching the end:

   * return `false`

This approach is efficient because every index is processed only once.

---

## Data Structures Used

| Data Structure           | Why It Was Used                                   |
| ------------------------ | ------------------------------------------------- |
| Queue                    | Used for BFS traversal of reachable indices       |
| Visited Array            | Prevents revisiting the same index multiple times |
| Integer Variable (`far`) | Tracks the farthest already-processed position    |

---

## Operations & Behavior Summary

The algorithm works like this:

1. Push index `0` into the queue.
2. Mark index `0` as visited.
3. Repeatedly pop indices from the queue.
4. From the current index:

   * compute the valid jump window
5. Scan only the new portion of the range.
6. For every valid `'0'` position:

   * mark visited
   * push into queue
7. Continue until:

   * last index is reached
   * or queue becomes empty

The `far` optimization is what keeps the solution fast.

---

## Complexity

| Type             | Complexity | Explanation                                         |
| ---------------- | ---------- | --------------------------------------------------- |
| Time Complexity  | `O(n)`     | Every index of the string is processed at most once |
| Space Complexity | `O(n)`     | Queue and visited array may store up to `n` indices |

Where:

* `n` = length of the binary string `s`

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool canReach(string s, int minJump, int maxJump) {
        int n = s.size();

        // Queue for BFS traversal
        queue<int> q;

        // Visited array to avoid revisiting indices
        vector<bool> visited(n, false);

        // Start from index 0
        q.push(0);
        visited[0] = true;

        // Keeps track of the farthest processed index
        int far = 0;

        while (!q.empty()) {
            int i = q.front();
            q.pop();

            // If we reached last index, answer is true
            if (i == n - 1)
                return true;

            // Calculate jump range
            int start = max(i + minJump, far + 1);
            int end = min(i + maxJump, n - 1);

            // Explore all valid next positions
            for (int j = start; j <= end; j++) {

                // Only visit positions containing '0'
                if (s[j] == '0' && !visited[j]) {
                    visited[j] = true;
                    q.push(j);
                }
            }

            // Update farthest processed position
            far = max(far, end);
        }

        return false;
    }
};
```

### Java

```java
class Solution {
    public boolean canReach(String s, int minJump, int maxJump) {

        int n = s.length();

        // Queue for BFS traversal
        Queue<Integer> queue = new LinkedList<>();

        // Visited array
        boolean[] visited = new boolean[n];

        // Start from index 0
        queue.offer(0);
        visited[0] = true;

        // Farthest processed index
        int far = 0;

        while (!queue.isEmpty()) {

            int i = queue.poll();

            // Reached destination
            if (i == n - 1) {
                return true;
            }

            // Valid jump range
            int start = Math.max(i + minJump, far + 1);
            int end = Math.min(i + maxJump, n - 1);

            // Check every possible next position
            for (int j = start; j <= end; j++) {

                // Only move to '0'
                if (s.charAt(j) == '0' && !visited[j]) {
                    visited[j] = true;
                    queue.offer(j);
                }
            }

            // Update processed boundary
            far = Math.max(far, end);
        }

        return false;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} s
 * @param {number} minJump
 * @param {number} maxJump
 * @return {boolean}
 */
var canReach = function(s, minJump, maxJump) {

    const n = s.length;

    // Queue for BFS
    const queue = [0];

    // Visited array
    const visited = new Array(n).fill(false);
    visited[0] = true;

    // Pointer for queue traversal
    let front = 0;

    // Farthest processed index
    let far = 0;

    while (front < queue.length) {

        const i = queue[front++];

        // Reached last index
        if (i === n - 1) {
            return true;
        }

        // Valid jump range
        const start = Math.max(i + minJump, far + 1);
        const end = Math.min(i + maxJump, n - 1);

        // Explore possible next positions
        for (let j = start; j <= end; j++) {

            // Only move to positions with '0'
            if (s[j] === '0' && !visited[j]) {
                visited[j] = true;
                queue.push(j);
            }
        }

        // Update farthest processed index
        far = Math.max(far, end);
    }

    return false;
};
```

### Python3

```python
class Solution:
    def canReach(self, s: str, minJump: int, maxJump: int) -> bool:

        n = len(s)

        # Queue for BFS traversal
        from collections import deque
        q = deque([0])

        # Visited array
        visited = [False] * n
        visited[0] = True

        # Farthest processed position
        far = 0

        while q:

            i = q.popleft()

            # If last index is reached
            if i == n - 1:
                return True

            # Valid jump range
            start = max(i + minJump, far + 1)
            end = min(i + maxJump, n - 1)

            # Explore all possible next positions
            for j in range(start, end + 1):

                # Only move to positions containing '0'
                if s[j] == '0' and not visited[j]:
                    visited[j] = True
                    q.append(j)

            # Update processed boundary
            far = max(far, end)

        return False
```

### Go

```go
func canReach(s string, minJump int, maxJump int) bool {

    n := len(s)

    // Queue for BFS traversal
    queue := []int{0}

    // Visited array
    visited := make([]bool, n)
    visited[0] = true

    // Pointer for queue front
    front := 0

    // Farthest processed index
    far := 0

    for front < len(queue) {

        i := queue[front]
        front++

        // If last index is reached
        if i == n-1 {
            return true
        }

        // Calculate valid jump range
        start := max(i+minJump, far+1)
        end := min(i+maxJump, n-1)

        // Explore all possible next positions
        for j := start; j <= end; j++ {

            // Only move to positions with '0'
            if s[j] == '0' && !visited[j] {
                visited[j] = true
                queue = append(queue, j)
            }
        }

        // Update farthest processed position
        if end > far {
            far = end
        }
    }

    return false
}

// Helper function for maximum
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Helper function for minimum
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic stays exactly the same in all five languages.

The only differences are syntax and built-in library usage.

### Step 1 — Initialize BFS

We begin from index `0`.

That position is guaranteed to contain `'0'`.

So we:

* push it into the queue
* mark it as visited

This becomes the starting point of traversal.

---

### Step 2 — Process Reachable Indices

The queue stores all indices that are currently reachable.

We repeatedly:

* remove one index from the queue
* explore all valid next jumps

This is standard BFS behavior.

---

### Step 3 — Check for Destination

Before processing jumps, we check:

* Did we reach the last index?

If yes:

* return `true` immediately

No more processing is needed.

---

### Step 4 — Compute Jump Window

From current index `i`, valid jumps are:

* `i + minJump`
  to
* `i + maxJump`

But scanning the full range every time would repeat work.

So we use:

* `far + 1`

This skips already-scanned positions.

That optimization is the key reason this solution passes large constraints.

---

### Step 5 — Visit Valid Positions

For every new position in the range:

We check:

* Is it inside bounds?
* Does it contain `'0'`?
* Was it already visited?

If all conditions are true:

* mark it visited
* push it into the queue

This ensures every index enters the queue at most once.

---

### Step 6 — Update Processed Boundary

After scanning a range, we update:

* `far`

This tells future iterations:

* "everything up to this point is already checked"

Without this optimization, the same ranges would get scanned repeatedly.

---

### Step 7 — Return False if BFS Ends

If the queue becomes empty before reaching the last index:

* the destination is unreachable

So we return `false`.

---

## Examples

### Example 1

Input:

```text
s = "011010"
minJump = 2
maxJump = 3
```

Output:

```text
true
```

Explanation:

* Start at index `0`
* Jump to index `3`
* Jump to index `5`
* Reached last index

---

### Example 2

Input:

```text
s = "01101110"
minJump = 2
maxJump = 3
```

Output:

```text
false
```

Explanation:

No valid sequence of jumps can reach the final index.

---

### Example 3

Input:

```text
s = "0000000"
minJump = 2
maxJump = 5
```

Output:

```text
true
```

Explanation:

Multiple paths are possible because all positions contain `'0'`.

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
javac Solution.java
```

Run:

```bash
java Solution
```

---

### JavaScript

Run using Node.js:

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

* A brute-force solution will timeout because it repeatedly scans overlapping ranges.
* BFS alone is not enough for this problem.
* The `far` pointer optimization is what reduces the complexity to `O(n)`.
* Every index is processed only once.
* This problem is a good example of combining:

  * BFS
  * greedy range processing
  * visited tracking

Alternative approaches:

* Dynamic Programming
* Prefix Sum optimization

But BFS with range optimization is usually the cleanest and fastest solution.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
