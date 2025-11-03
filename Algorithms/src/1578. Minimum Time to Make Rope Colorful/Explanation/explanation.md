# Minimum Time to Make Rope Colorful

## Table of Contents

* Problem Title

* ## Table of Contents

* ## Problem Summary

* ## Constraints

* ## Intuition

* ## Approach

* ## Data Structures Used

* ## Operations & Behavior Summary

* ## Complexity

* ## Multi-language Solutions

  * ### C++

  * ### Java

  * ### JavaScript

  * ### Python3

  * ### Go

* ## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

* ## Examples

* ## How to use / Run locally

* ## Notes & Optimizations

* ## Author

---

## Problem Summary

I am given a string `colors` where `colors[i]` is the color of the `i`-th balloon. I also have an integer array `neededTime` where `neededTime[i]` is the time required to remove the `i`-th balloon. I want no two adjacent balloons to have the same color. I can remove any balloons; removing balloon `i` costs `neededTime[i]`. I must return the **minimum total time** needed to make the rope "colorful" (i.e., no two adjacent balloons share the same color).

---

## Constraints

* `n == colors.length == neededTime.length`
* `1 <= n <= 10^5`
* `1 <= neededTime[i] <= 10^4`
* `colors` contains only lowercase English letters

---

## Intuition

I thought about contiguous groups of identical colors. For any contiguous block of the same color, to avoid having neighboring equal colors I must remove all but one balloon. To minimize the cost in that block, I should keep the balloon that costs the most to remove (so I *don’t* pay that high cost) and remove the rest. So the minimal cost for a block = (sum of times in block) − (maximum time in block). Summing this value for every block gives the global minimum.

---

## Approach

1. Scan `colors` left to right in one pass.
2. Maintain `block_sum` (sum of `neededTime` in the current same-color block) and `block_max` (maximum `neededTime` in the block).
3. When a color change occurs (or at end), add `block_sum - block_max` to result and reset block counters.
4. Continue until the end. The total is the minimal time.

This uses constant extra space and linear time.

---

## Data Structures Used

* Primitive counters and accumulators:

  * `ans` — accumulates the total minimum time
  * `block_sum` — sum of times in the current block
  * `block_max` — max time in the current block

No additional arrays or complex data structures are required.

---

## Operations & Behavior Summary

* Single pass iteration over `colors` (O(n)).
* For each character:

  * If it continues the block, update `block_sum` and `block_max`.
  * If it starts a new block, finalize the previous block by adding `block_sum - block_max` to `ans` and reset counters.
* Handle the last block after the loop.

---

## Complexity

* **Time Complexity:** O(n), where `n` = `colors.length`. We touch each balloon once.
* **Space Complexity:** O(1) extra space (only a few integer variables used).

---

## Multi-language Solutions

### C++

```c++
#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
    int minCost(string colors, vector<int>& neededTime) {
        long long ans = 0;               // total minimal time
        long long block_sum = 0;         // sum of times in current same-color block
        int block_max = 0;               // maximum time in current block
        int n = colors.size();
        
        for (int i = 0; i < n; ++i) {
            if (i > 0 && colors[i] != colors[i-1]) {
                // end of previous block -> remove all but the most expensive
                ans += block_sum - block_max;
                block_sum = 0;
                block_max = 0;
            }
            block_sum += neededTime[i];
            block_max = max(block_max, neededTime[i]);
        }
        ans += block_sum - block_max; // finalize last block
        return (int)ans;
    }
};
```

### Java

```java
class Solution {
    public int minCost(String colors, int[] neededTime) {
        long ans = 0L;          // accumulate answer here (long for safety)
        long blockSum = 0L;     // sum of times in current block
        int blockMax = 0;       // maximum time in current block
        int n = colors.length();
        
        for (int i = 0; i < n; ++i) {
            if (i > 0 && colors.charAt(i) != colors.charAt(i - 1)) {
                ans += blockSum - blockMax;
                blockSum = 0;
                blockMax = 0;
            }
            blockSum += neededTime[i];
            blockMax = Math.max(blockMax, neededTime[i]);
        }
        ans += blockSum - blockMax;
        return (int) ans;
    }
}
```

### JavaScript

```javascript
/**
 * @param {string} colors
 * @param {number[]} neededTime
 * @return {number}
 */
var minCost = function(colors, neededTime) {
    let ans = 0;
    let blockSum = 0;
    let blockMax = 0;
    const n = colors.length;
    
    for (let i = 0; i < n; ++i) {
        if (i > 0 && colors[i] !== colors[i-1]) {
            ans += blockSum - blockMax;
            blockSum = 0;
            blockMax = 0;
        }
        blockSum += neededTime[i];
        blockMax = Math.max(blockMax, neededTime[i]);
    }
    ans += blockSum - blockMax;
    return ans;
};
```

### Python3

