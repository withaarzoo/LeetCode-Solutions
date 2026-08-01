# Predict the Winner - LeetCode 486 Solution | Dynamic Programming Game Theory

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
- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

LeetCode 486, **Predict the Winner**, is a classic **Dynamic Programming** and **Game Theory** problem.

You are given an integer array where two players take turns choosing a number from either the beginning or the end of the array. Every chosen number is added to that player's score. Both players always play optimally, meaning they always make the best possible move.

Your task is to determine whether Player 1 can finish with a score greater than or equal to Player 2's score.

### Input

- An integer array `nums`.

### Output

- Return `true` if Player 1 can win or tie.
- Return `false` otherwise.

This problem is a great example of using **Dynamic Programming (DP)** to solve an optimal decision-making problem where every move changes the future state of the game.

---

## Constraints

| Constraint | Value |
| ---------- | ----- |
| `1 <= nums.length <= 20` |
| `0 <= nums[i] <= 10^7` |

---

## Intuition

My first thought was to simulate every possible game. That quickly became impossible because every move creates two more possible choices, making the number of games grow exponentially.

Then I noticed something important.

Instead of keeping track of two separate scores, I only need to know the **score difference** between the current player and the opponent.

For every subarray, I ask one simple question:

> What is the maximum advantage the current player can achieve?

That single observation turns the entire problem into a Dynamic Programming problem.

---

## Approach

I solve the problem using Dynamic Programming.

First, I create a DP table where each cell stores the maximum score difference the current player can achieve for a specific subarray.

For every range of numbers:

1. Try taking the left number.
2. Try taking the right number.
3. Calculate the score difference for both choices.
4. Keep whichever choice gives the larger advantage.

I begin with subarrays of length one because those answers are already known.

Then I gradually solve larger and larger subarrays until the whole array is covered.

Finally, if the score difference for the complete array is non-negative, Player 1 wins or ties.

---

## Data Structures Used

### 2D Dynamic Programming Table

A two-dimensional array stores the best score difference for every possible subarray.

It helps avoid solving the same subproblem multiple times and makes the overall solution efficient.

---

## Operations & Behavior Summary

The algorithm works in the following order:

1. Find the size of the array.
2. Create a DP table.
3. Initialize all single-element subarrays.
4. Solve subarrays of length two.
5. Continue increasing the subarray size.
6. For every range:
   - Try taking the left value.
   - Try taking the right value.
   - Store the better result.
7. Return whether the final score difference is at least zero.

This bottom-up DP approach ensures every smaller problem is solved before a larger one depends on it.

---

## Complexity

| Complexity | Value | Explanation |
| ---------- | ----- | ----------- |
| Time Complexity | **O(n²)** | Every possible subarray is computed exactly once. |
| Space Complexity | **O(n²)** | A DP table stores results for every subarray. |

