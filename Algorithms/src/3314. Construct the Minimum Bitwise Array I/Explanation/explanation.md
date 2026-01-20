# Construct the Minimum Bitwise Array I (LeetCode 3314)

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

I am given an array `nums` consisting of `n` **prime numbers**.

My task is to create a new array `ans` of the same length such that for every index `i`:

```bash
ans[i] OR (ans[i] + 1) == nums[i]
```

Important conditions:

* I must **minimize** the value of `ans[i]`
* If it is **not possible**, I should return `-1` for that index

Each index is **independent**, so I can solve them one by one.

---

## Constraints

* `1 <= nums.length <= 100`
* `2 <= nums[i] <= 1000`
* `nums[i]` is always a **prime number**

---

## Intuition

When I saw this problem, I didn’t jump into complex bit manipulation.

Instead, I noticed:

* The constraints are **very small**
* Each `nums[i]` is independent
* I only need the **smallest** valid value

So I thought:

> “Why not just try all possible values of `x` starting from `0` and stop when the condition is satisfied?”

Because `nums[i]` is at most `1000`, this approach is fast, safe, and very easy to understand.

---

## Approach

1. Create an empty result array `ans`
2. For each number `p` in `nums`:

   * Try all values of `x` from `0` to `p`
   * Check if `x OR (x + 1) == p`
   * If yes, store `x` and stop searching
3. If no valid `x` is found, store `-1`
4. Return the final array

This guarantees:

* Minimum value of `ans[i]`
* Correctness
* Simple and readable logic

---

## Data Structures Used

* Array / List
  (To store the result values)

No extra complex data structures are required.

---

## Operations & Behavior Summary

| Operation     | Purpose                              |                                 |
| ------------- | ------------------------------------ | ------------------------------- |
| Bitwise OR (` | `)                                   | To check the required condition |
| Looping       | To find the minimum valid value      |                                 |
| Break         | To stop once smallest value is found |                                 |

---

## Complexity

**Time Complexity:** `O(n × k)`

* `n` = length of `nums`
* `k` = maximum value of `nums[i]` (≤ 1000)

**Space Complexity:** `O(n)`

* Only the output array is used

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    vector<int> minBitwiseArray(vector<int>& nums) {
        vector<int> ans;

        for (int p : nums) {
            int found = -1;
            for (int x = 0; x <= p; x++) {
                if ((x | (x + 1)) == p) {
                    found = x;
                    break;
                }
            }
            ans.push_back(found);
        }
        return ans;
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
            int found = -1;

            for (int x = 0; x <= p; x++) {
                if ((x | (x + 1)) == p) {
                    found = x;
                    break;
                }
            }
            ans[i] = found;
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var minBitwiseArray = function(nums) {
    let ans = [];

    for (let p of nums) {
        let found = -1;
        for (let x = 0; x <= p; x++) {
            if ((x | (x + 1)) === p) {
                found = x;
                break;
            }
        }
        ans.push(found);
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def minBitwiseArray(self, nums: List[int]) -> List[int]:
        ans = []

        for p in nums:
            found = -1
            for x in range(p + 1):
                if (x | (x + 1)) == p:
                    found = x
                    break
            ans.append(found)

        return ans
```

---

### Go

```go
func minBitwiseArray(nums []int) []int {
    ans := make([]int, len(nums))

    for i, p := range nums {
        found := -1
        for x := 0; x <= p; x++ {
            if (x | (x + 1)) == p {
                found = x
                break
            }
        }
        ans[i] = found
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation

1. I process **one number at a time**
2. I try values of `x` starting from `0`
3. For each `x`, I calculate `x OR (x + 1)`
4. If it equals the given prime number:

   * I store `x`
   * I stop searching further
5. If no value matches:

   * I store `-1`

This guarantees the **smallest possible value** every time.

---

## Examples

### Example 1

**Input**

```bash
nums = [2, 3, 5, 7]
```

**Output**

```bash
[-1, 1, 4, 3]
```

---

### Example 2

**Input**

```bash
nums = [11, 13, 31]
```

**Output**

```bash
[9, 12, 15]
```

---

## How to use / Run locally

### C++

```bash
g++ solution.cpp && ./a.out
```

### Java

```bash
javac Solution.java && java Solution
```

### JavaScript

```bash
node solution.js
```

### Python

```bash
python solution.py
```

### Go

```bash
go run solution.go
```

---

## Notes & Optimizations

* Brute force is **intentional and optimal** here
* Constraints are small, so no extra bit hacks needed
* Code is interview-safe and beginner-friendly
* Very easy to debug and explain

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
