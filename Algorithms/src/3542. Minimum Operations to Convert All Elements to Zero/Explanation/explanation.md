# Minimum Operations to Convert All Elements to Zero (LeetCode 3542)

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

You are given an array `nums` of non-negative integers.
In one operation, you can choose a subarray `[i, j]` and set **all occurrences of the minimum non-negative integer inside that subarray** to `0`.
Return the **minimum number of operations** needed to make the whole array zeros.

---

## Constraints

* `1 ≤ n = nums.length ≤ 1e5`
* `0 ≤ nums[i] ≤ 1e5`

---

## Intuition

I asked: “When do I need a **new** operation as I move from left to right?”
If the current number is **lower** than what I was handling before, I can drop down (no new op).
If it is **equal** to the last height, still no new op.
If it is **higher** than the last height seen so far, I must start a **new layer** — that is exactly **one more operation**.

This screams **monotonic stack**: keep a non-decreasing stack of heights.

* Pop while the top is greater than the current number (we dropped).
* Ignore zeros.
* If the current number is greater than the top, push it and add one to the answer.

Each push represents starting a new “height segment” that needs exactly one operation.

---

## Approach

1. Initialize an empty stack `stk` and `ans = 0`.
2. For each number `x` in `nums`:

   * While `stk` not empty and `stk.top > x`, pop (we fell to a lower height).
   * If `x == 0`, continue (zeros don’t start an operation).
   * If `stk` is empty or `stk.top < x`, we just climbed to a higher height → `ans++` and push `x`.
   * If `stk.top == x`, do nothing.
3. Return `ans`.

This is linear because every element is pushed and popped at most once.

---

## Data Structures Used

* **Monotonic (non-decreasing) stack** of integers.

---

## Operations & Behavior Summary

* **Pop while decreasing:** close finished higher segments.
* **Skip zeros:** no new layer begins at height 0.
* **Push on strict increase:** start a new layer → +1 operation.
* **Do nothing on equal:** same active layer continues.

---

## Complexity

* **Time Complexity:** `O(n)`
  `n` = length of `nums`. Each value is pushed/popped at most once.
* **Space Complexity:** `O(n)` in the worst case (strictly increasing array), due to the stack.

---

## Multi-language Solutions

### C++

```c++
class Solution {
public:
    int minOperations(vector<int>& nums) {
        vector<int> stk;   // non-decreasing stack of heights
        int ans = 0;
        for (int x : nums) {
            while (!stk.empty() && stk.back() > x) stk.pop_back(); // drop down
            if (x == 0) continue;                                  // ignore zeros
            if (stk.empty() || stk.back() < x) {                   // new rise
                ans++;
                stk.push_back(x);
            }
        }
        return ans;
    }
};
```

### Java

```java
import java.util.*;

class Solution {
    public int minOperations(int[] nums) {
        Deque<Integer> stk = new ArrayDeque<>(); // non-decreasing stack
        int ans = 0;
        for (int x : nums) {
            while (!stk.isEmpty() && stk.peekLast() > x) stk.pollLast();
            if (x == 0) continue;
            if (stk.isEmpty() || stk.peekLast() < x) {
                ans++;
                stk.addLast(x);
            }
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
var minOperations = function(nums) {
  const stk = []; // non-decreasing
  let ans = 0;
  for (const x of nums) {
    while (stk.length && stk[stk.length - 1] > x) stk.pop();
    if (x === 0) continue;
    if (!stk.length || stk[stk.length - 1] < x) {
      ans++;
      stk.push(x);
    }
  }
  return ans;
};
```

### Python3

```python
from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        stk = []  # non-decreasing stack
        ans = 0
        for x in nums:
            while stk and stk[-1] > x:
                stk.pop()
            if x == 0:
                continue
            if not stk or stk[-1] < x:
                ans += 1
                stk.append(x)
        return ans
```

### Go

```go
package main

func minOperations(nums []int) int {
 ans := 0
 stk := make([]int, 0) // non-decreasing stack
 for _, x := range nums {
  for len(stk) > 0 && stk[len(stk)-1] > x {
   stk = stk[:len(stk)-1]
  }
  if x == 0 {
   continue
  }
  if len(stk) == 0 || stk[len(stk)-1] < x {
   ans++
   stk = append(stk, x)
  }
 }
 return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. **Create a stack** (vector/Deque/array/slice/list) to keep heights in **non-decreasing** order.
2. **Loop over `nums`:**

   * While the stack’s top is **greater than** `x`, `pop()` it.
     This models finishing those taller segments because we went down to a smaller number.
   * If `x` is **zero**, continue. Zeros never start a new layer/operation.
   * If the stack is empty **or** the top is **less than** `x`, it means we are climbing to a higher level that wasn’t active before.
     → **Increment** the answer and **push** `x` (start this new layer).
   * If the top equals `x`, do nothing — we are already running that layer.
3. After the loop, **the sum of all started layers** (`ans`) is exactly the minimal number of operations.

---

## Examples

* `nums = [0, 2]`

  * stack: push 2 → `ans = 1` → **Answer = 1**

* `nums = [3, 1, 2, 1]`

  * 3: push → `ans = 1`
  * 1: pop 3, push 1 → `ans = 2`
  * 2: push 2 → `ans = 3`
  * 1: pop 2 (since 1 < 2), top == 1 → no new op
    → **Answer = 3**

* `nums = [1,2,1,2,1,2]`

  * pushes at 1 (1), 2 (2), then repeat downs/equals → pushes at each fresh rise back to 2
    → **Answer = 4** (matches editorial example)

---

## How to use / Run locally

1. Copy the solution in your preferred language.
2. Compile/Run with the usual toolchain:

   * **C++:** `g++ -std=c++17 solution.cpp && ./a.out`
   * **Java:** `javac Solution.java && java Solution`
   * **JavaScript (Node):** `node solution.js`
   * **Python3:** `python3 solution.py`
   * **Go:** `go run solution.go`
3. Provide input via your own driver or test with hard-coded arrays.

---

## Notes & Optimizations

* The monotonic stack guarantees **linear time**.
* We never simulate the operations; we only count how many **new rises** we start.
* Memory can be reduced using an integer vector/list; no need for complex structures.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