```python3
from typing import List

class Solution:
    def minCost(self, colors: str, neededTime: List[int]) -> int:
        ans = 0                # total minimal removal time
        block_sum = 0          # sum of times in the current same-color block
        block_max = 0          # max time in the current block
        n = len(colors)
        
        for i in range(n):
            # if color changes, finalize previous block
            if i > 0 and colors[i] != colors[i-1]:
                ans += block_sum - block_max
                block_sum = 0
                block_max = 0
            block_sum += neededTime[i]
            block_max = max(block_max, neededTime[i])
        
        # add remaining block
        ans += block_sum - block_max
        return ans
```

### Go

```go
package main

// minCost returns the minimal time to make the rope colorful.
// It uses O(n) time and O(1) extra space.
func minCost(colors string, neededTime []int) int {
    var ans int64 = 0
    var blockSum int64 = 0
    var blockMax int64 = 0
    n := len(colors)
    
    for i := 0; i < n; i++ {
        if i > 0 && colors[i] != colors[i-1] {
            ans += blockSum - blockMax
            blockSum = 0
            blockMax = 0
        }
        blockSum += int64(neededTime[i])
        if int64(neededTime[i]) > blockMax {
            blockMax = int64(neededTime[i])
        }
    }
    ans += blockSum - blockMax
    return int(ans)
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain the logic and map it to code. The same sequence applies to each language.

1. **Initialize accumulators**

   * `ans = 0` — total minimum time I will return.
   * `block_sum = 0` — sum of `neededTime` for the current run of equal colors.
   * `block_max = 0` — maximum `neededTime` in the current run.

   In code: e.g. `long long ans = 0; long long block_sum = 0; int block_max = 0;`

2. **Iterate through balloons**

   * For each index `i` from `0` to `n-1`:

     * If `i > 0` and `colors[i] != colors[i-1]`, that means the contiguous block of the previous color ended.

       * Finalize that block by adding `block_sum - block_max` to `ans`. Why? Because I keep the balloon that costs the most to remove (so I don't pay that cost) and remove the rest. Removing all others costs `sum - max`.
       * Reset `block_sum = 0` and `block_max = 0` to start accumulating the new block.
     * Add `neededTime[i]` to `block_sum`.
     * Update `block_max = max(block_max, neededTime[i])`.

   In code (Python example):

   ```python
   for i in range(n):
       if i > 0 and colors[i] != colors[i-1]:
           ans += block_sum - block_max
           block_sum = 0
           block_max = 0
       block_sum += neededTime[i]
       block_max = max(block_max, neededTime[i])
   ```

3. **Finalize last block**

   * After the loop, the final group hasn't been added to `ans`. Add `block_sum - block_max`.

   In code: `ans += block_sum - block_max`

4. **Return `ans`**

   * This is the minimum total time.

---

## Examples

1. Example:

   * Input: `colors = "aabcc"`, `neededTime = [1,2,3,4,5]`
   * Blocks: `"aa"`, `"b"`, `"cc"`

     * `"aa"` → sum=1+2=3, max=2 → cost = 3-2 = 1
     * `"b"` → cost = 0 (single balloon)
     * `"cc"` → sum=4+5=9, max=5 → cost = 9-5 = 4
   * Total = 1 + 0 + 4 = 5

2. Example from prompt:

   * Input: `colors = "abaac"`, `neededTime = [1,2,3,4,5]`
   * Minimal total time = 3 (remove one balloon in the "aa" block using cost 3 in that test configuration).

3. Example:

   * Input: `colors = "abc"`, `neededTime = [1,2,3]`
   * Already colorful; cost = 0.

---

## How to use / Run locally

### C++

* Create `solution.cpp` with the C++ code inside a `main` wrapper or use the `Solution` class in an online judge.
* Compile:

  ```bash
  g++ -std=c++17 solution.cpp -O2 -o solution
  ./solution
  ```

### Java

* Save as `Solution.java`.
* Compile and run:

  ```bash
  javac Solution.java
  java Solution
  ```

### JavaScript (Node.js)

* Save the function in a file `solution.js` and add a small wrapper to read input and call `minCost`.
* Run:

  ```bash
  node solution.js
  ```

### Python3

* Save as `solution.py` and include a small test harness:

  ```bash
  python3 solution.py
  ```

### Go

* Save the function in a `main.go` file with a `main()` that calls the function and prints the result.
* Build and run:

  ```bash
  go run main.go
  ```

(Online judges like LeetCode will accept the function/class format directly; for local testing, add a main or wrapper.)

---

## Notes & Optimizations

* This is optimal: single pass, constant extra space.
* Use `long long` (C++, Go int64, Java long) if you're paranoid about overflow when summing many `neededTime` values — although given constraints `n <= 1e5` and `neededTime[i] <= 1e4`, `sum <= 1e9` fits in 32-bit signed int, but using 64-bit is safe.
* I keep resetting accumulators only when a color changes. That ensures we consider each contiguous block exactly once.
* Edge cases:

  * All different colors → ans = 0.
  * All same color → ans = sum(all) − max(all).

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
