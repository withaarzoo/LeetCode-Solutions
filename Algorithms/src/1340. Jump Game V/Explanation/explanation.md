# 1340. Jump Game V

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

LeetCode 1340: Jump Game V is a hard Dynamic Programming and DFS problem where we need to find the maximum number of indices we can visit in an array.

We are given:

* An integer array `arr`
* An integer `d`

From any index `i`, we can jump:

* Left up to distance `d`
* Right up to distance `d`

But there are conditions:

* We can only jump to smaller values
* We cannot jump over an element greater than or equal to the current value

The goal is to find the maximum number of indices that can be visited starting from any index.

This problem is a great example of:

* Dynamic Programming
* DFS with Memoization
* Array traversal
* State caching optimization

---

## Constraints

| Constraint                | Value                 |
| ------------------------- | --------------------- |
| `1 <= arr.length <= 1000` | Array size            |
| `1 <= arr[i] <= 10^5`     | Element value         |
| `1 <= d <= arr.length`    | Maximum jump distance |

---

## Intuition

The first thing I noticed was that jumps are only allowed toward smaller values.

That means:

* I can never come back to a larger value
* Cycles are impossible

This immediately made the problem feel like a DFS + Dynamic Programming problem.

For every index, I wanted to know:

"How many positions can I visit if I start from here?"

If I already know the best answer for neighboring smaller positions, I can reuse those results instead of recalculating them again and again.

That is exactly where memoization helps.

---

## Approach

I solved this problem using DFS with memoization.

Here is the step-by-step idea:

1. Create a DFS function for every index.
2. From the current index:

   * Try jumping left
   * Try jumping right
3. Stop moving in a direction if:

   * Distance exceeds `d`
   * A value greater than or equal to the current value appears
4. Recursively calculate the best answer from reachable indices.
5. Store results in a DP array to avoid repeated computation.
6. Run DFS from every index because the starting point can be anywhere.

This approach works efficiently because every index is solved only once.

---

## Data Structures Used

| Data Structure  | Purpose                             |
| --------------- | ----------------------------------- |
| Array / Vector  | Stores the input array              |
| DP Array        | Memoizes already calculated answers |
| Recursion Stack | Handles DFS traversal               |

### Why Memoization?

Without memoization:

* the same paths would be recalculated many times

With memoization:

* each index is processed only once

This reduces the overall time complexity significantly.

---

## Operations & Behavior Summary

The algorithm works like this:

1. Start DFS from an index.
2. Explore all valid left jumps.
3. Explore all valid right jumps.
4. Stop exploration if a blocking element appears.
5. Recursively calculate answers for smaller reachable values.
6. Store the best result in DP.
7. Reuse stored answers whenever needed.
8. Return the maximum result among all starting positions.

This creates an optimized top-down Dynamic Programming solution.

---

## Complexity

| Type             | Complexity | Explanation                                            |
| ---------------- | ---------- | ------------------------------------------------------ |
| Time Complexity  | `O(n * d)` | Every index checks at most `d` positions on both sides |
| Space Complexity | `O(n)`     | DP array + recursion stack                             |

Where:

* `n` = size of the array
* `d` = maximum jump distance

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> dp;

    // DFS function returns maximum jumps starting from index i
    int dfs(int i, vector<int>& arr, int d) {

        // If already computed, return stored answer
        if (dp[i] != -1)
            return dp[i];

        // Minimum answer is 1 because current index is counted
        int ans = 1;

        // Try jumping to the right
        for (int j = i + 1; j <= min((int)arr.size() - 1, i + d); j++) {

            // Cannot jump further if value is greater or equal
            if (arr[j] >= arr[i])
                break;

            // Take best possible path
            ans = max(ans, 1 + dfs(j, arr, d));
        }

        // Try jumping to the left
        for (int j = i - 1; j >= max(0, i - d); j--) {

            // Cannot jump further if value is greater or equal
            if (arr[j] >= arr[i])
                break;

            // Take best possible path
            ans = max(ans, 1 + dfs(j, arr, d));
        }

        // Store and return answer
        return dp[i] = ans;
    }

    int maxJumps(vector<int>& arr, int d) {

        int n = arr.size();

        // Initialize DP array with -1 meaning unvisited
        dp.assign(n, -1);

        int answer = 1;

        // Start DFS from every index
        for (int i = 0; i < n; i++) {
            answer = max(answer, dfs(i, arr, d));
        }

        return answer;
    }
};
```

### Java

```java
class Solution {

