# 3336. Find the Number of Subsequences With Equal GCD

A beginner-friendly Dynamic Programming (DP) solution for **LeetCode 3336 - Find the Number of Subsequences With Equal GCD**. This repository explains the intuition, approach, complexity analysis, and provides implementations in **C++, Java, JavaScript, Python3, and Go**.

---

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

You are given an integer array `nums`.

Your task is to count the number of pairs of **non-empty disjoint subsequences** where both subsequences have exactly the same Greatest Common Divisor (GCD).

Since the number of valid pairs can become very large, the final answer must be returned modulo **10^9 + 7**.

This is a challenging Dynamic Programming and Number Theory problem because every element can either belong to the first subsequence, the second subsequence, or neither. The goal is to efficiently track the GCD of both subsequences while processing the array.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= nums.length <= 200` | Array size |
| `1 <= nums[i] <= 200` | Element value |

---

## Intuition

The first thing I noticed was that every element has only three possible choices.

- Ignore it.
- Put it into the first subsequence.
- Put it into the second subsequence.

Instead of remembering the actual subsequences, I only need to know their current GCD values.

That observation completely changes the problem.

Since every value is at most `200`, the possible GCD values are also limited to `200`. This makes it possible to build a Dynamic Programming solution where each state stores only the two current GCD values.

---

## Approach

I process the array one element at a time.

For every processed element, I maintain a DP table where:

- The first dimension represents the GCD of the first subsequence.
- The second dimension represents the GCD of the second subsequence.

Initially, both subsequences are empty.

For every number, I try three different choices.

1. Skip the current number.
2. Add it to the first subsequence.
3. Add it to the second subsequence.

Whenever I add a number to a subsequence, I update its GCD.

If the subsequence is empty, its GCD simply becomes the current number.

Otherwise, the new GCD becomes:

`gcd(oldGCD, currentNumber)`

After processing every element, I simply count all states where:

- Both subsequences are non-empty.
- Both GCD values are equal.

---

## Data Structures Used

### 1. 2D Dynamic Programming Table

Stores the number of ways to reach every pair of GCD values.

Why?

Because the only information needed for future decisions is the current GCD of both subsequences.

---

### 2. Temporary DP Table

For every new element, I build a fresh DP table.

Why?

Updating the current table directly would cause the same element to be counted multiple times.

---

### 3. GCD Function

Used whenever a new number is inserted into one of the subsequences.

This efficiently updates the current GCD without recomputing everything.

---

## Operations & Behavior Summary

The algorithm follows these steps.

1. Create a DP table.
2. Start with both subsequences empty.
3. Process every number one by one.
4. For every DP state:
   - Skip the number.
   - Insert into the first subsequence.
   - Insert into the second subsequence.
5. Store every new state in a fresh DP table.
6. Replace the old DP table.
7. After processing all numbers, count every state where:
   - First GCD equals second GCD.
   - Neither subsequence is empty.
8. Return the answer modulo `10^9 + 7`.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | `O(n × M²)` | `n` is the array length and `M = 200` is the maximum possible GCD value. |
| Space Complexity | `O(M²)` | Two DP tables are maintained, each storing all possible GCD pairs. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int subsequencePairCount(vector<int>& nums) {
        const int MOD = 1000000007;
        const int MAX = 200;

        // Current DP table.
        vector<vector<int>> dp(MAX + 1, vector<int>(MAX + 1, 0));
        dp[0][0] = 1; // Both subsequences are initially empty.

        for (int x : nums) {
            // Next DP table after processing current number.
            vector<vector<int>> ndp(MAX + 1, vector<int>(MAX + 1, 0));

            for (int g1 = 0; g1 <= MAX; g1++) {
                for (int g2 = 0; g2 <= MAX; g2++) {
                    if (dp[g1][g2] == 0) continue;

                    long long ways = dp[g1][g2];

                    // Choice 1: Skip current number.
                    ndp[g1][g2] = (ndp[g1][g2] + ways) % MOD;

                    // Choice 2: Put current number into first subsequence.
                    int ng1 = (g1 == 0) ? x : gcd(g1, x);
                    ndp[ng1][g2] = (ndp[ng1][g2] + ways) % MOD;

                    // Choice 3: Put current number into second subsequence.
                    int ng2 = (g2 == 0) ? x : gcd(g2, x);
                    ndp[g1][ng2] = (ndp[g1][ng2] + ways) % MOD;
                }
            }

            dp = move(ndp);
        }

        long long ans = 0;

        // Count only states where both subsequences are non-empty
        // and have the same GCD.
        for (int g = 1; g <= MAX; g++) {
            ans = (ans + dp[g][g]) % MOD;
        }

        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int subsequencePairCount(int[] nums) {
        final int MOD = 1_000_000_007;
        final int MAX = 200;

        // Current DP table.
        int[][] dp = new int[MAX + 1][MAX + 1];
        dp[0][0] = 1;

        for (int x : nums) {
            // Next DP table after processing current number.
            int[][] next = new int[MAX + 1][MAX + 1];

            for (int g1 = 0; g1 <= MAX; g1++) {
                for (int g2 = 0; g2 <= MAX; g2++) {
                    if (dp[g1][g2] == 0) continue;

                    long ways = dp[g1][g2];

                    // Choice 1: Skip current number.
                    next[g1][g2] = (int)((next[g1][g2] + ways) % MOD);

                    // Choice 2: Put into first subsequence.
                    int ng1 = (g1 == 0) ? x : gcd(g1, x);
                    next[ng1][g2] = (int)((next[ng1][g2] + ways) % MOD);

                    // Choice 3: Put into second subsequence.
                    int ng2 = (g2 == 0) ? x : gcd(g2, x);
                    next[g1][ng2] = (int)((next[g1][ng2] + ways) % MOD);
                }
            }

            dp = next;
        }

        long ans = 0;

        // Count states where both GCDs are equal and non-zero.
        for (int g = 1; g <= MAX; g++) {
            ans = (ans + dp[g][g]) % MOD;
        }

        return (int)ans;
    }

    // Computes GCD using Euclid's algorithm.
    private int gcd(int a, int b) {
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var subsequencePairCount = function(nums) {
    const MOD = 1000000007;
    const MAX = 200;

    // Current DP table.
    let dp = Array.from({ length: MAX + 1 }, () => Array(MAX + 1).fill(0));
    dp[0][0] = 1;

    // Computes GCD using Euclid's algorithm.
    const gcd = (a, b) => {
        while (b !== 0) {
            const t = a % b;
            a = b;
            b = t;
        }
        return a;
    };

    for (const x of nums) {
        // Next DP table after processing current number.
        let next = Array.from({ length: MAX + 1 }, () => Array(MAX + 1).fill(0));

        for (let g1 = 0; g1 <= MAX; g1++) {
            for (let g2 = 0; g2 <= MAX; g2++) {
                if (dp[g1][g2] === 0) continue;

                const ways = dp[g1][g2];

                // Choice 1: Skip current number.
                next[g1][g2] = (next[g1][g2] + ways) % MOD;

                // Choice 2: Put into first subsequence.
                const ng1 = (g1 === 0) ? x : gcd(g1, x);
                next[ng1][g2] = (next[ng1][g2] + ways) % MOD;

                // Choice 3: Put into second subsequence.
                const ng2 = (g2 === 0) ? x : gcd(g2, x);
                next[g1][ng2] = (next[g1][ng2] + ways) % MOD;
            }
        }

        dp = next;
    }

    let ans = 0;

    // Count states where both GCDs are equal and non-zero.
    for (let g = 1; g <= MAX; g++) {
        ans = (ans + dp[g][g]) % MOD;
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def subsequencePairCount(self, nums: List[int]) -> int:
        from math import gcd

        MOD = 10 ** 9 + 7
        MAX = 200

        # Current DP table.
        dp = [[0] * (MAX + 1) for _ in range(MAX + 1)]
        dp[0][0] = 1

        for x in nums:
            # Next DP table after processing current number.
            ndp = [[0] * (MAX + 1) for _ in range(MAX + 1)]

            for g1 in range(MAX + 1):
                for g2 in range(MAX + 1):
                    if dp[g1][g2] == 0:
                        continue

                    ways = dp[g1][g2]

                    # Choice 1: Skip current number.
                    ndp[g1][g2] = (ndp[g1][g2] + ways) % MOD

                    # Choice 2: Put into first subsequence.
                    ng1 = x if g1 == 0 else gcd(g1, x)
                    ndp[ng1][g2] = (ndp[ng1][g2] + ways) % MOD

                    # Choice 3: Put into second subsequence.
                    ng2 = x if g2 == 0 else gcd(g2, x)
                    ndp[g1][ng2] = (ndp[g1][ng2] + ways) % MOD

            dp = ndp

        ans = 0

        # Count states where both GCDs are equal and non-zero.
        for g in range(1, MAX + 1):
            ans = (ans + dp[g][g]) % MOD

        return ans
```

