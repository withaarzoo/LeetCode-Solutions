# 3514. Number of Unique XOR Triplets II

A clean and optimized solution for **LeetCode 3514 - Number of Unique XOR Triplets II**.

This repository explains the thought process behind the solution, the algorithm used, time and space complexity, and provides implementations in **C++, Java, JavaScript, Python, and Go**. If you're preparing for coding interviews or improving your competitive programming skills, this guide walks through the problem in a simple and beginner-friendly way.

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
- [Step-by-step Detailed Explanation](#step-by-step-detailed-explanation-c-java-javascript-python3-go)
- [Examples](#examples)
- [How to Use / Run Locally](#how-to-use--run-locally)
- [Notes & Optimizations](#notes--optimizations)
- [Author](#author)

---

## Problem Summary

In this problem, we are given an integer array `nums`.

We need to consider every possible triplet of indices `(i, j, k)` where:

- `i ≤ j ≤ k`

For every valid triplet, we calculate:

```
nums[i] XOR nums[j] XOR nums[k]
```

Different triplets may produce the same XOR value. The goal is **not** to count the number of triplets.

Instead, we only need to count **how many distinct XOR values** can be produced.

The main challenge is finding all unique XOR results efficiently without checking every possible triplet, which would be far too slow for the given constraints.

This problem combines **bit manipulation**, **dynamic programming**, and **state transitions** to achieve an efficient solution.

---

## Constraints

| Constraint | Value |
|------------|-------|
| `1 <= nums.length <= 1500` | Array size |
| `1 <= nums[i] <= 1500` | Element value |

---

## Intuition

The first thing I noticed was that the indices satisfy `i ≤ j ≤ k`.

That means I am allowed to choose the same index multiple times.

Once I realized this, I no longer cared about how many times a value appears in the array. I only needed to know whether the value exists.

Since every number is at most `1500`, every XOR result stays within a very small range.

Instead of thinking about millions of triplets, I could think about reachable XOR states.

I gradually build every possible XOR value after choosing one number, then two numbers, and finally three numbers.

This keeps the solution fast while still finding every unique XOR result.

---

## Approach

I solve the problem in the following order.

1. Store every distinct value that exists in the array.
2. Create a DP array that represents every XOR value currently reachable.
3. Initially, only XOR value `0` is reachable because no numbers have been chosen yet.
4. Repeat the transition exactly three times.
5. During each transition:
   - Look at every reachable XOR value.
   - Try adding every value that exists in the array.
   - Mark the new XOR value as reachable.
6. After three rounds, every reachable state represents one valid XOR triplet.
7. Count all reachable XOR values.

This approach avoids checking every possible combination of indices and works within the small XOR state space.

---

## Data Structures Used

| Data Structure | Purpose |
|---------------|---------|
| Boolean Array | Stores which values exist in the input array. |
| Boolean DP Array | Tracks every XOR value that is currently reachable. |
| Temporary DP Array | Stores the next set of reachable XOR values during each transition. |

These structures keep memory usage small while allowing constant-time state updates.

---

## Operations & Behavior Summary

The algorithm performs the following steps:

1. Read every number from the input array.
2. Mark every unique value as available.
3. Initialize the DP state with XOR value `0`.
4. Perform three transition rounds.
5. During each round:
   - Visit every reachable XOR state.
   - Combine it with every available value.
   - Store the newly generated XOR state.
6. Replace the old DP array with the new one.
7. Count every reachable XOR value after the third round.
8. Return the final count.

---

## Complexity

| Complexity | Value | Explanation |
|------------|-------|-------------|
| Time Complexity | `O(n + 3 × 2048 × 2048)` | `n` is the array length. The XOR state space is limited to 2048 values, making the transitions efficient. |
| Space Complexity | `O(2048)` | Only a few fixed-size boolean arrays are used regardless of the input size. |

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int uniqueXorTriplets(vector<int>& nums) {
        const int MAX_XOR = 2048;

        // Store whether a value exists in the array.
        vector<bool> present(MAX_XOR, false);
        for (int x : nums) {
            present[x] = true;
        }

        // Initially only XOR = 0 is possible before picking any value.
        vector<bool> dp(MAX_XOR, false);
        dp[0] = true;

        // Pick exactly 3 values.
        for (int step = 0; step < 3; step++) {
            vector<bool> next(MAX_XOR, false);

            // Try extending every reachable XOR.
            for (int cur = 0; cur < MAX_XOR; cur++) {
                if (!dp[cur]) continue;

                // Add every value that exists in the array.
                for (int v = 0; v < MAX_XOR; v++) {
                    if (present[v]) {
                        next[cur ^ v] = true;
                    }
                }
            }

            dp = move(next);
        }

        // Count all unique XOR values.
        int ans = 0;
        for (bool ok : dp) {
            if (ok) ans++;
        }

        return ans;
    }
};
```

### Java

```java
class Solution {
    public int uniqueXorTriplets(int[] nums) {
        final int MAX_XOR = 2048;

        // Store whether a value exists in the array.
        boolean[] present = new boolean[MAX_XOR];
        for (int x : nums) {
            present[x] = true;
        }

        // Initially only XOR = 0 is reachable.
        boolean[] dp = new boolean[MAX_XOR];
        dp[0] = true;

        // Pick exactly 3 values.
        for (int step = 0; step < 3; step++) {
            boolean[] next = new boolean[MAX_XOR];

            // Extend every reachable XOR.
            for (int cur = 0; cur < MAX_XOR; cur++) {
                if (!dp[cur]) continue;

                // Try every value that exists.
                for (int v = 0; v < MAX_XOR; v++) {
                    if (present[v]) {
                        next[cur ^ v] = true;
                    }
                }
            }

            dp = next;
        }

        // Count unique XOR values.
        int ans = 0;
        for (boolean ok : dp) {
            if (ok) ans++;
        }

        return ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var uniqueXorTriplets = function(nums) {
    const MAX_XOR = 2048;

    // Store whether a value exists.
    const present = new Array(MAX_XOR).fill(false);
    for (const x of nums) {
        present[x] = true;
    }

    // Initially only XOR = 0 is reachable.
    let dp = new Array(MAX_XOR).fill(false);
    dp[0] = true;

    // Pick exactly 3 values.
    for (let step = 0; step < 3; step++) {
        const next = new Array(MAX_XOR).fill(false);

        // Extend every reachable XOR.
        for (let cur = 0; cur < MAX_XOR; cur++) {
            if (!dp[cur]) continue;

            // Try every existing value.
            for (let v = 0; v < MAX_XOR; v++) {
                if (present[v]) {
                    next[cur ^ v] = true;
                }
            }
        }

        dp = next;
    }

    // Count unique XOR values.
    let ans = 0;
    for (const ok of dp) {
        if (ok) ans++;
    }

    return ans;
};
```

### Python3

```python
class Solution:
    def uniqueXorTriplets(self, nums: List[int]) -> int:
        MAX_XOR = 2048

        # Store whether a value exists.
        present = [False] * MAX_XOR
        for x in nums:
            present[x] = True

        # Initially only XOR = 0 is reachable.
        dp = [False] * MAX_XOR
        dp[0] = True

        # Pick exactly 3 values.
        for _ in range(3):
            nxt = [False] * MAX_XOR

            # Extend every reachable XOR.
            for cur in range(MAX_XOR):
                if not dp[cur]:
                    continue

                # Try every existing value.
                for v in range(MAX_XOR):
                    if present[v]:
                        nxt[cur ^ v] = True

            dp = nxt

        # Count unique XOR values.
        return sum(dp)
```

### Go

```go
func uniqueXorTriplets(nums []int) int {
 const MAX_XOR = 2048

 // Store whether a value exists.
 present := make([]bool, MAX_XOR)
 for _, x := range nums {
  present[x] = true
 }

 // Initially only XOR = 0 is reachable.
 dp := make([]bool, MAX_XOR)
 dp[0] = true

 // Pick exactly 3 values.
 for step := 0; step < 3; step++ {
  next := make([]bool, MAX_XOR)

  // Extend every reachable XOR.
  for cur := 0; cur < MAX_XOR; cur++ {
   if !dp[cur] {
    continue
   }

   // Try every existing value.
   for v := 0; v < MAX_XOR; v++ {
    if present[v] {
     next[cur^v] = true
    }
   }
  }

  dp = next
 }

 // Count unique XOR values.
 ans := 0
 for _, ok := range dp {
  if ok {
   ans++
  }
 }

 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is exactly the same in every language.

Only the syntax changes.

The first step is creating a boolean array that records every value present in the input.

This removes duplicate values automatically because multiple copies of the same number never change the set of possible XOR values.

Next, another boolean array represents every XOR value that is currently reachable.

Initially, only XOR value `0` is reachable because no numbers have been selected.

The algorithm performs exactly three rounds.

Each round represents choosing one additional value.

During a round, every reachable XOR value is combined with every value that exists in the array.

The new XOR result becomes reachable in the next DP array.

A temporary array is used instead of modifying the current DP array directly.

If the current array were updated while iterating, newly created states could be reused in the same round, which would incorrectly simulate choosing more than one value at once.

After processing every reachable state, the temporary array becomes the current DP array.

Once three rounds finish, every reachable XOR state corresponds to one valid triplet.

The final answer is simply the number of reachable XOR values.

The C++, Java, JavaScript, Python, and Go implementations all follow this exact sequence of operations.

---

## Examples

### Example 1

**Input**

```text
nums = [1,3]
```

**Output**

```text
2
```

**Trace**

Possible XOR values are:

- 1
- 3

Only two distinct XOR values exist.

---

### Example 2

**Input**

```text
nums = [6,7,8,9]
```

**Output**

```text
4
```

**Trace**

The reachable XOR values become:

```
6
7
8
9
```

There are four distinct XOR values.

---

### Example 3

**Input**

```text
nums = [5]
```

**Output**

```text
1
```

**Trace**

The only possible triplet is:

```
5 XOR 5 XOR 5 = 5
```

Only one unique XOR value exists.

---

## How to Use / Run Locally

Clone the repository.

```bash
git clone <repository-url>
```

Move into the project directory.

```bash
cd <repository-folder>
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

Or build first:

```bash
go build solution.go
```

---

## Notes & Optimizations

- Duplicate values do not affect the final answer because choosing the same index multiple times is already allowed.
- The XOR value range is small, making dynamic programming over XOR states practical.
- A brute-force solution would require checking every triplet, which is far too slow.
- Using boolean arrays keeps both memory usage and transition time low.
- The same algorithm works in every programming language with only syntax differences.
- This is a good example of reducing a large combinatorial search into a small fixed-size state space.

---

## Author

[Md Aarzoo Islam](https://www.instagram.com/code.with.aarzoo/)
