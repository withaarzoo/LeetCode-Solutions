# 1872. Stone Game VIII - LeetCode Solution

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

This repository contains an optimized solution for LeetCode 1872, Stone Game VIII.

Alice and Bob play a game with stones arranged in a row. Alice starts first. On every turn, a player must remove more than one stone from the left side of the row.

The player adds the sum of all removed stones to their score. Then, one new stone with the same value as that sum is placed back at the left side.

The game ends when only one stone remains.

The goal is to find the maximum possible score difference between Alice and Bob when both players play optimally.

The input is an integer array `stones`, where each value represents one stone. The output is the final maximum score difference:

`Alice's score - Bob's score`

The main idea behind this Stone Game VIII solution is to use prefix sums and dynamic programming. Instead of simulating every possible game path, the algorithm works backward and keeps track of the best score difference that can be achieved.

## Constraints

| Constraint                   | Description                                  |
| ---------------------------- | -------------------------------------------- |
| `n == stones.length`         | `n` is the number of stones                  |
| `2 <= n <= 10^5`             | The array can contain up to 100,000 stones   |
| `-10^4 <= stones[i] <= 10^4` | Each stone value can be negative or positive |

Because `n` can be as large as `10^5`, checking every possible game state with a slow recursive solution would not be practical. An `O(n)` solution is needed.

## Intuition

My first thought was to focus on what changes after every move.

Whenever a player removes the first `x` stones, those stones are replaced by one new stone whose value is exactly the sum of the removed stones. That means the important value for every possible move is the prefix sum.

For example, if I know the sum of the first `i + 1` stones, I immediately know the score gained by removing that prefix.

I also noticed that after making a move, the opponent gets the next chance to create the best possible score difference. Since Alice wants to maximize the difference and Bob wants to minimize it, I can think of each state as a dynamic programming problem.

Instead of storing every game state separately, I only need the best result from the next possible state. This allows the solution to be reduced to a single backward pass.

## Approach

I use the following approach:

1. Convert the `stones` array into prefix sums.
2. After this conversion, `stones[i]` represents the sum of all original stones from index `0` to `i`.
3. Start from the total sum because one possible move is to remove all remaining stones.
4. Move backward through the prefix sums.
5. For every position, compare two choices:

   * Keep the best answer already found.
   * Take the current prefix sum and subtract the best result the opponent can achieve afterward.
6. Update the answer using the larger value.
7. Return the final maximum score difference.

The key dynamic programming transition is based on:

`best = max(best, prefixSum - best)`

This gives an optimized `O(n)` time solution for the Stone Game VIII problem.

## Data Structures Used

### Input Array

I reuse the given `stones` array to store prefix sums.

This avoids creating another array and keeps the extra space usage low.

### Integer Variable

I use one variable to store the current best score difference.

Since every state only depends on the previously calculated result, a full dynamic programming array is not necessary.

## Operations & Behavior Summary

The algorithm works in the following order:

1. Read the original stone values.
2. Replace each value with its prefix sum.
3. Treat the total sum as the initial best result.
4. Move from right to left through the valid prefix positions.
5. At every position, calculate what happens if the current player chooses that prefix.
6. Subtract the opponent's best possible result because the opponent will also play optimally.
7. Compare this result with the current best answer.
8. Keep the larger score difference.
9. Return the final result after processing all possible moves.

This avoids recursion, repeated calculations, and expensive game simulation.

## Complexity

| Complexity       | Value              | Explanation                                                                                                                                  |
| ---------------- | ------------------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Time Complexity  | `O(n)`             | `n` is the number of stones. The algorithm makes one pass to build prefix sums and one backward pass to calculate the best score difference. |
| Space Complexity | `O(1)` extra space | The input array is reused for prefix sums, and only a few integer variables are needed. No extra DP array is created.                        |

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int stoneGameVIII(vector<int>& stones) {
        for (int i = 1; i < stones.size(); ++i) {
            stones[i] += stones[i - 1];
        }

        int best = stones.back();

        for (int i = stones.size() - 2; i >= 1; --i) {
            best = max(best, stones[i] - best);
        }

        return best;
    }
};
```

### Java

```java
class Solution {
    public int stoneGameVIII(int[] stones) {
        for (int i = 1; i < stones.length; i++) {
            stones[i] += stones[i - 1];
        }

        int best = stones[stones.length - 1];

        for (int i = stones.length - 2; i >= 1; i--) {
            best = Math.max(best, stones[i] - best);
        }

        return best;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} stones
 * @return {number}
 */
var stoneGameVIII = function(stones) {
    for (let i = 1; i < stones.length; i++) {
        stones[i] += stones[i - 1];
    }

    let best = stones[stones.length - 1];

    for (let i = stones.length - 2; i >= 1; i--) {
        best = Math.max(best, stones[i] - best);
    }

    return best;
};
```

### Python3

```python
from typing import List

class Solution:
    def stoneGameVIII(self, stones: List[int]) -> int:
        for i in range(1, len(stones)):
            stones[i] += stones[i - 1]

        best = stones[-1]

        for i in range(len(stones) - 2, 0, -1):
            best = max(best, stones[i] - best)

