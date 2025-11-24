# 1018. Binary Prefix Divisible By 5

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

I’m given an array `nums` of binary digits (only `0` or `1`).
For every prefix of this array (from index `0` to `i`), I need to check whether the number formed by that prefix (interpreted as a **binary number**) is divisible by `5`.

I must return an array of booleans `answer` where:

* `answer[i]` is `true` if the binary number `nums[0..i]` is divisible by 5
* otherwise `answer[i]` is `false`.

---

## Constraints

* `1 <= nums.length <= 10^5`
* `nums[i]` is either `0` or `1`

So my solution must:

* Work in **linear time** (O(n))
* Avoid building huge integers directly (to prevent overflow and inefficiency)

---

## Intuition

At first, I thought of directly converting each prefix into a decimal number and checking `num % 5 == 0`.
But this is a bad idea:

* The number grows very fast with each bit.
* Recomputing the whole value for every prefix is too slow.
* It can overflow in some languages.

Then I remembered something from number theory:

> To know if a number is divisible by 5, I only need the **remainder modulo 5**, not the full number.

If I read the binary number from left to right:

* Let the current value be `X`.
* When I add a new bit `b` at the end, the new value becomes:
  `X' = 2 * X + b` (this is how binary works: shift left then add the bit)

So I realized I can track only the current remainder when divided by 5 and update it like this:

[
\text{rem}' = (2 * \text{rem} + b) \mod 5
]

If `rem` becomes `0` at some position, then that prefix is divisible by 5.

---

## Approach

1. Initialize:

   * `rem = 0` → this will store the remainder of the current prefix modulo `5`.
   * an empty result list `ans`.

2. Loop over each bit `b` in `nums` from left to right:

   * Update the remainder using:
     [
     rem = (rem * 2 + b) \mod 5
     ]
   * If `rem == 0`, push `true` into `ans`, else push `false`.

3. Return the result list `ans`.

I never need the full number; I only carry its remainder modulo 5.
This keeps the solution fast and memory-efficient.

---

## Data Structures Used

* **Main data structure for output:**

  * C++: `vector<bool>`
  * Java: `List<Boolean>` (usually `ArrayList<>`)
  * JavaScript: `boolean[]` (actually `Array<boolean>`)
  * Python3: `List[bool]`
  * Go: `[]bool`

* **Variables:**

  * A single integer `rem` to store the remainder modulo 5.

There are no complex data structures required.
Everything is done with simple arrays/lists and a few integer variables.

---

## Operations & Behavior Summary

For each element in `nums`:

1. **Binary shift + add bit:**

   * Conceptually: `X = X * 2 + bit`
   * But we do this **only on the remainder**: `rem = (rem * 2 + bit) % 5`.

2. **Check divisibility:**

   * If `rem == 0`, then the current prefix is divisible by 5.
   * Record this as `true` in the answer; otherwise `false`.

This process is repeated once for each bit, so the behavior is stable and predictable.

---

## Complexity

Let `n = nums.length`.

* **Time Complexity:**

  * `O(n)`
  * I go through the array once, and each step does constant work.

* **Space Complexity:**

  * Extra space: `O(1)` (ignoring the output array)
  * I only store:

    * `rem` (an integer)
    * and the output list/array which is required by the problem.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<bool> prefixesDivBy5(vector<int>& nums) {
        vector<bool> ans;
        ans.reserve(nums.size());              // small optimization to avoid reallocations
        
        int rem = 0;                           // remainder of current prefix modulo 5
        
        for (int bit : nums) {
            // Update remainder: newVal = oldVal * 2 + bit, then take modulo 5
            rem = (rem * 2 + bit) % 5;
            
            // If remainder is 0, this prefix is divisible by 5
            ans.push_back(rem == 0);
        }
        
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    public List<Boolean> prefixesDivBy5(int[] nums) {
        List<Boolean> ans = new ArrayList<>(nums.length);
        
        int rem = 0;                            // remainder of current prefix modulo 5
        
        for (int bit : nums) {
            // newValue = oldValue * 2 + bit (binary logic), keep only remainder
            rem = (rem * 2 + bit) % 5;
            
            // remainder 0 means divisible by 5
            ans.add(rem == 0);
        }
        
        return ans;
    }
}
```

---

### JavaScript

```javascript
/**
 * @param {number[]} nums
 * @return {boolean[]}
 */
