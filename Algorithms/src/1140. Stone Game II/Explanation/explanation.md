# 1140. Stone Game II | LeetCode Dynamic Programming Solution

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

LeetCode 1140, **Stone Game II**, is a dynamic programming problem based on optimal game strategy.

We are given an array `piles`, where each element represents the number of stones in one pile. The piles are arranged in a row.

Alice starts the game. On every turn, the current player can take all the stones from the first `X` remaining piles, where:

`1 <= X <= 2 * M`

Initially, `M = 1`.

After taking `X` piles, the value of `M` becomes:

`M = max(M, X)`

The game continues until there are no piles left.

Both Alice and Bob play optimally. The goal is to return the maximum number of stones Alice can collect.

For example:

```text
Input:  piles = [2,7,9,4,4]
Output: 10
```

Alice can play optimally and collect a maximum of `10` stones.

The main concepts used in this **Stone Game II solution** are **dynamic programming, memoization, recursion, game theory, state transitions, and suffix sums**.

## Constraints

| Constraint             | Value                      |
| ---------------------- | -------------------------- |
| Number of piles        | `1 <= piles.length <= 100` |
| Stones in each pile    | `1 <= piles[i] <= 10^4`    |
| Starting player        | Alice                      |
| Initial `M`            | `1`                        |
| Allowed piles per turn | `1 <= X <= 2 * M`          |

The input size is small enough for a dynamic programming solution with `O(n^3)` time complexity and `O(n^2)` space complexity.

## Intuition

My first thought was to simulate the game and try every possible move.

But there is a problem with doing that directly. After every move, the value of `M` can change, which changes how many piles the next player is allowed to take.

So the current position alone is not enough to describe the game.

I realized that every situation can be represented using two values:

`i` = the first remaining pile

`M` = the current value of `M`

So I can define a state as:

`dp(i, M)`

This state tells me the maximum number of stones the current player can collect from that point onward.

From a state `(i, M)`, I try every possible number of piles `X` from `1` to `2 * M`.

After taking `X` piles, the next state becomes:

`(i + X, max(M, X))`

The interesting part is calculating my final score after making a move.

If the total number of remaining stones is `suffix[i]`, and the next player can collect `dp(i + X, max(M, X))`, then I can collect:

`remaining stones - opponent's stones`

So the transition becomes:

`dp(i, M) = max(suffix[i] - dp(i + X, max(M, X)))`

I also store already calculated states so I do not solve the same state repeatedly.

## Approach

I use three main ideas.

First, I calculate suffix sums.

For every index `i`, `suffix[i]` stores the total number of stones from `i` to the end of the array.

This lets me find the total remaining stones in constant time.

Second, I use dynamic programming with the state `(i, M)`.

For every state, I try all legal values of `X`.

The current player can take:

`1 <= X <= 2 * M`

After taking `X` piles:

* The current player gets the stones from those `X` piles.
* The next player starts from index `i + X`.
* The new value of `M` is `max(M, X)`.

Instead of explicitly calculating the opponent's entire strategy, I calculate the opponent's best possible score from the next state.

If there are `suffix[i]` total stones remaining and the opponent can get `opponent`, then I get:

`suffix[i] - opponent`

Finally, I choose the move that gives me the largest result.

Memoization makes sure each `(i, M)` state is calculated only once.

## Data Structures Used

### Suffix Sum Array

I use a suffix sum array to store the total number of stones remaining from every index.

For example:

```text
piles  = [2, 7, 9, 4, 4]

suffix = [26, 24, 17, 8, 4, 0]
```

This allows me to get the total remaining stones using `suffix[i]` instead of calculating the sum again and again.

### 2D DP / Memoization Table

I use a two-dimensional table where:

* First dimension represents the current pile index `i`.
* Second dimension represents the current value of `M`.

The value stored at `dp[i][M]` is the maximum number of stones the current player can collect from that state.

### Recursion

The solution naturally follows the game's structure.

One move creates another smaller game starting from a later index, so recursion is a simple way to represent these state transitions.

Python uses `lru_cache` for memoization, while the other implementations use an explicit DP table.

## Operations & Behavior Summary

The algorithm works like this:

