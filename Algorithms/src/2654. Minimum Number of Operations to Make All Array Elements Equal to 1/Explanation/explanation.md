# Minimum Number of Operations to Make All Array Elements Equal to 1

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

You are given a 0-indexed array `nums` consisting of positive integers. You can perform the following operation any number of times:

* Select an index `i` such that `0 <= i < n - 1` and replace either `nums[i]` **or** `nums[i+1]` with `gcd(nums[i], nums[i+1])`.

Return the **minimum number of operations** needed to make all elements of `nums` equal to `1`. If it is impossible, return `-1`.

Key reasoning facts:

* `gcd(a, b)` never increases the values and always divides its operands.
* If there is at least one `1` in the array, we can spread it to neighbors with one operation per element.
* If array-wide gcd > 1, it's impossible to ever get `1`.

---

## Constraints

* `2 <= nums.length <= 50`
* `1 <= nums[i] <= 10^6`

---

## Intuition

I thought about what the operation does: replacing a number with the gcd of two adjacent numbers. Since gcd never increases and divides the operands:

* If I already have some `1`s, each other element can be converted to `1` in one operation by pairing with a neighbor `1`. So if `c1` is count of ones, answer is `n - c1`.
* If there are no `1`s, I must first create one. Creating a `1` requires finding a subarray whose gcd is `1`. Converting that subarray of length `L` into a single `1` takes `L - 1` operations. After that I still need `n - 1` operations to spread that `1` across all elements.
* If the gcd of the whole array is > 1, it's impossible to get `1` at all.

---

## Approach

1. Count how many elements equal `1`. If `c1 > 0`, return `n - c1`.
2. Compute gcd of the entire array. If gcd > 1, return `-1`.
3. Otherwise, find the shortest subarray with gcd `1`:

   * For every start index `i`, compute running gcd with elements `j = i..n-1`; stop when gcd becomes `1`. Track the smallest length `L`.
4. Answer = `(L - 1) + (n - 1)`:

   * `L - 1` to reduce that subarray to a single `1`.
   * `n - 1` to spread the created `1` across the array.

I used an O(n^2) scan over subarrays and gcd updates with early breaks.

---

## Data Structures Used

* Primitive arrays (or vectors / slices) for input.
* O(1) auxiliary variables (counters, running gcd, best length).

---

## Operations & Behavior Summary

* `gcd(a, b)`: compute greatest common divisor of two integers (Euclidean algorithm).
* Count ones: linear scan.
* Global gcd: linear scan accumulating gcd.
* Shortest subarray with gcd 1: nested loops with early break when gcd becomes 1.

---

## Complexity

* **Time Complexity:** `O(n^2 * log A)` where:

  * `n` = number of elements (`<= 50`).
  * `A` = maximum element value (`<= 10^6`) and `log A` is the cost of each gcd operation. Practically for constraints this is fast.
* **Space Complexity:** `O(1)` additional space (ignoring input).

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        int ones = 0;
        for (int x : nums) if (x == 1) ones++;
        if (ones > 0) return n - ones;

        int g = 0;
        for (int x : nums) g = std::gcd(g, x);
        if (g > 1) return -1;

        int best = INT_MAX;
        for (int i = 0; i < n; ++i) {
            int cur = 0;
            for (int j = i; j < n; ++j) {
                cur = std::gcd(cur, nums[j]);
                if (cur == 1) {
                    best = min(best, j - i + 1);
                    break;
                }
            }
        }
        return (best - 1) + (n - 1);
    }
};

// Example of usage:
// int main() { vector<int> a = {2,6,3,4}; Solution s; cout << s.minOperations(a) << endl; }
```

---

### Java

```java
import java.util.*;

class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int ones = 0;
        for (int x : nums) if (x == 1) ones++;
        if (ones > 0) return n - ones;

        int g = 0;
        for (int x : nums) g = gcd(g, x);
        if (g > 1) return -1;

        int best = Integer.MAX_VALUE;
        for (int i = 0; i < n; i++) {
            int cur = 0;
            for (int j = i; j < n; j++) {
                cur = gcd(cur, nums[j]);
                if (cur == 1) {
                    best = Math.min(best, j - i + 1);
                    break;
                }
            }
        }
        return (best - 1) + (n - 1);
    }

    private int gcd(int a, int b) {
        a = Math.abs(a);
        b = Math.abs(b);
        while (b != 0) {
            int t = a % b;
            a = b;
            b = t;
        }
        return a;
    }
}

// Example usage:
// public static void main(String[] args){ Solution s = new Solution(); int[] a = {2,6,3,4}; System.out.println(s.minOperations(a)); }
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {number}
 */
var minOperations = function(nums) {
  const n = nums.length;
  let ones = 0;
  for (const x of nums) if (x === 1) ones++;
  if (ones > 0) return n - ones;

  const gcd = (a, b) => {
    a = Math.abs(a); b = Math.abs(b);
    while (b !== 0) {
      const t = a % b;
      a = b;
      b = t;
    }
    return a;
  };

  let g = 0;
  for (const x of nums) g = gcd(g, x);
  if (g > 1) return -1;

  let best = Infinity;
  for (let i = 0; i < n; i++) {
    let cur = 0;
    for (let j = i; j < n; j++) {
      cur = gcd(cur, nums[j]);
      if (cur === 1) {
        best = Math.min(best, j - i + 1);
        break;
      }
    }
  }
  return (best - 1) + (n - 1);
};

