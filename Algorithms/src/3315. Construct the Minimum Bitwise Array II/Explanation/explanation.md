# Construct the Minimum Bitwise Array II (LeetCode 3315)

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
* How to Use / Run Locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given an array `nums` consisting of **prime numbers**.

Your task is to create another array `ans` such that for every index `i`:

```bash
ans[i] OR (ans[i] + 1) == nums[i]
```

Additional rules:

* Each `ans[i]` must be **as small as possible**
* If no such value exists, set `ans[i] = -1`
* Each index is independent (no relation between `ans[i]` and `ans[i+1]`)

---

## Constraints

* `1 <= nums.length <= 100`
* `2 <= nums[i] <= 10^9`
* `nums[i]` is a **prime number**

---

## Intuition

When I saw the condition
`x OR (x + 1) = p`
I immediately thought about **binary behavior**.

Key observations I made:

* Adding `1` to a number flips the **rightmost continuous 1s**
* OR operation restores bits that were flipped
* `x OR (x + 1)` is **always odd**
* So if `nums[i]` is even (like `2`), the answer is impossible

Then I realized something important:

For odd numbers, there is **exactly one smallest bit** that must be removed so that:

* `(x + 1)` flips it back
* OR reconstructs the original number

This means:

* No brute force
* No trying multiple values
* The answer can be calculated directly using bit tricks

---

## Approach

For each number `p` in `nums`:

1. Compute the removable bit:

   ``` bash
   removable = ((p + 1) & ~p) >> 1
   ```

2. If `removable == 0`

   * No valid answer exists
   * Set result to `-1`

3. Otherwise

   * Remove the bit using XOR

   ``` bash
   answer = p ^ removable
   ```

4. Do this in-place for all elements

This guarantees:

* Correct OR condition
* Minimum possible value
* Constant-time operation

---

## Data Structures Used

* No extra data structures
* Only integer bit operations
* In-place array modification

---

## Operations & Behavior Summary

| Operation | Purpose                                      |
| --------- | -------------------------------------------- |
| `p + 1`   | Finds the first zero bit after trailing ones |
| `~p`      | Identifies removable bit positions           |
| `&`       | Isolates the correct bit                     |
| `>> 1`    | Targets the bit to remove                    |
| `^`       | Removes that bit                             |
| Bit logic | Ensures minimal valid value                  |

---

## Complexity

* **Time Complexity:** `O(n)`

  * Each element is processed once
  * All operations are constant time

* **Space Complexity:** `O(1)`

  * No extra memory used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> minBitwiseArray(vector<int>& nums) {
        for (int &p : nums) {
            int removable = ((p + 1) & ~p) >> 1;
            p ^= removable | (-(removable == 0) & ~p);
        }
        return nums;
    }
};
```

---

### Java

```java
class Solution {
    public int[] minBitwiseArray(List<Integer> nums) {
        int[] ans = new int[nums.size()];

        for (int i = 0; i < nums.size(); i++) {
            int p = nums.get(i);
            int removable = ((p + 1) & ~p) >> 1;

            if (removable == 0) ans[i] = -1;
            else ans[i] = p ^ removable;
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var minBitwiseArray = function(nums) {
    for (let i = 0; i < nums.length; i++) {
        let p = nums[i];
        let removable = ((p + 1) & ~p) >> 1;
        nums[i] = removable === 0 ? -1 : (p ^ removable);
    }
    return nums;
};
```

---

### Python3

```python
class Solution:
    def minBitwiseArray(self, nums):
        for i in range(len(nums)):
            p = nums[i]
            removable = ((p + 1) & ~p) >> 1
            nums[i] = -1 if removable == 0 else p ^ removable
        return nums
```

---

### Go

```go
func minBitwiseArray(nums []int) []int {
    for i, p := range nums {
        removable := ((p + 1) & ^p) >> 1
        if removable == 0 {
            nums[i] = -1
        } else {
            nums[i] = p ^ removable
        }
    }
    return nums
}
```

---

## Step-by-step Detailed Explanation

### Example: `nums[i] = 13`

Binary:

```bash
13  -> 1101
14  -> 1110
~13 -> 0010
```

Step 1:

```bash
(p + 1) & ~p = 0010
```

Step 2:

```bash
removable = 0010 >> 1 = 0001
```

Step 3:

```bash
13 ^ 1 = 12
```

Final check:

```bash
12 OR 13 = 13
```

✔ Condition satisfied
✔ Minimum value achieved

---

### Why `2` fails

```bash
2  -> 10
3  -> 11
~2 -> 01
(p + 1) & ~p = 01
removable >> 1 = 0
```

No removable bit → answer is `-1`

---

## Examples

### Input

```bash
nums = [2, 3, 5, 7]
```

### Output

```bash
[-1, 1, 4, 3]
```

---

### Input

```bash
nums = [11, 13, 31]
```

### Output

```bash
[9, 12, 15]
```

---

## How to Use / Run Locally

1. Copy any language solution
2. Paste into LeetCode or your local compiler
3. Call `minBitwiseArray(nums)`
4. Print the result

No special libraries required.

---

## Notes & Optimizations

* No brute force
* No looping over candidates
* Pure bit manipulation
* Interview-grade solution
* Works for max constraints safely

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