1. Read the `piles` array.
2. Calculate suffix sums from right to left.
3. Start the game at index `0` with `M = 1`.
4. Check whether this state was already calculated.
5. If all piles have been taken, return `0`.
6. Try every possible `X` from `1` to `2 * M`.
7. Make sure `X` does not go beyond the remaining number of piles.
8. Calculate the next value of `M` using `max(M, X)`.
9. Move to the next state.
10. Calculate the current player's score as total remaining stones minus the opponent's best score.
11. Keep the maximum score among all possible moves.
12. Store the result for the current state.
13. Return the result for the initial state `(0, 1)`.

In simple pseudocode:

```text
solve(i, M):

    if no piles remain:
        return 0

    if state already exists:
        return stored answer

    best = 0

    for every valid X:
        nextM = max(M, X)
        opponent = solve(i + X, nextM)
        current = suffix[i] - opponent
        best = max(best, current)

    store best
    return best
```

## Complexity

| Type             | Complexity | Explanation                                                                                                  |
| ---------------- | ---------- | ------------------------------------------------------------------------------------------------------------ |
| Time Complexity  | `O(n^3)`   | There are `O(n^2)` possible `(i, M)` states, and each state can try up to `O(n)` possible values of `X`.     |
| Space Complexity | `O(n^2)`   | The memoization table stores results for `(i, M)` states. The suffix array requires `O(n)` additional space. |

Here, `n` is the number of piles.

Since `n <= 100`, the `O(n^3)` dynamic programming solution is practical for the given constraints.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int n;
    vector<int> suffix;
    vector<vector<int>> dp;

    int solve(int i, int M) {
        if (i == n) {
            return 0;
        }

        if (dp[i][M] != -1) {
            return dp[i][M];
        }

        int best = 0;

        for (int X = 1; X <= 2 * M && i + X <= n; X++) {
            int nextM = max(M, X);
            int current = suffix[i] - solve(i + X, nextM);
            best = max(best, current);
        }

        return dp[i][M] = best;
    }

    int stoneGameII(vector<int>& piles) {
        n = piles.size();

        suffix.assign(n + 1, 0);
        for (int i = n - 1; i >= 0; i--) {
            suffix[i] = suffix[i + 1] + piles[i];
        }

        dp.assign(n, vector<int>(n + 1, -1));

        return solve(0, 1);
    }
};
```

### Java

```java
class Solution {
    private int n;
    private int[] suffix;
    private int[][] dp;

    private int solve(int i, int m) {
        if (i == n) {
            return 0;
        }

        if (dp[i][m] != -1) {
            return dp[i][m];
        }

        int best = 0;

        for (int x = 1; x <= 2 * m && i + x <= n; x++) {
            int nextM = Math.max(m, x);
            int current = suffix[i] - solve(i + x, nextM);
            best = Math.max(best, current);
        }

        return dp[i][m] = best;
    }

    public int stoneGameII(int[] piles) {
        n = piles.length;
        suffix = new int[n + 1];

        for (int i = n - 1; i >= 0; i--) {
            suffix[i] = suffix[i + 1] + piles[i];
        }

        dp = new int[n][n + 1];

        for (int i = 0; i < n; i++) {
            java.util.Arrays.fill(dp[i], -1);
        }

        return solve(0, 1);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} piles
 * @return {number}
 */
var stoneGameII = function(piles) {
    const n = piles.length;
    const suffix = new Array(n + 1).fill(0);

    for (let i = n - 1; i >= 0; i--) {
        suffix[i] = suffix[i + 1] + piles[i];
    }

    const dp = Array.from({ length: n }, () => new Array(n + 1).fill(-1));

    const solve = (i, M) => {
        if (i === n) {
            return 0;
        }

        if (dp[i][M] !== -1) {
            return dp[i][M];
        }

        let best = 0;

        for (let X = 1; X <= 2 * M && i + X <= n; X++) {
            const nextM = Math.max(M, X);
            const current = suffix[i] - solve(i + X, nextM);
            best = Math.max(best, current);
        }

        dp[i][M] = best;
        return best;
    };

    return solve(0, 1);
};
```

### Python3

```python
from typing import List
from functools import lru_cache

class Solution:
    def stoneGameII(self, piles: List[int]) -> int:
        n = len(piles)
        suffix = [0] * (n + 1)

        for i in range(n - 1, -1, -1):
            suffix[i] = suffix[i + 1] + piles[i]

        @lru_cache(None)
        def solve(i, M):
            if i == n:
                return 0

            best = 0

            for X in range(1, min(2 * M, n - i) + 1):
                next_M = max(M, X)
                current = suffix[i] - solve(i + X, next_M)
                best = max(best, current)

