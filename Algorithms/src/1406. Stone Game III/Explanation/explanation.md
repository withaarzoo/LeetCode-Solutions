# 1406. Stone Game III - Dynamic Programming Solution | LeetCode Hard

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
- [Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

**Stone Game III** is a classic Dynamic Programming and Game Theory problem from LeetCode.

Alice and Bob are playing a game with a row of stones. Every stone has an integer value, which can be positive or negative.

Alice always plays first. During each turn, a player can take **1, 2, or 3 stones** from the beginning of the remaining row. The values of the stones they pick are added to their score.

Both players always make the best possible move.

The goal is to determine the final winner after every stone has been taken.

Return:

- `"Alice"` if Alice finishes with a higher score.
- `"Bob"` if Bob finishes with a higher score.
- `"Tie"` if both players end with the same score.

This problem is a great example of using **Dynamic Programming**, **Minimax thinking**, and **optimal game strategy** together.

---

## Constraints

| Constraint | Value |
|------------|-------|
| Number of stones | `1 <= stoneValue.length <= 5 × 10^4` |
| Stone value | `-1000 <= stoneValue[i] <= 1000` |

---

## Intuition

My first thought was to simulate every possible game.

That idea quickly became impossible because every move creates up to three new choices. The number of possibilities grows extremely fast.

Then I realized I don't actually need to know every sequence of moves.

Instead, I only need to know the **best score difference** that the current player can achieve from every position.

Once I know the answer for future positions, I can calculate the answer for the current position by trying all three possible moves.

This naturally leads to a Dynamic Programming solution.

---

## Approach

I solve the problem from the end of the array toward the beginning.

For every position, I consider taking:

- one stone
- two stones
- three stones

For every choice:

1. Add the values of the stones I pick.
2. Assume the opponent will also play perfectly.
3. Subtract the opponent's best possible score difference.
4. Keep the move that gives the largest final score difference.

After processing every position, the answer is stored at the first index.

Finally:

- positive score difference means Alice wins
- negative score difference means Bob wins
- zero means the game ends in a tie

---

## Data Structures Used

### Dynamic Programming Array

A one-dimensional DP array stores the maximum score difference starting from every index.

Why I used it:

- avoids solving the same subproblem repeatedly
- gives constant-time lookup for future states
- keeps the overall solution linear

### Integer Variables

A few integer variables are used to keep track of:

- current sum
- current index
- best score difference

No additional complex data structures are required.

---

## Operations & Behavior Summary

The algorithm works in the following order:

1. Create a DP array.
2. Start from the last stone.
3. Move backward toward the first stone.
4. Try taking 1 stone.
5. Try taking 2 stones.
6. Try taking 3 stones.
7. Calculate the score difference for every choice.
8. Store the best possible result.
9. Continue until the beginning of the array.
10. Decide whether Alice wins, Bob wins, or the game ends in a tie.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | **O(n)** | Each position checks at most three possible moves. |
| Space Complexity | **O(n)** | One DP array stores the best score difference for every index. |

Where:

- **n** = number of stones in the input array.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    string stoneGameIII(vector<int>& stoneValue) {
        int n = stoneValue.size();

        // dp[i] = maximum score difference current player can achieve
        // starting from index i
        vector<int> dp(n + 1, 0);

        // Process from back because current state depends on future states
        for (int i = n - 1; i >= 0; i--) {

            // Initialize with a very small value since we are maximizing
            dp[i] = INT_MIN;

            int sum = 0;

            // Try taking 1, 2 and 3 stones
            for (int j = i; j < min(n, i + 3); j++) {

                // Current player's collected score
                sum += stoneValue[j];

                // Current score difference =
                // current collected score - opponent's best difference
                dp[i] = max(dp[i], sum - dp[j + 1]);
            }
        }

        // Decide the winner using the final score difference
        if (dp[0] > 0) return "Alice";
        if (dp[0] < 0) return "Bob";
        return "Tie";
    }
};
```

### Java

```java
class Solution {
    public String stoneGameIII(int[] stoneValue) {

        int n = stoneValue.length;

        // dp[i] stores the maximum score difference from index i
        int[] dp = new int[n + 1];

        // Fill DP from right to left
        for (int i = n - 1; i >= 0; i--) {

            dp[i] = Integer.MIN_VALUE;
            int sum = 0;

            // Try taking 1, 2 and 3 stones
            for (int j = i; j < Math.min(n, i + 3); j++) {

                // Add current stone value
                sum += stoneValue[j];

                // Update the best possible score difference
                dp[i] = Math.max(dp[i], sum - dp[j + 1]);
            }
        }

        // Decide the winner
        if (dp[0] > 0) return "Alice";
        if (dp[0] < 0) return "Bob";
        return "Tie";
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} stoneValue
 * @return {string}
 */
var stoneGameIII = function(stoneValue) {

    const n = stoneValue.length;

    // dp[i] stores the best score difference from index i
    const dp = new Array(n + 1).fill(0);

    // Build DP from right to left
    for (let i = n - 1; i >= 0; i--) {

        dp[i] = -Infinity;
        let sum = 0;

        // Try taking 1, 2 and 3 stones
        for (let j = i; j < Math.min(n, i + 3); j++) {

            // Current collected score
            sum += stoneValue[j];

            // Choose the move with maximum score difference
            dp[i] = Math.max(dp[i], sum - dp[j + 1]);
        }
    }

    // Return the final result
    if (dp[0] > 0) return "Alice";
    if (dp[0] < 0) return "Bob";
    return "Tie";
};
```

### Python3

```python
class Solution:
    def stoneGameIII(self, stoneValue: List[int]) -> str:

