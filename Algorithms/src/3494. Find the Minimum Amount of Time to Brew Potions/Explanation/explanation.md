# 3494. Find the Minimum Amount of Time to Brew Potions

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
* [How to use / Run locally](#how-to-use--run-locally)
* [Notes & Optimizations](#notes--optimizations)
* [Author](#author)

---

## Problem Summary

I am given two integer arrays: `skill` of length `n` (the `n` wizards), and `mana` of length `m` (the `m` potions). A potion must pass through all `n` wizards in order. The time taken by wizard `i` to process potion `j` is `skill[i] * mana[j]`.

A potion must be passed immediately from wizard `i` to wizard `i+1` as soon as wizard `i` finishes it. The order of potions does not change (potions are brewed sequentially in the given order). I must compute the **minimum total time** until all potions are fully processed by the final wizard.

In short: schedule potions through a fixed pipeline of wizards so that the handing-off constraints are satisfied, then return the earliest time when the last potion finishes.

## Constraints

* `n == skill.length`.
* `m == mana.length`.
* `1 <= n, m <= 5000`.
* `1 <= skill[i], mana[j] <= 5000`.
* Use 64-bit integers (or language `int` with arbitrary precision) because products and sums can be large.

---

## Intuition

I thought about the setup as a pipeline of `n` machines (wizards) and `m` jobs (potions). Each potion `j` spends `skill[i] * mana[j]` time at wizard `i`.

If I let `S_j` be the time when potion `j` starts at wizard `0`, then the time potion `j` arrives at wizard `i` is `S_j + pref[i-1] * mana[j]` (where `pref[k]` is the sum of skill[0..k]).

To avoid collisions and ensure handoffs are immediate, the start times `S_j` must satisfy constraints derived from comparing potion `j`'s progress with potion `j-1`'s progress. After algebraic manipulation, I found a simple recurrence to compute `S_j` from `S_{j-1}`:

```
S_j = S_{j-1} + max_{0 <= i < n} ( pref[i] * mana[j-1] - pref[i-1] * mana[j] )
```

(Define `pref[-1] = 0`.)

I start with `S_0 = 0`. After computing all `S_j`, the finish time of the last potion (`m-1`) is `S_{m-1} + pref[n-1] * mana[m-1]`.

This leads to a straightforward `O(n * m)` algorithm.

---

## Approach

1. Compute `pref[i] = sum_{k=0..i} skill[k]` (use 64-bit integers).
2. Initialize `S = 0` which is `S_0` (start time for potion 0 at wizard 0).
3. For `j` from `1` to `m-1`:

   * Let `prev = mana[j-1]` and `cur = mana[j]`.
   * Compute `best = max_{i=0..n-1} ( pref[i] * prev - pref[i-1] * cur )` where `pref[-1] = 0`.
   * Set `S += best` (now `S` equals `S_j`).
4. Answer = `S + pref[n-1] * mana[m-1]`.

Notes:

* All multiplications and sums use 64-bit integers.
* Complexity is `O(n*m)` and space is `O(n)` for the prefix sums.

---

## Data Structures Used

* Plain arrays / vectors for `skill`, `mana`, and `pref` (prefix sums).
* Constant number of scalar variables (64-bit) for accumulation and comparisons.

---

## Operations & Behavior Summary

* The inner operation scans all `n` wizards once for each consecutive pair of potions (`j-1` and `j`) and computes a candidate value per wizard. We take the maximum across wizards and add it to the accumulated start time.
* The final finish time is the start time for the last potion plus the total work required by all `n` wizards on that potion.

Mathematical key formula used in the code (derived from pipeline constraints):

```
S_j = S_{j-1} + max_i (pref[i] * mana[j-1] - pref[i-1] * mana[j])
```

where `pref[-1] = 0` and `pref[i] = sum_{k=0..i} skill[k]`.

---

## Complexity

* **Time Complexity:** `O(n * m)` where `n = skill.length` and `m = mana.length`. For each of the `m-1` transitions I scan all `n` wizards.
* **Space Complexity:** `O(n)` to store `pref`. All other space is `O(1)`.

For the constraints `n, m <= 5000`, the worst-case number of iterations is `n*m = 25,000,000` which is reasonable in compiled languages (C++, Java, Go). Python needs micro-optimizations to stay fast (local variable binding, simple loops).

---

## Multi-language Solutions

### C++

```c++
#include <vector>
#include <climits>
using namespace std;

class Solution {
public:
    // Return minimum finish time as 64-bit
    long long minTime(vector<int>& skill, vector<int>& mana) {
        int n = skill.size();
        int m = mana.size();
        if (m == 0) return 0LL;

        // 1) prefix sums of skills
        vector<long long> pref(n);
        for (int i = 0; i < n; ++i) {
            pref[i] = skill[i] + (i ? pref[i-1] : 0LL);
        }

        // S holds current S_{j-1}. Start S_0 = 0
        long long S = 0LL;

        // 2) compute S_j iteratively using the derived formula
        for (int j = 1; j < m; ++j) {
            long long prev = (long long)mana[j-1];
            long long cur  = (long long)mana[j];
            long long best = LLONG_MIN;
            for (int i = 0; i < n; ++i) {
                long long prev_pref = (i ? pref[i-1] : 0LL);
                long long cand = pref[i] * prev - prev_pref * cur;
                if (cand > best) best = cand;
            }
            S += best;
        }

        // 3) final finish time
        long long ans = S + pref[n-1] * (long long)mana[m-1];
        return ans;
    }
};
```

---

### Java

```java
class Solution {
    public long minTime(int[] skill, int[] mana) {
        int n = skill.length;
        int m = mana.length;
        if (m == 0) return 0L;

        long[] pref = new long[n];
        for (int i = 0; i < n; ++i) {
            pref[i] = skill[i] + (i > 0 ? pref[i-1] : 0L);
        }

        long S = 0L;
        for (int j = 1; j < m; ++j) {
            long prev = (long) mana[j-1];
            long cur  = (long) mana[j];
            long best = Long.MIN_VALUE;
            for (int i = 0; i < n; ++i) {
                long prev_pref = (i > 0 ? pref[i-1] : 0L);
                long cand = pref[i] * prev - prev_pref * cur;
                if (cand > best) best = cand;
            }
            S += best;
        }

        return S + pref[n-1] * (long) mana[m-1];
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} skill
 * @param {number[]} mana
 * @return {number}
 */
var minTime = function(skill, mana) {
    const n = skill.length;
    const m = mana.length;
    if (m === 0) return 0;

    // Use BigInt to avoid intermediate overflow; final answer fits in JS Number for given constraints
    const pref = new Array(n);
    for (let i = 0; i < n; ++i) {
        pref[i] = BigInt(skill[i]) + (i ? pref[i-1] : 0n);
    }

    let S = 0n;
    for (let j = 1; j < m; ++j) {
        const prev = BigInt(mana[j-1]);
        const cur  = BigInt(mana[j]);
        let best = null;
        for (let i = 0; i < n; ++i) {
            const prev_pref = i ? pref[i-1] : 0n;
            const cand = pref[i] * prev - prev_pref * cur;
            if (best === null || cand > best) best = cand;
        }
        S += best;
    }

    const ans = S + pref[n-1] * BigInt(mana[m-1]);
    return Number(ans); // safe for constraints
};
```

---

### Python3

```python3
class Solution:
    def minTime(self, skill: List[int], mana: List[int]) -> int:
        n = len(skill)
        m = len(mana)
        if m == 0:
            return 0

        # prefix sums
        pref = [0] * n
        for i in range(n):
            pref[i] = skill[i] + (pref[i-1] if i > 0 else 0)

        S = 0
        # iterate over consecutive potions
        for j in range(1, m):
            prev = mana[j-1]
            cur  = mana[j]
            best = -10**30
            for i in range(n):
                prev_pref = pref[i-1] if i > 0 else 0
                cand = pref[i] * prev - prev_pref * cur
                if cand > best:
                    best = cand
            S += best

        return S + pref[-1] * mana[-1]
```

---

### Go

```go
package main

func minTime(skill []int, mana []int) int64 {
    n := len(skill)
    m := len(mana)
    if m == 0 {
        return 0
    }

    // prefix sums
    pref := make([]int64, n)
    for i := 0; i < n; i++ {
        if i == 0 {
            pref[0] = int64(skill[0])
        } else {
            pref[i] = pref[i-1] + int64(skill[i])
        }
    }

    S := int64(0)
    const NEG_INF int64 = -1 << 62

    for j := 1; j < m; j++ {
        prev := int64(mana[j-1])
        cur  := int64(mana[j])
        best := NEG_INF
        for i := 0; i < n; i++ {
            prev_pref := int64(0)
            if i > 0 {
                prev_pref = pref[i-1]
            }
            cand := pref[i]*prev - prev_pref*cur
            if cand > best {
                best = cand
            }
        }
        S += best
    }

    return S + pref[n-1]*int64(mana[m-1])
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm in small steps and tie these steps to the code above (the same logic applies to all five languages):

1. **Prefix sums (`pref`)**

   * I compute `pref[i] = sum_{k=0..i} skill[k]`.
   * This lets me compute how long a potion with mana `x` takes to reach wizard `i` as `pref[i] * x` (the total work done by wizards `0..i`).

2. **Start time `S_j` idea**

   * Let `S_j` be the time potion `j` begins at wizard `0`.
   * By the immediate-handoff rule, for every wizard `i` we must have the arrival time of potion `j` at wizard `i` no earlier than the time the previous potion `j-1` finishes at wizard `i`.
   * After algebra, the constraint reduces to the recurrence used in code.

3. **Compute transitions**

   * For every pair `(j-1, j)` we loop over all `i` to compute the candidate expression `pref[i] * mana[j-1] - pref[i-1] * mana[j]` (with `pref[-1] = 0`).
   * We pick the maximum candidate and add it to `S`.

4. **Final answer**

   * After finishing the loop, `S` equals `S_{m-1}` (the start time of the last potion at wizard 0).
   * The last potion finishes at `S_{m-1} + pref[n-1] * mana[m-1]`.

5. **Why max over `i`?**

   * Each wizard `i` imposes a lower bound on `S_j` when comparing potion `j` with `j-1`. The strongest (largest) lower bound determines how much later potion `j` must start.

6. **Edge cases**

   * If `m == 0`, answer is `0`.
   * For `n == 1`, pref is just `skill[0]` and the recurrence still works (the inner loop has one value).

7. **Language-specific micro notes**

   * **C++ / Java / Go:** use 64-bit (`long long` / `long` / `int64`) to avoid overflow.
   * **JavaScript:** I used `BigInt` for intermediate computations and returned `Number(ans)` at the end. This keeps correctness for the constraint range.
   * **Python:** Python `int` is unbounded; still take care to write loops efficiently when `n*m` is large.

---

## Examples

**Example 1**

```
Input: skill = [1,5,2,4], mana = [5,1,4,2]
Output: 110
```

**Example 2**

```
Input: skill = [1,1,1], mana = [1,1,1]
Output: 5
```

**Example 3**

```
Input: skill = [1,2,3,4], mana = [1,2]
Output: 21
```

(See code comments for more explanation of how the examples match the computed schedule.)

---

## How to use / Run locally

* **C++** (g++):

  * Put the `Solution` class into a `.cpp` file along with a `main()` that reads input or tests examples. Compile: `g++ -O2 solution.cpp -o solution` then `./solution`.
* **Java**:

  * Put the class in `Solution.java`, add a `main` to test, then `javac Solution.java` and `java Solution`.
* **JavaScript** (Node):

  * Save the function in a `.js` file and call it from a test harness. Run: `node solution.js`.
* **Python3**:

  * Save the `Solution` class in a `.py` file and add test code under `if __name__ == '__main__':`. Run: `python3 solution.py`.
* **Go**:

  * Put the function in a `main.go`, add `package main` and a `main()` for testing. Run: `go run main.go`.

For LeetCode, paste the language-specific `Solution` implementation into the editor and run the provided test cases.

---

## Notes & Optimizations

* The presented solution is simple and robust. For `n, m <= 5000`, the `O(n*m)` approach performs up to 25 million inner iterations. This is fine in C++, Java, and Go with optimizations enabled.
* In Python, keep the inner loop fast by:

  * Binding `pref` to a local variable (already used in example),
  * Avoiding repeated attribute lookups, and
  * Using PyPy if needed.
* There might exist faster algorithms using advanced techniques (Convex Hull Trick, divide-and-conquer optimizations) under specific distributions of `skill` and `mana`. However, deriving and implementing those is more complex and unnecessary for the given constraints.

---

## Author

[Aarzoo Islam](https://bento.me/withaarzoo)
