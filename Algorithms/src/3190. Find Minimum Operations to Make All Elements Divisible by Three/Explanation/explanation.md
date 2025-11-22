# 3190. Find Minimum Operations to Make All Elements Divisible by Three

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

## Problem Summary

You are given an integer array `nums`.

In one operation, you can **add or subtract 1** from **any one element** of `nums`.

Your goal is to find the **minimum number of operations** needed to make **every element** of the array **divisible by 3**.

You must return this minimum number of operations.

---

## Constraints

* `1 <= nums.length <= 50`
* `1 <= nums[i] <= 50`

The input size is small, but the logic should still be clean and optimal.

---

## Intuition

I started by asking a simple question:

> “If I look at just **one number**, what is the minimum number of operations to make it divisible by 3?”

Because one operation only affects **one element**, each number can be handled **independently** of the others.

For any number `x`:

* If `x % 3 == 0` → already divisible by 3 → **0 operations**.
* If `x % 3 == 1` →

  * `x - 1` is divisible by 3 → **1 operation**.
* If `x % 3 == 2` →

  * `x + 1` is divisible by 3 → **1 operation**.

So, any number that is **not divisible by 3** will always take **exactly 1 operation** to fix.

This leads to a very simple idea:

> **Count how many numbers are not divisible by 3**.
> That count is the minimum number of operations.

---

## Approach

1. Initialize a variable `operations = 0`.
2. Loop through every element `x` in `nums`.
3. For each `x`:

   * If `x % 3 != 0`, it means `x` is not divisible by 3.
   * Increase `operations` by 1 (because we need exactly one operation for this `x`).
4. After the loop ends, return `operations`.

No need to actually simulate adding or subtracting.
We only care that **one change is enough** per “bad” number.

---

## Data Structures Used

* Just a few **simple variables**:

  * `operations` to store the answer.
  * A loop variable to traverse the array.
* **No extra arrays, lists, stacks, or maps** are used.

So the memory usage is constant.

---

## Operations & Behavior Summary

* **Single allowed operation:**

  * Choose one index `i`.
  * Replace `nums[i]` with `nums[i] + 1` **or** `nums[i] - 1`.

* **Behavior on each number:**

  * If a number is already divisible by 3 → cost `0`.
  * If remainder is `1` or `2` when divided by 3 → cost `1`.
  * We never need more than one move for a number because the **nearest multiple of 3** is at distance **1** when `x % 3 != 0`.

* **Global behavior:**

  * Total answer = **number of elements with `x % 3 != 0`**.

---

## Complexity

* **Time Complexity:** `O(n)`

  * `n` = number of elements in `nums`.
  * We loop over the array once and do `O(1)` work for each element.