### Go

```go
func subsequencePairCount(nums []int) int {
 const MOD = 1000000007
 const MAX = 200

 // Computes GCD using Euclid's algorithm.
 gcd := func(a, b int) int {
  for b != 0 {
   a, b = b, a%b
  }
  return a
 }

 // Current DP table.
 dp := make([][]int, MAX+1)
 for i := range dp {
  dp[i] = make([]int, MAX+1)
 }
 dp[0][0] = 1

 for _, x := range nums {
  // Next DP table after processing current number.
  next := make([][]int, MAX+1)
  for i := range next {
   next[i] = make([]int, MAX+1)
  }

  for g1 := 0; g1 <= MAX; g1++ {
   for g2 := 0; g2 <= MAX; g2++ {
    if dp[g1][g2] == 0 {
     continue
    }

    ways := dp[g1][g2]

    // Choice 1: Skip current number.
    next[g1][g2] = (next[g1][g2] + ways) % MOD

    // Choice 2: Put into first subsequence.
    ng1 := x
    if g1 != 0 {
     ng1 = gcd(g1, x)
    }
    next[ng1][g2] = (next[ng1][g2] + ways) % MOD

    // Choice 3: Put into second subsequence.
    ng2 := x
    if g2 != 0 {
     ng2 = gcd(g2, x)
    }
    next[g1][ng2] = (next[g1][ng2] + ways) % MOD
   }
  }

  dp = next
 }

 ans := 0

 // Count states where both GCDs are equal and non-zero.
 for g := 1; g <= MAX; g++ {
  ans = (ans + dp[g][g]) % MOD
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The overall logic is identical in every language. Only the syntax changes.

### Step 1

Create a 2D DP table.

Each cell stores how many different ways exist to reach a pair of GCD values.

Initially both subsequences are empty.

---

### Step 2

Process each element one by one.

Instead of modifying the existing DP table, create a fresh one.

This avoids using the current element multiple times during the same iteration.

---

### Step 3

Visit every DP state.

Each state represents:

- Current GCD of the first subsequence.
- Current GCD of the second subsequence.

---

### Step 4

For every state, try all three possible choices.

First choice:

Skip the current element.

Nothing changes.

Second choice:

Insert the element into the first subsequence.

If the subsequence is empty, the GCD becomes the element itself.

Otherwise, update it using the GCD function.

Third choice:

Insert the element into the second subsequence.

The update rule is exactly the same.

---

### Step 5

After processing the current element, replace the old DP table with the newly created one.

Now continue with the next element.

---

### Step 6

After every number has been processed, examine every DP state.

Only keep states where:

- Both subsequences are non-empty.
- Both GCD values are identical.

Add all those counts together.

Return the final answer modulo `10^9 + 7`.

Since all five implementations follow the same Dynamic Programming strategy, only the programming syntax differs. The algorithm, transitions, and complexity remain exactly the same.

---

## Examples

### Example 1

**Input**

```text
nums = [1,2,3,4]
```

**Output**

```text
10
```

**Explanation**

The DP explores every possible assignment of each element into the first subsequence, second subsequence, or neither.

After processing all elements, exactly 10 valid pairs have equal non-zero GCD values.

---

### Example 2

**Input**

```text
nums = [10,20,30]
```

**Output**

```text
2
```

**Trace**

The algorithm builds all valid GCD states while processing each number.

Only two disjoint subsequence pairs end up with the same GCD of 10.

---

### Example 3

**Input**

```text
nums = [1,1,1,1]
```

**Output**

```text
50
```

**Trace**

Every non-empty subsequence has a GCD of 1.

The DP simply counts all valid disjoint assignments and returns the total number of valid pairs.

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

- The maximum value in the array is only `200`, which makes a GCD-state Dynamic Programming solution practical.
- Using only the current GCD instead of storing the entire subsequence keeps the state space small.
- Creating a fresh DP table for every iteration prevents counting the same element more than once.
- This solution is significantly faster than generating all subsequences, which would require exponential time.
- The algorithm combines Dynamic Programming and Number Theory to achieve an efficient solution within the given constraints.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
