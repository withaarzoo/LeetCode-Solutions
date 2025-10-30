# Minimum Number of Increments on Subarrays to Form a Target Array (LeetCode 1526)

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

You are given an integer array `target`. Initially you have an array `initial` of the same size as `target` with all elements equal to 0. In one operation you can choose any subarray of `initial` and increment every element in that subarray by 1. Return the **minimum number of operations** required to convert `initial` into `target`.

Example:
For `target = [1,2,3,2,1]`, the answer is `3`.

---

## Constraints

* `1 <= target.length <= 10^5`
* `1 <= target[i] <= 10^5`
* The answer fits into a signed 32-bit integer (per problem statement).

---

## Intuition

I thought about how the array grows from all zeros. Every time a position `i` needs to be higher than the previous position `i-1`, I must perform additional operations specifically to increase that position beyond the previous level. Those extra operations equal the **positive difference** `target[i] - target[i-1]`. Also, the very first element `target[0]` always requires `target[0]` operations. So the minimum operations are:

```
target[0] + sum_{i=1..n-1} max(0, target[i] - target[i-1])
```

This counts each needed rise and ignores falls or equal values because decreases or equal values can be achieved by stopping increments — they don't require new operations.

---

## Approach

1. If `target` is empty, return `0`.
2. Initialize `ans = target[0]`.
3. Iterate `i` from `1` to `n-1`:

   * If `target[i] > target[i-1]`, add `target[i] - target[i-1]` to `ans`.
   * Otherwise do nothing.
4. Return `ans`.

This is a single pass algorithm (left-to-right) and uses constant extra memory.

---

## Data Structures Used

* Plain arrays (input) and primitive variables (counters).
  No advanced data structures are required.

---

## Operations & Behavior Summary

* We examine adjacent elements to detect positive increases.
* Each positive increase indicates new operations needed to raise that index (and potentially a suffix) above previous heights.
* We accumulate those increases to get the minimum number of operations.

---

## Complexity

* **Time Complexity:** `O(n)` — we scan the array once, where `n = target.length`.
* **Space Complexity:** `O(1)` — we only use a fixed number of extra variables (no additional arrays).

---

## Multi-language Solutions

### C++

```c++
#include <vector>
using namespace std;

class Solution {
public:
    int minNumberOperations(vector<int>& target) {
        if (target.empty()) return 0;
        long long ans = target[0]; // operations for the first element
        for (int i = 1; i < (int)target.size(); ++i) {
            if (target[i] > target[i - 1]) {
                ans += (target[i] - target[i - 1]);
            }
        }
        return (int)ans; // fits in 32-bit per problem guarantee
    }
};
```

### Java

```java
class Solution {
    public int minNumberOperations(int[] target) {
        if (target == null || target.length == 0) return 0;
        long ans = target[0]; // operations required for index 0
        for (int i = 1; i < target.length; i++) {
            if (target[i] > target[i - 1]) {
                ans += (long)(target[i] - target[i - 1]);
            }
        }
        return (int) ans; // fits in 32-bit per problem guarantee
    }
}
```

### JavaScript

```javascript
/**
 * @param {number[]} target
 * @return {number}
 */
var minNumberOperations = function(target) {
    if (!target || target.length === 0) return 0;
    let ans = target[0];
    for (let i = 1; i < target.length; i++) {
        if (target[i] > target[i-1]) {
            ans += target[i] - target[i-1];
        }
    }
    return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def minNumberOperations(self, target: List[int]) -> int:
        if not target:
            return 0
        ans = target[0]
        for i in range(1, len(target)):
            if target[i] > target[i-1]:
                ans += target[i] - target[i-1]
        return ans
```

### Go

```go
package main

func minNumberOperations(target []int) int {
    if len(target) == 0 {
        return 0
    }
    ans := target[0]
    for i := 1; i < len(target); i++ {
        if target[i] > target[i-1] {
            ans += target[i] - target[i-1]
        }
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

I'll explain as if I'm teaching a friend. I'll use the same logical steps for all languages; the code differences are syntactic.

### Common idea

* I looked at how the `target` array differs from left to right.
* If `target[i]` is greater than `target[i-1]`, then `target[i] - target[i-1]` new operations are required — because these extra increments must include index `i` and cannot be covered by operations that only raised earlier indices.
* For the first element `target[0]`, we need `target[0]` operations from zero.

### Example walkthrough: `target = [1,2,3,2,1]`

* Start `ans = target[0] = 1`.
* i = 1: `2 > 1` → add `1` → `ans = 2`.
* i = 2: `3 > 2` → add `1` → `ans = 3`.
* i = 3: `2 <= 3` → add `0` → `ans = 3`.
* i = 4: `1 <= 2` → add `0` → `ans = 3`.
* Return `3`.

### Line-by-line (Python-style pseudocode mapping to each implementation)

```python
if not target: 
    return 0        # If input empty — no operations needed.

ans = target[0]     # We must raise index 0 from 0 to target[0] with target[0] ops.

for i in range(1, len(target)):
    if target[i] > target[i-1]:
        ans += target[i] - target[i-1]  # Only positive deltas add new ops.
return ans
```

* Each implementation follows the same three steps above. Types and syntax differ by language.

---

## Examples

1. `target = [1,2,3,2,1]` → `3`
   Reason: `1 + (2-1) + (3-2) = 3`.

2. `target = [3,1,1,2]` → `4`
   Reason: `3 + max(0,1-3) + max(0,1-1) + max(0,2-1) = 3 + 0 + 0 + 1 = 4`.

3. `target = [3,1,5,4,2]` → `7`
   Reason: `3 + 0 + (5-1) + 0 + 0 = 3 + 4 = 7`.

---

## How to use / Run locally

### Python

1. Create `solution.py` with the Python class above.
2. Add simple test harness:

```python
if __name__ == "__main__":
    sol = Solution()
    print(sol.minNumberOperations([1,2,3,2,1]))  # prints 3
```

3. Run `python solution.py`.

### JavaScript (Node)

1. Create `solution.js` with the JavaScript function.
2. Add:

```javascript
console.log(minNumberOperations([1,2,3,2,1])); // 3
```

3. Run `node solution.js`.

### C++

1. Put class into a file and write a `main()` to test the method.
2. Compile with `g++ -std=c++17 -O2 -o run solution.cpp`.
3. Run `./run`.

### Java

1. Put `Solution` class into `Solution.java` with a `main` that tests the method.
2. Compile: `javac Solution.java`.
3. Run: `java Solution`.

### Go

1. Place function in `main.go` and call it from `main`.
2. Build: `go run main.go`.

---

## Notes & Optimizations

* This solution is already optimal in time and space for the constraints. It is `O(n)` time and `O(1)` space.
* The reasoning is equivalent to viewing the target as a skyline: every time the skyline height increases from left to right we pay for that rise.
* Use 64-bit accumulator (e.g., `long long` in C++ or `long` in Java) if you worry about intermediate sums — but problem guarantees final answer fits in 32-bit. I used `long`/`long long` to be safe in code examples.

---

## Author

* [Md. Aarzoo Islam](https://bento.me/withaarzoo)