            return best

        return solve(0, 1)
```

### Go

```go
func stoneGameII(piles []int) int {
 n := len(piles)

 suffix := make([]int, n+1)
 for i := n - 1; i >= 0; i-- {
  suffix[i] = suffix[i+1] + piles[i]
 }

 dp := make([][]int, n)
 for i := 0; i < n; i++ {
  dp[i] = make([]int, n+1)
  for j := 0; j <= n; j++ {
   dp[i][j] = -1
  }
 }

 var solve func(int, int) int

 solve = func(i, m int) int {
  if i == n {
   return 0
  }

  if dp[i][m] != -1 {
   return dp[i][m]
  }

  best := 0

  for x := 1; x <= 2*m && i+x <= n; x++ {
   nextM := m
   if x > nextM {
    nextM = x
   }

   current := suffix[i] - solve(i+x, nextM)
   if current > best {
    best = current
   }
  }

  dp[i][m] = best
  return best
 }

 return solve(0, 1)
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in all five languages. Only the syntax and the way memoization is represented are different.

### 1. Store the number of piles

I first store the size of the input array as `n`.

This is useful throughout the solution because I need to know when I have reached the end of the piles.

The state always works with an index between `0` and `n`.

### 2. Build the suffix sum

I create an array of size `n + 1`.

The extra position represents the case where there are no piles remaining.

So:

`suffix[n] = 0`

Then I process the piles from right to left.

For each index:

`suffix[i] = piles[i] + suffix[i + 1]`

This means `suffix[i]` gives me the total number of stones available from index `i`.

For example:

```text
piles = [2, 7, 9, 4, 4]

suffix = [26, 24, 17, 8, 4, 0]
```

If I am currently at index `2`, I immediately know that there are:

`17`

stones remaining.

### 3. Define the DP state

I define a recursive function using two values:

`i`

and:

`M`

The state means:

> What is the maximum number of stones the current player can collect when the first remaining pile is `i` and the current value of `M` is `M`?

This is the key observation behind the entire solution.

I do not need to remember every previous move.

Once I know `i` and `M`, the future game is completely determined.

### 4. Handle the base case

When:

`i == n`

there are no piles left.

So the current player gets:

`0`

more stones.

Every language checks this condition before trying any moves.

### 5. Check memoization

Before solving a state, I check whether its answer already exists.

In C++, Java, JavaScript, and Go, I use a two-dimensional array.

In Python, I use `lru_cache`.

If the state was already calculated, I immediately return the saved answer.

This prevents the same game state from being solved many times.

### 6. Try every possible X

The rules allow the current player to take from:

`1`

to:

`2 * M`

piles.

I therefore loop through every valid `X`.

I also stop if `i + X` reaches beyond the end of the array.

This handles cases where fewer than `2 * M` piles remain.

For example, if there are only three piles left and I am allowed to take up to six, I can only take one, two, or three.

### 7. Calculate the next M

After taking `X` piles, the game changes the value of `M`:

`newM = max(M, X)`

This step is important.

I cannot simply use `X` as the new `M`.

If my current `M` is `4` and I take only `2` piles, the new value remains:

`max(4, 2) = 4`

If I take `6` piles, it becomes:

`max(4, 6) = 6`

### 8. Move to the opponent's state

After taking `X` piles, the next player starts at:

`i + X`

with:

`newM`

So I calculate:

`solve(i + X, newM)`

This tells me how many stones the next player can optimally collect.

### 9. Calculate my score

This is the main game-theory part.

Suppose there are `30` stones remaining before my move.

If the opponent can eventually collect `18` of those stones, then I must get:

`30 - 18 = 12`

stones.

Therefore:

`current score = suffix[i] - opponent score`

This works because every remaining stone eventually belongs to either the current player or the opponent.

I do not need to separately calculate the current player's immediate stones.

The suffix sum already contains all stones that are still available.

### 10. Choose the best move

For every possible `X`, I calculate my final score.

Then I keep the largest value.

This represents optimal play because the current player will always choose the move that gives them the highest possible final score.

### 11. Save the result

Once all possible moves have been tested, I store the best result for `(i, M)`.

This is the memoization step.

If another sequence of moves reaches the same state later, the algorithm can directly use this stored value.

### 12. Return the initial state

The game always begins at:

`i = 0`

and:

`M = 1`

So the final answer is the result of:

`solve(0, 1)`

That value represents the maximum number of stones Alice can collect when both players play optimally.

### C++ Implementation Notes

The C++ implementation uses:

* `vector<int>` for the suffix array.
* `vector<vector<int>>` for the DP table.
* A recursive member function for the state transition.

I initialize the DP table with `-1` so I can distinguish an uncalculated state from a state whose answer is `0`.

### Java Implementation Notes

The Java implementation uses:

* `int[]` for suffix sums.
* `int[][]` for memoization.
* A private recursive method for the DP calculation.

Before starting the recursion, I fill the DP table with `-1`.

Java's `Arrays.fill()` makes this initialization simple.

### JavaScript Implementation Notes

The JavaScript implementation uses:

* A normal array for suffix sums.
* An array of arrays for the memoization table.
* A nested function for recursion.

JavaScript arrays are dynamic, but I still create the required dimensions in advance so every state can be accessed directly.

### Python3 Implementation Notes

The Python implementation uses:

* A list for suffix sums.
* `lru_cache` for memoization.
* A nested recursive function for the DP state.

`lru_cache` automatically stores the result for each `(i, M)` pair.

This makes the Python version shorter while keeping the same dynamic programming logic.

### Go Implementation Notes

The Go implementation uses:

* A slice for suffix sums.
* A two-dimensional slice for the DP table.
* A recursive closure for the DP function.

Go does not have a built-in equivalent of Python's `lru_cache`, so I explicitly create and manage the memoization table.

## Examples

### Example 1

Input:

```text
piles = [2,7,9,4,4]
```

Expected output:

```text
10
```

Alice starts with `M = 1`.

She can take either one or two piles.

One useful path is:

```text
Alice takes 1 pile  -> gets 2
Bob takes 2 piles   -> gets 7 + 9
Alice takes 2 piles -> gets 4 + 4
```

Alice gets:

```text
2 + 4 + 4 = 10
```

The DP checks all possible choices and confirms that `10` is Alice's maximum possible score.

### Example 2

Input:

```text
piles = [1,2,3,4,5,100]
```

Expected output:

```text
104
```

The last pile contains `100` stones, so the players' decisions around the earlier piles become very important.

The dynamic programming solution considers every possible state and every valid choice of `X`.

With optimal play, Alice can collect:

```text
104
```

stones.

### Example 3

Input:

```text
piles = [1]
```

Expected output:

```text
1
```

There is only one pile.

Alice starts with `M = 1`, so she can take that pile immediately.

Therefore, Alice gets all `1` stone.

This also checks the smallest possible input size.

## How to Use / Run Locally

The code in this repository is written in the standard LeetCode `Solution` format. If you want to run it locally, you can place the solution inside a small driver program.

### C++

Save the solution as:

```text
main.cpp
```

Compile it with:

```bash
g++ -std=c++17 main.cpp -o main
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

For a standalone program, add a `main` method that creates the input array and calls `stoneGameII()`.

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

You can create a sample array and call the `stoneGameII()` function from the file to test the result.

### Python3

Save the solution as:

```text
solution.py
```

Run:

```bash
python3 solution.py
```

For local testing, create a sample `piles` list and call the `stoneGameII()` method.

### Go

Save the program as:

```text
main.go
```

Run it directly with:

```bash
go run main.go
```

You can also build it first:

```bash
go build main.go
```

Then execute the generated program.

## Notes & Optimizations

The most important optimization is memoization.

A plain recursive solution would repeatedly solve the same `(i, M)` states. That can lead to a very large number of repeated calculations.

By storing every state, I calculate each state only once.

The suffix sum is another important optimization. Without it, finding the total number of remaining stones could require another loop for every transition.

I also stop the `X` loop when `i + X > n`. There is no reason to consider taking more piles than actually remain.

A useful edge case is when all piles are available within the current `2 * M` limit. In that situation, the current player can take everything remaining and immediately get `suffix[i]`.

Another important detail is that `M` can grow during the game. It never decreases because:

`M = max(M, X)`

So the DP state must include `M`. Tracking only the current index would not be enough.

The solution uses top-down dynamic programming with memoization. A bottom-up DP approach is possible, but the recursive version matches the game structure more naturally and is easier to follow.

For the given constraint of `n <= 100`, the `O(n^3)` time and `O(n^2)` space solution is efficient enough.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
