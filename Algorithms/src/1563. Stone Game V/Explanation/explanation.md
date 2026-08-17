# 1563. Stone Game V — Dynamic Programming Solution

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

LeetCode 1563, Stone Game V, is a dynamic programming and game theory problem based on splitting an array of stones.

We are given an array called `stoneValue`. Each stone has a positive integer value, and all stones are arranged in a row.

Alice starts with a score of `0`.

In every round, Alice must split the current row into two non-empty parts. Bob calculates the sum of both parts.

* If the left part has a smaller sum, Bob removes the right part and Alice gains the left sum.
* If the right part has a smaller sum, Bob removes the left part and Alice gains the right sum.
* If both parts have the same sum, Alice can choose which part remains.

The game continues with the remaining stones until only one stone is left.

The goal is to return the maximum score Alice can obtain.

For example:

```text
Input:  [6,2,3,4,5,5]
Output: 18
```

The main challenge is that there are many possible ways to split the array. A straightforward recursive solution would repeat the same subproblems many times, so dynamic programming is needed.

This solution uses interval dynamic programming, prefix sums, and monotonic split pointers to reduce the runtime from `O(n³)` to `O(n²)`.

## Constraints

* `1 <= stoneValue.length <= 500`
* `1 <= stoneValue[i] <= 10⁶`
* All stone values are positive integers.
* The input array can contain only one stone, in which case the answer is `0`.

## Intuition

I first thought about trying every possible split of the current array.

For an interval `[l, r]`, I can split it after any position `k`. Then I compare the sum of the left and right parts and determine which part survives.

This gives a natural dynamic programming state:

`dp[l][r]`

where `dp[l][r]` represents the maximum score Alice can get from the stones between indices `l` and `r`.

The problem is that if I try every split for every interval, the solution becomes `O(n³)`. With `n` up to `500`, that is more work than necessary.

The important observation is that every stone value is positive.

Because of this, when I move the split point from left to right, the left-side sum always increases. So I can find the point where the left sum becomes larger than the right sum without starting the search from the beginning every time.

I combine this observation with prefix sums so every range sum can be calculated in constant time.

I also maintain helper DP tables that store the best score for valid left and right parts. This lets me avoid checking every split individually.

The result is an `O(n²)` interval dynamic programming solution.

## Approach

I solve the problem in the following steps.

### 1. Build prefix sums

I create a prefix sum array so that I can calculate the sum of any subarray in `O(1)`.

For an interval `[l, r]`:

`sum(l, r) = prefix[r + 1] - prefix[l]`

This is important because the algorithm needs interval sums repeatedly.

### 2. Define the main DP state

I use:

`dp[l][r]`

to store the maximum score Alice can obtain when the remaining stones are from index `l` to index `r`.

For a single stone, there is no possible split, so:

`dp[i][i] = 0`

### 3. Consider a split

Suppose I split `[l, r]` after index `k`.

The left side is `[l, k]`.

The right side is `[k + 1, r]`.

If:

`leftSum <= rightSum`

Alice can keep the left side and receive:

`leftSum + dp[l][k]`

If:

`rightSum <= leftSum`

Alice can keep the right side and receive:

`rightSum + dp[k + 1][r]`

When both sums are equal, both choices are possible, so I take the better result.

### 4. Avoid checking every split

The left sum increases as the split moves right because every stone has a positive value.

I therefore maintain two pointers:

* One points to the last split where the left side is not larger.
* The other points to the first split where the left side is not smaller.

These pointers only move forward.

### 5. Store the best left and right transitions

I use `leftBest` to store the best value of:

`dp[l][k] + sum(l, k)`

for possible left parts.

I use `rightBest` to store the best value of:

`dp[k][r] + sum(k, r)`

for possible right parts.

This means I can get the best valid transition without scanning every possible split.

### 6. Process intervals by length

I calculate intervals from smaller to larger.

For example:

`length = 2`

then:

`length = 3`

then:

`length = 4`

and so on.

This guarantees that all smaller DP states required by the current interval have already been calculated.

## Data Structures Used

| Data Structure | Purpose                                                                                |
| -------------- | -------------------------------------------------------------------------------------- |
| `prefix`       | Stores prefix sums so any interval sum can be found in `O(1)`.                         |
| `dp`           | Stores the maximum score Alice can obtain for every interval.                          |
| `leftBest`     | Stores the best transition when the left part survives.                                |
| `rightBest`    | Stores the best transition when the right part survives.                               |
| `leftPtr`      | Tracks the boundary where the left sum is still less than or equal to the right sum.   |
| `rightPtr`     | Tracks the boundary where the left sum becomes greater than or equal to the right sum. |

