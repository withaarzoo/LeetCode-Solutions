# 1510. Stone Game IV - Dynamic Programming Solution

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

LeetCode 1510, **Stone Game IV**, is a two-player game where Alice and Bob take turns removing stones from a pile.

Alice always plays first.

On each turn, a player can remove any non-zero perfect square number of stones. For example, a player can remove `1`, `4`, `9`, `16`, and so on, as long as that many stones are available.

The player who cannot make a valid move loses the game.

Given an integer `n`, the goal is to determine whether Alice can win if both players play optimally.

The function should return:

* `true` if Alice can force a win.
* `false` if Bob can force a win.

The main idea behind this **DSA problem** is to use **dynamic programming** to identify winning and losing game states.

## Constraints

* `1 <= n <= 10^5`
* `n` is a positive integer.
* A move can remove only a non-zero perfect square.
* Both players play optimally.

## Intuition

I started by looking at what happens when there are only a few stones.

With `1` stone, Alice can remove `1` and leave `0`. Bob cannot move, so Alice wins.

With `2` stones, Alice can only remove `1`. That leaves `1` stone for Bob. Bob removes it and wins, so Alice loses.

This made me realize that every number of stones can be treated as a separate game state.

For any number `i`, I can try removing every possible perfect square from it.

If I can make a move that leaves my opponent in a losing state, then the current state is winning.

For example, if I am at `4`, I can remove `4` directly:

```text
4 -> 0
```

The player at `0` stones cannot move, so `4` is a winning state.

This naturally leads to dynamic programming.

I store whether each number of stones is winning or losing and use those previously calculated states to determine the next state.

## Approach

I use a boolean DP array where `dp[i]` represents whether the player whose turn it is can win when there are exactly `i` stones.

The base state is:

```text
dp[0] = false
```

With zero stones, there is no valid move, so the current player loses.

Then I calculate the states from `1` through `n`.

For each `i`, I try all perfect squares that are less than or equal to `i`.

If I remove a square `j * j`, the opponent gets:

```text
i - j * j
```

If that remaining state is losing, then I have found a winning move.

So I mark `dp[i]` as `true`.

If none of the possible moves leads to a losing state, `dp[i]` remains `false`.

At the end, `dp[n]` gives the answer for Alice.

## Data Structures Used

### Boolean DP Array

I use a one-dimensional boolean array of size `n + 1`.

```text
dp[i] = true
```

means the current player can force a win with `i` stones.

```text
dp[i] = false
```

means the current player will lose if both players play optimally.

The DP array is enough because every state depends only on smaller numbers of stones.

No graph, stack, queue, map, or other complex data structure is required.

## Operations & Behavior Summary

The algorithm works in the following order:

1. Create a DP array from `0` to `n`.
2. Set `dp[0]` to `false` because a player with zero stones cannot move.
3. Start calculating states from `1` to `n`.
4. For each state `i`, generate perfect squares using `j * j`.
5. Check every square while `j * j <= i`.
6. Calculate the remaining stones after removing that square.
7. If the remaining state is losing for the opponent, mark the current state as winning.
8. Stop checking more squares once a winning move is found.
9. If no winning move exists, the state remains losing.
10. Return `dp[n]`.

In simple pseudocode, the logic is:

```text
dp[0] = false

for every number i from 1 to n:
    try every perfect square <= i

    if removing that square reaches a losing state:
        dp[i] = true
        stop checking

return dp[n]
```

## Complexity

| Complexity       |     Cost | Explanation                                                                         |
| ---------------- | -------: | ----------------------------------------------------------------------------------- |
| Time Complexity  | `O(n√n)` | There are `n` DP states, and for each state I may check up to `√n` perfect squares. |
| Space Complexity |   `O(n)` | I store one boolean value for every number of stones from `0` to `n`.               |

Here, `n` is the number of stones in the initial pile.