var prefixesDivBy5 = function(nums) {
    const ans = [];
    let rem = 0;                      // remainder of current prefix modulo 5
    
    for (const bit of nums) {
        // Binary shift: value = value * 2 + bit, then take modulo 5
        rem = (rem * 2 + bit) % 5;
        
        ans.push(rem === 0);
    }
    
    return ans;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def prefixesDivBy5(self, nums: List[int]) -> List[bool]:
        ans: List[bool] = []
        rem = 0                          # remainder of current prefix modulo 5
        
        for bit in nums:
            # Shift left (multiply by 2) and add bit, then mod 5
            rem = (rem * 2 + bit) % 5
            ans.append(rem == 0)
        
        return ans
```

---

### Go

```go
package main

func prefixesDivBy5(nums []int) []bool {
    ans := make([]bool, len(nums))
    rem := 0 // remainder of current prefix modulo 5

    for i, bit := range nums {
        // Update remainder with binary logic: value = value*2 + bit
        rem = (rem*2 + bit) % 5
        ans[i] = (rem == 0)
    }

    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic in all languages is the same; only syntax changes.

### 1. Initialize remainder and answer list

```text
rem = 0
ans = []
```

* `rem` starts at `0` because before reading any bits, the "number" is 0.
* `ans` will store the boolean result for each prefix.

### 2. Loop through each bit in `nums`

Pseudo-code:

```text
for each bit in nums:
    rem = (rem * 2 + bit) % 5
    ans.append(rem == 0)
```

**Explanation of the update:**

* Assume the current prefix (before adding this new bit) represents some number `X`.
* Adding a new bit in binary is like shifting left and adding the bit:

  * `X' = 2 * X + bit`
* Instead of storing `X` directly, I only store `rem = X % 5`.
* Using modular arithmetic:

  * `(2 * X + bit) % 5 == (2 * (X % 5) + bit) % 5`
  * So I can update:

    * `rem = (rem * 2 + bit) % 5`

This keeps `rem` always between `0` and `4` and avoids big numbers.

### 3. Check if divisible by 5

* If `rem == 0`, the current prefix is divisible by 5.
* Save that as `true`; otherwise, save `false`.

Language-wise:

* **C++**: `ans.push_back(rem == 0);`
* **Java**: `ans.add(rem == 0);`
* **JavaScript**: `ans.push(rem === 0);`
* **Python3**: `ans.append(rem == 0)`
* **Go**: `ans[i] = (rem == 0)`

### 4. Return the result

At the end of the loop, I return `ans` (list/array of booleans):

* **C++**: `return ans;`
* **Java**: `return ans;`
* **JavaScript**: `return ans;`
* **Python3**: `return ans`
* **Go**: `return ans`

---

## Examples

### Example 1

**Input:**

```text
nums = [0, 1, 1]
```

**Step-by-step:**

* Prefix 0 → binary `0` → decimal `0` → `0 % 5 == 0` → `true`
* Prefix 0..1 → `01` → decimal `1` → `1 % 5 != 0` → `false`
* Prefix 0..2 → `011` → decimal `3` → `3 % 5 != 0` → `false`

**Output:**

```text
[true, false, false]
```

---

### Example 2

**Input:**

```text
nums = [1, 1, 1]
```

**Step-by-step:**

* Prefix 0 → `1` → decimal `1` → not divisible by 5 → `false`
* Prefix 0..1 → `11` → decimal `3` → not divisible by 5 → `false`
* Prefix 0..2 → `111` → decimal `7` → not divisible by 5 → `false`

**Output:**

```text
[false, false, false]
```

---

## How to use / Run locally

You can copy the solution for your preferred language into a file and run it with your own test harness or use an online judge like LeetCode.

### C++

1. Save the solution class in `solution.cpp`.
2. Write a `main()` that constructs `Solution` and calls `prefixesDivBy5`.
3. Compile and run:

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

---

### Java

1. Save as `Solution.java`.
2. Add a `main` method for testing if needed.
3. Compile and run:

```bash
javac Solution.java
java Solution
```

---

### JavaScript (Node.js)

1. Save as `solution.js`.
2. Export or directly call `prefixesDivBy5`.
3. Run with Node:

```bash
node solution.js
```

---

### Python3

1. Save as `solution.py`.
2. Create a test at the bottom:

```python
if __name__ == "__main__":
    print(Solution().prefixesDivBy5([0,1,1]))
```

3. Run:

```bash
python3 solution.py
```

---

### Go

1. Save as `main.go`.
2. Add `package main` and a `main()` function calling `prefixesDivBy5`.
3. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* I never build the full integer from the binary array; I only keep track of the remainder modulo 5.
* This avoids:

  * Overflow issues.
  * Expensive big integer operations.
* The `% 5` operation ensures `rem` always stays in the range `[0, 4]`.
* `ans.reserve(nums.size())` in C++ and `new ArrayList<>(nums.length)` in Java are small micro-optimizations to avoid resizing the output container repeatedly.
* The logic is identical in all languages, which makes it easy to port and test.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