## Operations & Behavior Summary

The algorithm works like this:

1. Read the stone values.
2. Build the prefix sum array.
3. Initialize the DP and helper tables.
4. Treat every single stone as an interval with score `0`.
5. Process intervals from length `2` up to `n`.
6. Calculate the total sum of the current interval.
7. Move the left split pointer while the left sum is smaller than or equal to the right sum.
8. Move the right split pointer until the left sum becomes greater than or equal to the right sum.
9. Use `leftBest` for the case where the left part survives.
10. Use `rightBest` for the case where the right part survives.
11. Store the maximum result in `dp[l][r]`.
12. Update the helper tables for future larger intervals.
13. Return `dp[0][n - 1]`.

The important optimization is that I do not test every split for every interval. The split pointers move forward because all stone values are positive.

## Complexity

| Type             | Complexity | Explanation                                                                                                                         |
| ---------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n²)`    | There are `O(n²)` intervals, and the split pointers move forward instead of scanning all `O(n)` split positions for every interval. |
| Space Complexity | `O(n²)`    | The `dp`, `leftBest`, and `rightBest` tables each use `O(n²)` space. Prefix sums and pointers use `O(n)` additional space.          |

Here, `n` is the number of stones in the `stoneValue` array.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int stoneGameV(vector<int>& stoneValue) {
        int n = stoneValue.size();

        vector<long long> prefix(n + 1, 0);
        for (int i = 0; i < n; ++i) {
            prefix[i + 1] = prefix[i] + stoneValue[i];
        }

        vector<vector<int>> dp(n, vector<int>(n, 0));

        vector<vector<int>> leftBest(n, vector<int>(n, 0));

        vector<vector<int>> rightBest(n, vector<int>(n, 0));

        vector<int> leftPtr(n);

        vector<int> rightPtr(n);

        for (int i = 0; i < n; ++i) {
            leftBest[i][i] = stoneValue[i];
            rightBest[i][i] = stoneValue[i];

            leftPtr[i] = i - 1;

            rightPtr[i] = i;
        }

        for (int len = 2; len <= n; ++len) {
            for (int l = 0; l + len <= n; ++l) {
                int r = l + len - 1;

                long long total = prefix[r + 1] - prefix[l];

                while (leftPtr[l] + 1 <= r - 1) {
                    int k = leftPtr[l] + 1;
                    long long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum > total) {
                        break;
                    }

                    ++leftPtr[l];
                }

                while (rightPtr[l] <= r - 1) {
                    int k = rightPtr[l];
                    long long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum >= total) {
                        break;
                    }

                    ++rightPtr[l];
                }

                int best = 0;

                if (leftPtr[l] >= l) {
                    best = leftBest[l][leftPtr[l]];
                }

                if (rightPtr[l] <= r - 1) {
                    best = max(best, rightBest[rightPtr[l] + 1][r]);
                }

                dp[l][r] = best;

                leftBest[l][r] = max(
                    leftBest[l][r - 1],
                    dp[l][r] + static_cast<int>(total)
                );

                rightBest[l][r] = max(
                    rightBest[l + 1][r],
                    dp[l][r] + static_cast<int>(total)
                );
            }
        }

        return dp[0][n - 1];
    }
};
```

### Java