        return best
```

### Go

```go
func stoneGameVIII(stones []int) int {
    for i := 1; i < len(stones); i++ {
        stones[i] += stones[i-1]
    }

    best := stones[len(stones)-1]

    for i := len(stones) - 2; i >= 1; i-- {
        best = max(best, stones[i]-best)
    }

    return best
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in C++, Java, JavaScript, Python3, and Go. Only the syntax changes between languages.

The first important step is building prefix sums.

Suppose the input is:

`[-1, 2, -3, 4, -5]`

After converting it into prefix sums, it becomes:

`[-1, 1, -2, 2, -3]`

Now each position tells me the sum of all stones from the beginning up to that position.

For example, the value at index `3` is `2`, which represents:

`-1 + 2 + -3 + 4 = 2`

This matters because every move removes stones only from the left side. So every possible move is based on a prefix sum.

Next, I initialize the answer using the total prefix sum.

This represents the state where the current player takes all remaining stones. It gives me a starting value for the dynamic programming calculation.

Then I move backward from right to left.

I do this because the result at the current position depends on the best result that has already been calculated for a later state.

At each valid prefix position, I have two choices.

The first choice is to keep the current best answer. This means a larger prefix or another move already gives a better result.

The second choice is to remove the current prefix.

If the current prefix sum is `prefixSum`, the current player gains that value immediately. After that, the opponent gets the next move and can force their own optimal score difference.

Because the game is zero-sum, the opponent's best result is subtracted:

`prefixSum - best`

So the transition becomes:

`best = max(best, prefixSum - best)`

This single update replaces the need for a large dynamic programming table.

The loop stops before index `0`.

The reason is that the rules require a player to remove more than one stone. Choosing only the first stone is not a valid move.

The C++, Java, JavaScript, Python3, and Go versions all follow this same pattern:

1. Build prefix sums in place.
2. Store the total sum as the initial answer.
3. Traverse backward through the valid positions.
4. Update the best score difference.
5. Return the result.

The only language-specific difference is syntax, such as array access, loop structure, and maximum-value functions.

## Examples

### Example 1

**Input:**

```text
stones = [-1, 2, -3, 4, -5]
```

**Expected Output:**

```text
5
```

**How the algorithm processes it:**

First, I calculate the prefix sums:

`[-1, 1, -2, 2, -3]`

I start with the total sum:

`best = -3`

Then I process the valid prefix sums from right to left.

Using prefix sum `2`:

`best = max(-3, 2 - (-3)) = 5`

Using prefix sum `-2`:

`best = max(5, -2 - 5) = 5`

Using prefix sum `1`:

`best = max(5, 1 - 5) = 5`

The final answer is:

`5`

### Example 2

**Input:**

```text
stones = [7, -6, 5, 10, 5, -2, -6]
```

**Expected Output:**

```text
13
```

**How the algorithm processes it:**

The total sum of all stones is `13`.

Alice can remove all stones in one move and gain `13`.

Since Bob does not get another move after only one stone remains, the final score difference is:

`13 - 0 = 13`

So the answer is:

`13`

### Example 3

**Input:**

```text
stones = [-10, -12]
```

**Expected Output:**

```text
-22
```

**How the algorithm processes it:**

There are only two stones, so Alice must remove both of them.

Their sum is:

`-10 + -12 = -22`

Alice receives `-22` points, and Bob receives `0` points.

The final score difference is:

`-22`

## How to Use / Run Locally

Create a separate source file for the language you want to run and paste the solution into that file.

### C++

Save the solution as:

```text
main.cpp
```

Compile it using:

```bash
g++ -std=c++17 main.cpp -o main
```

Run it using:

```bash
./main
```

### Java

Save the solution in a Java file such as:

```text
Solution.java
```

Compile it using:

```bash
javac Solution.java
```

Run it using:

```bash
java Solution
```

Depending on how the test code is written, you may need a separate `main` method to run the solution locally.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it using Node.js:

```bash
node solution.js
```

You may need to add your own test input when running outside the LeetCode environment.

### Python3

Save the solution as:

```text
solution.py
```

Run it using:

```bash
python3 solution.py
```

For local testing, add your own input and print the returned result.

### Go

Save the solution as:

```text
solution.go
```

Run it using:

```bash
go run solution.go
```

If the solution is copied directly from LeetCode, add a `main` function when testing it locally.

## Notes & Optimizations

The main optimization in this solution is reusing the original `stones` array for prefix sums.

This removes the need for an additional prefix sum array and keeps the extra space complexity at `O(1)`.

A recursive minimax solution could describe the game naturally, but it would explore many repeated states. Without careful optimization, that approach can become much slower.

A full dynamic programming array would also work, but it is unnecessary here because each transition only needs one previously calculated best value.

Negative stone values are important in this problem. The algorithm must not assume that taking more stones always improves the score. The backward dynamic programming transition handles positive and negative values correctly.

The solution also correctly handles the smallest possible input size of two stones. In that case, the first player has only one valid move: remove both stones.

This prefix sum and dynamic programming approach is the most efficient way to solve LeetCode 1872 Stone Game VIII within the given constraints.

## Author

[Md Aarzoo Islam](https://www.instagram.com/codewithaarzoo.in/)
