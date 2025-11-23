# 1262. Greatest Sum Divisible by Three

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

> **LeetCode 1262 – Greatest Sum Divisible by Three**

Given an integer array `nums`, I need to **find the maximum possible sum of a subset** of elements such that this sum is **divisible by 3**.

I can choose any subset:

* I may use all numbers, some of them, or even none of them.
* The goal is simply to maximize the sum that is divisible by 3.

---

## Constraints

* `1 <= nums.length <= 4 * 10^4`
* `1 <= nums[i] <= 10^4`

---

## Intuition

When I first saw the problem, I noticed one key thing:

* For divisibility by 3, only the **remainder modulo 3** matters.

So I thought:

1. If I sum all numbers:

   * If the total sum `% 3 == 0`, I am already done.
2. If the sum is **not** divisible by 3, I don’t need to change much.

   * I just need to **remove a small total value** so that the remaining sum becomes divisible by 3.
3. Every number has remainder `0`, `1`, or `2` when divided by 3.
4. That means:

   * If `sum % 3 == 1` →
     I must remove either:

     * one small number with `remainder 1`, or
     * two small numbers with `remainder 2`.
   * If `sum % 3 == 2` →
     I must remove either:

     * one small number with `remainder 2`, or
     * two small numbers with `remainder 1`.

So the problem converts to:

> **Track a few smallest numbers by remainder and subtract the cheapest valid combination from the total sum.**

---

## Approach

Step-by-step, my approach is:

1. **Compute the total sum** of all elements in `nums`.
2. While looping, I also track:

   * The **two smallest numbers** where `num % 3 == 1`.
   * The **two smallest numbers** where `num % 3 == 2`.
3. After processing the array:

   * If `sum % 3 == 0` → return `sum`.
4. If `sum % 3 == 1`:

   * Option A: remove the **smallest remainder-1** number.
   * Option B: remove the **two smallest remainder-2** numbers.
   * Choose the option with **minimum removal cost**.
5. If `sum % 3 == 2`:

   * Option A: remove the **smallest remainder-2** number.
   * Option B: remove the **two smallest remainder-1** numbers.
   * Again, choose the minimum removal cost.
6. The answer is:

   ```text
   result = sum - best_removal_cost
   ```

7. If there is no valid option (for example we don’t have enough numbers of a certain remainder), then the best we can do is return `0` (i.e., choose no elements).

This gives me an **O(n)** scan with **O(1)** extra memory.

---

## Data Structures Used

I intentionally keep it very light:

* A few scalar variables:

  * `sum` – total sum of all elements.
  * `r1_min1`, `r1_min2` – smallest and second smallest numbers with remainder 1.
  * `r2_min1`, `r2_min2` – smallest and second smallest numbers with remainder 2.
* No arrays, no DP tables, no hash maps.

Everything fits into basic variables → **constant space**.

---

## Operations & Behavior Summary

For each element `x` in `nums`:

1. Add `x` to `sum`.
2. Compute `r = x % 3`.
3. If `r == 1`:

   * Update the two smallest remainder-1 values.
4. If `r == 2`:

   * Update the two smallest remainder-2 values.
5. Ignore `r == 0` for tracking, because removing them doesn’t help fix divisibility (they don’t change `sum % 3`).

At the end:

* Decide how to adjust the total sum based on its remainder (`sum % 3`).
* Carefully choose the minimal removal cost from the stored candidates.
* Return `sum - removal_cost` (or `0` if impossible).

---

## Complexity

* **Time Complexity:** `O(n)`

  * `n` = length of `nums`.
  * I only scan the array once and do constant work for each element.

