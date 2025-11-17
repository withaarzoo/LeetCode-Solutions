# 1437. Check If All 1's Are at Least Length K Places Away

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

Given a binary array `nums` and an integer `k`, return `true` if every two `1`s in the array are at least `k` places (i.e., there are at least `k` zeros) apart from each other. Otherwise, return `false`.

In other words: for every pair of consecutive indices `i` and `j` with `nums[i] == nums[j] == 1` and `i < j`, we must have `j - i - 1 >= k`.

---

## Constraints

* `1 <= nums.length <= 10^5`
* `0 <= k <= nums.length`
* `nums[i]` is either `0` or `1`

These constraints suggest an O(n) one-pass solution with O(1) extra space is ideal.

---

## Intuition

I thought about how to check distances between `1`s. I realized I only need to compare each `1` to the *previous* `1` I saw while scanning left-to-right. If I store the index of the last `1` (`prev`) and at any new `1` check `i - prev - 1`, I can validate the condition. This leads to a single pass over `nums` and constant extra space.

---

## Approach

1. Initialize `prev = -1` to represent "no previous `1` seen yet".
2. Iterate the array indices `i` from `0` to `n-1`.
3. When `nums[i] == 1`:

   * If `prev != -1` then compute `gap = i - prev - 1`.
   * If `gap < k`, then return `false` immediately.
   * Update `prev = i`.
4. If the loop completes without detecting any violation, return `true`.

This approach is simple and efficient: one traversal, constant memory.

---

## Data Structures Used

* Primitive integer variables:

  * `prev` (index of previous `1`)
  * loop index `i`
* No extra arrays, lists, or maps are needed.

---

## Operations & Behavior Summary

* Single linear scan of the array.
* On encountering `1`, compute distance to previous `1` using subtraction.
* Early exit (return `false`) on the first violation.
* If no violation, return `true` at the end.

---

## Complexity

* **Time Complexity:** **O(n)**, where `n = nums.length`.
  We examine each element once in a single loop.

* **Space Complexity:** **O(1)**.
  We only store a few integer variables (`prev`, `i`, etc.), no additional data structures proportional to `n`.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    bool kLengthApart(vector<int>& nums, int k) {
        int prev = -1; // index of previous 1, -1 means none seen yet
        for (int i = 0; i < (int)nums.size(); ++i) {
            if (nums[i] == 1) {
                if (prev != -1) {
                    // zeros between them = i - prev - 1
                    if (i - prev - 1 < k) return false;
                }
                prev = i;
            }
        }
        return true;
    }
};
```

---

### Java

```java
public class Solution {
    public boolean kLengthApart(int[] nums, int k) {
        int prev = -1; // index of last seen 1; -1 means none
        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                if (prev != -1) {
                    // zeros between two 1s = i - prev - 1
                    if (i - prev - 1 < k) return false;
                }
                prev = i;
            }
        }
        return true;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var kLengthApart = function(nums, k) {
    let prev = -1; // index of last 1 seen
    for (let i = 0; i < nums.length; ++i) {
        if (nums[i] === 1) {
            if (prev !== -1) {
                // zeros between current and previous 1 = i - prev - 1
                if (i - prev - 1 < k) return false;
            }
            prev = i;
        }
    }
    return true;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        prev = -1  # index of last seen 1; -1 means none yet
        for i, val in enumerate(nums):
            if val == 1:
                if prev != -1:
                    # zeros between = i - prev - 1
                    if i - prev - 1 < k:
                        return False
                prev = i
        return True
```

---

### Go

```go
package main

func kLengthApart(nums []int, k int) bool {
    prev := -1 // index of last seen 1; -1 means none yet
    for i, v := range nums {
        if v == 1 {
            if prev != -1 {
                // zeros between current and previous 1 = i - prev - 1
                if i - prev - 1 < k {
                    return false
                }
            }
            prev = i
        }
    }
    return true
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm conceptually, then map it to each language's important lines.

**Core Idea:** For every `1` in the array, check distance from the previous `1`. That distance (number of zeros between them) is `current_index - previous_index - 1`. If it is less than `k` for any pair of consecutive `1`s, fail.

### Pseudocode

```
prev = -1
for i from 0 to n-1:
    if nums[i] == 1:
        if prev != -1 and (i - prev - 1) < k:
            return false
        prev = i
return true
```

### Mapping to languages — highlights

**C++**

* `int prev = -1;` — store index of last `1`.
* Loop `for (int i = 0; i < nums.size(); ++i)` to inspect each element.
* `if (i - prev - 1 < k) return false;` — early exit on violation.
* `prev = i;` — update last index.

**Java**

* `int prev = -1;` same purpose.
* Standard `for` loop over `i`.
* Check `if (i - prev - 1 < k)` and `return false`.

**JavaScript**

* `let prev = -1;` and use `for (let i = 0; i < nums.length; ++i)`.
* Strict equality `nums[i] === 1`.
* Update `prev = i`.

**Python3**

* `prev = -1`
* `for i, val in enumerate(nums):`
* `if val == 1: ... if i - prev - 1 < k: return False`
* `prev = i`

**Go**

* `prev := -1`
* `for i, v := range nums { if v == 1 { ... } }`
* `if i - prev - 1 < k { return false }`

Each implementation follows the same arithmetic and control flow; only syntax differs.

---

## Examples

**Example 1**

* Input: `nums = [1,0,0,0,1,0,0,1]`, `k = 2`
* Explanation: Distances between adjacent `1`s are at least 2 zeros.
* Output: `true`

**Example 2**

* Input: `nums = [1,0,0,1,0,1]`, `k = 2`
* Explanation: Second and third `1` are only 1 zero apart (less than `k`).
* Output: `false`

---

## How to use / Run locally

**Prerequisites:** Install appropriate compilers or interpreters for the language you want to run.

### C++

1. Save the solution in `solution.cpp` (wrap in `main` if you want to run custom input).
2. Compile:

   ```bash
   g++ -std=c++17 solution.cpp -O2 -o solution
   ```

3. Run:

   ```bash
   ./solution
   ```

### Java

1. Save class as `Solution.java`.
2. Compile:

   ```bash
   javac Solution.java
   ```

3. Run (if `main` provided):

   ```bash
   java Solution
   ```

### JavaScript (Node.js)

1. Save as `solution.js`.
2. Run:

   ```bash
   node solution.js
   ```

### Python3

1. Save as `solution.py`.
2. Run:

   ```bash
   python3 solution.py
   ```

### Go

1. Save as `solution.go`.
2. Run:

   ```bash
   go run solution.go
   ```

If you want to test with custom inputs, add a small `main` or driver code around the provided function/class to parse and print results.

---

## Notes & Optimizations

* This solution uses a single pass and constant memory — optimal under the constraints.
* Early return stops further work as soon as a violating pair is found.
* Edge cases:

  * If array has 0 or 1 occurrences of `1`, the function trivially returns `true`.
  * `k = 0` means any spacing is allowed — algorithm still correct (always true when encountering `1`s because `i - prev - 1 >= 0`).
* If you need to report which pair violates the rule, you could return the pair of indexes instead of `false`. That would require minor modification.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