* **Space Complexity:** `O(1)`

  * We use a constant amount of extra memory (just integers), no extra data structures.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minimumOperations(vector<int>& nums) {
        int operations = 0;

        // Traverse each number in the array
        for (int x : nums) {
            // If x is not divisible by 3, we need exactly 1 operation
            if (x % 3 != 0) {
                operations++;
            }
        }

        return operations;
    }
};
```

---

### Java

```java
class Solution {
    public int minimumOperations(int[] nums) {
        int operations = 0;

        // Go through each element in the array
        for (int x : nums) {
            // If x % 3 is not 0, one operation is needed
            if (x % 3 != 0) {
                operations++;
            }
        }

        return operations;
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
var minimumOperations = function(nums) {
    let operations = 0;

    // Loop through the array
    for (const x of nums) {
        // If remainder when divided by 3 is not zero, need 1 operation
        if (x % 3 !== 0) {
            operations++;
        }
    }

    return operations;
};
```

---

### Python3

```python
from typing import List

class Solution:
    def minimumOperations(self, nums: List[int]) -> int:
        operations = 0

        # Check each number in nums
        for x in nums:
            # If x is not divisible by 3, we need one operation
            if x % 3 != 0:
                operations += 1

        return operations
```

---

### Go

```go
package main

func minimumOperations(nums []int) int {
    operations := 0

    // Iterate over all numbers in the slice
    for _, x := range nums {
        // If x % 3 != 0, it takes exactly 1 operation to fix x
        if x%3 != 0 {
            operations++
        }
    }

    return operations
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

Here I’ll walk through the logic in a language-agnostic way.
All implementations above follow this same pattern.

### 1. Initialize the counter

We start with:

```text
operations = 0
```

* This variable will count how many numbers need to be changed.
* Initially, we assume nothing needs to be changed.

### 2. Loop over the array

Pseudo-code:

```text
for each x in nums:
    ...
```

* We visit each number one by one.
* Every number is processed independently.

### 3. Check divisibility by 3

Inside the loop:

```text
if x % 3 != 0:
    operations += 1
```

* `x % 3` gives the remainder when `x` is divided by 3.
* If the remainder is:

  * `0` → divisible by 3 → no action needed.
  * `1` or `2` → not divisible by 3 → we need **one operation**.
* Why exactly one?

  * If `x % 3 == 1`, then `x - 1` is divisible by 3.
  * If `x % 3 == 2`, then `x + 1` is divisible by 3.
  * Both are just 1 step away.

So for every “bad” number (not divisible by 3), we increment `operations` by 1.

### 4. Return the final answer

After the loop finishes:

```text
return operations
```

* `operations` now equals the count of numbers that were not divisible by 3.
* And this count is the **minimum** number of operations needed because:

  * Each such number must be changed at least once.
  * One change is always enough.
  * Operations don’t overlap between elements.

So this is both **correct** and **optimal**.

---

## Examples

### Example 1

**Input:**

```text
nums = [1, 2, 3, 4]
```

Check each element:

* `1 % 3 = 1` → needs 1 operation (e.g., `1 - 1 = 0`)
* `2 % 3 = 2` → needs 1 operation (e.g., `2 + 1 = 3`)
* `3 % 3 = 0` → needs 0 operations
* `4 % 3 = 1` → needs 1 operation (e.g., `4 - 1 = 3`)

Total operations = `1 + 1 + 0 + 1 = 3`

**Output:**

```text
3
```

---

### Example 2

**Input:**

```text
nums = [3, 6, 9]
```

All elements are already divisible by 3:

* `3 % 3 = 0`
* `6 % 3 = 0`
* `9 % 3 = 0`

Total operations = `0`

**Output:**

```text
0
```

---

## How to use / Run locally

Below are simple commands for running the solutions locally.
(Assuming you save each code snippet into its own file.)

### C++

1. Save the solution in `solution.cpp`.
2. Compile and run:

```bash
g++ -std=c++17 solution.cpp -o solution
./solution
```

You can create a `main()` function to test:

```cpp
int main() {
    Solution sol;
    vector<int> nums = {1, 2, 3, 4};
    cout << sol.minimumOperations(nums) << endl; // Should print 3
    return 0;
}
```

---

### Java

1. Save the solution in `Solution.java`.
2. Add a `main` method for testing:

```java
public class Main {
    public static void main(String[] args) {
        Solution sol = new Solution();
        int[] nums = {1, 2, 3, 4};
        System.out.println(sol.minimumOperations(nums)); // Should print 3
    }
}
```

3. Compile and run:

```bash
javac Main.java Solution.java
java Main
```

---

### JavaScript (Node.js)

1. Save the function in `solution.js`.
2. Add test code:

```javascript
const nums = [1, 2, 3, 4];
console.log(minimumOperations(nums)); // Should print 3
```

3. Run:

```bash
node solution.js
```

---

### Python3

1. Save the solution in `solution.py`.
2. Add test code:

```python
if __name__ == "__main__":
    sol = Solution()
    nums = [1, 2, 3, 4]
    print(sol.minimumOperations(nums))  # Should print 3
```

3. Run:

```bash
python3 solution.py
```

---

### Go

1. Save the solution in `main.go`.

Example:

```go
package main

import "fmt"

func minimumOperations(nums []int) int {
    operations := 0
    for _, x := range nums {
        if x%3 != 0 {
            operations++
        }
    }
    return operations
}

func main() {
    nums := []int{1, 2, 3, 4}
    fmt.Println(minimumOperations(nums)) // Should print 3
}
```

2. Run:

```bash
go run main.go
```

---

## Notes & Optimizations

* The solution is already **optimal**:

  * Each element is processed once → `O(n)`.
  * No extra memory → `O(1)`.
* We don’t actually modify the array values because:

  * It’s enough to know **how many** need changes.
* This pattern (checking remainder with `%`) is very common for questions about divisibility.

If the constraints were huge (like `10^7` elements), this solution would still scale well because it’s linear in the size of the input.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
