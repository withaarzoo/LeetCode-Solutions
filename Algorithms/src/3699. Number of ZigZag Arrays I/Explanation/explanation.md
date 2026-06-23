# Number of ZigZag Arrays I - LeetCode 3699 Solution

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
* [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
* [Examples](#examples)
* [How to Use / Run Locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

LeetCode 3699. Number of ZigZag Arrays I asks us to count how many valid ZigZag arrays can be formed.

We are given three integers:

* `n` → length of the array
* `l` → minimum allowed value
* `r` → maximum allowed value

Every element in the array must belong to the range `[l, r]`.

A valid ZigZag array must satisfy the following conditions:

1. Adjacent elements cannot be equal.
2. No three consecutive elements can form a strictly increasing sequence.
3. No three consecutive elements can form a strictly decreasing sequence.

The goal is to return the total number of valid ZigZag arrays modulo `10^9 + 7`.

This problem is a great example of Dynamic Programming, Prefix Sum Optimization, Alternating Sequences, and Combinatorial Counting.

---

## Constraints

| Constraint       | Value        |
| ---------------- | ------------ |
| 3 ≤ n ≤ 2000     | Array length |
| 1 ≤ l < r ≤ 2000 | Value range  |
| Answer modulo    | 10^9 + 7     |

---

## Intuition

My first observation was that the actual values do not matter as much as their relative order.

The ZigZag condition only cares about whether numbers are larger or smaller than each other.

If three consecutive values become:

```text
a < b < c
```

or

```text
a > b > c
```

the sequence becomes invalid.

That means the direction must keep alternating.

Once the sequence goes up, the next move must go down.

Once the sequence goes down, the next move must go up.

After noticing this pattern, I realized that Dynamic Programming can be used to count valid endings while efficiently maintaining transitions with prefix sums.

---

## Approach

1. Compute the total number of available values:

   ```text
   m = r - l + 1
   ```

2. Create a DP array where each position represents a possible ending value rank.

3. Initialize all positions with `1` because every value can form a valid sequence of length `1`.

4. Build the answer length by length.

5. For every new length:

   * Use prefix-sum style transitions.
   * Reverse the DP array before processing.
   * Update each position using accumulated counts from previous positions.

6. Continue until sequences of length `n` are generated.

7. Sum all valid ending states.

8. Multiply by `2` because both alternating directions are symmetric:

   * Up → Down → Up ...
   * Down → Up → Down ...

9. Return the answer modulo `10^9 + 7`.

---

## Data Structures Used

### 1. Dynamic Programming Array

A one-dimensional array stores the number of valid sequences ending at each rank.

Why I used it:

* Compact memory usage
* Fast updates
* Easy prefix-sum optimization

### 2. Running Prefix Sum Variable

A single integer tracks cumulative counts while updating DP states.

Why I used it:

* Avoids expensive nested loops
* Reduces transition cost dramatically

---

## Operations & Behavior Summary

The algorithm performs these steps:

1. Determine how many distinct values are available.
2. Initialize all possible endings.
3. Repeatedly extend sequence length.
4. Reverse the DP state.
5. Build new states using running prefix sums.
6. Store updated counts modulo `10^9 + 7`.
7. Sum all final states.
8. Multiply by two to include both ZigZag directions.
9. Return the final count.

---

## Complexity

| Metric           | Complexity | Explanation                                                         |
| ---------------- | ---------- | ------------------------------------------------------------------- |
| Time Complexity  | O(n × m)   | For every array length, all possible value ranks are processed once |
| Space Complexity | O(m)       | Only one DP array is maintained                                     |

Where:

* `n` = length of the ZigZag array
* `m = r - l + 1` = number of available values

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    static constexpr int MOD = 1000000007;

    int zigZagArrays(int n, int l, int r) {
        int m = r - l + 1;

        // Length 1: every rank can be chosen once
        vector<int> dp(m, 1);

        for (int len = 2; len <= n; len++) {
            // Reversing allows the same prefix-sum logic
            // to act as alternating prefix/suffix transitions
            reverse(dp.begin(), dp.end());

            long long pref = 0;

            for (int i = 0; i < m; i++) {
                int old = dp[i];     // Previous DP value

                // New value becomes sum of all earlier values
                dp[i] = pref;

                // Update running prefix sum
                pref = (pref + old) % MOD;
            }
        }

        long long ans = 0;

        // Sum all ending ranks
        for (int x : dp) {
            ans = (ans + x) % MOD;
        }

        // Count both starting directions
        return (ans * 2) % MOD;
    }
};
```

### Java

```java
class Solution {
    private static final int MOD = 1000000007;

    public int zigZagArrays(int n, int l, int r) {
        int m = r - l + 1;

        // Length 1: every rank is a valid sequence
        int[] dp = new int[m];
        java.util.Arrays.fill(dp, 1);

        for (int len = 2; len <= n; len++) {

            // Even length -> prefix transition
            if ((len & 1) == 0) {
                long pref = 0;

                for (int i = 0; i < m; i++) {
                    int old = dp[i];

                    // Sum of all smaller ranks
                    dp[i] = (int) pref;

                    pref = (pref + old) % MOD;
                }
            }
            // Odd length -> suffix transition
            else {
                long suff = 0;

                for (int i = m - 1; i >= 0; i--) {
                    int old = dp[i];

                    // Sum of all larger ranks
                    dp[i] = (int) suff;

                    suff = (suff + old) % MOD;
                }
            }
        }

        long ans = 0;

        for (int x : dp) {
            ans = (ans + x) % MOD;
        }

        return (int) ((ans * 2) % MOD);
    }
}
```

### JavaScript

```javascript
/**
 * @param {number} n
 * @param {number} l
 * @param {number} r
 * @return {number}
 */
var zigZagArrays = function(n, l, r) {
    const MOD = 1000000007;
    const m = r - l + 1;

    // Length 1 DP
    const dp = Array(m).fill(1);

    for (let len = 2; len <= n; len++) {

        // Reverse to reuse the same prefix-sum transition
        dp.reverse();

        let pref = 0;

        for (let i = 0; i < m; i++) {
            const old = dp[i];

            // New state = accumulated prefix
            dp[i] = pref;

            pref = (pref + old) % MOD;
        }
    }

    let ans = 0;

    for (const x of dp) {
        ans = (ans + x) % MOD;
    }

    // Both possible starting directions
    return (ans * 2) % MOD;
};
```

### Python3

```python
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        MOD = 1000000007
        m = r - l + 1

        # Length 1: every rank is valid
        dp = [1] * m

        for _ in range(2, n + 1):
            # Reverse so one prefix-sum pass handles
            # the alternating transition automatically
            dp.reverse()

            pref = 0

            for i in range(m):
                old = dp[i]

                # New state gets all previous contributions
                dp[i] = pref

                pref = (pref + old) % MOD

        # Sum all ending positions
        ans = sum(dp) % MOD

        # Count both zigzag directions
        return (ans * 2) % MOD
```

### Go

```go
func zigZagArrays(n int, l int, r int) int {
 const MOD int = 1000000007

 m := r - l + 1

 // Length 1 DP
 dp := make([]int, m)
 for i := 0; i < m; i++ {
  dp[i] = 1
 }

 for length := 2; length <= n; length++ {

  // Reverse the array so the same prefix-sum
  // update works for alternating directions
  for left, right := 0, m-1; left < right; left, right = left+1, right-1 {
   dp[left], dp[right] = dp[right], dp[left]
  }

  pref := 0

  for i := 0; i < m; i++ {
   old := dp[i]

   // Sum of all earlier states
   dp[i] = pref

   pref += old
   pref %= MOD
  }
 }

 ans := 0

 for _, x := range dp {
  ans += x
  ans %= MOD
 }

 // Both starting directions
 return (ans * 2) % MOD
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic remains exactly the same across all five languages.

### Step 1: Calculate Available Values

First, calculate:

```text
m = r - l + 1
```

This tells us how many distinct values can appear in the array.

Instead of working with actual numbers, we only work with their relative positions.

---

### Step 2: Initialize DP

Every value can independently form a sequence of length one.

So:

```text
dp[i] = 1
```

for every valid rank.

This becomes the starting state.

---

### Step 3: Extend the Sequence

For every additional position:

* Reverse the DP array.
* Maintain a running prefix sum.
* Replace each state with the accumulated count before it.

This efficiently computes all valid alternating transitions.

Without this optimization, we would need nested loops and the solution would be too slow.

---

### Step 4: Maintain Modulo Arithmetic

The number of valid arrays grows very quickly.

To prevent overflow:

```text
MOD = 1,000,000,007
```

Every update is performed modulo this value.

---

### Step 5: Sum Final States

After processing length `n`, every DP position represents a valid ending value.

Add all of them together.

---

### Step 6: Count Both ZigZag Directions

The DP counts one alternating pattern.

The opposite pattern produces the same number of valid arrays.

Therefore:

```text
answer = answer × 2
```

Finally return the result modulo `10^9 + 7`.

---

## Examples

### Example 1

Input

```text
n = 3
l = 4
r = 5
```

Output

```text
2
```

Valid arrays:

```text
[4,5,4]
[5,4,5]
```

Only two ZigZag arrays exist.

---

### Example 2

Input

```text
n = 3
l = 1
r = 3
```

Output

```text
10
```

The range contains three values.

Many alternating patterns become possible, producing ten valid ZigZag arrays.

---

### Example 3

Input

```text
n = 4
l = 1
r = 2
```

Output

```text
2
```

Valid arrays:

```text
[1,2,1,2]
[2,1,2,1]
```

With only two values available, the sequence must alternate continuously.

---

## How to Use / Run Locally

### C++

Compile

```bash
g++ solution.cpp -o solution
```

Run

```bash
./solution
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

---

### Go

Run

```bash
go run solution.go
```

---

## Notes & Optimizations

* The key optimization comes from using prefix sums.
* A naive dynamic programming solution would require nested transitions and become too slow.
* The reverse-and-prefix technique allows all transitions to be processed in linear time.
* Only one DP array is needed, reducing memory usage.
* Modulo arithmetic prevents integer overflow.
* This approach comfortably handles the maximum constraints.

Alternative approaches exist, but they typically require more memory or more expensive transitions.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
