# 877. Stone Game - LeetCode Solution | Greedy Mathematical Observation | O(1) Time

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

The **Stone Game** is a classic game theory and competitive programming problem where two players, Alice and Bob, take turns picking an entire pile of stones from either the beginning or the end of a row.

Alice always makes the first move, and both players play perfectly. The goal is to determine whether Alice can always finish the game with more stones than Bob.

The input is an array named `piles`, where each element represents the number of stones in a pile.

The output is a boolean value:

* `true` if Alice wins.
* `false` otherwise.

Although the problem may look like a Dynamic Programming problem at first, there is a much simpler mathematical observation that leads to an optimal **O(1)** solution.

---

## Constraints

| Constraint             | Value                      |
| ---------------------- | -------------------------- |
| Number of piles        | `2 <= piles.length <= 500` |
| Piles count            | Always even                |
| Stones in each pile    | `1 <= piles[i] <= 500`     |
| Total number of stones | Always odd                 |

---

## Intuition

When I first read this problem, my instinct was to solve it using Dynamic Programming because both players make optimal choices.

After looking more carefully at the constraints, I noticed something much more interesting.

The number of piles is always even.

That means Alice can always decide whether she wants to collect all piles that originally sit at even indices or all piles that originally sit at odd indices.

Since the total number of stones is guaranteed to be odd, the sum of stones in these two groups can never be the same.

One group must always contain more stones than the other.

Alice simply chooses the larger group before the game even begins and follows that strategy throughout the game.

This observation completely removes the need for recursion, memoization, or Dynamic Programming.

---

## Approach

I solved this problem using a mathematical observation instead of simulating the game.

The steps are straightforward.

1. Notice that the number of piles is always even.
2. Realize that Alice can force herself to pick either all even-indexed piles or all odd-indexed piles.
3. Since the total number of stones is odd, one of these groups always has a larger total.
4. Alice follows the larger group during the game.
5. Because of this guaranteed strategy, Alice always wins.
6. Return `true`.

This gives the most optimized solution possible.

---

## Data Structures Used

This solution does not require any additional data structures.

| Data Structure | Purpose                                                   |
| -------------- | --------------------------------------------------------- |
| None           | The mathematical proof is enough to determine the answer. |

---

## Operations & Behavior Summary

The algorithm performs only one logical operation.

* Read the input.
* Use the mathematical property of the game.
* Conclude that Alice always has a winning strategy.
* Return `true`.

No loops are needed.

No recursion is needed.

No Dynamic Programming table is created.

No extra memory is allocated.

---

## Complexity

| Metric           | Complexity | Explanation                                                 |
| ---------------- | ---------- | ----------------------------------------------------------- |
| Time Complexity  | **O(1)**   | The answer is always the same regardless of the input size. |
| Space Complexity | **O(1)**   | No extra memory or data structures are used.                |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    bool stoneGame(vector<int>& piles) {
        // Alice always has a winning strategy.
        // So the answer is always true.
        return true;
    }
};
```

### Java

```java
class Solution {
    public boolean stoneGame(int[] piles) {
        // Alice always has a winning strategy.
        // So the answer is always true.
        return true;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} piles
 * @return {boolean}
 */
var stoneGame = function(piles) {
    // Alice always has a winning strategy.
    // So the answer is always true.
    return true;
};
```

### Python3

```python
class Solution:
    def stoneGame(self, piles: List[int]) -> bool:
        # Alice always has a winning strategy.
        # So the answer is always True.
        return True
```

### Go

```go
func stoneGame(piles []int) bool {
    // Alice always has a winning strategy.
    // So the answer is always true.
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in every programming language because the solution is based entirely on mathematics rather than implementation details.

The first thing the function receives is the input array of stone piles.

Normally, we might expect to simulate turns or calculate the best possible score for both players.

However, the constraints allow us to skip all of that.

Because the number of piles is always even, Alice can control which parity of piles she collects throughout the game.

She can either collect every pile that originally appears at an even index or every pile that originally appears at an odd index.

The total number of stones is always odd, so the sums of these two groups can never be equal.

One group must contain strictly more stones.

Alice simply follows the larger group, guaranteeing that she finishes with more stones than Bob.

Since this strategy always works, the function immediately returns `true`.

This behavior is exactly the same in C++, Java, JavaScript, Python3, and Go.

The only difference between the implementations is the syntax used by each programming language.

---

## Examples

### Example 1

**Input**

```text
piles = [5,3,4,5]
```

**Output**

```text
true
```

**Explanation**

Alice chooses the strategy that guarantees she collects the larger parity group.

No matter how Bob responds, Alice finishes with more stones.

---

### Example 2

**Input**

```text
piles = [3,7,2,3]
```

**Output**

```text
true
```

**Explanation**

Again, Alice follows the parity with the larger total.

Since Bob cannot prevent this strategy, Alice wins.

---

### Example 3

**Input**

```text
piles = [1,100,2,99]
```

**Output**

```text
true
```

**Explanation**

Even though the values vary significantly, the mathematical property of the game still holds.

Alice can always force the better parity and therefore wins.

---

## How to Use / Run Locally

### C++

Compile the program.

```bash
g++ solution.cpp -o solution
```

Run it.

```bash
./solution
```

---

### Java

Compile the file.

```bash
javac Solution.java
```

Run it.

```bash
java Solution
```

---

### JavaScript

Run the solution using Node.js.

```bash
node solution.js
```

---

### Python3

Run the file with Python.

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run the solution.

```bash
go run solution.go
```

---

## Notes & Optimizations

* This problem is often categorized under Dynamic Programming because it has an optimal substructure, and a DP solution can be written.
* A recursive or memoized solution works but is unnecessary for the given constraints.
* The mathematical observation provides the best possible solution.
* The algorithm runs in constant time and constant space.
* No extra arrays, stacks, queues, hash maps, or recursion are required.
* The proof depends on the problem guarantees that the number of piles is even and the total number of stones is odd. Without these guarantees, this shortcut would not always be valid.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