    int[] dp;

    // DFS function returns maximum jumps starting from index i
    private int dfs(int i, int[] arr, int d) {

        // Return stored answer if already computed
        if (dp[i] != -1)
            return dp[i];

        // At least current index can be visited
        int ans = 1;

        // Move right
        for (int j = i + 1; j <= Math.min(arr.length - 1, i + d); j++) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update best answer
            ans = Math.max(ans, 1 + dfs(j, arr, d));
        }

        // Move left
        for (int j = i - 1; j >= Math.max(0, i - d); j--) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update best answer
            ans = Math.max(ans, 1 + dfs(j, arr, d));
        }

        // Store result
        return dp[i] = ans;
    }

    public int maxJumps(int[] arr, int d) {

        int n = arr.length;

        // DP array initialized with -1
        dp = new int[n];

        for (int i = 0; i < n; i++) {
            dp[i] = -1;
        }

        int answer = 1;

        // Try starting from every index
        for (int i = 0; i < n; i++) {
            answer = Math.max(answer, dfs(i, arr, d));
        }

        return answer;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} arr
 * @param {number} d
 * @return {number}
 */
var maxJumps = function(arr, d) {

    const n = arr.length;

    // dp[i] stores maximum jumps starting from i
    const dp = new Array(n).fill(-1);

    // DFS function
    const dfs = (i) => {

        // Return stored answer
        if (dp[i] !== -1)
            return dp[i];

        // Current index counts as 1
        let ans = 1;

        // Move right
        for (let j = i + 1; j <= Math.min(n - 1, i + d); j++) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update answer
            ans = Math.max(ans, 1 + dfs(j));
        }

        // Move left
        for (let j = i - 1; j >= Math.max(0, i - d); j--) {

            // Stop if blocked
            if (arr[j] >= arr[i])
                break;

            // Update answer
            ans = Math.max(ans, 1 + dfs(j));
        }

        // Store answer
        dp[i] = ans;

        return ans;
    };

    let answer = 1;

    // Start from every index
    for (let i = 0; i < n; i++) {
        answer = Math.max(answer, dfs(i));
    }

    return answer;
};
```

### Python3

```python
class Solution:
    def maxJumps(self, arr: List[int], d: int) -> int:

        n = len(arr)

        # dp[i] stores maximum jumps starting from i
        dp = [-1] * n

        # DFS function
        def dfs(i):

            # Return stored answer if already calculated
            if dp[i] != -1:
                return dp[i]

            # Minimum answer is 1 (current index)
            ans = 1

            # Move right
            for j in range(i + 1, min(n, i + d + 1)):

                # Stop if blocked
                if arr[j] >= arr[i]:
                    break

                # Update best answer
                ans = max(ans, 1 + dfs(j))

            # Move left
            for j in range(i - 1, max(-1, i - d - 1), -1):

                # Stop if blocked
                if arr[j] >= arr[i]:
                    break

                # Update best answer
                ans = max(ans, 1 + dfs(j))

            # Store answer
            dp[i] = ans

            return ans

        answer = 1

        # Try every starting index
        for i in range(n):
            answer = max(answer, dfs(i))

        return answer
