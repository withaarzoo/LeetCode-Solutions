# Smallest Missing Non-negative Integer After Operations (LeetCode 2598)

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation (for each language)
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given a 0-indexed integer array `nums` and an integer `value`. In one operation I can add or subtract `value` from any element of `nums` (I may do this any number of times on any elements). The MEX (minimum excluded value) of an array is the smallest non-negative integer not present in it. I must return the **maximum MEX** I can obtain after applying the allowed operations any number of times.

Important observation: adding or subtracting `value` changes numbers by multiples of `value`, so each element can only become integers sharing the same remainder modulo `value`. Thus residues modulo `value` are the invariant buckets I must work with.

---

## Constraints

* `1 <= nums.length, value <= 10^5`
* `-10^9 <= nums[i] <= 10^9`

---

## Intuition

I thought: each number can only change within its modulo-`value` residue class. So if I want to ensure `0, 1, 2, ...` up to some `k-1` are all present, for each integer `x` I need one element whose residue equals `x % value`. Therefore I can count how many elements I have for each residue `r` in `[0, value-1]`. Then I greedily attempt to form `x = 0, 1, 2, ...` in order; for each `x` I check if the residue `x % value` has any available elements left. If not, `x` is the MEX. If yes, I consume one and continue. This yields the maximum possible MEX.

---

## Approach

