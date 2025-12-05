# 3432. Count Partitions with Even Sum Difference

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

You are given an integer array `nums` of length `n`.

A partition is defined as an index `i` where `0 < i < n`, splitting the array into two **non-empty** subarrays:

* Left subarray: `nums[0..i-1]`
* Right subarray: `nums[i..n-1]`

For each partition, we compute:

* `leftSum`  = sum of the left subarray
* `rightSum` = sum of the right subarray
* `diff`     = `leftSum - rightSum`

We need to count how many partitions have `diff` **even**.

Return the number of such partitions.

---

## Constraints

* `2 <= n == nums.length <= 100`
* `1 <= nums[i] <= 100`

The constraints are very small, but the solution we use is still optimal and works for much larger sizes.

---

## Intuition

I started by thinking:
“For every cut in the array, I need the difference between the sum of the left part and the sum of the right part to be even.”

For a partition at index `i`:

* `leftSum`  = sum of `nums[0..i-1]`
* `rightSum` = sum of `nums[i..n-1]`
* `total`    = sum of the whole array

Then:

```text
rightSum = total - leftSum
diff = leftSum - rightSum
     = leftSum - (total - leftSum)
     = 2 * leftSum - total
```

Now I only care if `diff` is **even or odd**.

* `2 * leftSum` is always even.
* So the parity of `diff` depends only on `total`:

  * If `total` is **even** → `even - even` = even → every partition gives **even difference**.
  * If `total` is **odd**  → `even - odd`  = odd  → every partition gives **odd difference**.

So surprisingly:

> Either **all** partitions are valid or **none** are valid.

That means I don’t need to check each partition one-by-one at all.

---

## Approach

1. **Compute total sum** of the array:
   `total = sum(nums)`

2. **Check parity of total**:

   * If `total` is **odd**:

     * For every possible partition, `diff = 2 * leftSum - total` will always be odd.
     * So there is **no valid partition**.
     * Return `0`.
   * If `total` is **even**:

     * For every possible partition, `diff` will always be even.
     * The number of possible cut positions in an array of length `n` is `n - 1`.
     * So return `n - 1`.

Only one pass for the sum, one parity check, and one simple formula.

---

## Data Structures Used

* Just simple variables:

  * `total` – to store the sum of all elements.
  * `n` – length of the array.
* No extra arrays, no complex data structures.
* Memory usage stays constant (`O(1)` extra).

---

## Operations & Behavior Summary

1. **Summation**

   * Traverse the array once.
   * Add each element to `total`.

2. **Parity Check**

   * If `total % 2 != 0` → return `0`.

3. **Count Valid Partitions** (when total is even)

   * Number of valid partitions = number of ways to cut the array into two non-empty parts.
   * That is simply `n - 1`.

Behavior:

* If array sum is odd → answer is `0`.
* If array sum is even → answer is `n - 1`.

---

## Complexity

* **Time Complexity:**
  `O(n)` – I scan the array once to calculate the total sum.
  Here `n` is the length of `nums`.