// Example:
// console.log(minOperations([2,6,3,4])); // 4
```

---

### Python3

```python
from math import gcd
from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        n = len(nums)
        ones = sum(1 for x in nums if x == 1)
        if ones > 0:
            return n - ones

        g = 0
        for x in nums:
            g = gcd(g, x)
        if g > 1:
            return -1

        best = 10**9
        for i in range(n):
            cur = 0
            for j in range(i, n):
                cur = gcd(cur, nums[j])
                if cur == 1:
                    best = min(best, j - i + 1)
                    break
        return (best - 1) + (n - 1)

# Example:
# s = Solution(); print(s.minOperations([2,6,3,4]))  # 4
```

---

### Go

```go
package main

import (
 "fmt"
 "math"
)

func gcd(a, b int) int {
 a = int(math.Abs(float64(a)))
 b = int(math.Abs(float64(b)))
 for b != 0 {
  a, b = b, a%b
 }
 return a
}

func minOperations(nums []int) int {
 n := len(nums)
 ones := 0
 for _, x := range nums {
  if x == 1 { ones++ }
 }
 if ones > 0 { return n - ones }

 g := 0
 for _, x := range nums { g = gcd(g, x) }
 if g > 1 { return -1 }

 best := int(1e9)
 for i := 0; i < n; i++ {
  cur := 0
  for j := i; j < n; j++ {
   cur = gcd(cur, nums[j])
   if cur == 1 {
    if j-i+1 < best { best = j-i+1 }
    break
   }
  }
 }
 return (best - 1) + (n - 1)
}

func main() {
 fmt.Println(minOperations([]int{2,6,3,4})) // expected 4
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the algorithm line-by-line in a general sense that maps to each language implementation.

1. **Count existing 1s**

   * I scan through the array and count elements equal to `1`.
   * If count `c1 > 0`, every other element can be turned into `1` by pairing with an adjacent `1`, costing `n - c1` operations total.
   * This is an O(n) step.

2. **Global gcd impossibility check**

   * I compute `g = gcd(nums[0], nums[1], ..., nums[n-1])`.
   * If `g > 1`, then every number shares a common divisor greater than 1, so no sequence of pairwise gcd operations can ever produce `1`. Return `-1`.
   * This is O(n * log A) for gcd accumulation.

3. **Find shortest subarray with gcd == 1**

   * For each starting index `i` from `0` to `n-1`, I keep a running gcd `cur = gcd(cur, nums[j])` as `j` moves right from `i`.
   * If `cur` becomes `1`, I record the subarray length `L = j - i + 1`, update `best = min(best, L)`, and break the inner loop (no need to extend further for this `i` because length increases).
   * This nested loop is O(n^2) in the worst case but `n <= 50`, so it's acceptable.

4. **Combine costs**

   * To make that subarray (length `best`) collapsed into `1` takes `best - 1` operations.
   * After creating one `1`, spreading it to the entire array takes `n - 1` operations (or simply `n - best` additional operations to convert the rest, but it’s commonly written as `n - 1`).
   * So final answer is `(best - 1) + (n - 1)`.

**Important implementation notes**

* Use Euclidean algorithm for gcd for speed and correctness.
* Early breaks when `cur == 1` improve constant factors.
* If a code returns a very large `best` (never updated), that case should be impossible after the global gcd check, but keep guards if you like.

---

## Examples

1. Input: `[2, 6, 3, 4]`

   * There are no `1`s. Global gcd = 1. Shortest subarray with gcd 1 is length 2 or 3 depending on scan; result = 4.
   * Output: `4`

2. Input: `[2, 10, 6, 14]`

   * Global gcd is 2 (>1), impossible to make any `1`.
   * Output: `-1`

3. Input: `[1, 2, 3]`

   * One `1` exists. Convert the two other elements in `n - c1 = 3 - 1 = 2` operations.
   * Output: `2`

---

## How to use / Run locally

### Python

1. Save `solution.py` with the Python code above (wrap call in `if __name__ == "__main__":` to test).
2. Run:

```bash
python3 solution.py
```

### C++

1. Save the C++ code as `solution.cpp`. Add a `main()` that tests the method if needed.
2. Compile & run:

```bash
g++ -std=c++17 solution.cpp -O2 -o solution
./solution
```

### Java

1. Save the Java code in `Solution.java` and add a `main` method if you want to run locally.
2. Compile & run:

```bash
javac Solution.java
java Solution
```

### JavaScript (Node.js)

1. Save code to `solution.js`.
2. Run:

```bash
node solution.js
```

### Go

1. Save code to `main.go`.
2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The algorithm is simple and robust for constraints (`n <= 50`). For much larger `n` you’d need more advanced methods (e.g., segment tree of gcds and binary search to find minimum subarray gcd = 1 in O(n log n log A)).
* Early exit: if `ones > 0` we return immediately — this is the most common fast path.
* Global gcd check quickly eliminates impossible cases.
* Avoid recalculating gcd from scratch inside the inner loop by using the running gcd value.
* The current approach is easy to read and less error-prone — tradeoff: O(n^2) loops, but acceptable given constraints.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