1. Build an array `freq` of length `value` where `freq[r]` counts numbers in `nums` whose remainder (normalized) modulo `value` is `r`.
2. Iterate `x` starting at `0`. For each `x`:

   * Let `r = x % value`.
   * If `freq[r] == 0`: return `x` (can't form `x`).
   * Else `freq[r]--` (use one element to form `x`) and continue to `x+1`.
3. The first `x` that cannot be formed is the answer (maximum MEX).

This is greedy and correct because residues are independent and each target `x` only cares about residue `x % value`.

---

## Data Structures Used

* Fixed-size array (or vector) `freq` of length `value`.
* Simple counters and loop variables.

---

## Operations & Behavior Summary

* Normalizing modulo for negative numbers: `r = ((a % value) + value) % value` (or language equivalent).
* Counting: single pass over `nums` to populate `freq`.
* Greedy consumption: loop `x = 0,1,2,...` using `freq[x % value]` until a bucket is empty.
* The loop stops and returns the first missing `x`.

---

## Complexity

* **Time Complexity:** `O(n + M)` where `n = nums.length` and `M` is the MEX found. In the worst case `M ≤ n + value - 1`, so the worst-case runtime is `O(n + value)`. Practically `O(n)`.
* **Space Complexity:** `O(value)` for the `freq` array.

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    int findSmallestInteger(vector<int>& nums, int value) {
        // freq[r] = count of numbers with remainder r (0 <= r < value)
        vector<int> freq(value, 0);
        for (int a : nums) {
            int r = ((a % value) + value) % value; // normalize negative values
            freq[r]++;
        }
        // Try to form 0,1,2,... greedily using residues
        int x = 0;
        while (true) {
            int r = x % value;
            if (freq[r] == 0) return x; // can't form x
            freq[r]--; // use one element with residue r to form x
            x++;
        }
        return -1; // unreachable
    }
};
```

### Java

```java
class Solution {
    public int findSmallestInteger(int[] nums, int value) {
        int[] freq = new int[value];
        for (int a : nums) {
            int r = a % value;
            if (r < 0) r += value;
            freq[r]++;
        }
        int x = 0;
        while (true) {
            int r = x % value;
            if (freq[r] == 0) return x;
            freq[r]--;
            x++;
        }
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @param {number} value
 * @return {number}
 */
var findSmallestInteger = function(nums, value) {
    const freq = new Array(value).fill(0);
    for (const a of nums) {
        let r = a % value;
        if (r < 0) r += value;
        freq[r]++;
    }
    let x = 0;
    while (true) {
        const r = x % value;
        if (freq[r] === 0) return x;
        freq[r]--;
        x++;
    }
};
```

### Python3

```python3
from typing import List

class Solution:
    def findSmallestInteger(self, nums: List[int], value: int) -> int:
        freq = [0] * value
        for a in nums:
            r = a % value
            if r < 0:
                r += value
            freq[r] += 1

        x = 0
        while True:
            r = x % value
            if freq[r] == 0:
                return x
            freq[r] -= 1
            x += 1
```

### Go

```go
package main

func findSmallestInteger(nums []int, value int) int {
    freq := make([]int, value)
    for _, a := range nums {
        r := a % value
        if r < 0 {
            r += value
        }
        freq[r]++
    }
    x := 0
    for {
        r := x % value
        if freq[r] == 0 {
            return x
        }
        freq[r]--
        x++
    }
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic line-by-line conceptually (the same logic applies in all languages):

### 1) Build residue frequency

* I create an array `freq` of length `value`. Each index represents a residue class `r` where `0 <= r < value`.
* For each number `a` in `nums`, I compute its residue `r = a % value`. Because different languages treat negative modulo differently, I normalize so `r` is in `[0, value-1]` (e.g., `if r < 0 then r += value`).
* I increment `freq[r]`. Now `freq[r]` equals how many elements can be converted into any integer congruent to `r (mod value)`.

### 2) Greedily attempt to produce 0, 1, 2, ... in order

* I start with `x = 0`.
* For each `x`, compute `r = x % value`.
* Check `freq[r]`:

  * If `freq[r] > 0`: I can use one element from residue `r` and shift it (by adding/subtracting multiples of `value`) to exactly match `x`. So I decrement `freq[r]` and increment `x` to attempt the next number.
  * If `freq[r] == 0`: There's no element left that can be converted to `x`. Therefore `x` is the smallest missing non-negative integer (MEX). I return `x`.
* This loop will find the first missing `x` and terminate. The total number of successful consumptions cannot exceed `nums.length`, so the loop is finite.

### Example walkthrough (short)

* `nums = [1, -10, 7, 13, 6, 8]`, `value = 5`.
* Normalize residues: 1->1, -10->0, 7->2, 13->3, 6->1, 8->3
* freq = [1,2,1,2,0]
* x=0: r=0 -> freq[0]>0 -> freq[0]=0 -> x=1
* x=1: r=1 -> freq[1]>0 -> freq[1]=1 -> x=2
* x=2: r=2 -> freq[2]>0 -> freq[2]=0 -> x=3
* x=3: r=3 -> freq[3]>0 -> freq[3]=1 -> x=4
* x=4: r=4 -> freq[4]==0 -> return 4 (MEX)

---

## Examples

**Example 1**

```
Input: nums = [1,-10,7,13,6,8], value = 5
Output: 4
Explanation: After operations we can ensure 0,1,2,3 exist but 4 cannot be formed.
```

**Example 2**

```
Input: nums = [1,-10,7,13,6,8], value = 7
Output: 2
Explanation: Frequencies result in inability to form 2.
```

---

## How to use / Run locally

1. Clone your repository or create a file (e.g., `solution.py`, `Solution.java`, `solution.cpp`, `solution.js`, `solution.go`).
2. Paste the corresponding language code into the file.
3. Add a small driver / main to run sample tests (languages vary):

   * **Python**: instantiate `Solution()` and call `findSmallestInteger`.
   * **C++**: add `main()` and create `Solution` object.
   * **Java**: add a `public static void main` and call the method.
   * **JS (node)**: call `findSmallestInteger` and `console.log` the result.
   * **Go**: add `main()` and `fmt.Println(findSmallestInteger(...))`.
4. Run:

   * Python: `python3 solution.py`
   * C++: `g++ -std=c++17 solution.cpp && ./a.out`
   * Java: `javac Solution.java && java Solution`
   * JavaScript: `node solution.js`
   * Go: `go run solution.go`

---

## Notes & Optimizations

* `freq` uses `O(value)` space; if `value` is large but `n` is much smaller, you could use a hash map keyed by residues instead of a full-size array to save memory. However, array is typically faster and more memory-predictable given `value <= 1e5`.
* The greedy method is optimal because residues are independent and every target `x` requires exactly one element from `x % value`.
* This approach avoids simulated operations and works with counts only — very efficient.

---

## Author

* Created by me — [Md. Aarzoo Islam](https://bento.me/withaarzoo)
* Problem: LeetCode 2598 — "Smallest Missing Non-negative Integer After Operations"