```java
class Solution {
    public int stoneGameV(int[] stoneValue) {
        int n = stoneValue.length;

        long[] prefix = new long[n + 1];

        for (int i = 0; i < n; i++) {
            prefix[i + 1] = prefix[i] + stoneValue[i];
        }

        int[][] dp = new int[n][n];

        int[][] leftBest = new int[n][n];

        int[][] rightBest = new int[n][n];

        int[] leftPtr = new int[n];

        int[] rightPtr = new int[n];

        for (int i = 0; i < n; i++) {
            leftBest[i][i] = stoneValue[i];
            rightBest[i][i] = stoneValue[i];

            leftPtr[i] = i - 1;

            rightPtr[i] = i;
        }

        for (int len = 2; len <= n; len++) {
            for (int l = 0; l + len <= n; l++) {
                int r = l + len - 1;

                long total = prefix[r + 1] - prefix[l];

                while (leftPtr[l] + 1 <= r - 1) {
                    int k = leftPtr[l] + 1;
                    long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum > total) {
                        break;
                    }

                    leftPtr[l]++;
                }

                while (rightPtr[l] <= r - 1) {
                    int k = rightPtr[l];
                    long leftSum = prefix[k + 1] - prefix[l];

                    if (2 * leftSum >= total) {
                        break;
                    }

                    rightPtr[l]++;
                }

                int best = 0;

                if (leftPtr[l] >= l) {
                    best = leftBest[l][leftPtr[l]];
                }

                if (rightPtr[l] <= r - 1) {
                    best = Math.max(best, rightBest[rightPtr[l] + 1][r]);
                }

                dp[l][r] = best;

                leftBest[l][r] = Math.max(
                    leftBest[l][r - 1],
                    dp[l][r] + (int) total
                );

                rightBest[l][r] = Math.max(
                    rightBest[l + 1][r],
                    dp[l][r] + (int) total
                );
            }
        }

        return dp[0][n - 1];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} stoneValue
 * @return {number}
 */
var stoneGameV = function(stoneValue) {
    const n = stoneValue.length;

    const prefix = new Array(n + 1).fill(0);

    for (let i = 0; i < n; i++) {
        prefix[i + 1] = prefix[i] + stoneValue[i];
    }

    const dp = Array.from({ length: n }, () => new Array(n).fill(0));

    const leftBest = Array.from(
        { length: n },
        () => new Array(n).fill(0)
    );

    const rightBest = Array.from(
        { length: n },
        () => new Array(n).fill(0)
    );

    const leftPtr = new Array(n);

    const rightPtr = new Array(n);

    for (let i = 0; i < n; i++) {
        leftBest[i][i] = stoneValue[i];
        rightBest[i][i] = stoneValue[i];

        leftPtr[i] = i - 1;

        rightPtr[i] = i;
    }

    for (let len = 2; len <= n; len++) {
        for (let l = 0; l + len <= n; l++) {
            const r = l + len - 1;

            const total = prefix[r + 1] - prefix[l];

            while (leftPtr[l] + 1 <= r - 1) {
                const k = leftPtr[l] + 1;
                const leftSum = prefix[k + 1] - prefix[l];

                if (2 * leftSum > total) {
                    break;
                }

                leftPtr[l]++;
            }

            while (rightPtr[l] <= r - 1) {
                const k = rightPtr[l];
                const leftSum = prefix[k + 1] - prefix[l];

                if (2 * leftSum >= total) {
                    break;
                }

                rightPtr[l]++;
            }

            let best = 0;

            if (leftPtr[l] >= l) {
                best = leftBest[l][leftPtr[l]];
            }

            if (rightPtr[l] <= r - 1) {
                best = Math.max(
                    best,
                    rightBest[rightPtr[l] + 1][r]
                );
            }

            dp[l][r] = best;

            leftBest[l][r] = Math.max(
                leftBest[l][r - 1],
                dp[l][r] + total
            );

            rightBest[l][r] = Math.max(
                rightBest[l + 1][r],
                dp[l][r] + total
            );
        }
    }

    return dp[0][n - 1];
};
```

### Python3

```python
from typing import List

class Solution:
    def stoneGameV(self, stoneValue: List[int]) -> int:
        n = len(stoneValue)

        prefix = [0] * (n + 1)

        for i in range(n):
            prefix[i + 1] = prefix[i] + stoneValue[i]

        dp = [[0] * n for _ in range(n)]

        left_best = [[0] * n for _ in range(n)]

        right_best = [[0] * n for _ in range(n)]

        left_ptr = [0] * n

        right_ptr = list(range(n))

        for i in range(n):
            left_best[i][i] = stoneValue[i]
            right_best[i][i] = stoneValue[i]

            left_ptr[i] = i - 1

            right_ptr[i] = i

        for length in range(2, n + 1):
            for l in range(n - length + 1):
                r = l + length - 1

                total = prefix[r + 1] - prefix[l]

                while left_ptr[l] + 1 <= r - 1:
                    k = left_ptr[l] + 1
                    left_sum = prefix[k + 1] - prefix[l]

                    if 2 * left_sum > total:
                        break

                    left_ptr[l] += 1

                while right_ptr[l] <= r - 1:
                    k = right_ptr[l]
                    left_sum = prefix[k + 1] - prefix[l]

                    if 2 * left_sum >= total:
                        break

                    right_ptr[l] += 1

                best = 0

                if left_ptr[l] >= l:
                    best = left_best[l][left_ptr[l]]

                if right_ptr[l] <= r - 1:
                    best = max(
                        best,
                        right_best[right_ptr[l] + 1][r]
                    )

                dp[l][r] = best

                left_best[l][r] = max(
                    left_best[l][r - 1],
                    dp[l][r] + total
                )

                right_best[l][r] = max(
                    right_best[l + 1][r],
                    dp[l][r] + total
                )

        return dp[0][n - 1]
```