* **Space Complexity:** `O(1)`

  * I only store a few integer variables, regardless of input size.
  * No extra arrays or complex data structures.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int maxSumDivThree(vector<int>& nums) {
        long long sum = 0;
        // A large number to represent "infinity" / not set
        const int INF = 1e9;

        // Two smallest numbers with remainder 1
        int r1_min1 = INF, r1_min2 = INF;
        // Two smallest numbers with remainder 2
        int r2_min1 = INF, r2_min2 = INF;

        for (int x : nums) {
            sum += x;
            int r = x % 3;

            if (r == 1) {
                // Maintain two smallest remainder-1 values
                if (x < r1_min1) {
                    r1_min2 = r1_min1;
                    r1_min1 = x;
                } else if (x < r1_min2) {
                    r1_min2 = x;
                }
            } else if (r == 2) {
                // Maintain two smallest remainder-2 values
                if (x < r2_min1) {
                    r2_min2 = r2_min1;
                    r2_min1 = x;
                } else if (x < r2_min2) {
                    r2_min2 = x;
                }
            }
        }

        int mod = sum % 3;
        if (mod == 0) return (int)sum;

        long long removeCost = 1e18; // very large number

        if (mod == 1) {
            // Remove one remainder-1 OR two remainder-2 numbers
            if (r1_min1 != INF) removeCost = min(removeCost, (long long)r1_min1);
            if (r2_min2 != INF) removeCost = min(removeCost, (long long)r2_min1 + r2_min2);
        } else { // mod == 2
            // Remove one remainder-2 OR two remainder-1 numbers
            if (r2_min1 != INF) removeCost = min(removeCost, (long long)r2_min1);
            if (r1_min2 != INF) removeCost = min(removeCost, (long long)r1_min1 + r1_min2);
        }

        if (removeCost >= 1e18) return 0; // no valid subset
        return (int)(sum - removeCost);
    }
};
```

---

### Java

```java
class Solution {
    public int maxSumDivThree(int[] nums) {
        long sum = 0;
        int INF = (int)1e9;

        int r1_min1 = INF, r1_min2 = INF; // remainder-1
        int r2_min1 = INF, r2_min2 = INF; // remainder-2

        for (int x : nums) {
            sum += x;
            int r = x % 3;

            if (r == 1) {
                if (x < r1_min1) {
                    r1_min2 = r1_min1;
                    r1_min1 = x;
                } else if (x < r1_min2) {
                    r1_min2 = x;
                }
            } else if (r == 2) {
                if (x < r2_min1) {
                    r2_min2 = r2_min1;
                    r2_min1 = x;
                } else if (x < r2_min2) {
                    r2_min2 = x;
                }
            }
        }

        int mod = (int)(sum % 3);
        if (mod == 0) return (int)sum;

        long removeCost = (long)1e18;

        if (mod == 1) {
            if (r1_min1 != INF) removeCost = Math.min(removeCost, (long)r1_min1);
            if (r2_min2 != INF) removeCost = Math.min(removeCost, (long)r2_min1 + r2_min2);
        } else { // mod == 2
            if (r2_min1 != INF) removeCost = Math.min(removeCost, (long)r2_min1);
            if (r1_min2 != INF) removeCost = Math.min(removeCost, (long)r1_min1 + r1_min2);
        }

        if (removeCost >= (long)1e18) return 0;
        return (int)(sum - removeCost);
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSumDivThree = function(nums) {
    let sum = 0;
    const INF = 1e9;

    let r1_min1 = INF, r1_min2 = INF; // two smallest with remainder 1
    let r2_min1 = INF, r2_min2 = INF; // two smallest with remainder 2

    for (const x of nums) {
        sum += x;
        const r = x % 3;

        if (r === 1) {
            if (x < r1_min1) {
                r1_min2 = r1_min1;
                r1_min1 = x;
            } else if (x < r1_min2) {
                r1_min2 = x;
            }
        } else if (r === 2) {
            if (x < r2_min1) {
                r2_min2 = r2_min1;
                r2_min1 = x;
            } else if (x < r2_min2) {
                r2_min2 = x;
            }
        }
    }

    const mod = sum % 3;
    if (mod === 0) return sum;

    let removeCost = 1e18;

    if (mod === 1) {
        if (r1_min1 !== INF) removeCost = Math.min(removeCost, r1_min1);
        if (r2_min2 !== INF) removeCost = Math.min(removeCost, r2_min1 + r2_min2);
    } else { // mod === 2
        if (r2_min1 !== INF) removeCost = Math.min(removeCost, r2_min1);
        if (r1_min2 !== INF) removeCost = Math.min(removeCost, r1_min1 + r1_min2);
    }

    if (removeCost >= 1e18) return 0;
    return sum - removeCost;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def maxSumDivThree(self, nums: List[int]) -> int:
        total = 0
        INF = 10**9

        # Two smallest numbers with remainder 1
        r1_min1 = INF
        r1_min2 = INF
        # Two smallest numbers with remainder 2
        r2_min1 = INF
        r2_min2 = INF

        for x in nums:
            total += x
            r = x % 3

            if r == 1:
                if x < r1_min1:
                    r1_min2 = r1_min1
                    r1_min1 = x
                elif x < r1_min2:
                    r1_min2 = x
            elif r == 2:
                if x < r2_min1:
                    r2_min2 = r2_min1
                    r2_min1 = x
                elif x < r2_min2:
                    r2_min2 = x

        mod = total % 3
        if mod == 0:
            return total

        remove_cost = 10**18

        if mod == 1:
            if r1_min1 != INF:
                remove_cost = min(remove_cost, r1_min1)
            if r2_min2 != INF:
                remove_cost = min(remove_cost, r2_min1 + r2_min2)
        else:  # mod == 2
            if r2_min1 != INF:
                remove_cost = min(remove_cost, r2_min1)
            if r1_min2 != INF:
                remove_cost = min(remove_cost, r1_min1 + r1_min2)

        if remove_cost >= 10**18:
            return 0
        return total - remove_cost
```

---

### Go

```go
package main

func maxSumDivThree(nums []int) int {
    sum := 0
    const INF = int(1e9)

    // Two smallest numbers with remainder 1
    r1Min1, r1Min2 := INF, INF
    // Two smallest numbers with remainder 2
    r2Min1, r2Min2 := INF, INF

    for _, x := range nums {
        sum += x
        r := x % 3

        if r == 1 {
            if x < r1Min1 {
                r1Min2 = r1Min1
                r1Min1 = x
            } else if x < r1Min2 {
                r1Min2 = x
            }
        } else if r == 2 {
            if x < r2Min1 {
                r2Min2 = r2Min1
                r2Min1 = x
            } else if x < r2Min2 {
                r2Min2 = x
            }
        }
    }

    mod := sum % 3
    if mod == 0 {
        return sum
    }

    removeCost := int64(1e18)

    if mod == 1 {
        if r1Min1 != INF && int64(r1Min1) < removeCost {
            removeCost = int64(r1Min1)
        }
        if r2Min2 != INF && int64(r2Min1+r2Min2) < removeCost {
            removeCost = int64(r2Min1 + r2Min2)
        }
    } else { // mod == 2
        if r2Min1 != INF && int64(r2Min1) < removeCost {
            removeCost = int64(r2Min1)
        }
        if r1Min2 != INF && int64(r1Min1+r1Min2) < removeCost {
            removeCost = int64(r1Min1 + r1Min2)
        }
    }

    if removeCost >= int64(1e18) {
        return 0
    }
    return sum - int(removeCost)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all languages. I’ll explain the core steps generically:

1. **Initialize total sum and helpers**

   * Set `sum = 0`.
   * Set `INF` to a large value.
   * Initialize:

     * `r1_min1 = r1_min2 = INF`
     * `r2_min1 = r2_min2 = INF`

2. **Loop through each number `x` in `nums`**

   * Add `x` to `sum`.
   * Compute `r = x % 3`.

3. **Update the smallest remainder-1 numbers**

   * If `r == 1`:

     * If `x` is smaller than the current `r1_min1`:

       * Move `r1_min1` to `r1_min2`.
       * Set `r1_min1 = x`.
     * Else if `x` is smaller than `r1_min2`:

       * Set `r1_min2 = x`.

4. **Update the smallest remainder-2 numbers**

   * If `r == 2`:

     * If `x` is smaller than the current `r2_min1`:

       * Move `r2_min1` to `r2_min2`.
       * Set `r2_min1 = x`.
     * Else if `x` is smaller than `r2_min2`:

       * Set `r2_min2 = x`.

5. **Check the total sum’s remainder**

   * After the loop, compute `mod = sum % 3`.
   * If `mod == 0`, return `sum` — we are already perfectly divisible by 3.

6. **Calculate minimal removal cost**

   * Initialize `removeCost` to a very large number (e.g., `1e18`).
   * If `mod == 1`:

     * Option A: remove `r1_min1` if it exists.
     * Option B: remove `r2_min1 + r2_min2` if both exist.
     * `removeCost = min(optionA, optionB)`.
   * If `mod == 2`:

     * Option A: remove `r2_min1` if it exists.
     * Option B: remove `r1_min1 + r1_min2` if both exist.
     * `removeCost = min(optionA, optionB)`.

7. **Handle impossible cases**

   * If `removeCost` is still extremely large:

     * That means there was no valid combination of numbers to fix the remainder.
     * Return `0` (choosing no elements gives sum `0`, which is divisible by 3).

8. **Return final answer**

   * Otherwise, return `sum - removeCost`.

---

## Examples

### Example 1

**Input:**

```text
nums = [3, 6, 5, 1, 8]
```

**Total sum:**

```text
sum = 3 + 6 + 5 + 1 + 8 = 23
23 % 3 = 2
```

So I need to remove:

* either one number with remainder 2,
* or two numbers with remainder 1.

Remainders:

* 3 % 3 = 0
* 6 % 3 = 0
* 5 % 3 = 2
* 1 % 3 = 1
* 8 % 3 = 2

Candidates:

* remainder-1: [1]
* remainder-2: [5, 8]

Options:

* Remove one remainder-2 → min is 5 → new sum = 23 - 5 = 18.
* Remove two remainder-1 → not possible (only one such element).

**Answer:** `18`

---

### Example 2

**Input:**

```text
nums = [4]
```

**Total sum:** `4`, and `4 % 3 = 1`.

We need to remove:

* one remainder-1 (4 itself),
* or two remainder-2 (none exist).

So best we can do is remove 4 → sum becomes 0.

**Answer:** `0`

---

### Example 3

**Input:**

```text
nums = [1, 2, 3, 4, 4]
```

Total sum = `1 + 2 + 3 + 4 + 4 = 14`, and `14 % 3 = 2`.

Remainders:

* 1 → remainder 1
* 2 → remainder 2
* 3 → remainder 0
* 4 → remainder 1
* 4 → remainder 1

Remainder-1 numbers: [1, 4, 4] → two smallest are 1 and 4.
Remainder-2 numbers: [2].

To fix remainder 2:

* Option A: remove one remainder-2 → remove 2 → new sum = 12.
* Option B: remove two remainder-1 → remove 1 + 4 = 5 → new sum = 9.

We choose the **smaller removal** (2), so sum becomes 12.

**Answer:** `12`

---

## How to use / Run locally

### C++

1. Save the solution in a file, e.g., `solution.cpp`.
2. Make sure to include the proper headers and a `main()` if you want to test manually.
3. Compile:

   ```bash
   g++ -std=c++17 solution.cpp -o solution
   ./solution
   ```

### Java

1. Save the class in `Solution.java`.
2. Add a `main` method for local testing if needed.
3. Compile and run:

   ```bash
   javac Solution.java
   java Solution
   ```

### JavaScript (Node.js)

1. Save the function in `solution.js`.
2. Add some test calls at the bottom like:

   ```javascript
   console.log(maxSumDivThree([3,6,5,1,8]));
   ```

3. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save the class in `solution.py`.
2. At the bottom, create an instance and test:

   ```python
   print(Solution().maxSumDivThree([3,6,5,1,8]))
   ```

3. Run:

   ```bash
   python3 solution.py
   ```

### Go

1. Save the code in `main.go` with a `package main` and a `main()` function that calls `maxSumDivThree`.
2. Run:

   ```bash
   go run main.go
   ```

---

## Notes & Optimizations

* This solution is already optimal in terms of:

  * **Time:** single pass → `O(n)`.
  * **Space:** constant extra variables → `O(1)`.
* We avoid dynamic programming over all sums, which would be too slow (`O(n * sum)`).
* The trick is realizing that **only remainders modulo 3 matter**, so we only track **a handful of smallest candidates**.
* This logic can be generalized for similar “maximum sum divisible by k” problems by tracking more remainder groups.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