The `√n` factor comes from the number of perfect squares up to `n`.

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool winnerSquareGame(int n) {
        // dp[i] tells me whether the current player can win with i stones.
        vector<bool> dp(n + 1, false);

        // With 0 stones, the current player has no move and loses.
        dp[0] = false;

        // Calculate the result for every number of stones from 1 to n.
        for (int i = 1; i <= n; ++i) {
            // Try removing every perfect square that is at most i.
            for (int j = 1; j * j <= i; ++j) {
                // If removing j*j gives the opponent a losing state,
                // then I can make this move and win.
                if (!dp[i - j * j]) {
                    dp[i] = true;

                    // One winning move is enough, so I stop checking squares.
                    break;
                }
            }
        }

        // Return whether Alice can force a win from n stones.
        return dp[n];
    }
};
```

### Java

```java
class Solution {
    public boolean winnerSquareGame(int n) {
        // dp[i] tells me whether the current player can win with i stones.
        boolean[] dp = new boolean[n + 1];

        // With 0 stones, the current player has no valid move and loses.
        dp[0] = false;

        // Calculate the result for every number of stones from 1 to n.
        for (int i = 1; i <= n; i++) {
            // Try every perfect square that can be removed from i.
            for (int j = 1; j * j <= i; j++) {
                // If the remaining state is losing for the opponent,
                // this move lets me force a win.
                if (!dp[i - j * j]) {
                    dp[i] = true;

                    // I only need one winning move, so I can stop here.
                    break;
                }
            }
        }

        // Return whether Alice can win with n stones.
        return dp[n];
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @return {boolean}
 */
var winnerSquareGame = function(n) {
    // dp[i] tells me whether the current player can win with i stones.
    const dp = new Array(n + 1).fill(false);

    // With 0 stones, there is no possible move, so the player loses.
    dp[0] = false;

    // Calculate the winning or losing state for every number of stones.
    for (let i = 1; i <= n; i++) {
        // Try removing every perfect square that is at most i.
        for (let j = 1; j * j <= i; j++) {
            // If the remaining state is losing for the opponent,
            // I can remove this square and force a win.
            if (!dp[i - j * j]) {
                dp[i] = true;

                // One winning move is enough, so stop checking.
                break;
            }
        }
    }

    // Return whether Alice can force a win from n stones.
    return dp[n];
};
```

### Python3

```python
class Solution:
    def winnerSquareGame(self, n: int) -> bool:
        # dp[i] tells me whether the current player can win with i stones.
        dp = [False] * (n + 1)

        # With 0 stones, the current player has no move and loses.
        dp[0] = False

        # Calculate the result for every number of stones from 1 to n.
        for i in range(1, n + 1):
            # Try every perfect square that can be removed from i.
            j = 1
            while j * j <= i:
                # If the remaining state is losing for the opponent,
                # I can make this move and force a win.
                if not dp[i - j * j]:
                    dp[i] = True

                    # One winning move is enough, so stop checking squares.
                    break

                # Move to the next possible perfect square.
                j += 1

        # Return whether Alice can force a win with n stones.
        return dp[n]
```

### Go

```go
func winnerSquareGame(n int) bool {
 // dp[i] tells me whether the current player can win with i stones.
 dp := make([]bool, n+1)

 // With 0 stones, the current player has no move and loses.
 dp[0] = false

 // Calculate the result for every number of stones from 1 to n.
 for i := 1; i <= n; i++ {
  // Try every perfect square that can be removed from i.
  for j := 1; j*j <= i; j++ {
   // If the remaining state is losing for the opponent,
   // I can remove this square and force a win.
   if !dp[i-j*j] {
    dp[i] = true

    // One winning move is enough, so stop checking squares.
    break
   }
  }
 }

 // Return whether Alice can force a win with n stones.
 return dp[n]
}
```

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The core algorithm is the same in all five languages. Only the syntax and array handling are different.

### 1. Create the DP array

I create a boolean array with `n + 1` positions.

I need `n + 1` because I want to store states from `0` through `n`.

For example, when `n = 5`, the states are:

```text
0, 1, 2, 3, 4, 5
```

Each position answers whether the player can win from that state.

### 2. Handle the zero-stone state

The state `0` is a losing state.

There are no stones left, so the player whose turn it is cannot remove anything.

Therefore:

```text
dp[0] = false
```

This state is also important because many other states can directly reach `0`.

For example:

```text
1 -> 0
4 -> 0
9 -> 0
16 -> 0
```

All of these are immediately winning states for the player making the move.

### 3. Process states from small to large

I process every number from `1` to `n`.

This order matters because the result for a larger state depends on smaller states.

For example, when calculating `dp[10]`, I may need:

```text
dp[9]
dp[6]
dp[1]
```

depending on which perfect square I remove.

All of these states have already been calculated.

### 4. Generate perfect squares

For every state `i`, I generate squares using:

```text
1 * 1
2 * 2
3 * 3
...
```

I continue as long as the square is not greater than `i`.

For `i = 10`, the possible moves are:

```text
remove 1
remove 4
remove 9
```

I cannot remove `16` because there are not enough stones.

### 5. Check the resulting state

Suppose I have `10` stones and remove `4`.

The opponent gets:

```text
10 - 4 = 6
```

So I check `dp[6]`.

If:

```text
dp[6] = false
```

then the opponent is in a losing state.

That means removing `4` was a winning move for me, so:

```text
dp[10] = true
```

### 6. Stop after finding one winning move

I do not need to find every possible winning move.

One winning move is enough.

Once I find a square that sends the opponent to a losing state, I immediately stop checking the remaining squares.

This does not change the worst-case complexity, but it avoids unnecessary work in many cases.

### 7. What if no winning move exists?

Suppose I check every possible perfect square and every resulting state is winning for the opponent.

Then there is no move that lets me force a win.

The state stays:

```text
dp[i] = false
```

This means the player whose turn it is loses if both players make optimal decisions.

### 8. Final result

After calculating every state, I return:

```text
dp[n]
```

If it is `true`, Alice has a winning strategy.

If it is `false`, Bob has a winning strategy.

### C++ Behavior

In C++, the DP state can be stored in a boolean vector.

The nested loops calculate the states and test perfect squares.

The integer square generation avoids creating a separate collection of all perfect squares.

### Java Behavior

Java uses a boolean array for the same DP table.

The logic is identical to C++.

The array is initialized with `false`, which also naturally gives the correct default value for states that remain losing.

### JavaScript Behavior

JavaScript uses an array containing boolean values.

The array is initialized with `false` so every state starts as a losing state.

A state is changed to `true` only when a winning move is found.

### Python3 Behavior

Python uses a list of boolean values.

The DP list is created with `n + 1` positions.

For every state, a loop checks the possible perfect squares and updates the state when a winning move is found.

### Go Behavior

Go uses a boolean slice with `n + 1` elements.

Boolean values in a newly created slice start as `false`, so the initial state naturally represents losing positions.

The nested loops then update the states that have at least one winning move.

## Examples

### Example 1

Input:

```text
n = 1
```

The only possible move is to remove `1`.

```text
1 -> 0
```

The next player has no move.

Expected output:

```text
true
```

Alice wins.

---

### Example 2

Input:

```text
n = 2
```

Alice can only remove `1`.

```text
2 -> 1
```

Bob now has `1` stone and removes it:

```text
1 -> 0
```

Alice has no move left.

Expected output:

```text
false
```

Alice loses.

---

### Example 3

Input:

```text
n = 4
```

`4` itself is a perfect square.

Alice can remove all four stones:

```text
4 -> 0
```

Bob has no valid move.

Expected output:

```text
true
```

Alice wins immediately.

## How to Use / Run Locally

The repository can contain the same algorithm implemented in C++, Java, JavaScript, Python3, and Go.

Since the code blocks above are intentionally left empty, paste the corresponding solution into the appropriate source file before running it.

### C++

Save the solution as:

```text
solution.cpp
```

Compile it with:

```text
g++ -std=c++17 solution.cpp -o solution
```

Run it with:

```text
./solution
```

If you are using Windows with MinGW, the generated executable can be run as:

```text
solution.exe
```

### Java

Save the solution as:

```text
Solution.java
```

Compile it with:

```text
javac Solution.java
```

Run it with:

```text
java Solution
```

For a LeetCode submission, keep the required `Solution` class and method signature.

### JavaScript

Save the solution as:

```text
solution.js
```

Run it with Node.js:

```text
node solution.js
```

Make sure Node.js is installed and available in your terminal.

### Python3

Save the solution as:

```text
solution.py
```

Run it with:

```text
python3 solution.py
```

On some Windows installations, the command may be:

```text
python solution.py
```

### Go

Save the solution as:

```text
solution.go
```

Run it directly with:

```text
go run solution.go
```

You can also build an executable with:

```text
go build solution.go
```

## Notes & Optimizations

The most important observation is that I do not need to simulate the entire game between Alice and Bob.

Instead, I only need to know whether each number of stones is a winning or losing state.

A state is winning if I can make at least one move that sends the opponent to a losing state.

A state is losing if every possible move sends the opponent to a winning state.

The base case `dp[0] = false` is essential because a player with zero stones cannot make a move.

For `n = 1`, the answer is immediately `true` because `1` is a perfect square.

For values that are themselves perfect squares, the current player can remove all stones in one move, so those states are always winning.

The solution uses `O(n)` memory. This is reasonable for the given constraint of `n <= 10^5`.

The straightforward DP solution runs in `O(n√n)` time, which is efficient enough for the given constraints.

Another possible approach is to use game theory and study the pattern of winning and losing states, but the DP solution is much easier to understand and directly proves the result for every state.

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
