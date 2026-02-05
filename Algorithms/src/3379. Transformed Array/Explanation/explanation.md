# Transformed Array (LeetCode 3379)

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
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given an integer array `nums` that behaves like a **circular array**.
My task is to create a new array `result` of the same size.

For each index `i`:

* If `nums[i] > 0`, I move `nums[i]` steps **to the right**
* If `nums[i] < 0`, I move `abs(nums[i])` steps **to the left**
* If `nums[i] == 0`, I stay at the same index

After moving, I copy the value from the landing index into `result[i]`.

Because the array is circular, moving past the ends wraps around.

---

## Constraints

* 1 ≤ `nums.length` ≤ 100
* -100 ≤ `nums[i]` ≤ 100

---

## Intuition

When I read the problem, I realized something important:

Each index works **independently**.

So I don’t need to simulate step-by-step movement.
Instead, I can directly calculate **where I will land** using math.

Since the array is circular, **modulo (`%`)** is enough to handle wrapping.

Once I know the final index, I just copy the value from `nums`.

---

## Approach

1. Store the size of the array as `n`
2. Create a new array `result` of size `n`
3. Loop through every index `i`
4. If `nums[i] == 0`, store it directly
5. Otherwise:

   * Calculate the target index using `(i + nums[i]) % n`
   * Fix negative index if needed
6. Store `nums[targetIndex]` in `result[i]`
7. Return the result array

---

## Data Structures Used

* Array / Vector
* No extra data structures required

---

## Operations & Behavior Summary

* Right movement → positive value
* Left movement → negative value
* Circular behavior handled using modulo
* One pass solution

---

## Complexity

**Time Complexity:** `O(n)`

* I process each index exactly once

**Space Complexity:** `O(n)`

* Extra array used to store the result

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> constructTransformedArray(vector<int>& nums) {
        int n = nums.size();
        vector<int> result(n);

        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) {
                result[i] = nums[i];
            } else {
                int target = (i + nums[i]) % n;
                if (target < 0) target += n;
                result[i] = nums[target];
            }
        }
        return result;
    }
};
```

---

### Java

```java
class Solution {
    public int[] constructTransformedArray(int[] nums) {
        int n = nums.length;
        int[] result = new int[n];

        for (int i = 0; i < n; i++) {
            if (nums[i] == 0) {
                result[i] = nums[i];
            } else {
                int target = (i + nums[i]) % n;
                if (target < 0) target += n;
                result[i] = nums[target];
            }
        }
        return result;
    }
}
```

---

### JavaScript

```javascript
var constructTransformedArray = function(nums) {
    const n = nums.length;
    const result = new Array(n);

    for (let i = 0; i < n; i++) {
        if (nums[i] === 0) {
            result[i] = nums[i];
        } else {
            let target = (i + nums[i]) % n;
            if (target < 0) target += n;
            result[i] = nums[target];
        }
    }
    return result;
};
```

---

### Python3

```python
class Solution:
    def constructTransformedArray(self, nums):
        n = len(nums)
        result = [0] * n

        for i in range(n):
            if nums[i] == 0:
                result[i] = nums[i]
            else:
                target = (i + nums[i]) % n
                result[i] = nums[target]

        return result
```

---

### Go

```go
func constructTransformedArray(nums []int) []int {
    n := len(nums)
    result := make([]int, n)

    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            result[i] = nums[i]
        } else {
            target := (i + nums[i]) % n
            if target < 0 {
                target += n
            }
            result[i] = nums[target]
        }
    }
    return result
}
```

---

## Step-by-step Detailed Explanation

For each index `i`:

```cpp
int target = (i + nums[i]) % n;
```

* Adds movement to current index
* `% n` keeps the index inside bounds

```cpp
if (target < 0) target += n;
```

* Fixes negative index when moving left

```cpp
result[i] = nums[target];
```

* Copies value from landing index

This logic is identical in all languages.

---

## Examples

**Input**

```bash
nums = [3, -2, 1, 1]
```

**Output**

```bash
[1, 1, 1, 3]
```

**Explanation**

* Index 0 → move right 3 → index 3 → value 1
* Index 1 → move left 2 → index 3 → value 1
* Index 2 → move right 1 → index 3 → value 1
* Index 3 → move right 1 → index 0 → value 3

---

## How to use / Run locally

1. Copy the solution code
2. Paste it into your preferred language environment
3. Call the function with the input array
4. Print or return the result

---

## Notes & Optimizations

* No nested loops
* No simulation of steps
* Modulo handles circular movement efficiently
* Works for both positive and negative values

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
