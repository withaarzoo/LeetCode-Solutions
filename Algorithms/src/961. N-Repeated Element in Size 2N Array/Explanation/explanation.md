# 961. N-Repeated Element in Size 2N Array

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

I am given an integer array `nums` of size `2N`.
The array contains `N + 1` unique elements.
Out of these, **only one element is repeated exactly `N` times**, and all other elements appear only once.

My task is to **find and return the element that is repeated `N` times**.

---

## Constraints

* `2 <= N <= 5000`
* `nums.length = 2 * N`
* `0 <= nums[i] <= 10^4`
* Exactly one number is repeated `N` times
* The answer is always guaranteed

---

## Intuition

When I looked at the problem, I realized something important.

If there are `2N` elements but only `N + 1` unique values, then **one value must appear again and again**.

So I thought:
If I scan the array from left to right and remember what I have already seen,
the **first number that appears again must be the repeated one**.

There is no need to count frequencies fully.

---

## Approach

1. I create an empty set.
2. I traverse the array one element at a time.
3. For each element:

   * If it already exists in the set, I return it immediately.
   * Otherwise, I add it to the set.
4. The problem guarantees an answer, so I will always find it.

---

## Data Structures Used

* **Set / HashSet / Map**

  * Used to store numbers that I have already seen
  * Helps detect duplicates in constant time

---

## Operations & Behavior Summary

* Traverse array once
* Check presence in set
* Return on first duplicate
* No sorting
* No nested loops

---

## Complexity

**Time Complexity:** `O(n)`
I traverse the array only once, where `n` is the length of the array.

**Space Complexity:** `O(n)`
I use a set to store visited elements.

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int repeatedNTimes(vector<int>& nums) {
        unordered_set<int> seen;

        for (int x : nums) {
            if (seen.count(x)) {
                return x;
            }
            seen.insert(x);
        }
        return -1;
    }
};
```

---

### Java

```java
class Solution {
    public int repeatedNTimes(int[] nums) {
        HashSet<Integer> seen = new HashSet<>();

        for (int x : nums) {
            if (seen.contains(x)) {
                return x;
            }
            seen.add(x);
        }
        return -1;
    }
}
```

---

### JavaScript

```javascript
var repeatedNTimes = function(nums) {
    const seen = new Set();

    for (let x of nums) {
        if (seen.has(x)) {
            return x;
        }
        seen.add(x);
    }
};
```

---

### Python3

```python
class Solution:
    def repeatedNTimes(self, nums):
        seen = set()

        for x in nums:
            if x in seen:
                return x
            seen.add(x)
```

---

### Go

```go
func repeatedNTimes(nums []int) int {
    seen := make(map[int]bool)

    for _, x := range nums {
        if seen[x] {
            return x
        }
        seen[x] = true
    }
    return -1
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

1. I create a set to store elements I have already seen.
2. I start looping through the array.
3. For each element:

   * I check if it already exists in the set.
   * If yes, I immediately return it.
   * If not, I store it in the set.
4. Since only one element repeats `N` times, the first duplicate is always correct.
5. I do not need to check the rest of the array.

This logic is the same in all languages.

---

## Examples

**Example 1**

```bash
Input: [1,2,3,3]
Output: 3
```

**Example 2**

```bash
Input: [2,1,2,5,3,2]
Output: 2
```

**Example 3**

```bash
Input: [5,1,5,2,5,3,5,4]
Output: 5
```

---

## How to use / Run locally

1. Copy the code for your preferred language.
2. Paste it into your editor or online compiler.
3. Provide the input array as required by the platform.
4. Run the program to get the repeated element.

This solution works directly on LeetCode and local environments.

---

## Notes & Optimizations

* This solution is already optimal for clarity and reliability.
* An `O(1)` space solution exists using math or index tricks, but it is harder to read.
* For interviews and real-world use, this approach is **clean, safe, and recommended**.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
