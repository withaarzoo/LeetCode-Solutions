# Minimum Removals to Balance Array

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

I am given an integer array `nums` and an integer `k`.

An array is called **balanced** if:

```bash
maximum element ≤ minimum element × k
```

I can remove any number of elements from the array, but the array must **not become empty**.

My task is to return the **minimum number of elements** I need to remove so that the remaining array becomes balanced.

A single-element array is always considered balanced.

---

## Constraints

* 1 ≤ nums.length ≤ 10⁵
* 1 ≤ nums[i] ≤ 10⁹
* 1 ≤ k ≤ 10⁵

---

## Intuition

When I first looked at this problem, I realized that removing elements one by one would be inefficient.

So instead of thinking **what to remove**, I changed my thinking to **what to keep**.

If I can find the **largest group of elements** that already satisfies the balance condition, then:

```bash
minimum removals = total elements − kept elements
```

To make checking min and max easier, sorting the array felt like the natural first step.

---

## Approach

1. I sort the array.
2. I use a sliding window with two pointers.
3. The left pointer represents the minimum element.
4. The right pointer expands the window.
5. I keep expanding while:

   ```bash
   nums[right] ≤ nums[left] × k
   ```

6. If the condition breaks, I move the left pointer forward.
7. I track the maximum size of a valid window.
8. Final answer is:

   ```bash
   n − maximum valid window size
   ```

---

## Data Structures Used

* Array (after sorting)
* Two pointers (sliding window technique)

No extra data structures are required.

---

## Operations & Behavior Summary

* Sorting ensures minimum and maximum values are easy to track.
* Sliding window guarantees linear traversal after sorting.
* Each element is visited at most twice.
* The algorithm always keeps the largest valid balanced subarray.

---

## Complexity

**Time Complexity:**
`O(n log n)`

* `n log n` for sorting
* `O(n)` for sliding window traversal

**Space Complexity:**
`O(1)`

* No extra space used apart from sorting

---

## Multi-language Solutions

### C++

```cpp
class Solution {
public:
    int minRemoval(vector<int>& nums, int k) {
        sort(nums.begin(), nums.end());

        int left = 0;
        int maxKeep = 1;

        for (int right = 0; right < nums.size(); right++) {
            while ((long long)nums[right] > (long long)nums[left] * k) {
                left++;
            }
            maxKeep = max(maxKeep, right - left + 1);
        }

        return nums.size() - maxKeep;
    }
};
```

---

### Java

```java
class Solution {
    public int minRemoval(int[] nums, int k) {
        Arrays.sort(nums);

        int left = 0;
        int maxKeep = 1;

        for (int right = 0; right < nums.length; right++) {
            while ((long) nums[right] > (long) nums[left] * k) {
                left++;
            }
            maxKeep = Math.max(maxKeep, right - left + 1);
        }

        return nums.length - maxKeep;
    }
}
```

---

### JavaScript

```javascript
var minRemoval = function(nums, k) {
    nums.sort((a, b) => a - b);

    let left = 0;
    let maxKeep = 1;

    for (let right = 0; right < nums.length; right++) {
        while (nums[right] > nums[left] * k) {
            left++;
        }
        maxKeep = Math.max(maxKeep, right - left + 1);
    }

    return nums.length - maxKeep;
};
```

---

### Python3

```python
class Solution:
    def minRemoval(self, nums, k):
        nums.sort()
        left = 0
        max_keep = 1

        for right in range(len(nums)):
            while nums[right] > nums[left] * k:
                left += 1
            max_keep = max(max_keep, right - left + 1)

        return len(nums) - max_keep
```

---

### Go

```go
func minRemoval(nums []int, k int) int {
    sort.Ints(nums)

    left := 0
    maxKeep := 1

    for right := 0; right < len(nums); right++ {
        for nums[right] > nums[left]*k {
            left++
        }
        if right-left+1 > maxKeep {
            maxKeep = right - left + 1
        }
    }

    return len(nums) - maxKeep
}
```

---

## Step-by-step Detailed Explanation

1. I sort the array so that minimum and maximum values are easy to compare.
2. I start a sliding window from index 0.
3. The left pointer always represents the minimum element.
4. I expand the right pointer to include more elements.
5. If the balance condition breaks, I move the left pointer.
6. I calculate the window size after every valid step.
7. The largest valid window gives me the maximum elements I can keep.
8. Removing the rest gives the minimum removals.

---

## Examples

**Example 1**

```bash
Input: nums = [2,1,5], k = 2
Output: 1
```

I remove `5`.
Remaining array `[1,2]` is balanced.

---

**Example 2**

```bash
Input: nums = [1,6,2,9], k = 3
Output: 2
```

I remove `1` and `9`.
Remaining array `[2,6]` is balanced.

---

## How to Use / Run Locally

1. Copy the code for your preferred language.
2. Paste it into the LeetCode editor or your local IDE.
3. Provide the input array and value of `k`.
4. Run and get the minimum removals.

---

## Notes & Optimizations

* Brute force approaches will time out.
* Sorting + sliding window is the optimal strategy.
* Use `long long` or `long` to avoid overflow.
* Single-element arrays are always valid.

---

## Author

**Md Aarzoo Islam**
🔗 [https://bento.me/withaarzoo](https://bento.me/withaarzoo)