* **Space Complexity:**
  `O(1)` – I only use a few integer variables, no extra data structures whose size depends on `n`.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int countPartitions(vector<int>& nums) {
        long long total = 0;
        // 1. Compute the total sum of the array
        for (int x : nums) {
            total += x;
        }
        
        // 2. If total sum is odd, no partition can give an even difference
        if (total % 2 != 0) return 0;
        
        // 3. If total is even, all (n - 1) partitions are valid
        int n = (int)nums.size();
        return n - 1;
    }
};
```

---

### Java

```java
class Solution {
    public int countPartitions(int[] nums) {
        long total = 0;
        // 1. Compute the total sum of the array
        for (int x : nums) {
            total += x;
        }
        
        // 2. If total sum is odd, no partition can give an even difference
        if ((total & 1L) == 1L) return 0;
        
        // 3. If total is even, all (n - 1) partitions are valid
        int n = nums.length;
        return n - 1;
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
var countPartitions = function(nums) {
    let total = 0;
    // 1. Compute the total sum of the array
    for (const x of nums) {
        total += x;
    }
    
    // 2. If total sum is odd, no partition can give an even difference
    if (total % 2 !== 0) return 0;
    
    // 3. If total is even, all (n - 1) partitions are valid
    const n = nums.length;
    return n - 1;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def countPartitions(self, nums: List[int]) -> int:
        # 1. Compute the total sum of the array
        total = sum(nums)
        
        # 2. If total sum is odd, no partition can give an even difference
        if total % 2 == 1:
            return 0
        
        # 3. If total is even, all (n - 1) partitions are valid
        n = len(nums)
        return n - 1
```

---

### Go

```go
package main

func countPartitions(nums []int) int {
    // 1. Compute the total sum of the array
    total := 0
    for _, x := range nums {
        total += x
    }

    // 2. If total sum is odd, no partition can give an even difference
    if total%2 != 0 {
        return 0
    }

    // 3. If total is even, all (n - 1) partitions are valid
    n := len(nums)
    return n - 1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is identical in all languages, so I’ll explain it once and point to how it matches each version.

### 1. Calculate the total sum

C++:

```cpp
long long total = 0;
for (int x : nums) {
    total += x;
}
```

* I create a variable `total` and start from `0`.
* I loop through each element `x` in `nums`.
* I add `x` to `total`.
* After the loop, `total` contains the sum of all elements.

Java / JS / Python / Go do exactly the same thing with their own syntax.

### 2. Check if the total sum is odd

C++:

```cpp
if (total % 2 != 0) return 0;
```

* `% 2` gives the remainder when dividing by 2.
* If remainder is `1`, then the number is odd.
* If `total` is odd:

  * For any partition, `diff = 2 * leftSum - total`.
  * `2 * leftSum` is always even.
  * `even - odd` is odd → difference is always odd.
* So there is **no** valid partition.
* I immediately return `0`.

Same idea:

* Java: `if ((total & 1L) == 1L) return 0;`
* JS: `if (total % 2 !== 0) return 0;`
* Python: `if total % 2 == 1: return 0`
* Go: `if total%2 != 0 { return 0 }`

### 3. Count the number of valid partitions when total is even

If the code didn’t return in step 2, that means `total` is even.

C++:

```cpp
int n = (int)nums.size();
return n - 1;
```

Why `n - 1`?

* Suppose `nums = [a0, a1, a2, ..., a_{n-1}]`.
* Valid partitions must split the array into **two non-empty** parts.
* Possible cut positions are:

  * after index 0 → between elements 0 and 1 (i = 1)
  * after index 1 → between elements 1 and 2 (i = 2)
  * ...
  * after index n-2 → between elements n-2 and n-1 (i = n-1)
* So there are `(n - 1)` ways to cut.

Since `total` is even, for any cut position:

```text
diff = 2 * leftSum - total
```

* `2 * leftSum` is even.
* `total` is even.
* `even - even = even`, so all cuts are valid.

Therefore the answer is simply `n - 1`.

Each language returns the same thing:

* Java: `int n = nums.length; return n - 1;`
* JS: `const n = nums.length; return n - 1;`
* Python: `n = len(nums); return n - 1`
* Go: `n := len(nums); return n - 1`

---

## Examples

### Example 1

```text
Input:  nums = [10, 10, 3, 7, 6]
total = 10 + 10 + 3 + 7 + 6 = 36 (even)

Number of possible partitions = n - 1 = 5 - 1 = 4

All 4 partitions will have even difference.

Output: 4
```

### Example 2

```text
Input:  nums = [1, 2, 2]
total = 1 + 2 + 2 = 5 (odd)

Since total is odd, no partition can have an even difference.

Output: 0
```

### Example 3

```text
Input:  nums = [2, 4, 6, 8]
total = 2 + 4 + 6 + 8 = 20 (even)

Number of possible partitions = n - 1 = 4 - 1 = 3

All 3 partitions are valid.

Output: 3
```

---

## How to use / Run locally

Assume you have cloned this repo and are inside the project folder.

### C++

```bash
g++ -std=c++17 -O2 solution.cpp -o solution
./solution
```

Make sure `main()` in your local test file calls `Solution().countPartitions(nums)` with test input.

---

### Java

```bash
javac Solution.java
java Solution
```

Where `Solution.java` contains the `Solution` class and a `main` method that tests `countPartitions`.

---

### JavaScript (Node.js)

```bash
node solution.js
```

In `solution.js`, you can require or define the function and log output like:

```javascript
console.log(countPartitions([10, 10, 3, 7, 6]));
```

---

### Python3

```bash
python3 solution.py
```

In `solution.py`, create an instance and call it:

```python
print(Solution().countPartitions([10, 10, 3, 7, 6]))
```

---

### Go

```bash
go run main.go
```

In `main.go`, call `countPartitions` and print the result in `main()`.

---

## Notes & Optimizations

* The big insight is the algebraic simplification:

  ```text
  diff = leftSum - rightSum
       = 2 * leftSum - total
  ```

* Because `2 * leftSum` is always even, the parity of `diff` is fully controlled by the parity of `total`.

* This turns a potentially `O(n^2)` idea (checking all partitions and computing sums) into a **very simple** `O(n)` solution.

* No prefix sums array is needed. A single pass and one conditional is enough.

* This pattern is useful in many competitive programming problems:

  * When you care only about **parity** (even/odd), always try to reduce expressions and see what really matters.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