        n = len(stoneValue)

        # dp[i] stores the maximum score difference from index i
        dp = [0] * (n + 1)

        # Fill DP from back to front
        for i in range(n - 1, -1, -1):

            dp[i] = float("-inf")
            total = 0

            # Try taking 1, 2 and 3 stones
            for j in range(i, min(n, i + 3)):

                # Add current stone value
                total += stoneValue[j]

                # Update the best score difference
                dp[i] = max(dp[i], total - dp[j + 1])

        # Decide the winner
        if dp[0] > 0:
            return "Alice"
        elif dp[0] < 0:
            return "Bob"
        else:
            return "Tie"
```

### Go

```go
func stoneGameIII(stoneValue []int) string {

 n := len(stoneValue)

 // dp[i] stores the maximum score difference from index i
 dp := make([]int, n+1)

 // Process from right to left
 for i := n - 1; i >= 0; i-- {

  // Very small initial value
  dp[i] = -(1 << 30)

  sum := 0

  // Try taking 1, 2 and 3 stones
  for j := i; j < n && j < i+3; j++ {

   // Add current stone value
   sum += stoneValue[j]

   // Keep the maximum score difference
   if sum-dp[j+1] > dp[i] {
    dp[i] = sum - dp[j+1]
   }
  }
 }

 // Decide the winner
 if dp[0] > 0 {
  return "Alice"
 }
 if dp[0] < 0 {
  return "Bob"
 }
 return "Tie"
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in all five implementations.

Only the programming language syntax changes.

### Step 1

Create a DP array.

Each position stores the maximum score difference the current player can achieve from that index.

The last position represents the situation where no stones remain.

Its value is zero because no player can gain additional points.

---

### Step 2

Process the array from right to left.

Every answer depends on future positions.

Working backward guarantees those answers already exist when they are needed.

---

### Step 3

For every index, try taking one stone.

Calculate the current score.

Subtract the opponent's best possible score difference.

Store the result.

---

### Step 4

Repeat the same process after taking two stones.

Sometimes taking more stones immediately gives a better result.

Sometimes it creates a worse future position.

The DP formula automatically compares both situations.

---

### Step 5

Try taking three stones.

Again, compare the resulting score difference with previous choices.

Keep whichever option produces the largest advantage.

---

### Step 6

Store the maximum score difference for the current position.

Once this value is saved, future calculations can reuse it directly.

---

### Step 7

Continue until reaching the first stone.

The value stored at index zero represents the entire game.

---

### Step 8

Determine the winner.

- Positive value means Alice earns more points.
- Negative value means Bob earns more points.
- Zero means both players finish with the same score.

The implementation is identical across C++, Java, JavaScript, Python3, and Go.

Only the syntax for arrays, loops, and return statements changes.

---

## Examples

### Example 1

**Input**

```text
stoneValue = [1,2,3,7]
```

**Output**

```text
Bob
```

**Trace**

Alice tries every possible first move.

No matter what she chooses, Bob can respond with a better strategy and finish with a higher score.

---

### Example 2

**Input**

```text
stoneValue = [1,2,3,-9]
```

**Output**

```text
Alice
```

**Trace**

Alice takes the first three stones.

Bob is forced to take the negative stone.

Alice finishes with the larger score.

---

### Example 3

**Input**

```text
stoneValue = [1,2,3,6]
```

**Output**

```text
Tie
```

**Trace**

Both players can force the game into a draw by making optimal decisions.

Neither player can gain a final advantage.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone https://github.com/your-username/your-repository.git
```

Move into the project folder.

```bash
cd your-repository
```

### C++

Compile

```bash
g++ solution.cpp -std=c++17
```

Run

```bash
./a.out
```

---

### Java

Compile

```bash
javac Solution.java
```

Run

```bash
java Solution
```

---

### JavaScript

Run

```bash
node solution.js
```

---

### Python3

Run

```bash
python solution.py
```

or

```bash
python3 solution.py
```

---

### Go

Run

```bash
go run solution.go
```

---

## Notes & Optimizations

- This is a Dynamic Programming solution with optimal time complexity.
- Every state is calculated only once.
- Only three transitions are checked for every position.
- A recursive memoization solution also works, but the iterative DP version avoids recursion depth issues on large inputs.
- The algorithm handles positive numbers, negative numbers, and mixed values correctly.
- This approach is commonly used in competitive programming, coding interviews, and LeetCode Hard Dynamic Programming problems involving Game Theory.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