Here, `n` represents the number of elements in the input array.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool predictTheWinner(vector<int>& nums) {
        int n = nums.size();

        // dp[i][j] stores the maximum score difference the current player
        // can achieve over the opponent for subarray [i...j]
        vector<vector<int>> dp(n, vector<int>(n, 0));

        // Base case:
        // If only one number exists, the current player takes it.
        for (int i = 0; i < n; i++) {
            dp[i][i] = nums[i];
        }

        // Build the DP table for increasing subarray lengths.
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;

                // Choose the left element.
                int takeLeft = nums[i] - dp[i + 1][j];

                // Choose the right element.
                int takeRight = nums[j] - dp[i][j - 1];

                // Keep the better option.
                dp[i][j] = max(takeLeft, takeRight);
            }
        }

        // Non-negative means Player 1 can win or tie.
        return dp[0][n - 1] >= 0;
    }
};
```

### Java

```java
class Solution {
    public boolean predictTheWinner(int[] nums) {
        int n = nums.length;

        // dp[i][j] stores the maximum score difference
        // current player can achieve for subarray [i...j].
        int[][] dp = new int[n][n];

        // Base case:
        // One element means the current player takes it.
        for (int i = 0; i < n; i++) {
            dp[i][i] = nums[i];
        }

        // Fill the table from smaller ranges to larger ranges.
        for (int len = 2; len <= n; len++) {
            for (int i = 0; i + len - 1 < n; i++) {
                int j = i + len - 1;

                // Pick the left number.
                int takeLeft = nums[i] - dp[i + 1][j];

                // Pick the right number.
                int takeRight = nums[j] - dp[i][j - 1];

                // Store the better result.
                dp[i][j] = Math.max(takeLeft, takeRight);
            }
        }

        // Player 1 wins or ties if score difference is non-negative.
        return dp[0][n - 1] >= 0;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {boolean}
 */
var predictTheWinner = function(nums) {
    const n = nums.length;

    // dp[i][j] stores the maximum score difference
    // for subarray [i...j].
    const dp = Array.from({ length: n }, () => Array(n).fill(0));

    // Base case:
    // Only one number remains.
    for (let i = 0; i < n; i++) {
        dp[i][i] = nums[i];
    }

    // Build answers for larger subarrays.
    for (let len = 2; len <= n; len++) {
        for (let i = 0; i + len - 1 < n; i++) {
            const j = i + len - 1;

            // Take the left number.
            const takeLeft = nums[i] - dp[i + 1][j];

            // Take the right number.
            const takeRight = nums[j] - dp[i][j - 1];

            // Store the better choice.
            dp[i][j] = Math.max(takeLeft, takeRight);
        }
    }

    // Player 1 wins or ties.
    return dp[0][n - 1] >= 0;
};
```

### Python3

```python
class Solution:
    def predictTheWinner(self, nums: List[int]) -> bool:
        n = len(nums)

        # dp[i][j] stores the maximum score difference
        # the current player can achieve for subarray [i...j].
        dp = [[0] * n for _ in range(n)]

        # Base case:
        # One number is left, so the player takes it.
        for i in range(n):
            dp[i][i] = nums[i]

        # Fill the table from smaller ranges to larger ranges.
        for length in range(2, n + 1):
            for i in range(n - length + 1):
                j = i + length - 1

                # Pick the left number.
                take_left = nums[i] - dp[i + 1][j]

                # Pick the right number.
                take_right = nums[j] - dp[i][j - 1]

                # Store the better option.
                dp[i][j] = max(take_left, take_right)

        # Non-negative means Player 1 wins or ties.
        return dp[0][n - 1] >= 0
```

### Go

```go
func predictTheWinner(nums []int) bool {
 n := len(nums)

 // dp[i][j] stores the maximum score difference
 // current player can achieve for subarray [i...j].
 dp := make([][]int, n)
 for i := range dp {
  dp[i] = make([]int, n)
 }

 // Base case:
 // Only one number remains.
 for i := 0; i < n; i++ {
  dp[i][i] = nums[i]
 }

 // Fill DP for larger subarrays.
 for length := 2; length <= n; length++ {
  for i := 0; i+length-1 < n; i++ {
   j := i + length - 1

   // Take the left number.
   takeLeft := nums[i] - dp[i+1][j]

   // Take the right number.
   takeRight := nums[j] - dp[i][j-1]

   // Store the better choice.
   if takeLeft > takeRight {
    dp[i][j] = takeLeft
   } else {
    dp[i][j] = takeRight
   }
  }
 }

 // Player 1 wins or ties.
 return dp[0][n-1] >= 0
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The algorithm behaves exactly the same in all five programming languages. Only the syntax changes.

### Step 1

Read the input array and determine its length.

This tells us how large the DP table needs to be.

---

### Step 2

Create a two-dimensional DP table.

Each cell represents the best score difference the current player can achieve for one specific subarray.

Instead of storing actual scores, storing score differences makes the transitions much simpler.

---

### Step 3

Initialize every single-element subarray.

If only one number remains, the current player simply takes it.

So the score difference is exactly that number.

These values become the base cases for all larger subarrays.

---

### Step 4

Process subarrays from the shortest length to the longest.

This order is important because every larger answer depends on smaller ranges that must already be calculated.

---

### Step 5

For every subarray, evaluate both available moves.

If the player chooses the left value, the opponent gets the remaining subarray.

If the player chooses the right value, the opponent again receives the remaining subarray.

The DP table already knows the opponent's best possible advantage.

Subtracting that value gives the current player's final score difference.

---

### Step 6

Compare both possible moves.

Whichever move gives the larger score difference is stored in the DP table.

Since both players always play optimally, choosing anything smaller would never lead to the best result.

---

### Step 7

Look at the value stored for the entire array.

If the score difference is zero or greater, Player 1 can either win or tie.

Otherwise, Player 2 finishes with a higher score.

---

## Examples

### Example 1

**Input**

```text
nums = [1,5,2]
```

**Output**

```text
false
```

**Explanation**

If Player 1 chooses `1`, Player 2 chooses `5`.

If Player 1 chooses `2`, Player 2 still chooses `5`.

Player 2 always finishes with a higher score.

---

### Example 2

**Input**

```text
nums = [1,5,233,7]
```

**Output**

```text
true
```

**Explanation**

No matter which move Player 2 makes, Player 1 can eventually collect `233`.

That gives Player 1 the winning score.

---

### Example 3

**Input**

```text
nums = [4]
```

**Output**

```text
true
```

**Explanation**

Only one number exists.

Player 1 takes it immediately and wins.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project directory.

```bash
cd your-repository
```

### C++

Compile:

```bash
g++ solution.cpp -o solution
```

Run:

```bash
./solution
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

or

```bash
python3 solution.py
```

---

### Go

Run:

```bash
go run solution.go
```

---

## Notes & Optimizations

- This is a classic Dynamic Programming and Game Theory problem.
- The DP solution avoids exponential recursion by solving every subproblem only once.
- An alternative solution uses recursion with memoization and produces the same time complexity.
- The current solution is already optimal for the given constraints.
- Since the maximum array length is only 20, the DP table easily fits into memory.
- The same idea can be extended to many interval Dynamic Programming problems where players alternately make optimal choices.

---

## Author

**Md Aarzoo Islam**

Instagram: <https://www.instagram.com/code.with.aarzoo/>