```

### Go

```go
func maxJumps(arr []int, d int) int {

 n := len(arr)

 // dp[i] stores maximum jumps starting from i
 dp := make([]int, n)

 // Initialize with -1 meaning unvisited
 for i := 0; i < n; i++ {
  dp[i] = -1
 }

 // DFS function
 var dfs func(int) int

 dfs = func(i int) int {

  // Return stored answer
  if dp[i] != -1 {
   return dp[i]
  }

  // Current index itself counts as 1
  ans := 1

  // Move right
  for j := i + 1; j <= min(n-1, i+d); j++ {

   // Stop if blocked
   if arr[j] >= arr[i] {
    break
   }

   // Update best answer
   ans = max(ans, 1+dfs(j))
  }

  // Move left
  for j := i - 1; j >= max(0, i-d); j-- {

   // Stop if blocked
   if arr[j] >= arr[i] {
    break
   }

   // Update best answer
   ans = max(ans, 1+dfs(j))
  }

  // Store result
  dp[i] = ans

  return ans
 }

 answer := 1

 // Try every starting index
 for i := 0; i < n; i++ {
  answer = max(answer, dfs(i))
 }

 return answer
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

The logic remains the same across all five languages.

### Step 1 — Create a DP Array

I first create a DP array where:

* `dp[i]` stores the maximum jumps possible starting from index `i`

Initially:

* every value is unvisited

This helps avoid recalculating the same state multiple times.

---

### Step 2 — Define DFS

The DFS function solves the problem for one index at a time.

For a given index:

* check all valid left moves
* check all valid right moves

Then return the best possible answer.

---

### Step 3 — Explore Right Side

I move from:

* `i + 1`
  to
* `i + d`

For every position:

* if the value is smaller, continue
* otherwise stop immediately

Stopping is important because:

* larger or equal values block further movement

---

### Step 4 — Explore Left Side

The exact same logic is repeated toward the left side.

I continue checking positions until:

* distance exceeds `d`
* or a blocking value appears

---

### Step 5 — Memoization

Before solving an index:

* check whether the answer already exists

If yes:

* directly return it

This optimization prevents repeated DFS calls.

---

### Step 6 — Compute Final Answer

Since we can start from any index:

* run DFS from every position

Then return the maximum answer found.

---

### Language Notes

#### C++

* `vector<int>` is used for DP storage
* recursion is fast and clean
* very efficient for competitive programming

#### Java

* uses integer arrays for memoization
* recursion logic remains identical
* `Math.max()` simplifies comparisons

#### JavaScript

* arrays are dynamic
* recursion works naturally
* memoization avoids timeout issues

#### Python3

* concise syntax makes DFS very readable
* list-based memoization is simple to implement

#### Go

* slices are used for DP
* helper functions are usually created for `max()` and `min()`

---

## Examples

### Example 1

Input:

```text
arr = [6,4,14,6,8,13,9,7,10,6,12]
d = 2
```

Output:

```text
4
```

Explanation:

One optimal path is:

```text
10 -> 8 -> 6 -> 7
```

Visited indices count:

* 4

---

### Example 2

Input:

```text
arr = [3,3,3,3,3]
d = 3
```

Output:

```text
1
```

Explanation:

No jumps are possible because:

* all values are equal
* jumps require strictly smaller values

---

### Example 3

Input:

```text
arr = [7,6,5,4,3,2,1]
d = 1
```

Output:

```text
7
```

Explanation:

Every next position is smaller, so we can visit every index.

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

* DFS + Memoization is the most common optimized solution for this problem.
* A pure brute-force DFS would be too slow because of repeated calculations.
* Since jumps only go toward smaller values, cycles are impossible.
* The blocking condition is very important:

  * once a larger or equal value appears, movement must stop completely
* Top-down DP fits naturally here because every state depends on smaller states.

### Alternative Approaches

Some people also solve this using:

* Bottom-up Dynamic Programming
* Monotonic Stack optimizations

But DFS + memoization is easier to understand and implement during interviews and contests.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