### Go

```go
func stoneGameV(stoneValue []int) int {
 n := len(stoneValue)

 prefix := make([]int64, n+1)

 for i := 0; i < n; i++ {
  prefix[i+1] = prefix[i] + int64(stoneValue[i])
 }

 dp := make([][]int, n)

 leftBest := make([][]int, n)

 rightBest := make([][]int, n)

 for i := 0; i < n; i++ {
  dp[i] = make([]int, n)
  leftBest[i] = make([]int, n)
  rightBest[i] = make([]int, n)

  leftBest[i][i] = stoneValue[i]
  rightBest[i][i] = stoneValue[i]
 }

 leftPtr := make([]int, n)

 rightPtr := make([]int, n)

 for i := 0; i < n; i++ {
  leftPtr[i] = i - 1

  rightPtr[i] = i
 }

 for length := 2; length <= n; length++ {
  for l := 0; l+length <= n; l++ {
   r := l + length - 1

   total := prefix[r+1] - prefix[l]

   for leftPtr[l]+1 <= r-1 {
    k := leftPtr[l] + 1
    leftSum := prefix[k+1] - prefix[l]

    if 2*leftSum > total {
     break
    }

    leftPtr[l]++
   }

   for rightPtr[l] <= r-1 {
    k := rightPtr[l]
    leftSum := prefix[k+1] - prefix[l]

    if 2*leftSum >= total {
     break
    }

    rightPtr[l]++
   }

   best := 0

   if leftPtr[l] >= l {
    best = leftBest[l][leftPtr[l]]
   }

   if rightPtr[l] <= r-1 {
    candidate := rightBest[rightPtr[l]+1][r]

    if candidate > best {
     best = candidate
    }
   }

   dp[l][r] = best

   currentSum := int(total)
   leftCandidate := dp[l][r] + currentSum

   if leftBest[l][r-1] > leftCandidate {
    leftBest[l][r] = leftBest[l][r-1]
   } else {
    leftBest[l][r] = leftCandidate
   }

   rightCandidate := dp[l][r] + currentSum

   if rightBest[l+1][r] > rightCandidate {
    rightBest[l][r] = rightBest[l+1][r]
   } else {
    rightBest[l][r] = rightCandidate
   }
  }
 }

 return dp[0][n-1]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax and memory-management details are different.

### C++

I first create a `prefix` vector containing cumulative sums. I use `long long` for prefix sums because the total value of the stones can be larger than what a normal 32-bit integer safely represents during intermediate calculations.

Next, I create the three main two-dimensional DP tables.

`dp[l][r]` stores the final answer for an interval.

`leftBest[l][r]` stores the best result when an interval ending somewhere inside `[l, r]` is used as the left surviving part.

`rightBest[l][r]` does the same thing for the right surviving part.

I initialize the helper tables for single stones with their individual values. Their `dp` values remain zero because a single stone cannot be split.

While processing every interval, I calculate its total sum using the prefix array.

I then move `leftPtr` until the next split would make the left side larger than the right side.

I also move `rightPtr` until the left side becomes at least as large as the right side.

The two pointer positions cover the only boundary cases I need.

I then take the best result from the left and right possibilities and store it in `dp[l][r]`.

Finally, I update the helper tables so this interval can be used when solving a larger interval.

### Java

The Java solution follows exactly the same interval DP idea.

I use a `long[]` for prefix sums because Java's `int` type can overflow when calculating large intermediate sums.

The DP tables use `int[][]` because the final Alice score fits within the expected result range.

The arrays `leftPtr` and `rightPtr` store the moving split boundaries for each starting index.

Java arrays give direct indexed access, which is useful here because all DP states are identified by two indices.

The intervals are processed by increasing length, so every required smaller state is already available.

### JavaScript

In JavaScript, I use normal arrays for the prefix sums, DP tables, and pointer arrays.

JavaScript numbers use double-precision floating-point values, but the values involved in this problem remain within the exact integer range of JavaScript's `Number` type.

I create the two-dimensional arrays before processing the intervals.

For every interval, I calculate the total sum from the prefix array.

The split pointers move forward in the same way as the C++, Java, Python3, and Go implementations.

The final answer is stored at:

`dp[0][n - 1]`

### Python3

The Python implementation uses lists to represent the two-dimensional DP tables.

I use Python integers for prefix sums and DP calculations. Python integers automatically handle large integer values, so I do not need a separate 64-bit integer type.

The interval length loop ensures that all smaller intervals have already been solved.

The two pointers are stored in normal Python lists.

Although Python has more overhead than C++ for large two-dimensional lists, the `O(n²)` solution is still designed around the maximum constraint of `n = 500`.

### Go

The Go implementation uses slices to create the two-dimensional DP tables.

I use `int64` for the prefix sums so that intermediate sums are safely represented.

The final DP values can remain `int`.

Go does not have built-in two-dimensional arrays with dynamic sizes in the same way as some other languages, so I create a slice for each row.

The pointer arrays are simple `[]int` slices.

The algorithm itself is unchanged: process intervals from short to long, maintain the split boundaries, calculate the best transition, and update the helper tables.

## Examples

### Example 1

Input:

```text
[6,2,3,4,5,5]
```

Output:

```text
18
```

One optimal sequence of choices is:

```text
[6,2,3] [4,5,5]
```

The sums are:

```text
11 and 14
```

The left side is smaller, so Alice keeps `[6,2,3]` and gains `11`.

Then she splits:

```text
[6,2] [3]
```

The sums are:

```text
8 and 3
```

The right side is smaller, so Alice keeps `[3]` and gains `3`.

The optimal sequence considered by the DP produces a total score of:

```text
18
```

### Example 2

Input:

```text
[7,7,7,7,7,7,7,7]
```

Output:

```text
28
```

All values are equal, so many splits produce equal sums.

Because Alice can choose which side survives when both sums are equal, the DP considers both possibilities and keeps the better one at every state.

The final maximum score is:

```text
28
```

### Example 3

Input:

```text
[4]
```

Output:

```text
0
```

There is only one stone.

Alice cannot divide one stone into two non-empty rows, so the game ends immediately and her score remains `0`.

## How to Use / Run Locally

### C++

Save the solution in a C++ file such as:

```text
main.cpp
```

The LeetCode-style `Solution` class can be tested by adding a small `main()` function around it.

Compile with:

```bash
g++ -std=c++17 -O2 main.cpp -o main
```

Run it with:

```bash
./main
```

On Windows, the executable can be run with:

```bash
main.exe
```

### Java

Save the solution in:

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

If you are testing locally, add a `main` method that creates a `Solution` object and calls `stoneGameV`.

### JavaScript

Save the solution as:

```text
solution.js
```

Make sure Node.js is installed.

Run:

```bash
node solution.js
```

For local testing, add a small test case that calls `stoneGameV` with an array of stone values.

### Python3

Save the solution as:

```text
solution.py
```

Run:

```bash
python3 solution.py
```

You can add a small test section at the bottom of the file to call the method with sample inputs.

### Go

Save the code in:

```text
main.go
```

For a complete standalone Go program, add the required `package main`, imports, and a `main()` function for testing.

Run:

```bash
go run main.go
```

You can also build the program with:

```bash
go build main.go
```

## Notes & Optimizations

The most important optimization is avoiding the straightforward `O(n³)` interval DP.

A basic solution can use:

`dp[l][r] = max over every split k`

That approach is easy to understand, but it checks too many split points.

The optimized solution takes advantage of the fact that every `stoneValue[i]` is positive. This means the left-side sum always increases when the split moves right.

Because of that property, the split pointers never need to move backward.

Prefix sums are also important. Without them, calculating the sum of every left and right interval would require additional work.

The equality case needs special attention. When:

`leftSum == rightSum`

Alice can choose either side. The solution therefore considers both possibilities.

The single-element interval is another important edge case. It cannot be split, so its DP value is `0`.

The input constraint `n <= 500` makes `O(n²)` space practical, while the optimization is useful for keeping the runtime low.

An `O(n³)` solution is easier to write and can be useful for understanding the basic recurrence first. However, the optimized `O(n²)` approach is a better fit for the given constraints.

This problem is a useful example of interval dynamic programming, prefix sums, game theory, and monotonic pointer optimization working together.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)

---
